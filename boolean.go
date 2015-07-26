package amf0

import (
	"io"
)

type Boolean struct {
	isTrue byte
}

var _ AmfType = &Boolean{}

// Creates a new Boolean type, with an optional initial value.
func NewBoolean(bol ...bool) *Boolean {
	b := &Boolean{}
	if len(bol) == 1 {
		b.Set(bol[0])
	}

	return b
}

// Implements AmfType.Decode
func (n *Boolean) Decode(r io.Reader) error {
	bytes, err := readBytes(r, 1)
	if err != nil {
		return err
	}

	n.isTrue = bytes[0]
	return nil
}

// Gets the contained boolean
func (n *Boolean) True() bool {
	return n.isTrue > 0
}

// Sets the contained boolean.
func (n *Boolean) Set(isTrue bool) {
	if isTrue {
		n.isTrue = 1
	} else {
		n.isTrue = 0
	}
}

// Implements AmfType.Encode
func (n *Boolean) Encode(w io.Writer) (int, error) {
	return w.Write(n.EncodeBytes())
}

// Implements AmfType.EncodeBytes
func (n *Boolean) EncodeBytes() []byte {
	return []byte{MARKER_BOOLEAN, n.isTrue}
}

// Implements AmfType.Marker
func (b *Boolean) Marker() byte {
	return MARKER_BOOLEAN
}
