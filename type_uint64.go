package marshal

import (
	"encoding/binary"
	"io"
)

type Uint64 struct {
	Encoding binary.ByteOrder
	Value    *uint64
}

var _ Type = (*Uint64)(nil)

func (u Uint64) Marshal(w io.Writer) error {
	return binary.Write(w, u.Encoding, *u.Value)
}

func (u Uint64) Unmarshal(r io.Reader) error {
	return binary.Read(r, u.Encoding, u.Value)
}
