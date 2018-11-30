.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/binaryToDecimal networking/binaryToDecimal/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/binaryToHexadecimal networking/binaryToHexadecimal/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/decimalToBinary networking/decimalToBinary/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/decimalToHexadecimal networking/decimalToHexadecimal/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/hexadecimalToBinary networking/hexadecimalToBinary/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/hexadecimalToDecimal networking/hexadecimalToDecimal/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
