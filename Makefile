# Default - top level rule is what gets run when you run just 'make' without specifying a goal/target.
.DEFAULT_GOAL := help

# Make will delete the target of a rule if it has changed and its recipe exits with a nonzero exit status, just as it
# does when it receives a signal.
.DELETE_ON_ERROR:

# When a target is built, all lines of the recipe will be given to a single invocation of the shell rather than each
# line being invoked separately.
.ONESHELL:

# If this variable is not set, the program '/bin/sh' is used as the shell.
SHELL := bash

# The default value of .SHELLFLAGS is -c normally, or -ec in POSIX-conforming mode.
# Extra options are set for Bash:
#   -e             Exit immediately if a command exits with a non-zero status.
#   -u             Treat unset variables as an error when substituting.
#   -o pipefail    The return value of a pipeline is the status of the last command to exit with a non-zero status,
#                  or zero if no command exited with a non-zero status.
.SHELLFLAGS := -euo pipefail -c

# Eliminate use of Make's built-in implicit rules.
MAKEFLAGS += --no-builtin-rules

# Issue a warning message whenever Make sees a reference to an undefined variable.
MAKEFLAGS += --warn-undefined-variables

# Check that the version of Make running this file supports the .RECIPEPREFIX special variable.
# We set it to '>' to clarify inlined scripts and disambiguate whitespace prefixes.
# All script lines start with "> " which is the angle bracket and one space, with no tabs.
ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later.)
endif

.RECIPEPREFIX = >

# GNU Make knows how to execute several recipes at once.
# Normally, Make will execute only one recipe at a time, waiting for it to finish before executing the next.
# However, the '-j' or '--jobs' option tells Make to execute many recipes simultaneously.
# With no numeric argument to '--jobs', Make runs as many recipes simultaneously as possible.
MAKEFLAGS += --jobs

# Define a list of adapters as relative module paths.
ADAPTERS := $(shell find . -name 'go.mod' -exec sh -c 'echo {} | sed -e "s/go.mod/.../"' \;)

# Configure an 'all' target to cover the bases.
all: generate test lint build ## Generate and test and lint and build.
.PHONY: all

