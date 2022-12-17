package gocek

import "testing"

func TestToBeFalsy(action *testing.T) {
	action.Run("Should be ToBeFalsy - value false is to be success", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(false).ToBeFalsy()
	})

	action.Run("Should be ToBeFalsy - value false is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeFalsy()
	})
}
