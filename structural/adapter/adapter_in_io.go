package adapter

import (
	"io"
	"strconv"
)

type Counter struct {
	Writer io.Writer
}

func (f *Counter) Count(n uint64) (uint64, error) {
	if n == 0 {
		_, err := f.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		if err != nil {
			return 0, err
		}
		return 0, nil
	}

	cur := n
	_, err := f.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	if err != nil {
		return 0, err
	}
	return f.Count(n - 1)
}
