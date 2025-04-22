set -e

# shellcheck disable=SC2046

PATH=${PATH}:/opt/bin:/root/go/bin

PROTOVALIDATE_DIR=protovalidate/proto/protovalidate
# Find all proto files in the src directory,
# but exclude the vendored google and grafeas directories
# https://stackoverflow.com/a/4210072
SOURCE_PROTO_FILES=$(find src -type d \( -path src/google -o -path src/grafeas -o  -path src/buf \) -prune -false -o -name '*.proto')
# shellcheck disable=SC2086
protoc -I${GOOGLEAPIS_DIR} -I${PROTOVALIDATE_DIR} \
    --proto_path=src \
    --go_out ./generated/go \
    --go_opt paths=source_relative \
    --go-grpc_out ./generated/go \
    --go-grpc_opt paths=source_relative $SOURCE_PROTO_FILES \
    --go-streamhandler_out packages=upm_beacon_stream:./generated/go \
    --go-streamhandler_opt paths=source_relative $SOURCE_PROTO_FILES
printf "Done\n\n"