package payload

import (
	"github.com/fxamacker/cbor/v2"
	"github.com/paulwizviz/learn-mqtt/internal/misc"
)

type Data struct {
	id   misc.Identifier
	data []byte
}

func (i Data) ID() misc.Identifier {
	return i.id
}

func (i Data) Data() []byte {
	return i.data
}

func (i Data) MarshalCBOR() ([]byte, error) {
	return cbor.Marshal(&struct {
		ID   misc.Identifier `cbor:"1"`
		Data []byte          `cbor:"2"`
	}{
		ID:   i.id,
		Data: i.data,
	})
}

func (i *Data) UnmarshalCBOR(data []byte) error {
	var aux struct {
		ID   misc.Identifier `cbor:"1"`
		Data []byte          `cbor:"2"`
	}
	if err := cbor.Unmarshal(data, &aux); err != nil {
		return err
	}
	i.id = aux.ID
	i.data = aux.Data
	return nil
}

func MustSerialize(item *Data) []byte {
	strm, _ := cbor.Marshal(item)
	return strm
}

func MustDeserialize(data []byte) *Data {
	var i Data
	cbor.Unmarshal(data, &i)
	return &i
}

func NewData(data []byte) *Data {
	return &Data{
		id:   misc.CreateID(),
		data: data,
	}
}
