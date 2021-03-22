TARGET = ./bin/tjsj

${TARGET}: $(shell find ./src -type f)
	$(info Building ${TARGET}...)
	@go build -o $@ ./src

.PHONY: run
run: ${TARGET}
	$(info Running ${TARGET}...)
	@${TARGET}

.PHONY: clean
clean:
	$(info Cleaning...)
	@rm -rf $(dir ${TARGET})
