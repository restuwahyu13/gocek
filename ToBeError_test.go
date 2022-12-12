package gocek

import (
	"fmt"
	"testing"
)

func TestTobeError(action *testing.T) {

	action.Run("Should be ToBeError - value is to be error success", func(t *testing.T) {
		err := fmt.Errorf("error detected")
		// var data string = "xx"
		// x := &data

		test := NewGocek(t)
		test.Expect(err).ToBeError()
	})

	action.Run("Should be ToBeError - value is not to be error failed", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).Not().ToBeError()
	})
}
