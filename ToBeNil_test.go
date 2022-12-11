package gocek

import (
	"testing"
	"time"
)

func TestToBeNil(action *testing.T) {
	action.Run("Should be ToBeNil - value is to be null", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(nil).ToBeNil()
	})

	action.Run("Should be ToBeNil - value is not to be null", func(t *testing.T) {
		test := NewGocek(t)
		test.Expect(time.Now().Year()).Not().ToBeNil()
	})
}
