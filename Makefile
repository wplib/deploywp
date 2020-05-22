PKGREFLECT := $(shell which pkgreflect)
GODOCMD := $(shell which godocdown)
VERSION := $(shell tools/getVersion.sh)
COMMENT := $(shell tools/getComment.sh)

all:
	@echo "Current launch version is:	v$(VERSION)"
	@echo "Last commit message is:		'$(COMMENT)'"
	@#echo git tag -a v$(VERSION) -m '"$(COMMENT)"'
	@#echo git push origin v$(VERSION)
	@echo ""
	@echo "build		- Build for local testing."
	@echo "release		- Build for published release."
	@echo "push		- Push repo to GitHub."
	@echo ""
	@echo "build-docker	- Build Docker cotainer."
	@echo "doc		- Build GoLang documentation."
	@echo "test-run		- Run a test using native GoLang."
	@echo "test-template	- Run a test using GoLang template."
	@echo "test-docker	- Run a test using native GoLang within container."
	@echo "pkgreflect	- Install 'pkgreflect' - Used to auto-discover exported functions."


build:
	@make pkgreflect
	@goreleaser --snapshot --skip-publish --rm-dist


release:
	@make pkgreflect
	@goreleaser --rm-dist


push:
	@echo "Pushing to: $(shell git branch)"
	@git config core.hooksPath .git-hooks
	@git add .
	@git commit .
	@git push


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
	@$(PKGREFLECT) -notests jsonTemplate/helpers


doc:
ifeq ($(GODOCMD),)
	@echo "godocdown - Installing"
	@go install github.com/robertkrimen/godocdown/godocdown
else
	@echo "godocdown - already installed here $(GODOCMD)"
endif
	@$(GODOCMD) 

