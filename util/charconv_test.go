package util

import (
	"bytes"
	"io"
	"testing"
)

func TestToUTF8(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    string
		wantErr bool
	}{
		{
			name:  "utf8",
			input: []byte{0xe3, 0x81, 0xa6, 0xe3, 0x81, 0x99, 0xe3, 0x81, 0xa8},
			want:  "てすと",
		},
		{
			name:  "shift-jis",
			input: []byte{0x82, 0xc4, 0x82, 0xb7, 0x82, 0xc6},
			want:  "てすと",
		},
		{
			name: "euc-jp",
			// chardet cannot detect euc-jp when reads short string.
			// https://go.dev/play/p/QUyyvZT789M
			input: []byte{0xb8, 0xe3, 0xc7, 0xda, 0xa4, 0xcf, 0xa4, 0xab, 0xa4, 0xc4, 0xa4, 0xc6, 0xc7, 0xad, 0xa4, 0xc7, 0xa4, 0xa2, 0xa4, 0xc3, 0xa4, 0xbf, 0xa1, 0xa3, 0xa1, 0xa3, 0xa1, 0xa3},
			want:  "吾輩はかつて猫であった。。。",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ToUTF8(bytes.NewReader(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("ToUTF8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			bytesResult, err := io.ReadAll(got)
			if err != nil {
				t.Fatalf("failure to retrieve result:%v", err)
			}

			strResult := string(bytesResult)
			if strResult != tt.want {
				t.Errorf("ToUTF8 want:%q got:%q", tt.want, strResult)
			}
		})
	}
}
