################################################################################
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

all:
	@echo "build		- Build for local testing."
	@echo "release		- Build for published release."
	@echo "push		- Push repo to GitHub."
	@echo ""
	@echo "build-docker	- Build Docker cotainer."
	@echo "test-run		- Run a test using native GoLang."
	@echo "test-template	- Run a test using GoLang template."
	@echo "test-docker	- Run a test using native GoLang within container."
	@echo ""
	@$(BUILDTOOL) get all

build:
	@make pkgreflect
	@$(BUILDTOOL) build

release:
	@make pkgreflect
	@$(BUILDTOOL) release

push:
	@make pkgreflect
	@$(BUILDTOOL) push

pkgreflect:
	@$(BUILDTOOL) pkgreflect jtc/helpers

build-docker:
	@make -C docker

test-run:
	@./bin/deploywp run --chdir --json tests/deploywp.json prod

test-template:
	#@./bin/deploywp load tests/deploywp.json prod
	@./tests/deploywp.tmpl prod

test-docker:
	@launch -q -p tests -m tests shell deploywp:0.9 deploywp run

