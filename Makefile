PKGREFLECT := $(shell which pkgreflect)
GODOCMD := $(shell which godocdown)

all:
	@cat Makefile
	@echo git tag -a v1.4 -m "Better error handling"
	@echo git push origin v1.4

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


test:
	@./bin/deploywp release --chdir --json tests/deploywp.json


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

