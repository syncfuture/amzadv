protoc -I . --go_out=plugins=grpc,paths=source_relative:. campaignsmodel/*.proto
protoc-go-tags --dir=campaignsmodel