RootDir = /home/rinai/PROJECTS/Some-WORKs/2025/04April/20250404/go-zero-test/service

UserDir = $(RootDir)/user
DocDir = $(RootDir)/doc

update:
	goctl rpc protoc $(UserDir)/rpc/desc/user.proto --proto_path=$(UserDir)/rpc/desc --go_out=$(UserDir)/rpc --go-grpc_out=$(UserDir)/rpc --zrpc_out=$(UserDir)/rpc --style=goZero && \
	goctl api go --api $(UserDir)/api/desc/user.api --dir $(UserDir)/api --style=goZero

model:
	mkdir -p $(UserDir)/rpc/internal/model && \
	goctl model mysql ddl -src $(UserDir)/rpc/internal/model/user.sql  --dir $(UserDir)/rpc/internal/model/
	
	