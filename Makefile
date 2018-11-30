.PHONY: build clean deploy gomodgen

build: gomodgen
	export GO111MODULE=on
	# networking category
	env GOOS=linux go build -ldflags="-s -w" -o bin/networking/binaryToDecimal category/networking/binaryToDecimal/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/networking/binaryToHexadecimal category/networking/binaryToHexadecimal/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/networking/decimalToBinary category/networking/decimalToBinary/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/networking/decimalToHexadecimal category/networking/decimalToHexadecimal/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/networking/hexadecimalToBinary category/networking/hexadecimalToBinary/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/networking/hexadecimalToDecimal category/networking/hexadecimalToDecimal/main.go
	# numbers category
	env GOOS=linux go build -ldflags="-s -w" -o bin/numbers/highestCommonFactor category/numbers/highestCommonFactor/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/numbers/isPrime category/numbers/isPrime/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/numbers/lowestCommonMultiple category/numbers/lowestCommonMultiple/main.go
	# percentages category
	env GOOS=linux go build -ldflags="-s -w" -o bin/percentages/changeByPercentage category/percentages/changeByPercentage/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/percentages/numberFromPercentage category/percentages/numberFromPercentage/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/percentages/percentageChange category/percentages/percentageChange/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/percentages/percentageFromNumber category/percentages/percentageFromNumber/main.go
	# tsa category
	env GOOS=linux go build -ldflags="-s -w" -o bin/tsa/tsaCone category/tsa/tsaCone/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/tsa/tsaCube category/tsa/tsaCube/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/tsa/tsaCylinder category/tsa/tsaCylinder/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/tsa/tsaPythagoreanTheorem category/tsa/tsaPythagoreanTheorem/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/tsa/tsaRectangularPrism category/tsa/tsaRectangularPrism/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/tsa/tsaSphere category/tsa/tsaSphere/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/tsa/tsaSquareBaseTriangle category/tsa/tsaSquareBaseTriangle/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

gomodgen:
	chmod u+x gomod.sh
	./gomod.sh
