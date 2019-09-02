# How to work with Go modules with 2 different major versions of a module
Demonstrating Go Modules in action imports of different major versions

(This project assumes the user is using a Go version which includes support for Go modules)

The setup of these projects is the following:

* Project A (This GitHub repository) is the main project which consumes package B in 2 different major versions and uses an exposed function which has a breaking change.

* Project B exposes a very simple function in 2 different setups, these are released in a certain way. I believe that setting up a specific branch for every major version above v1 is the way to go and is less error prone, allows for easier maintenance and on the plus side it doesn't break other depedency management tools other to Go modules like `dep`, `vendor` and others..

This main project consists of 2 functions which consume the different implementation of the function from B
both usages are covered by tests to ensure it's working from within this project aka - the consumer.

To run the tests, clone this repository to your local machine and then in a terminal run the following:

```
  go test -v
```

You should see something similar to the following:

```
➜  go-modules-project-a git:(master) go test -v
go: finding github.com/stretchr/testify v1.4.0
go: finding github.com/itamararjuan/go-modules-project-b v1.0.0
go: finding github.com/itamararjuan/go-modules-project-b/v2 v2.0.0
go: finding github.com/stretchr/objx v0.1.0
go: finding github.com/pmezard/go-difflib v1.0.0
go: finding github.com/davecgh/go-spew v1.1.0
go: finding gopkg.in/yaml.v2 v2.2.2
go: finding gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: downloading github.com/itamararjuan/go-modules-project-b v1.0.0
go: extracting github.com/itamararjuan/go-modules-project-b v1.0.0
go: downloading github.com/itamararjuan/go-modules-project-b/v2 v2.0.0
go: extracting github.com/itamararjuan/go-modules-project-b/v2 v2.0.0
=== RUN   TestSpeakerKnowsHowToSpeaks
--- PASS: TestSpeakerKnowsHowToSpeaks (0.00s)
=== RUN   TestSpeakerKnowsHowToSayMyName
--- PASS: TestSpeakerKnowsHowToSayMyName (0.00s)
PASS
ok  	github.com/itamararjuan/go-modules-project-a	0.012s
```

