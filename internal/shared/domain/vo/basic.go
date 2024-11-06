package vo

import "fmt"

type BasicType interface {
	string | int | float64
}

type IBasicValueObject[T BasicType] interface {
	Value() T
	ToString() string
}

type BasicValueObject[T BasicType] struct {
	value T
}

func NewBasicValueObject[T BasicType](s T) BasicValueObject[T] {
	return BasicValueObject[T]{value: s}
}

func (this BasicValueObject[T]) Value() T {
	return this.value
}

func (this BasicValueObject[T]) ToString() string {
	return fmt.Sprintf("%v", this.value)
}
