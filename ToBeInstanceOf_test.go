package gocek

import (
	"reflect"
	"testing"
)

func TestToBeInstanceOf(action *testing.T) {
	action.Run("Should be ToBeInstanceOf - value string is instance to be match", func(t *testing.T) {
		data := "hello wordl"

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.String)
	})

	action.Run("Should be ToBeInstanceOf - value string is instance not to be match", func(t *testing.T) {
		data := 22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.String)
	})

	action.Run("Should be ToBeInstanceOf - value slice is instance to be match", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Slice)
	})

	action.Run("Should be ToBeInstanceOf - value slice is instance not to be match", func(t *testing.T) {
		data := map[string]interface{}{"name": "john doe"}

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Slice)
	})

	action.Run("Should be ToBeInstanceOf - value int is instance to be match", func(t *testing.T) {
		data := 2022

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Int)
	})

	action.Run("Should be ToBeInstanceOf - value int is instance not to be match", func(t *testing.T) {
		data := "2022"

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Int)
	})

	action.Run("Should be ToBeInstanceOf - value map is instance to be match", func(t *testing.T) {
		data := map[string]interface{}{"name": "john doe"}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Map)
	})

	action.Run("Should be ToBeInstanceOf - value map is instance not to be match", func(t *testing.T) {
		data := 2022

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Map)
	})

	action.Run("Should be ToBeInstanceOf - value float is instance to be match", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Float64)
	})

	action.Run("Should be ToBeInstanceOf - value float is instance not to be match", func(t *testing.T) {
		data := 20

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Float32)
	})

	action.Run("Should be ToBeInstanceOf - value boolean is instance to be match", func(t *testing.T) {
		data := true

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Bool)
	})

	action.Run("Should be ToBeInstanceOf - value boolean is instance not to be match", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Bool)
	})

	action.Run("Should be ToBeInstanceOf - value array is instance to be match", func(t *testing.T) {
		data := [2]int{1, 2}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Array)
	})

	action.Run("Should be ToBeInstanceOf - value array is instance not to be match", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Array)
	})

	action.Run("Should be ToBeInstanceOf - value func is instance to be match", func(t *testing.T) {
		data := func() {}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Func)
	})

	action.Run("Should be ToBeInstanceOf - value func is instance not to be match", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Func)
	})

	action.Run("Should be ToBeInstanceOf - value struct is instance to be match", func(t *testing.T) {
		data := struct {
			Name string
		}{
			Name: "john doe",
		}

		test := NewGocek(t)
		test.Expect(data).ToBeInstanceOf(reflect.Struct)
	})

	action.Run("Should be ToBeInstanceOf - value struct is instance not to be match", func(t *testing.T) {
		data := 20.22

		test := NewGocek(t)
		test.Expect(data).Not().ToBeInstanceOf(reflect.Struct)
	})
}
