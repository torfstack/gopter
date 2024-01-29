package pkg

import (
	"github.com/stretchr/testify/assert"
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
