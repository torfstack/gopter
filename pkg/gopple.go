package pkg

import "errors"

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
