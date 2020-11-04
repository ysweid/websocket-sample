BIN_PATH := ../bin/
PROJECT := $(shell basename $(CURDIR))
PACKAGE_EXT := .zip
PACKAGE := $(PROJECT)$(PACKAGE_EXT)

$(BIN_PATH)$(PACKAGE): $(BIN_PATH)$(PROJECT)
	@zip -u -j $@ $<

$(BIN_PATH)$(PROJECT):
	env GOOS=linux go build -ldflags="-s -w" -o $@

deploy: $(BIN_PATH)$(PACKAGE)
	@cd ..; sls deploy --verbose --function $(PROJECT)

clean:
	rm -f $(BIN_PATH)$(PROJECT) $(BIN_PATH)$(PACKAGE)

.PHONY: clean $(BIN_PATH)$(PROJECT)
