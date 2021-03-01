package main

import (
	"bytes"
	"fmt"
	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/russross/blackfriday/v2"
)

// siteDir is the target directory into which the HTML gets generated. Its
// default is set here but can be changed by an argument passed into the
// program.
var siteDir = "./public"

func verbose() bool {
	return len(os.Getenv("VERBOSE")) > 0
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ensureDir(dir string) {
	err := os.MkdirAll(dir, 0755)
	check(err)
}

func copyFile(src, dst string) {
	dat, err := ioutil.ReadFile(src)
	check(err)
	err = ioutil.WriteFile(dst, dat, 0644)
	check(err)
}

func pipe(bin string, arg []string, src string) []byte {
	fmt.Printf("cmd: %s - %v\n", bin, arg)
	cmd := exec.Command(bin, arg...)
	in, err := cmd.StdinPipe()
	check(err)
	out, err := cmd.StdoutPipe()
	check(err)
	stderr, err := cmd.StderrPipe()
	check(err)
	err = cmd.Start()
	check(err)
	_, err = in.Write([]byte(src))
	check(err)
	err = in.Close()
	check(err)
	fileBytes, err := ioutil.ReadAll(out)
	check(err)
	out.Close()
	stderrB, err := ioutil.ReadAll(stderr)
	check(err)
	stderr.Close()
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("ERROR: %s - %s\n", stderrB, err)
		check(err)
	}
	//check(err)
	return fileBytes
}

func mustReadFile(path string) string {
	fileBytes, err := ioutil.ReadFile(path)
	check(err)
	return string(fileBytes)
}

func markdown(src string) string {
	return string(blackfriday.Run([]byte(src)))
}

func readLines(path string) []string {
	src := mustReadFile(path)
	return strings.Split(src, "\n")
}

func mustGlob(glob string) []string {
	paths, err := filepath.Glob(glob)
	check(err)
	return paths
}

func whichLexer(path string) string {
	if strings.HasSuffix(path, ".go") {
		return "go"
	} else if strings.HasSuffix(path, ".sh") {
		return "console"
	} else if strings.HasSuffix(path, ".lua") {
		return "lua"
	}
	panic("No lexer for " + path)
}

func debug(msg string) {
	if os.Getenv("DEBUG") == "1" {
		_, err := fmt.Fprintln(os.Stderr, msg)
		check(err)
	}
}

var docsPat = regexp.MustCompile(`^\s*(//|#|--)\s`) //"^\\s*(\\/\\/|#)\\s"
var dashPat = regexp.MustCompile(`-+"`)
var promptPat = regexp.MustCompile(`^\$\slua\s`)

// Seg is a segment of an example
type Seg struct {
	Docs         string
	DocsRendered string
	Code         string
	CodeRendered string
	CodeForJs    string
	CodeEmpty    bool
	CodeLeading  bool
	CodeRun      bool
}

// Example is info extracted from an example file
type Example struct {
	ID          string
	Name        string
	GoCode      string
	GoCodeHash  string
	URLHash     string
	Segs        [][]*Seg
	PrevExample *Example
	NextExample *Example
}

