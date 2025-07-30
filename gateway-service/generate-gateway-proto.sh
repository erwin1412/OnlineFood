set -e

echo "Generating proto files for ALL services..."

echo " Auth Service"
protoc --go_out=. --go-grpc_out=. proto/auth.proto

echo " Courier Service"
protoc --go_out=. --go-grpc_out=. proto/courier.proto

echo " Food Service"
protoc --go_out=. --go-grpc_out=. proto/food.proto

echo " Merchant Service"
protoc --go_out=. --go-grpc_out=. proto/merchant.proto

echo " Transaction Service"
protoc --go_out=. --go-grpc_out=. proto/transaction.proto
protoc --go_out=. --go-grpc_out=. proto/cart.proto

echo "All proto files generated successfully!"
