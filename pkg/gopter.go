package pkg

import (
	"encoding/json"
	"errors"
)

type Optional[T any] struct {
	v *T
}

func Of[T any](t T) Optional[T] {
	return Optional[T]{&t}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{nil}
}

func (o *Optional[T]) Get() (*T, error) {
	v := o.v
	if v == nil {
		return nil, errors.New("unwrapped empty optional")
	}
	return v, nil
}

func (o *Optional[T]) IsEmpty() bool {
	return o.v == nil
}

func (o *Optional[T]) MarshalJSON() ([]byte, error) {
	if o.v == nil {
		return []byte("null"), nil
	} else {
		m, ok := interface{}(o.v).(json.Marshaler)
		if ok {
			return m.MarshalJSON()
		}
		return json.Marshal(o.v)
	}
}
