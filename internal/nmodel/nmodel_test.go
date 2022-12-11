package nmodel

import (
	"fmt"

	"github.com/fxamacker/cbor/v2"
)

func Example_cborMarshaller() {
	d := Data{
		ID:      Identifier(123),
		Content: []byte("Hello"),
	}

	cb, err := cbor.Marshal(d)
	fmt.Println(cb, err)

	var result Data
	err = cbor.Unmarshal(cb, &result)
	fmt.Println(result, err)

	// Output:
	// [162 97 49 24 123 97 50 69 72 101 108 108 111] <nil>
	// {123 [72 101 108 108 111]} <nil>
}
