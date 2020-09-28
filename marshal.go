package marshal

import "io"

func Marshal(w io.Writer, o Marshallable) error {
	for _, v := range o.MarshalOrder() {
		if err := v.Marshal(w); err != nil {
			return err
		}
	}
	return nil
}

func Unmarshal(r io.Reader, o Marshallable) error {
	for _, v := range o.MarshalOrder() {
		if err := v.Unmarshal(r); err != nil {
			return err
		}
	}
	return nil
}