func parseSegs(sourcePath string) ([]*Seg, string) {

	log.Printf("====> Processing: %s", sourcePath)
	var (
		lines  []string
		source []string
	)

	eval := strings.HasSuffix(sourcePath, ".sh")
	wd, _ := os.Getwd()


	// Convert tabs to spaces for uniform rendering.
	for _, line := range readLines(sourcePath) {
		lines = append(lines, strings.Replace(line, "\t", "    ", -1))
		source = append(source, line)
		if eval && promptPat.MatchString(line) {
			l := promptPat.ReplaceAllString(line, "")
			log.Printf("Path: %s", path.Dir(sourcePath))
			os.Chdir(path.Dir(sourcePath))
			pwd, _ := os.Getwd()
			fmt.Printf("exe: |%s| pwd: %s\n", l, pwd)

			res := pipe("lua", strings.Split(l, " "), "")
			fmt.Printf("RESULT: %s", res)
			for _, out := range strings.Split(string(res), "\n") {
				lines = append(lines, out)
			}
		}
	}
	os.Chdir(wd)
	fmt.Printf("----------------\n\n")
	fileContent := strings.Join(source, "\n")
	segments := []*Seg{}
	lastSeen := ""
	for _, line := range lines {
		if line == "" {
			lastSeen = ""
			continue
		}
		matchDocs := docsPat.MatchString(line)
		matchCode := !matchDocs
		newDocs := (lastSeen == "") || ((lastSeen != "docs") && (segments[len(segments)-1].Docs != ""))
		newCode := (lastSeen == "") || ((lastSeen != "code") && (segments[len(segments)-1].Code != ""))
		if newDocs || newCode {
			debug("NEWSEG")
		}
		if matchDocs {
			trimmed := docsPat.ReplaceAllString(line, "")
			if newDocs {
				newSeg := Seg{Docs: trimmed, Code: ""}
				segments = append(segments, &newSeg)
			} else {
				segments[len(segments)-1].Docs = segments[len(segments)-1].Docs + "\n" + trimmed
			}
			debug("DOCS: " + line)
			lastSeen = "docs"
		} else if matchCode {
			if newCode {
				newSeg := Seg{Docs: "", Code: line}
				segments = append(segments, &newSeg)
			} else {
				segments[len(segments)-1].Code = segments[len(segments)-1].Code + "\n" + line
			}
			debug("CODE: " + line)
			lastSeen = "code"
		}
	}
	for i, seg := range segments {
		seg.CodeEmpty = seg.Code == ""
		seg.CodeLeading = i < (len(segments) - 1)

		seg.CodeRun = strings.Contains(seg.Code, "package main")
	}
	return segments, fileContent
}

func chromaFormat(code, filePath string) string {

	lexer := lexers.Get(filePath)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	if strings.HasSuffix(filePath, ".sh") {
		lexer = SimpleShellOutputLexer
	}

	lexer = chroma.Coalesce(lexer)

	style := styles.Get("swapoff")
	if style == nil {
		style = styles.Fallback
	}
	formatter := html.New(html.WithClasses(true))
	iterator, err := lexer.Tokenise(nil, string(code))
	check(err)
	buf := new(bytes.Buffer)
	err = formatter.Format(buf, style, iterator)
	check(err)
	return buf.String()
}

func parseAndRenderSegs(sourcePath string) ([]*Seg, string) {
	segs, filecontent := parseSegs(sourcePath)
	lexer := whichLexer(sourcePath)
	for _, seg := range segs {
		if seg.Docs != "" {
			//fmt.Printf("docs: |%s|\n", seg.Docs)
			seg.DocsRendered = markdown(seg.Docs)
			//fmt.Printf("docs: |%s| - |%s|\n", seg.Docs, seg.DocsRendered)
		}
		if seg.Code != "" {
			seg.CodeRendered = chromaFormat(seg.Code, sourcePath)

			// adding the content to the js code for copying to the clipboard
			if strings.HasSuffix(sourcePath, ".lua") {
				seg.CodeForJs = strings.Trim(seg.Code, "\n") + "\n"
			}
		}
	}
	// we are only interested in the 'go' code to pass to play.golang.org
	if lexer != "go" {
		filecontent = ""
	}
	return segs, filecontent
}

