# Go by Example

Content and build toolchain for [Lua by Example](https://luabyexample.techplexlabs.com/),
a site that teaches Lua via annotated example programs.

### Overview

The Lua by Example site is built by extracting code and
comments from source files in `examples` and rendering
them via the `templates` into a static `public`
directory. The programs implementing this build process
are in `tools`, along with dependencies specified in
the `go.mod`file.

The built `public` directory can be served by any
static content system. The production site uses github pages.

### Building

To build the site you'll need Go and Lua installed. Run:

```console
$ tools/build
```

To build continuously in a loop:

```console
$ tools/build-loop
```

To see the site locally:

```console
$ tools/serve
```

and open `http://127.0.0.1:8000/` in your browser.

### Publishing

The site is published via a Github action

### License

This work is licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

### Thanks


Thanks to [Mark McGranaghan](https://markmcgranaghan.com/) for 
[GoByExample](https://gobyexample.com/) and
[Jeremy Ashkenas](https://github.com/jashkenas)
for [Docco](http://jashkenas.github.com/docco/), both of which
inspired this project.
