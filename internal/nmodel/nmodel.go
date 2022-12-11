package nmodel

import (
	"math/rand"
	"time"

	"github.com/fxamacker/cbor/v2"
)

type Identifier int64

func CreateID() Identifier {
	rand.Seed(time.Now().Local().UnixNano())
	return Identifier(time.Now().Unix() + rand.Int63n(256))
}

type Data struct {
	ID      Identifier
	Content []byte
}

func (d Data) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(&struct {
		ID   Identifier `cbor:"1"`
		Data []byte     `cbor:"2"`
	}{
		ID:   d.ID,
		Data: d.Content,
	})
}

func (d *Data) UnmarshalCBOR(data []byte) error {
	var aux struct {
		ID      Identifier `cbor:"1"`
		Content []byte     `cbor:"2"`
	}
	if err := cbor.Unmarshal(data, &aux); err != nil {
		return err
	}
	d.ID = aux.ID
	d.Content = aux.Content
	return nil
}