help: Makefile ## Display this list of available Make targets.
> @grep -E '(^[a-zA-Z/_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | sort | awk 'BEGIN { FS = ":.*?## " }; { printf "\033[32m%-30s\033[0m %s\n", $$1, $$2 }' | sed -e 's/\[32m##/[33m/'
.PHONY: help

# Set up some lazy initialisation functions to find code files, so that targets using the output of '$(shell ...)' only
# execute their respective shell commands when they need to, rather than every single instance of '$(shell ...)' being
# executed every single time 'make' is run for any target and wasting a lot of time.
# Further reading at https://www.oreilly.com/library/view/managing-projects-with/0596006101/ch10.html under the 'Lazy
# Initialization' heading.
find-go-files = $(shell find $1 -type f \( -iname '*.go' -or -name go.mod -or -name go.sum \))

GO_FILES = $(redefine-go-files) $(GO_FILES)

redefine-go-files = $(eval GO_FILES := $(call find-go-files, .))

# Automatic variables set by Make in use in this Makefile:
#   $@        - The file name of the target of the rule.
#   $(@D)     - The directory part of the file name of the target, with the trailing slash removed.
#   $(CURDIR) - Set to the absolute pathname of the current working directory.

# Tools used for various things further down.
hack/bin/gci:
> mkdir -p $(@D)
> GOBIN=$(CURDIR)/hack/bin go install github.com/daixiang0/gci@latest

hack/bin/go-junit-report:
> mkdir -p $(@D)
> GOBIN=$(CURDIR)/hack/bin go install github.com/jstemmer/go-junit-report/v2@f50ae24655f6484f175ceb0672505dfec6565637 # Pin to a specific version that fixes a parsing bug, but has not been released yet.

hack/bin/gofumpt:
> mkdir -p $(@D)
> GOBIN=$(CURDIR)/hack/bin go install mvdan.cc/gofumpt@latest

hack/bin/golangci-lint:
> mkdir -p $(@D)
> curl --fail --location --show-error --silent \
  https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
  | sh -s -- -b $(CURDIR)/hack/bin

hack/bin/mockery:
> mkdir -p $(@D)
> GOBIN=$(CURDIR)/hack/bin go install github.com/vektra/mockery/v2@latest

hack/bin/yq:
> mkdir -p $(@D)
> GOBIN=$(CURDIR)/hack/bin go install github.com/mikefarah/yq/v4@v4.34.2

# Tests look for sentinel files to determine whether or not they need to be run again.
# If any Go code file has been changed since the sentinel file was last touched, it will trigger a retest.
test: tmp/.tests-passed.sentinel ## Run tests. Will also generate.
test-cover: tmp/.cover-tests-passed.sentinel ## Run all tests with the race detector, and output a coverage profile.
bench: tmp/.benchmarks-ran.sentinel ## Run enough iterations of each benchmark to take ten seconds each.
report: tmp/.report-ran.sentinel ## Run tests, and produce a JUnit report.
.PHONY: test test-cover bench report

# Linter checks look for sentinel files to determine whether or not they need to check again.
# If any Go code file has been changed since the sentinel file was last touched, it will trigger a rerun.
lint: tmp/.linted.sentinel ## Lint all of the Go code. Will also generate and test.
.PHONY: lint

# If any Go code file has been changed since the binary was last built, it will trigger a rebuild.
build: tmp/.built.sentinel ## Build the library. Will also generate, test, and lint.
.PHONY: build

clean: ## Clean up Go's output (if any), test coverage data, Go workspace checksums, and output and temp sub-directories.
> go clean -x -v
> rm -rf cover.out go.work.sum out tmp
.PHONY: clean

clean-hack: ## Deletes all binaries under 'hack'.
> rm -rf hack/bin
.PHONY: clean-hack

clean-all: clean clean-hack ## Clean all of the things.
.PHONY: clean-all

tmp/.generated.sentinel: hack/bin/gofumpt hack/bin/mockery $(GO_FILES)
> mkdir -p $(@D)
> hack/bin/mockery
> hack/bin/gofumpt -w . # Mockery output needs to be gofumpt'd otherwise this check will fail.
> touch $@

# Tests - re-run if any Go files have changes since 'tmp/.tests-passed.sentinel' was last touched.
tmp/.tests-passed.sentinel: tmp/.generated.sentinel $(GO_FILES)
> mkdir -p $(@D)
> go test -count=1 -v ./... $(ADAPTERS)
> touch $@

tmp/.cover-tests-passed.sentinel: tmp/.generated.sentinel $(GO_FILES)
> mkdir -p $(@D)
> go test -count=1 -covermode=atomic -coverprofile=cover.out -race -v ./... $(ADAPTERS)
> touch $@

tmp/.benchmarks-ran.sentinel: tmp/.generated.sentinel $(GO_FILES)
> mkdir -p $(@D)
> go test -bench=. -benchmem -benchtime=10s -count=1 -run='^DoNotRunTests$$' -v ./... $(ADAPTERS)
> touch $@

tmp/.report-ran.sentinel: tmp/.generated.sentinel hack/bin/go-junit-report $(GO_FILES)
> mkdir -p $(@D)
> go test -count=1 -v ./... $(ADAPTERS) 2>&1 | hack/bin/go-junit-report -iocopy -out report.xml -set-exit-code
> touch $@

# Lint - re-run if the tests have been re-run (and so, by proxy, whenever the source files have changed).
# These checks are all read-only and will not make any changes.
tmp/.linted.sentinel: tmp/.linted.gofumpt.sentinel tmp/.linted.go.vet.sentinel tmp/.linted.golangci-lint.sentinel
> mkdir -p $(@D)
> touch $@

tmp/.linted.gofumpt.sentinel: hack/bin/gofumpt tmp/.tests-passed.sentinel
> mkdir -p $(@D)
> hack/bin/gofumpt -l . \
  | awk '{ print } END { if (NR != 0) { print "Please run \"make lint-fix-gofumpt\" to fix these issues!"; exit 1 } }'
> touch $@

tmp/.linted.go.vet.sentinel: tmp/.tests-passed.sentinel
> mkdir -p $(@D)
> go vet ./...
> touch $@

tmp/.linted.golangci-lint.sentinel: .golangci.yaml hack/bin/golangci-lint tmp/.tests-passed.sentinel
> mkdir -p $(@D)
> hack/bin/golangci-lint run ./... $(ADAPTERS)
> touch $@

lint-fix-gci: hack/bin/gci hack/bin/yq $(GO_FILES) ## Runs 'gci' to make imports deterministic using config from '.golangci.yaml'.
> sections="$(shell hack/bin/yq '.linters-settings.gci.sections | join(" -s ")' .golangci.yaml )"
> hack/bin/gci write . -s $${sections}
.PHONY: lint-fix-gci

lint-fix-gofumpt: hack/bin/gofumpt $(GO_FILES) ## Runs 'gofumpt -w' to format and simplify all Go code.
> hack/bin/gofumpt -w .
.PHONY: lint-fix-gofumpt

lint-fix-golangci-lint: hack/bin/golangci-lint $(GO_FILES) ## Runs 'golangci-lint run --fix' to auto-fix lint issues where supported.
> hack/bin/golangci-lint run --timeout=10m --fix ./... $(ADAPTERS)
.PHONY: lint-fix-golangci-lint

lint-fix: lint-fix-gci lint-fix-gofumpt lint-fix-golangci-lint ## Runs 'gci', 'gofumpt', and 'golangci-lint'.
.PHONY: lint-fix

# Re-build if the lint output is re-run (and so, by proxy, whenever the source files have changed).
tmp/.built.sentinel: tmp/.linted.sentinel
> mkdir -p $(@D)
> go build -v ./... $(ADAPTERS)
> touch $@

generate: tmp/.generated.sentinel ## Generate mocks.
.PHONY: generate

tidy: ## Run 'go mod tidy' on all Go modules.
> find . -name 'go.mod' -execdir go mod tidy \;
.PHONY: tidy
