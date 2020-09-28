package marshal

import "io"

type Marshallable interface {
	MarshalOrder() []Type
}

type Type interface {
	Marshal(io.Writer) error
	Unmarshal(io.Reader) error
}
