package codec

import "io"

type Codec interface {
	Encode(r io.Reader, v interface{}) error
	Decode(w io.Writer, v interface{}) error
}
