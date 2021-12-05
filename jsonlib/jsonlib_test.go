package jsonlib

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

// go get -u github.com/cweill/gotests/...
// gotests -all XXX.go
// go test -v
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func TestEncoding(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Encoding",
			args: args{
				User{Name: "iskra", Age: 10},
			},
			want:    []byte(`{"name":"iskra","age":10}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encoding(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encoding() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestEncodingIndent(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodingIndent(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodingIndent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EncodingIndent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodingStream(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := EncodingStream(w, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodingStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("EncodingStream() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestEncodingIndentStream(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := EncodingIndentStream(w, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("EncodingIndentStream() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("EncodingIndentStream() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestDecoding(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Decoding(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Decoding() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDecodingStream(t *testing.T) {
	type args struct {
		r io.Reader
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DecodingStream(tt.args.r, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("DecodingStream() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
