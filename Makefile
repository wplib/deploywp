PKGREFLECT := $(shell which pkgreflect)
GODOCMD := $(shell which godocdown)
BINARY := $(shell tools/getBinary.sh)
VERSION := $(shell tools/getVersion.sh)
COMMENT := $(shell tools/getComment.sh)


all:
	@echo "Current launch version is:	v$(VERSION)"
	@echo "Last commit message is:		'$(COMMENT)'"
	@echo ""
	@echo "build		- Build for local testing."
	@echo "release		- Build for published release."
	@echo "push		- Push repo to GitHub."
	@echo "doc		- Build GoLang documentation."
	@echo ""
	@echo "build-docker	- Build Docker cotainer."
	@echo "test-run		- Run a test using native GoLang."
	@echo "test-template	- Run a test using GoLang template."
	@echo "test-docker	- Run a test using native GoLang within container."
	@echo "pkgreflect	- Install 'pkgreflect' - Used to auto-discover exported functions."


build:
	@echo "Current $(BINARY) version is v$(VERSION)"
	@make pkgreflect
	@goreleaser --snapshot --skip-publish --rm-dist


release:
	@echo "Current $(BINARY) version is v$(VERSION)"
	@make pkgreflect
	@make push
	@git tag -a v$(VERSION) -m '"Release v$(VERSION)"'
	@git push origin v$(VERSION)
	@goreleaser --rm-dist


push:
	@echo "Pushing to: $(shell git branch)"
	@git config core.hooksPath .git-hooks
	@git add .
	@git commit -m '"$(COMMENT)"' .
	@git push


doc:
ifeq ($(GODOCMD),)
	@echo "godocdown - Installing"
	@go install github.com/robertkrimen/godocdown/godocdown
else
	@echo "godocdown - already installed here $(GODOCMD)"
endif
	@$(GODOCMD) 

build-docker:
	@make -C docker


test-run:
	@./bin/deploywp run --chdir --json tests/deploywp.json prod


test-template:
	#@./bin/deploywp load tests/deploywp.json prod
	@./tests/deploywp.tmpl prod


test-docker:
	@launch -q -p tests -m tests shell deploywp:0.9 deploywp run


pkgreflect:
ifeq ($(PKGREFLECT),)
	@echo "pkgreflect - Installing"
	@go install github.com/ungerik/pkgreflect
else
	@echo "pkgreflect - already installed here $(PKGREFLECT)"
endif
	@$(PKGREFLECT) -notests jtc/helpers

