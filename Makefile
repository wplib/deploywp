################################################################################
SHELL=/bin/bash
ifeq (, $(shell which buildtool))
$(warning "Installing buildtool...")
$(warning "go get github.com/gearboxworks/buildtool")
$(shell go get github.com/gearboxworks/buildtool)
endif
BUILDTOOL := $(shell which buildtool)
ifeq (, $(BUILDTOOL))
$(error "No buildtool found...")
endif
################################################################################

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

all:
	@:

%:
	@:

################################################################################

help:
	@$(BUILDTOOL) $@ $(args)
	@echo ""
	@echo "build-docker	- Build Docker cotainer."
	@echo "test-run		- Run a test using native GoLang."
	@echo "test-template	- Run a test using GoLang template."
	@echo "test-docker	- Run a test using native GoLang within container."
	@echo ""
	@$(BUILDTOOL) get all

build:
	@$(BUILDTOOL) $@ $(args)

clone:
	@$(BUILDTOOL) $@ $(args)

commit:
	@$(BUILDTOOL) $@ $(args)

get:
	@$(BUILDTOOL) $@ $(args)

ghr:
	@$(BUILDTOOL) $@ $(args)

go:
	@$(BUILDTOOL) $@ $(args)

pkgreflect:
	@$(BUILDTOOL) $@ $(args)

pull:
	@$(BUILDTOOL) $@ $(args)

push:
	@$(BUILDTOOL) $@ $(args)

release:
	@$(BUILDTOOL) $@ $(args)

selfupdate:
	@$(BUILDTOOL) $@ $(args)

set:
	@$(BUILDTOOL) $@ $(args)

sync:
	@$(BUILDTOOL) $@ $(args)

version:
	@$(BUILDTOOL) $@ $(args)

vfsgen:
	@$(BUILDTOOL) $@ $(args)


################################################################################
build-docker:
	@make -C docker

test-run:
	@./bin/deploywp run --chdir --json tests/deploywp.json prod

test-template:
	#@./bin/deploywp load tests/deploywp.json prod
	@./tests/deploywp.tmpl prod

test-docker:
	@launch -q -p tests -m tests shell deploywp:0.9 deploywp run

