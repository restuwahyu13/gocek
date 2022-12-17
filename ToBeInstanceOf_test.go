package gocek

import (
	"reflect"
	"testing"
)

func TestToBeInstanceOf(action *testing.T) {
	action.Run("Should be ToBeInstanceOf - value string instanceOf is to be suc]", func(t *testing.T) {
		data := "hello wordl"

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.String)
	})

	action.Run("Should be ToBeInstanceOf - value string instanceOf is not to be failed", func(t *testing.T) {
		data := 22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.String)
	})

	action.Run("Should be ToBeInstanceOf - value slice instanceOf is to be success", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Slice)
	})

	action.Run("Should be ToBeInstanceOf - value slice instanceOf is to be failed", func(t *testing.T) {
		data := map[string]interface{}{"name": "john doe"}

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Slice)
	})

	action.Run("Should be ToBeInstanceOf - value int instanceOf to be success", func(t *testing.T) {
		data := 2022

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Int)
	})

	action.Run("Should be ToBeInstanceOf - value int instanceOf is to be failed", func(t *testing.T) {
		data := "2022"

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Int)
	})

	action.Run("Should be ToBeInstanceOf - value map instanceOf to be success", func(t *testing.T) {
		data := map[string]interface{}{"name": "john doe"}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Map)
	})

	action.Run("Should be ToBeInstanceOf - value map instanceOf is to be success", func(t *testing.T) {
		data := 2022

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Map)
	})

	action.Run("Should be ToBeInstanceOf - value float instanceOf to be success", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Float64)
	})

	action.Run("Should be ToBeInstanceOf - value float instanceOf is to be success", func(t *testing.T) {
		data := 20

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Float32)
	})

	action.Run("Should be ToBeInstanceOf - value boolean instanceOf to be success", func(t *testing.T) {
		data := true

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Bool)
	})

	action.Run("Should be ToBeInstanceOf - value boolean instanceOf is to be success", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Bool)
	})

	action.Run("Should be ToBeInstanceOf - value array instanceOf to be success", func(t *testing.T) {
		data := [2]int{1, 2}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Array)
	})

	action.Run("Should be ToBeInstanceOf - value array instanceOf is to be success", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Array)
	})

	action.Run("Should be ToBeInstanceOf - value func instanceOf to be success", func(t *testing.T) {
		data := func() {}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Func)
	})

	action.Run("Should be ToBeInstanceOf - value func instanceOf is to be success", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Func)
	})

	action.Run("Should be ToBeInstanceOf - value struct instanceOf to be success", func(t *testing.T) {
		data := struct {
			Name string
		}{
			Name: "john doe",
		}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Struct)
	})

	action.Run("Should be ToBeInstanceOf - value struct instanceOf is to be success", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Struct)
	})
}
