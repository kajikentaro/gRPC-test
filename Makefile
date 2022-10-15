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
init_grpc_front:
	cd frontend && yarn
	cd /tmp && wget https://github.com/grpc/grpc-web/releases/download/1.4.1/protoc-gen-grpc-web-1.4.1-linux-x86_64
	mv /tmp/protoc-gen-grpc-web-1.4.1-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web
	chmod +x /usr/local/bin/protoc-gen-grpc-web



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

start_dev_init:
	docker exec -it grpc_dev_front sh -c 'apk add git openssh-client'
	docker exec -it grpc_dev_front sh -c 'cd /root/ && git clone https://github.com/kajikentaro/gRPC-test /root/app'

start_dev_front:
	docker exec -itd grpc_dev_front sh -c 'cd /root/app/frontend && yarn && yarn dev'

start_dev_back:
	docker exec -itd grpc_dev_back sh -c 'cd /root/app/backend && go run main.go'

danger_rm:
	docker compose down --rmi all --volumes --remove-orphans