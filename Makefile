TJSJ_TARGET = ./bin/tjsj
SRC_FILES := $(shell find ./cmd ./pkg -type f)

${TJSJ_TARGET}: $(SRC_FILES)
	$(info Building ${TJSJ_TARGET}...)
	@go build -o $@ ./cmd/tjsj

.PHONY: clean
clean:
	$(info Cleaning...)
	@rm -rf $(dir ${TJSJ_TARGET})
