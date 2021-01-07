package adapter

import (
	"io"
	"os"
	"testing"
)

func TestIO(t *testing.T) {
	f, err := os.OpenFile("/tmp/pipe", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}

	pr, pw := io.Pipe()
	defer pr.Close()
	defer pw.Close()

	counter := &Counter{pw}

	tee := io.TeeReader(pr, f)
	//
	go func() {
		io.Copy(os.Stdout, tee)
	}()

	_, err = counter.Count(5)
	if err != nil {
		t.Fatal(err)
	}
}
