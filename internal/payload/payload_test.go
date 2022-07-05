package payload_test

import (
	"fmt"

	"github.com/paulwizviz/learn-mqtt/internal/payload"
)

func Example_serializeCBOR() {
	item := payload.NewData([]byte("Hello"))
	b := payload.MustSerialize(item)
	result := payload.MustDeserialize(b)
	fmt.Println(string(result.Data()))

	// Output:
	// Hello
}
