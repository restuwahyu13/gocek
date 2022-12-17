package gocek

import "testing"

func TestToContain(action *testing.T) {
	action.Run("Should be ToContain - value contain is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect("hello world").ToContain("world")
	})

	action.Run("Should be ToContain - value contain is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToContain("dear")
	})
}
