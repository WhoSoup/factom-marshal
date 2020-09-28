package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"

	marshal "github.com/WhoSoup/factom-marshal"
)

type Obj struct {
	FieldA []byte
	Other  *Obj2
	FieldB uint64
}

func (o *Obj) MarshalOrder() []marshal.Type {
	var m marshal.Marshallable = o.Other
	return []marshal.Type{
		marshal.BytesVarInt{Value: &o.FieldA},
		marshal.Struct{Value: &m, Init: func() { o.Other = new(Obj2) }},
		marshal.Uint64{Value: &o.FieldB, Encoding: binary.BigEndian},
	}
}

type Obj2 struct {
	One uint64
	Two uint64
}

func (o *Obj2) MarshalOrder() []marshal.Type {
	return []marshal.Type{
		marshal.Uint64{Value: &o.One, Encoding: binary.BigEndian},
		marshal.Uint64{Value: &o.Two, Encoding: binary.BigEndian},
	}
}

func main() {

	o := new(Obj)
	o.FieldA = make([]byte, 64)
	for i := 0; i < len(o.FieldA); i++ {
		o.FieldA[i] = byte(i)
	}
	o.FieldB = 0x666

	o.Other = new(Obj2)
	o.Other.One = 0x1234
	o.Other.Two = 0x5678

	buf := bytes.NewBuffer(nil)

	if err := marshal.Marshal(buf, o); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%x\n", buf.Bytes())

	o2 := new(Obj)
	if err := marshal.Unmarshal(buf, o2); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("FieldA: %x\nFieldB: %x\nOther.One = %x\nOther.Two = %x\n", o2.FieldA, o2.FieldB, o2.Other.One, o2.Other.Two)

}
