package gocek

import (
	"fmt"
	"testing"
)

func TestTobeError(action *testing.T) {
	action.Run("Should be ToBeError - value error is to be success", func(t *testing.T) {
		err := fmt.Errorf("error detected")

		test := NewGocek(t)
		test.Expect(err).ToBeError()
	})

	action.Run("Should be ToBeError - value error is to be failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeError()
	})
}
