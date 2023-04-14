CURRENT_DIR=$(shell pwd)

gen-proto-module:
	./scripts/gen_proto.sh ${CURRENT_DIR} && chmod 744 ./scripts/rm_omit_empty.sh && ./scripts/rm_omit_empty.sh ${CURRENT_DIR}

swag-init:
	swag init -g api/api.go -o api/docs