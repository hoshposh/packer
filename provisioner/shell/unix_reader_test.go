package shell

import (
	"bytes"
	"io"
	"testing"
)

func TestUnixReader_impl(t *testing.T) {
	var raw interface{}
	raw = new(UnixReader)
	if _, ok := raw.(io.Reader); !ok {
		t.Fatal("should be reader")
	}
}

func TestUnixReader(t *testing.T) {
	input := "one\r\ntwo\n\r\nthree\r\n"
	expected := "one\ntwo\n\nthree\n"

	r := &UnixReader{
		Reader: bytes.NewReader([]byte(input)),
	}

	result := new(bytes.Buffer)
	if _, err := io.Copy(result, r); err != nil {
		t.Fatalf("err: %s", err)
	}

	if result.String() != expected {
		t.Fatalf("bad: %#v", result.String())
	}
}

func TestUnixReader_unixOnly(t *testing.T) {
	input := "\none\n\ntwo\nthree\n\n"
	expected := "\none\n\ntwo\nthree\n\n"

	r := &UnixReader{
		Reader: bytes.NewReader([]byte(input)),
	}

	result := new(bytes.Buffer)
	if _, err := io.Copy(result, r); err != nil {
		t.Fatalf("err: %s", err)
	}

	if result.String() != expected {
		t.Fatalf("bad: %#v", result.String())
	}
}
