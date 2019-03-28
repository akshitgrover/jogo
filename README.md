<p align="center">
<img src="https://github.com/akshitgrover/jogo/blob/master/logo.png" alt="LOGO">
<br><br>
<a href="http://godoc.org/github.com/akshitgrover/jogo/jogo"><img src="http://godoc.org/github.com/akshitgrover/jogo/jogo?status.svg" alt="LOGO"></a>
<a href="https://goreportcard.com/report/github.com/akshitgrover/jogo"><img src="https://goreportcard.com/badge/github.com/akshitgrover/jogo"></a>
<a href="https://travis-ci.org/akshitgrover/jogo"><img src="https://travis-ci.org/akshitgrover/jogo.svg?branch=master"></a>
<br><br>
<b>JSON o Golang | Forget static types, No more complex structure definitions, Focus on code. Go Reflect!</b>
</p>

JoGO uses memoization to return results faster. JoGO facilitates handling of large and complex JSON structures by making use of go reflections and type assertions.

Take a look at [benchmarks](#Benchmarks)

# Installing
Type the following in **Command Line**

`go get -u github.com/akshitgrover/jogo`

# Usage
Import ***JoGO*** in .go source files as follows

`import "github.com/akshitgrover/jogo/jogo"`

***Note:***

***JoGO*** is distributed as a module with one package, To use ***JoGO*** package packed within the ***JoGO*** module above import path is to be used.

To find more about go modules, Read the [wiki](https://github.com/golang/go/wiki/Modules)

## Export Method

Export method is used to parse underlying json and return [*ExportedJson*](#ExportedJson), [*ResultJson*](#ResultJson) and [*Error*](#Error) objects.

### Usage

```go
package main

import (
    "github.com/akshitgrover/jogo/jogo"
    "fmt"
)

func main() {

    exp, r, err := jogo.Export(`{"hello":"world"}`)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(r.Type) // OBJECT
        _, _ = exp.Get("hello")
    }

}
```

## Get Method

Get method is used to fetch value from an [*ExportedJson*](#ExportedJson) object. It returns [*ResultJson*](#ResultJson) and [*Error*](#Error) objects.

```go
package main

import (
    "github.com/akshitgrover/jogo/jogo"
    "fmt"
)

func main() {

    exp, r, err := jogo.Export(`{"name":{"firstname":"akshit", "lastname":"grover"}}`)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(r.Type) //OBJECT
        r2, _ := exp.Get("name.firstname")
        r3, _ := exp.Get("name.lastname")
        fmt.Println(r2.Type, r3.Type) //STRING STRING
        fmt.Println(r2.String() + r3.String())
    }

}
```

# R method

R method is used to convert any interface to ResultJson struct.
It accepts `interface{}` as an argument and returns `ResultJson{}`.

***Note:*** R method makes it intuitive to iterate over Objects and Slices.

```go
package main

import (
    "github.com/akshitgrover/jogo/jogo"
    "fmt"
)

func main() {

    exp, r, err := jogo.Export(`{"name":{"firstname":"akshit", "lastname":"grover"}}`)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(r.Type) //OBJECT
        r2, _ := exp.Get("name")
        for k, v := range r2 {
            fmt.Println("Key: " + k)
            fmt.Println("Value: " + jogo.R(v).String())
        }
    }
    /* Output

	OBJECT
	Key: firstname
	Value: akshit
	--------
	Key: lastname
	Value: grover
	--------

	*/
}
```

# ExportedJson

Exported Json object holds parsed ***JSON***, If an underlying json represents an OBJECT (javascript alike),

*{ExportedJson Object}.Get("{prop1}.{prop2}.{....}.{propN}")* 

Is used to access any value in that json.

**Any access to the value of a JSON property is to be done using ExportedJson object's Get Method.**

# ResultJson

Go being statically typed, Value fetched from GET method holds an underlying representation of value in the form of an interface, To convert it into a native type, Type assertion is to be used.

***JoGO*** provides various method to do type assertion.

ResultJson object has ***Type*** attribute.

## Types supported by JoGO

* NUMBER
* STRING
* LIST
* OBJECT
* BOOLEAN

***Note:*** Type attribute holds one of the above.

## ResultJson type assertion methods

* {ResultJson Object}.Int()     # Returns: int64
* {ResultJson Object}.Float()   # Returns: float64
* {ResultJson Object}.String()  # Returns: string 
* {ResultJson Object}.Bool()    # Returns: bool
* {ResultJson Object}.Object()  # Returns: map[string]interface {}
* {ResultJson Object}.List()    # Returns: []interface {}

***Note:*** If underlying interface value is not the same as the one it is asserted as, go program panics.

To avoid panic state, Following methods are included.

These methods checks for interface type, If it does not match, an error is returned.

* {ResultJson Object}.IntStrict()     # Returns: int64, error
* {ResultJson Object}.FloatStrict()   # Returns: float64, error
* {ResultJson Object}.StringStrict()  # Returns: string, error 
* {ResultJson Object}.BoolStrict()    # Returns: bool, error
* {ResultJson Object}.ObjectStrict()  # Returns: map[string]interface {}, error
* {ResultJson Object}.ListStrict()    # Returns: []interface {}, error

***Note:*** These methods are slower than non-strict methods, It is advised to first check type using **Type** attribute then call non-strict type assert methods.

## Example

```go
package main

import (
    "github.com/akshitgrover/jogo/jogo"
    "fmt"
)

func main() {

    exp, _, err := jogo.Export(`{"name":{"firstname":"akshit", "lastname":"grover"}}`)
    if err != nil {
        fmt.Println(err)
    } else {
        r2, _ := exp.Get("name.firstname")
        if r2.Type == "STRING" {
            fmt.Println(r2.String()) //akshit
        }
    }

}
```

# Error

Error object is native error interface provided by go.

# Benchmarks

```
BenchmarkJoGOGet-4   	15000000	       315 ns/op	      61 B/op	       2 allocs/op

BenchmarkGJSONGet-8                  3000000        372 ns/op          0 B/op         0 allocs/op
BenchmarkGJSONUnmarshalMap-8          900000       4154 ns/op       1920 B/op        26 allocs/op
BenchmarkJSONUnmarshalMap-8           600000       9019 ns/op       3048 B/op        69 allocs/op
BenchmarkJSONDecoder-8                300000      14120 ns/op       4224 B/op       184 allocs/op
BenchmarkFFJSONLexer-8               1500000       3111 ns/op        896 B/op         8 allocs/op
BenchmarkEasyJSONLexer-8             3000000        887 ns/op        613 B/op         6 allocs/op
BenchmarkJSONParserGet-8             3000000        499 ns/op         21 B/op         0 allocs/op
BenchmarkJSONIterator-8              3000000        812 ns/op        544 B/op         9 allocs/op
```

JoGO benchmark testing was done on 2.3 GHz Intel Core i5 , Whereas other benchmarks were performed on MacBook Pro 15" 2.8 GHz Intel Core i7 and were taken from [here](https://github.com/tidwall/gjson-benchmarks).

# Copyright & License

[MIT License](https://opensource.org/licenses/MIT)

Copyright (c) 2018 Akshit Grover