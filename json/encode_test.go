package json

import (
	"testing"
)

func TestEncode(t *testing.T) {
	encodeMessage()
}

func TestDecode(t *testing.T) {
	decodeMessage()
}

func TestEncodeData(t *testing.T) {
	encodeDatav1()
	encodeDatav2()
}

func TestEncodeV3(t *testing.T) {
	decodeMessagev2()
}

func Test_encodeWithInnerNodes(t *testing.T) {
	encodeWithInnerNodes()

}
