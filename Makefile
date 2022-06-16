run:
	go run cmd/app/main.go

run-server:
	go run cmd/server/main.go

tidy:
	go mod tidy
	go mod vendor

compile-contract:
	solc --optimize --abi ./contracts/MySmartContract.sol -o build
	solc --optimize --bin ./contracts/MySmartContract.sol -o build
	abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/MySmartContract.go

clean-contract:
	rm -rf build
	rm -rf api/*