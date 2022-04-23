package payload_test

import (
	"fmt"

	"github.com/paulwizviz/mqtt-mock/internal/payload"
)

func Example_serializeCBOR() {
	item := payload.NewData([]byte("Hello"))
	b := payload.MustSerialize(item)
	result := payload.MistDeserialize(b)
	fmt.Println(string(result.Data()))

	// Output:
	// Hello
}
