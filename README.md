# OnlineFood

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=. --go-grpc_out=. proto/auth.proto

cd /d %KAFKA_HOME%
cd bin\windows
zookeeper-server-start.bat ..\..\config\zookeeper.properties


cd /d %KAFKA_HOME%
cd bin\windows
kafka-server-start.bat ..\..\config\server.properties
