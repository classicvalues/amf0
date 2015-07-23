package amf0

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var objTestData = []byte{0, 3, 97, 112, 112, 2, 0, 5, 109,
	121, 97, 112, 112, 0, 4, 116, 121, 112, 101, 2, 0, 10, 110, 111,
	110, 112, 114, 105, 118, 97, 116, 101, 0, 8, 102, 108, 97, 115,
	104, 86, 101, 114, 2, 0, 31, 70, 77, 76, 69, 47, 51, 46, 48, 32,
	40, 99, 111, 109, 112, 97, 116, 105, 98, 108, 101, 59, 32, 70, 77,
	83, 99, 47, 49, 46, 48, 41, 0, 6, 115, 119, 102, 85, 114, 108, 2,
	0, 22, 114, 116, 109, 112, 58, 47, 47, 108, 111, 99, 97, 108, 104,
	111, 115, 116, 47, 109, 121, 97, 112, 112, 0, 5, 116, 99, 85, 114,
	108, 2, 0, 22, 114, 116, 109, 112, 58, 47, 47, 108, 111, 99, 97,
	108, 104, 111, 115, 116, 47, 109, 121, 97, 112, 112, 0, 0, 9}

func TestObjectDecodes(t *testing.T) {
	o := NewObject()
	n, err := o.DecodeFrom(objTestData, 0)

	assert.Equal(t, len(objTestData), n)
	assert.Nil(t, err)
	assert.Equal(t, 5, o.Size())

	s, _ := o.String("app")
	assert.Equal(t, "myapp", s.GetBody())
	s, _ = o.String("type")
	assert.Equal(t, "nonprivate", s.GetBody())

	_, err = o.Boolean("app")
	assert.Equal(t, WrongTypeError, err)
	_, err = o.Boolean("foo")
	assert.Equal(t, NotFoundError, err)
}

func BenchmarkObjectDecode(b *testing.B) {
	out := NewObject()

	for i := 0; i < b.N; i++ {
		out.DecodeFrom(objTestData, 0)
	}
}

func BenchmarkObjectLookup(b *testing.B) {
	out := NewObject()
	out.DecodeFrom(objTestData, 0)

	for i := 0; i < b.N; i++ {
		out.String("app")
	}
}