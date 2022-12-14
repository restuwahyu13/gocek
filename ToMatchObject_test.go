package gocek

import "testing"

func TestToMatchObject(action *testing.T) {
	action.Run("Should be ToMatchObject - value slice int is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect([]int{1, 2, 3, 4, 5}).ToMatchObject([]int{1, 2, 3, 4, 5})
	})

	action.Run("Should be ToMatchObject - value slice string is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect([]string{"one", "two", "three"}).ToMatchObject([]string{"one", "two", "three"})
	})

	action.Run("Should be ToMatchObject - value map is to be success", func(t *testing.T) {
		firstVal := make(map[string]interface{})
		firstVal["name"] = "john doe"
		firstVal["age"] = 20

		secondVal := make(map[string]interface{})
		secondVal["name"] = "john doe"
		secondVal["age"] = 20

		test := NewGocek(t)
		test.Expect(firstVal).ToMatchObject(secondVal)
	})

	action.Run("Should be ToMatchObject - value slice map is to be success", func(t *testing.T) {
		firstVal := []map[string]interface{}{
			{"name": "john doex", "age": 20},
			{"name": "jane doex", "age": 20},
		}

		secondVal := []map[string]interface{}{
			{"name": "john doex", "age": 20},
			{"name": "jane doex", "age": 20},
		}

		test := NewGocek(t)
		test.Expect(firstVal).ToMatchObject(secondVal)
	})

	action.Run("Should be ToMatchObject - value 2d slice map is to be success", func(t *testing.T) {
		input := []map[string]interface{}{
			{"name": "paman", "age": 20},
			{"name": "bibi", "age": 20},
		}

		inputslice2dFirst := []interface{}{input}
		inputslice2dSecond := []interface{}{input}

		test := NewGocek(t)
		test.Expect(inputslice2dFirst).ToMatchObject(inputslice2dSecond)
	})

	action.Run("Should be ToMatchObject - value slice int is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect([]int{1, 2, 3}).Not().ToMatchObject([]string{"one", "two", "three"})
	})

	action.Run("Should be ToMatchObject - value slice string is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect([]string{"one", "two", "three"}).Not().ToMatchObject([]int{1, 2, 3})
	})

	action.Run("Should be ToMatchObject - value map is to be failed", func(t *testing.T) {
		firstVal := make(map[string]interface{})
		firstVal["name"] = "john doe"
		firstVal["age"] = 20

		secondVal := make(map[string]interface{})
		secondVal["name"] = "jane doe"
		secondVal["age"] = 21

		test := NewGocek(t)
		test.Expect(firstVal).Not().ToMatchObject(secondVal)
	})

	action.Run("Should be ToMatchObject - value slice map is to be failed", func(t *testing.T) {
		firstVal := []map[string]interface{}{
			{"name": "john doex", "age": 20},
			{"name": "jane doex", "age": 20},
		}

		secondVal := []map[string]interface{}{
			{"name": "john doe", "age": 21},
			{"name": "jane doe", "age": 21},
		}

		test := NewGocek(t)
		test.Expect(firstVal).Not().ToMatchObject(secondVal)
	})

	action.Run("Should be ToMatchObject - value 2d slice struct is to be failed", func(t *testing.T) {
		input := []struct{ Name string }{
			{Name: "john doe"},
		}

		input2 := []struct{ Namex string }{
			{Namex: "john doex"},
		}

		inputslice2dFirst := []interface{}{input}
		inputslice2dSecond := []interface{}{input2}

		test := NewGocek(t)
		test.Expect(inputslice2dFirst).Not().ToMatchObject(inputslice2dSecond)
	})
}
