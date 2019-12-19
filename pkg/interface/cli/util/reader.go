package util

import (
	"bufio"
	"io"
)

type LineIterator struct {
	reader *bufio.Reader
}

func NewLineIterator(rd io.Reader) *LineIterator {
	return &LineIterator{
		reader: bufio.NewReader(rd),
	}
}

func (ln *LineIterator) Next() ([]byte, error) {
	var bytes []byte
	for {
		line, isPrefix, err := ln.reader.ReadLine()
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, line...)
		if !isPrefix {
			break
		}
	}
	return bytes, nil
}
