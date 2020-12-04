
# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

init:
	go mod vendor

test:
	go test ./... -cover

build:
	go build -o adventofcode.out cmd/main.go

run: build
run:
	[ -n "$(RUN_ARGS)" ] || exit 1
	@echo "Running Advent of Code for December $(RUN_ARGS)"
	./adventofcode.out -day $(RUN_ARGS)