func parseExamples() []*Example {
	var exampleNames []string
	for _, line := range readLines("examples.txt") {
		if strings.HasPrefix(line, "###END") {
			break
		}
		if line != "" && !strings.HasPrefix(line, "#") {
			exampleNames = append(exampleNames, line)
		}
	}
	examples := make([]*Example, 0)
	for i, exampleName := range exampleNames {
		if verbose() {
			fmt.Printf("Processing %s [%d/%d]\n", exampleName, i+1, len(exampleNames))
		}
		example := Example{Name: exampleName}
		exampleID := strings.ToLower(exampleName)
		exampleID = strings.Replace(exampleID, " ", "-", -1)
		exampleID = strings.Replace(exampleID, "/", "-", -1)
		exampleID = strings.Replace(exampleID, "'", "", -1)
		exampleID = dashPat.ReplaceAllString(exampleID, "-")
		example.ID = exampleID
		example.Segs = make([][]*Seg, 0)
		sourcePaths := mustGlob("examples/" + exampleID + "/*")
		for _, sourcePath := range sourcePaths {

			sourceSegments, fileContent := parseAndRenderSegs(sourcePath)
			if fileContent != "" {
				example.GoCode = fileContent
			}
			example.Segs = append(example.Segs, sourceSegments)

		}
		examples = append(examples, &example)
	}
	for i, example := range examples {
		if i > 0 {
			example.PrevExample = examples[i-1]
		}
		if i < (len(examples) - 1) {
			example.NextExample = examples[i+1]
		}
	}
	return examples
}

func renderIndex(examples []*Example) {
	if verbose() {
		fmt.Println("Rendering index")
	}
	indexTmpl := template.New("index")
	_, err := indexTmpl.Parse(mustReadFile("templates/index.html"))
	check(err)
	indexF, err := os.Create(siteDir + "/index.html")
	check(err)
	err = indexTmpl.Execute(indexF, examples)
	check(err)
}

func renderExamples(examples []*Example) {
	if verbose() {
		fmt.Println("Rendering examples")
	}
	exampleTmpl := template.New("example")
	_, err := exampleTmpl.Parse(mustReadFile("templates/example.html"))
	check(err)
	for _, example := range examples {
		dir := siteDir + "/" + example.ID
		check(os.MkdirAll(dir, 0755))
		exampleF, err := os.Create(path.Join(dir, "index.html"))
		check(err)
		check(exampleTmpl.Execute(exampleF, example))
	}
}

func main() {
	if len(os.Args) > 1 {
		siteDir = os.Args[1]
	}
	ensureDir(siteDir)

	copyFile("templates/site.css", siteDir+"/site.css")
	copyFile("templates/site.js", siteDir+"/site.js")
	copyFile("templates/favicon.ico", siteDir+"/favicon.ico")
	copyFile("templates/404.html", siteDir+"/404.html")
	copyFile("templates/play.png", siteDir+"/play.png")
	copyFile("templates/clipboard.png", siteDir+"/clipboard.png")
	examples := parseExamples()
	renderIndex(examples)
	renderExamples(examples)
}

var SimpleShellOutputLexer = chroma.MustNewLexer(
	&chroma.Config{
		Name:      "Shell Output",
		Aliases:   []string{"console"},
		Filenames: []string{"*.sh"},
		MimeTypes: []string{},
	},
	chroma.Rules{
		"root": {
			// $ or > triggers the start of prompt formatting
			{`^\$`, chroma.GenericPrompt, chroma.Push("prompt")},
			{`^>`, chroma.GenericPrompt, chroma.Push("prompt")},

			// empty lines are just text
			{`^$\n`, chroma.Text, nil},

			// otherwise its all output
			{`[^\n]+$\n?`, chroma.GenericOutput, nil},
		},
		"prompt": {
			// when we find newline, do output formatting rules
			{`\n`, chroma.Text, chroma.Push("output")},
			// otherwise its all text
			{`[^\n]+$`, chroma.Text, nil},
		},
		"output": {
			// sometimes there isn't output so we go right back to prompt
			{`^\$`, chroma.GenericPrompt, chroma.Pop(1)},
			{`^>`, chroma.GenericPrompt, chroma.Pop(1)},
			// otherwise its all output
			{`[^\n]+$\n?`, chroma.GenericOutput, nil},
		},
	},
)
