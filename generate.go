package ceal

//go:generate java -jar antlr-4.12.0-complete.jar -Dlanguage=Go -package parser -visitor C.g4 -o parser

//go:generate go run cmd/spec/main.go -spec langspec.json --out ceal.json

//go:generate go run cmd/sdk/main.go -spec ceal.json --out avm.hpp
//go:generate go run cmd/gen/main.go -spec ceal.json --out compiler/gen.go
//go:generate go run cmd/teal/main.go -spec ceal.json --out teal/gen.go

//go:generate go run cmd/examples/main.go -path examples
