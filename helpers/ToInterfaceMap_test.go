package helpers

import "testing"

func TestToInterfaceMap(action *testing.T) {
	action.Run("Should be ToInterfaceMap - map int to map interface success", func(t *testing.T) {
		input := make(map[string]int)
		input["january"] = 1
		input["febuary"] = 2
		input["march"] = 3

		output := ToInterfaceMap(input)

		if len(input) != len(output) {
			t.FailNow()
		}
	})

	action.Run("Should be ToInterfaceMap - map int to map interface failed", func(t *testing.T) {
		input := make(map[string]int)
		input["january"] = 1
		input["febuary"] = 2
		input["march"] = 3

		output := ToInterfaceMap(input)

		if len(input) == (len(output) - 1) {
			t.FailNow()
		}
	})

	action.Run("Should be ToInterfaceMap - map string to map interface success", func(t *testing.T) {
		input := make(map[string]string)
		input["month"] = "january"
		input["month"] = "febuary"
		input["month"] = "march"

		output := ToInterfaceMap(input)

		if len(input) != len(output) {
			t.FailNow()
		}
	})

	action.Run("Should be ToInterfaceMap - map string to map interface failed", func(t *testing.T) {
		input := make(map[string]string)
		input["month"] = "january"
		input["month"] = "febuary"
		input["month"] = "march"

		output := ToInterfaceMap(input)

		if len(input) == (len(output) - 1) {
			t.FailNow()
		}
	})
}
