TARGET = ./bin/tjsj
SRC := $(wildcard ./src/*.go)

${TARGET}: ${SRC}
	$(info Building ${TARGET}...)
	@mkdir -p $(dir ${TARGET})
	@go build -o $@ $^

.PHONY: run
run: ${TARGET}
	$(info Running ${TARGET}...)
	@${TARGET}

.PHONY: clean
clean:
	$(info Cleaning...)
	@rm -rf $(dir ${TARGET})
