all:
	@cat Makefile
	@echo git tag -a v1.4 -m "Better error handling"
	@echo git push origin v1.4

build:
	@goreleaser --snapshot --skip-publish --rm-dist

release:
	@goreleaser --rm-dist

push:
	@echo "Pushing to: $(shell git branch)"
	@git config core.hooksPath .git-hooks
	@git add .
	@git commit .
	@git push

