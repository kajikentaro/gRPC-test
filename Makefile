PROTOC_GEN_TS_PATH="./node_modules/.bin/protoc-gen-ts"
OUT_DIR="./grpc_out"

# improbable-eng/grpc-web
gen_grpc_front_old:
	cd frontend && rm -rf $(OUT_DIR)
	cd frontend && mkdir $(OUT_DIR)
	cd frontend && protoc \
	--plugin="protoc-gen-ts=${PROTOC_GEN_TS_PATH}" \
	--js_out="import_style=commonjs,binary:${OUT_DIR}" \
	--ts_out="service=grpc-web:${OUT_DIR}" \
	--proto_path=../grpc/ \
	../grpc/*.proto

# grpc/grpc-web
gen_grpc_front:
	cd frontend && rm -rf $(OUT_DIR)
	cd frontend && mkdir $(OUT_DIR)
	cd frontend && protoc \
		-I=../grpc \
		../grpc/*.proto \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:$(OUT_DIR) \
		--js_out=import_style=commonjs,binary:$(OUT_DIR)

gen_grpc_back:
	cd grpc && protoc \
		*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative
