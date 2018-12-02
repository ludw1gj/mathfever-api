.PHONY: build clean deploy

build:
	# categories GET
	env GOOS=linux go build -ldflags="-s -w" -o bin/categories src/categories/main.go

	# networking category
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/networking/binaryToDecimal src/calculations/networking/binaryToDecimal/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/networking/binaryToHexadecimal src/calculations/networking/binaryToHexadecimal/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/networking/decimalToBinary src/calculations/networking/decimalToBinary/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/networking/decimalToHexadecimal src/calculations/networking/decimalToHexadecimal/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/networking/hexadecimalToBinary src/calculations/networking/hexadecimalToBinary/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/networking/hexadecimalToDecimal src/calculations/networking/hexadecimalToDecimal/main.go
	# numbers category
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/numbers/highestCommonFactor src/calculations/numbers/highestCommonFactor/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/numbers/isPrime src/calculations/numbers/isPrime/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/numbers/lowestCommonMultiple src/calculations/numbers/lowestCommonMultiple/main.go
	# percentages category
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/percentages/changeByPercentage src/calculations/percentages/changeByPercentage/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/percentages/numberFromPercentage src/calculations/percentages/numberFromPercentage/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/percentages/percentageChange src/calculations/percentages/percentageChange/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/percentages/percentageFromNumber src/calculations/percentages/percentageFromNumber/main.go
	# tsa category
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/tsa/tsaCone src/calculations/tsa/tsaCone/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/tsa/tsaCube src/calculations/tsa/tsaCube/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/tsa/tsaCylinder src/calculations/tsa/tsaCylinder/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/tsa/tsaPythagoreanTheorem src/calculations/tsa/tsaPythagoreanTheorem/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/tsa/tsaRectangularPrism src/calculations/tsa/tsaRectangularPrism/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/tsa/tsaSphere src/calculations/tsa/tsaSphere/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/calculations/tsa/tsaSquareBaseTriangle src/calculations/tsa/tsaSquareBaseTriangle/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy --verbose
	