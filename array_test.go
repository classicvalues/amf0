package amf0

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var arrTestData = []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x03, 0x66,
	0x6f, 0x6f, 0x02, 0x00, 0x03, 0x62, 0x61, 0x72}

func TestArrayDecodes(t *testing.T) {
	o := NewArray()
	err := o.Decode(&reluctantReader{src: arrTestData})

	assert.Nil(t, err)
	assert.Equal(t, 1, o.Size())

	s, _ := o.String("foo")
	assert.Equal(t, "bar", s.GetBody())

	_, err = o.Boolean("app")
	assert.Equal(t, WrongTypeError, err)
	_, err = o.Boolean("foo")
	assert.Equal(t, NotFoundError, err)
}

func TestArrayBuildsAndEncodes(t *testing.T) {
	s := NewArray()
	s.Add("foo", NewString("bar"))

	assert.Equal(t, append([]byte{MARKER_ECMA_ARRAY}, arrTestData...), s.EncodeBytes())
}

func BenchmarkArrayDecode(b *testing.B) {
	out := NewArray()

	for i := 0; i < b.N; i++ {
		out.Decode(bytes.NewReader(arrTestData))
	}
}

func BenchmarkArrayLookup(b *testing.B) {
	out := NewArray()
	out.Decode(bytes.NewReader(arrTestData))

	for i := 0; i < b.N; i++ {
		out.String("foo")
	}
}

func BenchmarkArrayBuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewArray().Add("app", NewString("myapp")).
			Add("type", NewString("nonprivate")).
			Add("flashVer", NewString("FMLE/3.0 (compatible; FMSc/1.0)")).
			Add("swfUrl", NewString("rtmp://localhost/myapp")).
			Add("tcUrl", NewString("rtmp://localhost/myapp")).
			EncodeBytes()
	}
}
