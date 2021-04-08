# protoc -I . --go_out=plugins=grpc,paths=source_relative:. campaignsmodel/*.proto
# protoc-go-tags --dir=campaignsmodel

# protoc --go_out=. --go_opt=paths=source_relative ./neworder/*.proto
# protoc-go-inject-tag --input=./neworder/neworder.pb.go
protoc -I . --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative campaignsmodel/*.proto
protoc-go-tags --dir=campaignsmodel
