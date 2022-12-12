package gocek

import "testing"

func TestToContain(action *testing.T) {
	action.Run("Should be value contain is exist", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect("hello world").ToContain("world")
	})

	action.Run("Should be value contain is not exist using not", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToContain("dear")
	})
}
