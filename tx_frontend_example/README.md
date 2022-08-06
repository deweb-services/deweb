
## Install it

```sh
npm install
npm install ts-proto protoc --save-dev
```

## Prepare messages

Load libs proto for SDK 0.45.4
```

curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.45.4/proto/cosmos/base/query/v1beta1/pagination.proto -o ./proto/cosmos/base/query/v1beta1/pagination.proto

mkdir -p ./proto/google/api
curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.45.4/third_party/proto/google/api/annotations.proto -o ./proto/google/api/annotations.proto
curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.45.4/third_party/proto/google/api/http.proto -o ./proto/google/api/http.proto
mkdir -p ./proto/gogoproto
curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.45.4/third_party/proto/gogoproto/gogo.proto -o ./proto/gogoproto/gogo.proto

```

Generate messages for Deweb. Before generate update proto files from PROJECT_ROOT to ./proto
```
mkdir -p client/src/types/generated

ls ./proto/deweb | xargs -I {} ./node_modules/protoc/protoc/bin/protoc \
--plugin="./node_modules/.bin/protoc-gen-ts_proto" \
--ts_proto_out="./client/src/types/generated" \
--proto_path="./proto" \
--ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
deweb/{}

ls ./proto/dns_module | xargs -I {} ./node_modules/protoc/protoc/bin/protoc \
--plugin="./node_modules/.bin/protoc-gen-ts_proto" \
--ts_proto_out="./client/src/types/generated" \
--proto_path="./proto" \
--ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
dns_module/{}
```

## Prepare config

Update file ```testnet.alice.mnemonic.key``` with mnemonic for account for test.

## Run it

```sh
npm run experiment
```
