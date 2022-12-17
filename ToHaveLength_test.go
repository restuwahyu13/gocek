package gocek

import "testing"

func TestToHaveLength(action *testing.T) {
	action.Run("Should be ToHaveLength - value length is to be success", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}

		test := NewGocek(t)
		test.Expect(len(data)).ToHaveLength(5)
	})

	action.Run("Should be ToHaveLength - value length is to be failed", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}

		test := NewGocek(t)
		test.Expect(len(data)).Not().ToHaveLength(0)
	})
}
