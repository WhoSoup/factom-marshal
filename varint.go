package marshal

import (
	"bytes"
	"io"
)

func WriteVarInt(w io.Writer, v uint64) error {
	buf := bytes.NewBuffer(nil)
	if v == 0 {
		buf.WriteByte(0)
	}
	start := false

	if 0x8000000000000000&v != 0 { // Deal with the high bit set; Zero
		buf.WriteByte(0x81) // doesn't need this, only when set.
		start = true        // Going the whole 10 byte path!
	}

	for i := 0; i < 9; i++ {
		b := byte(v >> 56) // Get the top 7 bits
		if b != 0 || start {
			start = true
			if i != 8 {
				b = b | 0x80
			} else {
				b = b & 0x7F
			}
			buf.WriteByte(b)
		}
		v = v << 7
	}

	_, err := w.Write(buf.Bytes())
	return err
}

func ReadVarInt(r io.Reader) (uint64, error) {
	var v uint64
	buf := make([]byte, 1)

	for {
		if _, err := r.Read(buf); err != nil {
			return 0, err
		}
		v = v << 7
		v += uint64(buf[0]) & 0x7F
		if buf[0] < 0x80 {
			break
		}
	}
	return v, nil
}
