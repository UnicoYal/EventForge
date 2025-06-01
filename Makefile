.PHONY: wire


wire: bin/wire
	bin/wire ./internal/event_generator/wire

bin/wire:
	go build -mod=mod -o bin/wire github.com/google/wire/cmd/wire