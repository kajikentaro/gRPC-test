PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
OUT_DIR="./generated"

gen_grpc_front:
	protoc \
		--plugin="protoc-gen-ts=$(PROTOC_GEN_TS_PATH)" \
		--js_out="import_style=commonjs,binary:$(OUT_DIR)" \
		--ts_out="$(OUT_DIR)" \
		helloworld.proto