PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
OUT_DIR="./grpc_out"

gen_grpc_front:
	cd frontend && rm -rf $(OUT_DIR)
	cd frontend && mkdir $(OUT_DIR)
	cd frontend && protoc \
	--plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
	--js_out="import_style=commonjs,binary:${OUT_DIR}" \
	--ts_out="service=grpc-web:${OUT_DIR}" \
	--proto_path=../grpc/ \
	../grpc/*.proto