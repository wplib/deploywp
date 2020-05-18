PKGREFLECT := $(shell which pkgreflect)

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
	@pkgreflect jsonTemplate/helpers


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

