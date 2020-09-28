package marshal

import (
	"io"
)

type Struct struct {
	Value *Marshallable
	Init  func()
}

func (s Struct) Marshal(w io.Writer) error {
	return Marshal(w, *s.Value)
}

func (s Struct) Unmarshal(r io.Reader) error {
	s.Init()
	if err := Unmarshal(r, *s.Value); err != nil {
		return err
	}
	return nil
}
