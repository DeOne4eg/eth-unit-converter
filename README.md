<img align="right" width="300px" src="./images/go.png">

# Ethereum Unit Converter

It is a simple library for conversion units between themselves. The library is written in Golang.

## Usage

```
// first
amount := ethUnit.NewWei(big.NewInt(1000000000000000000))

fmt.Printf("Wei: %v\n", amount.Wei())
fmt.Printf("GWei: %v\n", amount.GWei())
fmt.Printf("Ether: %v\n", amount.Ether())

// second (the same of first)
amount, err := ethUnit.ParseUnit(big.NewFloat(1000000000000000000), "wei")
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Wei: %v\n", amount.Wei())
fmt.Printf("GWei: %v\n", amount.GWei())
fmt.Printf("Ether: %v\n", amount.Ether())
```

## Tests
+ run with `make test`