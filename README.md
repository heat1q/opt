# Option

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/heat1q/opt)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/heat1q/opt)
[![Go Report Card](https://goreportcard.com/badge/github.com/heat1q/opt)](https://goreportcard.com/report/github.com/heat1q/opt)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)

This package implements a Rust styled `Option` type for optional values in Go. 
`Option` provides a wrapper around a value that may or may not be initialized 
and a set of methods to extract the inner value or handle nil cases.


## Usage
### Installation
```
go get -u github.com/heat1q/opt
```
### Example
```go
package main

import (
	"fmt"

	"github.com/heat1q/opt"
)

func main() {
	o := opt.New("potato")
	value, ok := o.Some()
	fmt.Println(ok) // true
	fmt.Println(value) // potato
}
```

## Marshalling JSON
`Option` solves the nullable issue for values where the 
default of a value is also considered valid. For instance, consider 
the scenario where the `false` value of a `bool` is a valid instance of the nullable field.


```go
package main

import (
	"fmt"

	"github.com/heat1q/opt"
)

type data struct {
	Flag opt.Option[bool] `json:"flag,omitempty"`
}

func main() {
	var d data

	_ = json.Unmarshal(`{}`, &d)
	_, ok := d.Value.Some() 
	fmt.Println(ok) // false
	
	_ = json.Unmarshal(`{"flag": false}`, &d)
	_, ok = d.Value.Some() 
	fmt.Println(ok) // true
}
```

