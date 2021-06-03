.PHONY: default all
APPS     := generator model-builder sequence-analysis
BLDDIR   ?= bin
BLDOPT   ?=

.EXPORT_ALL_VARIABLES:
GO111MODULE  = on

default: build

build: clean $(APPS)

$(BLDDIR)/%:
	go build $(BLDOPT) -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

clean:
	@mkdir -p $(BLDDIR)
	@for app in $(APPS) ; do \
		rm -f $(BLDDIR)/$$app ; \
	done

deps:
	@echo 'Installing go modules...'
	@go mod download

test:
	go test -cover ./...

lint-revive:
	@echo 'Linting with revive...'
	@revive -formatter stylish -config=revive.toml ./...

lint-golangci:
	@echo 'Linting with golangci...'
	@golangci-lint run ./pkg/...

lint: format lint-revive lint-golangci
