package tools

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/russross/blackfriday/v2"
)

func verbose() bool {
	return len(os.Getenv("VERBOSE")) > 0
}

func ensureDir(dir string) {
	err := os.MkdirAll(dir, 0755) //nolint:gomnd
	check(err)
}

func copyFile(src, dst string) {
	dat, err := ioutil.ReadFile(src)
	check(err)
	err = ioutil.WriteFile(dst, dat, 0644) //nolint:gomnd,gosec
	check(err)
}

func pipe(bin string, arg []string, stdin string) []byte {
	debug(fmt.Sprintf("cmd: %s - %v\n", bin, arg))
	cmd := exec.Command(bin, arg...)
	in, err := cmd.StdinPipe()
	check(err)
	out, err := cmd.StdoutPipe()
	check(err)
	stderr, err := cmd.StderrPipe()
	check(err)
	check(cmd.Start())

	_, err = in.Write([]byte(stdin))
	check(err)
	check(in.Close())

	fileBytes, err := ioutil.ReadAll(out)
	check(err)

	check(out.Close())

	stderrB, err := ioutil.ReadAll(stderr)
	check(err)

	check(stderr.Close())

	if err = cmd.Wait(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() != 3 {
			log.Printf("ERROR: %s - %s\n", stderrB, err)
			check(err)
		}
	}

	return fileBytes
}

func markdown(src string) string {
	return string(blackfriday.Run([]byte(src)))
}

func mustGlob(glob string) []string {
	paths, err := filepath.Glob(glob)
	check(err)

	return paths
}

func whichLexer(path string) string {
	switch {
	case strings.HasSuffix(path, ".go"):
		return "go"
	case strings.HasSuffix(path, ".sh"):
		return "console"
	case strings.HasSuffix(path, ".lua"):
		return "lua"
	default:
		panic("No lexer for " + path)
	}
}

func debug(msg string) {
	if os.Getenv("DEBUG") == "1" {
		_, err := fmt.Fprintln(os.Stderr, msg)
		check(err)
	}
}

var docsPat = regexp.MustCompile(`^\s*(//|#|--)\s`)
var dashPat = regexp.MustCompile(`-+"`)
var promptPat = regexp.MustCompile(`^\$\s`)

// Seg is a segment of an example.
type Seg struct {
	Docs         string
	DocsRendered template.HTML
	Code         string
	CodeRendered template.HTML
	CodeForJs    template.JSStr
	CodeEmpty    bool
	CodeLeading  bool
	CodeRun      bool
}

