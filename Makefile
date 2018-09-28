.PHONY: clean all deps version bindir gitconfig build test run-dev
GO ?= go
DEP = dep

ECHO ?= @echo
REMOVE ?= rm -rvf
MKDIR ?= mkdir -p
COPY ?= cp -v

EXPORTER_NAME=exporter
BIN_NAME=translate-bot
BIN_FOLDER=bin
REVISION_NAME=REVISION

GITURL ?= git@github.com

all: test build

deps: gitconfig
	$(DEP) ensure

bindir:
	$(MKDIR) $(BIN_FOLDER)

version:
	$(ECHO) `git describe --tags $(git log -n1 --pretty='%h')` > ./$(REVISION_NAME)

build: deps version bindir
	$(COPY) ./$(REVISION_NAME) $(BIN_FOLDER)/$(REVISION_NAME)
	$(GO) build -o $(BIN_FOLDER)/$(BIN_NAME)

test: deps
	$(GO) test ./...

clean:
	$(REMOVE) $(BIN_FOLDER)
