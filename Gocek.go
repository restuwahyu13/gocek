package gocek

import (
	"reflect"
	"testing"
)

type gocekAssertion struct {
	v interface{}
	o string
	t *testing.T
}

type gocekInterface interface {
	Expect(value interface{}) *gocekAssertion
	Not() *gocekAssertion
	ToBe(value interface{})
	ToBeNil()
	ToBeTruthy()
	ToBeFalsy()
	ToHaveLength(value int64)
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
	ToMatchObject(value interface{})
	ToMatch(regex string)
	ToBeError()
	ToHaveReturnedTimes(value int64)
}

func NewGocek(t *testing.T) gocekInterface {
	return &gocekAssertion{t: t}
}

func (h *gocekAssertion) Expect(value interface{}) *gocekAssertion {
	h.o = "=="
	h.v = value
	return h
}

func (h *gocekAssertion) Not() *gocekAssertion {
	h.o = "!="
	return h
}
