package gocek

import (
	"reflect"
	"testing"
)

type assertion struct {
	v interface{}
	o string
	t *testing.T
}

type iAssertion interface {
	Expect(value interface{}) *assertion
	Not() *assertion
	ToBe(value interface{})
	ToBeNil()
	ToBeTruthy()
	ToBeFalsy()
	ToHaveLength(value int)
	ToBeInstanceOf(value reflect.Kind)
	ToBeMinus()
	ToBeZero()
	ToContain(value string)
	ToHaveReturned()
	ToBeGreaterThan(value interface{})
	ToBeGreaterThanOrEqual(value interface{})
	ToBeLessThan(value interface{})
	ToBeLessThanOrEqual(value interface{})
	ToEqual(value interface{})
}

func NewAssertion(t *testing.T) iAssertion {
	return &assertion{t: t}
}

func (h *assertion) Expect(value interface{}) *assertion {
	h.o = "=="
	h.v = value
	return h
}

func (h *assertion) Not() *assertion {
	h.o = "!="
	return h
}
