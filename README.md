The `Go` Programming Language
========================

![Go's mascot, designed by Renée French.](https://upload.wikimedia.org/wikipedia/commons/2/23/Golang.png)
Go's mascot, designed by Renée French.

History
------
Go originated as an experiment by Robert Griesemer, Rob Pike and Ken Thompson at Google, to design a new systems programming language, with the following desiderata; the new language was to[16]

* be statically typed, scalable to large systems (as Java and C++);
* be "light on the page" (like dynamic languages);
* support networking and multiprocessing.

In later interviews, all three of the language designers cited their shared dislike of C++'s complexity as a reason to design a new language.

Notable Users
------------

Some notable open-source applications in Go include:

* Docker, a set of tools for deploying Linux containers
* Doozer, a lock service by managed hosting provider Heroku
* Juju, a service orchestration tool by Canonical, packagers of Ubuntu Linux
* Syncthing, an open-source file synchronization client/server application
* Packer, a tool for creating identical machine images for multiple platforms from a single source configuration

Other companies and sites using Go (generally together with other languages, not exclusively) include:

* Google, for many projects, notably including download server dl.google.com
* Dropbox, migrated some of their critical components from Python to Go
* CloudFlare, for their delta-coding proxy Railgun, their distributed DNS service, as well as tools for cryptography, logging, stream processing, and accessing SPDY sites.
* SoundCloud, for "dozens of systems"
* Splice, for the entire backend (API and parsers) of their online music collaboration platform.
* Cloud Foundry, a platform as a service
* MongoDB, tools for administering MongoDB instances
* Chango, a programmatic advertising company uses Go in its real-time bidding systems.
* SendGrid, a Boulder, Colorado-based transactional email delivery and management service.
* Plug.dj, an interactive online social music streaming website.
* ThoughtWorks, some tools and applications around continuous delivery and instant messages.

Language Tools
--------------
Go includes the same sort of debugging, testing, and code-vetting tools as many language distributions. The Go distribution includes, among other tools,

* `go build`, which builds Go binaries using only information in the source files themselves, no separate makefiles
* `go test`, for unit testing and microbenchmarks
* `go fmt`, for formatting code
* `go get`, for retrieving and installing remote packages
* `go vet`, a static analyzer looking for potential errors in code
* `go run`, a shortcut for building and executing code
* `godoc`, for displaying documentation or serving it via HTTP
* `gorename`, for renaming variables, functions, and so on in a type-safe way
* `go generate`, a standard way to invoke code generators

Features
------------
Go was designed with the intent of improving the conciseness, simplicity and safety of C

* A syntax and environment adopting patterns more common in dynamic languages:
  * Optional concise variable declaration and initialization through type inference
    * (`x := 0` not `int x = 0;`).
  * Fast compilation times.
  * Remote package management (`go get`: similar to `npm install`) and online package documentation.
* Distinctive approaches to particular problems:
  * Built-in concurrency primitives: light-weight processes (goroutines), channels, and the select statement.
  * An interface system in place of virtual inheritance, and type embedding instead of non-virtual inheritance.
  * A toolchain that, by default, produces statically linked native binaries without external dependencies.
* A desire to keep the language specification simple enough to hold in a programmer's head, in part by omitting features common to similar languages (see below)

Omissions
----------
The designers of Go deliberately omitted certain features common to other languages, most notably,

* implementation inheritance (sub-/super-class relationship)
* compile-time generics
* pointer arithmetic
* implicit type conversions

Criticisms
-----------
Frequent criticisms assert that:

* lack of compile-time generics leads to repetition or excessive dynamic code
* lack of language extensibility (through, for instance, operator overloading) makes certain tasks more verbose

Version History
------------

| Version | Date |
| -: | - |
| 1.0 | 2012-03-28 |
| 1.1 | 2013-05-13 |
| 1.2 | 2013-12-01 |
| 1.3 | 2014-06-18 |
| 1.4 | 2014-12-10 |
| 1.5 | 2015-08-19 |
