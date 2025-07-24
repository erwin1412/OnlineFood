set -e

echo "Generating proto files for ALL services..."

echo " Auth Service"
cd auth-service
protoc --go_out=. --go-grpc_out=. proto/auth.proto
cd ..

echo " Courier Service"
cd courier-service
protoc --go_out=. --go-grpc_out=. proto/courier.proto
cd ..

echo " Food Service"
cd food-service
protoc --go_out=. --go-grpc_out=. proto/food.proto
cd ..

echo " Merchant Service"
cd merchant-service
protoc --go_out=. --go-grpc_out=. proto/merchant.proto
cd ..

echo " Transaction Service"
cd transaction-service
protoc --go_out=. --go-grpc_out=. proto/transaction.proto
protoc --go_out=. --go-grpc_out=. proto/transaction_detail.proto
protoc --go_out=. --go-grpc_out=. proto/cart.proto
cd ..

echo "All proto files generated successfully!"
