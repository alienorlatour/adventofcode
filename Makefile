init:
	go mod vendor

test:
	go test ./...

build:
	go build -o adventofcode.out cmd/main.go

run: build
run:
	[ -n "$(filter-out $@,$(MAKECMDGOALS))" ] || exit 1
	@echo "Running Advent of Code for December $(filter-out $@,$(MAKECMDGOALS))"
	./adventofcode.out -day $(filter-out $@,$(MAKECMDGOALS))
