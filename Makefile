TARGET = ./bin/tjsj.dev
SRC_FILES := $(shell find ./cmd ./pkg -type f)

${TARGET}: $(SRC_FILES)
	$(info Building ${TARGET}...)
	@go build -o $@ ./cmd/tjsj.dev

.PHONY: clean
clean:
	$(info Cleaning...)
	@rm -rf $(dir ${TARGET})
