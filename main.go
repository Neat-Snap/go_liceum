package main

import (
	"testing"
	"fmt"
)

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name  string
		input []byte
		out   int
		err   error
	}{
		{
			name:  "1-Valid ASCII string",
			input: []byte("hello"),
			out:   5,
			err:   nil,
		},
		{
			name:  "2-Valid UTF-8 string",
			input: []byte("Привет"), // Русские буквы (UTF-8)
			out:   12,               // Каждая буква занимает 2 байта в UTF-8
			err:   nil,
		},
		{
			name:  "3-Valid mixed ASCII and UTF-8",
			input: []byte("hello Привет"),
			out:   17, // 5 байт ASCII и 12 байт UTF-8
			err:   nil,
		},
		{
			name:  "4-Empty input",
			input: []byte(""),
			out:   0,
			err:   nil,
		},
		{
			name:  "5-Invalid UTF-8",
			input: []byte{0xff, 0xfe, 0xfd}, // Некорректные байты
			out:   0,
			err:   ErrInvalidUTF8,           // Ожидаем ошибку
		},
		{
			name:  "6-Valid multi-byte UTF-8 character",
			input: []byte("😊"),
			out:   4, // Символ занимает 4 байта в UTF-8
			err:   nil,
		},
	}

	for _, cs := range cases {
		cs := cs
		t.Run(cs.name, func (t *testing.T) {
			output, err := GetUTFLength(cs.input)
			if cs.err != err || cs.out != output {
				fmt.Errorf("GetUTFLength(%v) = %v, %v. Want %v", cs.input, output, err, cs.out)
			}
		})
	}
}