// Example is info extracted from an example file.
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
	debug(fmt.Sprintf("====> Processing: %s", sourcePath))
	var (
		lines  []string
		source []string
	)

	eval := strings.HasSuffix(sourcePath, ".sh")
	wd, _ := os.Getwd()

	// Convert tabs to spaces for uniform rendering.
	for _, line := range readLines(sourcePath) {
		lines = append(lines, strings.ReplaceAll(line, "\t", "    "))
		source = append(source, line)

		if eval && promptPat.MatchString(line) {
			l := promptPat.ReplaceAllString(line, "")

			debug(fmt.Sprintf("Path: %s", path.Dir(sourcePath)))

			_ = os.Chdir(path.Dir(sourcePath))
			pwd, _ := os.Getwd()
			debug(fmt.Sprintf("exe: |%s| pwd: %s\n", l, pwd))

			res := pipe("bash", []string{"-c", l}, "")
			debug(fmt.Sprintf("RESULT: %s", res))

			for _, out := range strings.Split(string(res), "\n") {
				lines = append(lines, strings.ReplaceAll(out, "\t", "    "))
			}
		}
	}

	_ = os.Chdir(wd)

	debug("----------------\n\n")

	fileContent := strings.Join(source, "\n")

	segments := make([]*Seg, 0)
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
				newSeg := Seg{ //nolint:exhaustivestruct
					Docs: trimmed,
					Code: "",
				}
				segments = append(segments, &newSeg)
			} else {
				segments[len(segments)-1].Docs = segments[len(segments)-1].Docs + "\n" + trimmed
			}

			debug("DOCS: " + line)

			lastSeen = "docs"
		} else if matchCode {
			if newCode {
				newSeg := Seg{ //nolint:exhaustivestruct
					Docs: "",
					Code: line,
				}
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
	SimpleShellOutputLexer := chroma.MustNewLexer(&chroma.Config{ //nolint:exhaustivestruct
		Name:      "Shell Output",
		Aliases:   []string{"console"},
		Filenames: []string{"*.sh"},
		MimeTypes: []string{},
	},
		chroma.Rules{
			"root": {
				// $ or > triggers the start of prompt formatting
				{Pattern: `^\$`, Type: chroma.GenericPrompt, Mutator: chroma.Push("prompt")},
				{Pattern: `^>`, Type: chroma.GenericPrompt, Mutator: chroma.Push("prompt")},

				// empty lines are just text
				{Pattern: `^$\n`, Type: chroma.Text, Mutator: nil},

				// otherwise its all output
				{Pattern: `[^\n]+$\n?`, Type: chroma.GenericOutput, Mutator: nil},
			},
			"prompt": {
				// when we find newline, do output formatting rules
				{Pattern: `\n`, Type: chroma.Text, Mutator: chroma.Push("output")},
				// otherwise its all text
				{Pattern: `[^\n]+$`, Type: chroma.Text, Mutator: nil},
			},
			"output": {
				// sometimes there isn't output so we go right back to prompt
				{Pattern: `^\$`, Type: chroma.GenericPrompt, Mutator: chroma.Pop(1)},
				{Pattern: `^>`, Type: chroma.GenericPrompt, Mutator: chroma.Pop(1)},
				// otherwise its all output
				{Pattern: `[^\n]+$\n?`, Type: chroma.GenericOutput, Mutator: nil},
			},
		},
	)

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
	iterator, err := lexer.Tokenise(nil, code)
	check(err)

	buf := new(bytes.Buffer)
	check(formatter.Format(buf, style, iterator))

	return buf.String()
}

func parseAndRenderSegs(sourcePath string) ([]*Seg, string) {
	segs, filecontent := parseSegs(sourcePath)
	lexer := whichLexer(sourcePath)

	for _, seg := range segs {
		if seg.Docs != "" {
			debug(fmt.Sprintf("docs: |%s|\n", seg.Docs))
			seg.DocsRendered = template.HTML(markdown(seg.Docs)) //nolint:gosec
			debug(fmt.Sprintf("docs: |%s| - |%s|\n", seg.Docs, seg.DocsRendered))
		}

		if seg.Code != "" {
			seg.CodeRendered = template.HTML(chromaFormat(seg.Code, sourcePath)) //nolint:gosec

			// adding the content to the js code for copying to the clipboard
			if strings.HasSuffix(sourcePath, ".lua") {
				seg.CodeForJs = template.JSStr(strings.Trim(seg.Code, "\n") + "\n")
			}
		}
	}
	// we are only interested in the 'go' code to pass to play.golang.org
	if lexer != "go" {
		filecontent = ""
	}

	return segs, filecontent
}

func ParseExamples() []*Example {
	exampleNames := GetListOfExamples("examples.txt")
	examples := make([]*Example, 0)

	for i, exampleName := range exampleNames {
		if verbose() {
			debug(fmt.Sprintf("Processing %s [%d/%d]\n", exampleName, i+1, len(exampleNames)))
		}

		example := ParseExample(exampleName)
		examples = append(examples, example)
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

func GetPrevNextExample(exampleID string) (prev *Example, next *Example) {
	exampleNames := GetListOfExamples("examples.txt")
	for i, exampleName := range exampleNames {
		if exampleID == BuildExampleId(exampleName) {
			if i-1 > 0 {
				prev = &Example{ //nolint:exhaustivestruct
					ID:   BuildExampleId(exampleNames[i-1]),
					Name: exampleNames[i-1],
				}
			}

			if i+1 < len(exampleNames)-1 {
				next = &Example{ //nolint:exhaustivestruct
					ID:   BuildExampleId(exampleNames[i+1]),
					Name: exampleNames[i+1],
				}
			}
		}
	}

	return // prev and/or next may be nil. They are set above using the named returns
}

func ParseExample(exampleName string) *Example {
	example := Example{Name: exampleName} //nolint:exhaustivestruct
	exampleID := BuildExampleId(exampleName)
	example.ID = exampleID
	example.Segs = make([][]*Seg, 0)
	sourcePaths := mustGlob("examples/" + exampleID + "/*")

	if len(sourcePaths) == 0 {
		log.Printf("No source found for exampleID %s\n", exampleID)
		os.Exit(1)
	}

	for _, sourcePath := range sourcePaths {
		fileInfo, err := os.Stat(sourcePath)
		check(err)
		// skip example subdirectories to allow for library files that are not published
		if fileInfo.IsDir() {
			continue
		}

		sourceSegments, fileContent := parseAndRenderSegs(sourcePath)
		if fileContent != "" {
			example.GoCode = fileContent
		}

		example.Segs = append(example.Segs, sourceSegments)
	}

	return &example
}

func GetListOfExamples(examplesList string) []string {
	var exampleNames []string

	for _, line := range readLines(examplesList) {
		if strings.HasPrefix(line, "###END") {
			break
		}

		if line != "" && !strings.HasPrefix(line, "#") {
			exampleNames = append(exampleNames, line)
		}
	}

	return exampleNames
}

func BuildExampleId(exampleName string) string {
	exampleID := strings.ToLower(exampleName)
	exampleID = strings.Replace(exampleID, " ", "-", -1)
	exampleID = strings.Replace(exampleID, "/", "-", -1)
	exampleID = strings.Replace(exampleID, "'", "", -1)
	exampleID = dashPat.ReplaceAllString(exampleID, "-")
	return exampleID
}

func GetExampleNamesAndIds(exampleListFile string) []*Example {
	result := make([]*Example, 0)
	for _, exampleName := range GetListOfExamples(exampleListFile) {
		result = append(result, &Example{
			Name: exampleName,
			ID:   BuildExampleId(exampleName),
		})
	}
	return result
}

func RenderIndex(indexFileHandle io.Writer, examples []*Example) {
	if verbose() {
		fmt.Println("Rendering index")
	}

	indexTmpl, err := template.ParseFiles("templates/index.html")
	check(err)
	err = indexTmpl.Execute(indexFileHandle, examples)
	check(err)
}

func RenderExamples(siteDir string, examples []*Example) {
	if verbose() {
		fmt.Println("Rendering examples")
	}

	exampleTmpl, err := template.ParseFiles("templates/example.html")
	check(err)
	for _, example := range examples {
		dir := siteDir + "/" + example.ID
		check(os.MkdirAll(dir, 0755))
		exampleF, err := os.Create(path.Join(dir, "index.html"))
		check(err)
		check(exampleTmpl.Execute(exampleF, example))
	}
}

func RenderExample(exampleFileHandle io.Writer, example *Example) {
	exampleTmpl, err := template.ParseFiles("templates/example.html")
	check(err)

	check(exampleTmpl.Execute(exampleFileHandle, example))
}

// Build the site into siteDir
// siteDir is the target directory into which the HTML gets generated.
// The default will be used if a string of zero length is passed.
func Generate(siteDir string) {

	if len(siteDir) == 0 {
		siteDir = "./public"
	}

	ensureDir(siteDir)

	copyFile("static/site.css", siteDir+"/site.css")
	copyFile("static/monokai.css", siteDir+"/monokai.css")
	copyFile("static/site.js", siteDir+"/site.js")
	copyFile("static/favicon.ico", siteDir+"/favicon.ico")
	copyFile("static/404.html", siteDir+"/404.html")
	copyFile("static/play.png", siteDir+"/play.png")
	copyFile("static/clipboard.png", siteDir+"/clipboard.png")

	examples := ParseExamples()

	indexFileHandle, err := os.Create(siteDir + "/index.html")
	check(err)
	RenderIndex(indexFileHandle, examples)
	RenderExamples(siteDir, examples)
}
