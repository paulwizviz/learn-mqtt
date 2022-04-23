package misc

import (
	"math/rand"
	"time"
)

type Identifier int64

func CreateID() Identifier {
	rand.Seed(time.Now().Local().UnixNano())
	return Identifier(time.Now().Unix() + rand.Int63n(256))
}
