package pkg

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOptional_Get(t *testing.T) {
	type testCase[T any] struct {
		name     string
		o        Optional[T]
		wantFunc func(*testing.T, *T)
		wantErr  bool
	}
	tests := []testCase[any]{
		{
			name: "get on empty yields error and nil value",
			o:    Empty[any](),
			wantFunc: func(t *testing.T, a *any) {
				assert.Nil(t, a)
			},
			wantErr: true,
		},
		{
			name: "get on string some yields no error and string value",
			o:    Of[any]("test"),
			wantFunc: func(t *testing.T, a *any) {
				assert.Equal(t, *a, "test")
			},
			wantErr: false,
		},
		{
			name: "get on integer some yields no error and integer value",
			o:    Of[any](42),
			wantFunc: func(t *testing.T, a *any) {
				assert.Equal(t, *a, 42)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantFunc != nil {
				tt.wantFunc(t, got)
			}
		})
	}
}

func Test_JsonMarshall(t *testing.T) {
	type testCase[T any] struct {
		name     string
		o        Optional[T]
		wantFunc func(*testing.T, string)
		wantErr  bool
	}
	tests := []testCase[any]{
		{
			name: "marshall on some opt containing string works",
			o:    Of[any]("test"),
			wantFunc: func(t *testing.T, s string) {
				assert.Equal(t, "\"test\"", s)
			},
		},
		{
			name: "marshall on some opt containing struct works",
			o:    Of[any](struct {
                I int `json:"i"`
                S string `json:"s"`
                B bool `json:"b"`
            }{}),
			wantFunc: func(t *testing.T, s string) {
                assert.Equal(t, "{\"i\":0,\"s\":\"\",\"b\":false}", s)
			},
		},
		{
			name: "marshall on empty opt works",
			o:    Empty[any](),
			wantFunc: func(t *testing.T, s string) {
				assert.Equal(t, "null", s)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(&tt.o)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantFunc != nil {
				tt.wantFunc(t, trim(string(got)))
			}
		})
	}
}

func trim(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
