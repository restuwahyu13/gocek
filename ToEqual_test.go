package gocek

import (
	"testing"
	"time"
)

func TestToEqual(action *testing.T) {
	action.Run("Should be ToEqual - value slice is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect([]int{1, 2}).ToEqual([]int{1, 2})
	})

	action.Run("Should be ToEqual - value array is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect([2]int{1, 2}).ToEqual([2]int{1, 2})
	})

	action.Run("Should be ToEqual - value int is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(time.Now().Year()).ToEqual(time.Now().Year())
	})

	action.Run("Should be ToEqual - value string is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(time.November.String()).ToEqual(time.November.String())
	})

	action.Run("Should be ToEqual - value boolean is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(true).ToEqual(true)
	})

	action.Run("Should be ToEqual - value map is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(map[string]interface{}{"name": "john doe"}).ToEqual(map[string]interface{}{"name": "john doe"})
	})

	action.Run("Should be ToEqual - struct is to be success", func(t *testing.T) {
		firstValue := struct{ Name string }{Name: "john doe"}
		secondValue := struct{ Name string }{Name: "john doe"}

		test := NewGocek(t)
		test.Expect(firstValue).ToEqual(secondValue)
	})

	action.Run("Should be ToEqual - value slice map is to be success", func(t *testing.T) {
		firstVal := []map[string]interface{}{
			{"name": "john doe", "age": 20},
			{"name": "jane doe", "age": 20},
		}

		secondVal := []map[string]interface{}{
			{"name": "john doe", "age": 20},
			{"name": "jane doe", "age": 20},
		}

		test := NewGocek(t)
		test.Expect(firstVal).ToEqual(secondVal)
	})

	action.Run("Should be ToEqual - value slice struct is to be success", func(t *testing.T) {
		firstValue := []struct{ Name string }{
			{Name: "john doe"},
		}

		secondValue := []struct{ Name string }{
			{Name: "john doe"},
		}

		test := NewGocek(t)
		test.Expect(firstValue).ToEqual(secondValue)
	})

	action.Run("Should be ToEqual - value 2d slice map is to be success", func(t *testing.T) {
		input := []map[string]interface{}{
			{"name": "paman", "age": 20},
			{"name": "bibi", "age": 11},
		}

		input2 := []map[string]interface{}{
			{"name": "paman", "age": 20},
			{"name": "bibi", "age": 11},
		}

		inputslice2dFirst := []interface{}{input}
		inputslice2dSecond := []interface{}{input2}

		test := NewGocek(t)
		test.Expect(inputslice2dFirst).ToEqual(inputslice2dSecond)
	})

	action.Run("Should be ToEqual - value 2d slice struct is to be success", func(t *testing.T) {
		input := []struct{ Name string }{
			{Name: "john doe"},
		}

		input2 := []struct{ Name string }{
			{Name: "john doe"},
		}

		inputslice2dFirst := []interface{}{input}
		inputslice2dSecond := []interface{}{input2}

		test := NewGocek(t)
		test.Expect(inputslice2dFirst).ToEqual(inputslice2dSecond)
	})

	action.Run("Should be ToEqual - value slice is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect([]int{1, 2}).Not().ToEqual([]int{1, 1})
	})

	action.Run("Should be ToEqual - value array is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect([2]int{1, 2}).Not().ToEqual(1)
	})

	action.Run("Should be ToEqual - value int is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(time.Now().Year()).Not().ToEqual(time.November.String())
	})

	action.Run("Should be ToEqual - value string is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(time.November.String()).Not().ToEqual(time.Now().Year())
	})

	action.Run("Should be ToEqual - value boolean is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(true).Not().ToEqual(time.November.String())
	})

	action.Run("Should be ToEqual - value map is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(map[string]interface{}{"name": "john doe"}).Not().ToEqual(map[string]interface{}{"namex": "john doex"})
		test.Expect(map[string]interface{}{"name": "john doe"}).Not().ToEqual(2022)
	})

	action.Run("Should be ToEqual - value slice map is to be failed", func(t *testing.T) {
		firstVal := []map[string]interface{}{
			{"name": "john doe", "age": 20},
			{"name": "jane doe", "age": 20},
		}

		secondVal := []map[string]interface{}{
			{"namex": "john doe", "agex": 20},
			{"namex": "jane doe", "agex": 20},
		}

		test := NewGocek(t)
		test.Expect(firstVal).Not().ToEqual(secondVal)
		test.Expect(firstVal).Not().ToEqual(time.Now().Year)
	})

	action.Run("Should be ToEqual - value 2d slice map is to be failed", func(t *testing.T) {
		input := []map[string]interface{}{
			{"name": "paman", "age": 20},
			{"name": "bibi", "age": 11},
		}

		input2 := []map[string]interface{}{
			{"namex": "paman", "agex": 201},
			{"namex": "bibi", "agex": 111},
		}

		inputslice2dFirst := []interface{}{input}
		inputslice2dSecond := []interface{}{input2}

		test := NewGocek(t)
		test.Expect(inputslice2dFirst).Not().ToEqual(inputslice2dSecond)
		test.Expect(inputslice2dFirst).Not().ToEqual(time.Now().Year())
	})

	action.Run("Should be ToEqual - struct is to be failed", func(t *testing.T) {
		firstValue := struct{ Name string }{Name: "john doe"}
		secondValue := struct{ Namex string }{Namex: "john doex"}

		test := NewGocek(t)
		test.Expect(firstValue).Not().ToEqual(secondValue)
	})

	action.Run("Should be ToEqual - value slice struct is to be failed", func(t *testing.T) {
		firstValue := []struct{ Name string }{
			{Name: "john doe"},
		}

		secondValue := []struct{ Namex string }{
			{Namex: "john doe"},
		}

		test := NewGocek(t)
		test.Expect(firstValue).Not().ToEqual(secondValue)
	})

	action.Run("Should be ToEqual - value 2d slice struct is to be failed", func(t *testing.T) {
		input := []struct{ Name string }{
			{Name: "john doe"},
		}

		input2 := []struct{ Namex string }{
			{Namex: "john doex"},
		}

		inputslice2dFirst := []interface{}{input}
		inputslice2dSecond := []interface{}{input2}

		test := NewGocek(t)
		test.Expect(inputslice2dFirst).Not().ToEqual(inputslice2dSecond)
	})
}
