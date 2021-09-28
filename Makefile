LINKER_FLAGS = "-X 'go.dev.pztrn.name/periodicator/internal/config.Version=$(shell scripts/get_version.sh)'"
CONFIG ?= "./config.example.yaml"

build:
	go build -ldflags $(LINKER_FLAGS) -o periodicator .

generate-version:
	scripts/get_version.sh generate

run:
	go build -ldflags $(LINKER_FLAGS) -o periodicator .
	GPT_CONFIG=$(CONFIG) ./periodicator
	rm periodicator

run-version:
	go build -ldflags $(LINKER_FLAGS) -o periodicator .
	GPT_CONFIG=$(CONFIG) ./periodicator -version
	rm periodicator
