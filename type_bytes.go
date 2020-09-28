package marshal

import "io"

type BytesVarInt struct {
	Value *[]byte
}

func (b BytesVarInt) Marshal(w io.Writer) error {
	if err := WriteVarInt(w, uint64(len(*b.Value))); err != nil {
		return err
	}
	_, err := w.Write(*b.Value)
	return err
}

func (b BytesVarInt) Unmarshal(r io.Reader) error {
	l, err := ReadVarInt(r)
	if err != nil {
		return err
	}

	*b.Value = make([]byte, l)
	_, err = r.Read(*b.Value)
	return err
}
