package helpers

import "testing"

func TestToInterfaceSlice(action *testing.T) {
	action.Run("Should be TestToInterfaceSlice - slice int to slice interface success", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		output := ToInterfaceSlice(input)

		if len(input) != len(output) {
			t.FailNow()
		}
	})

	action.Run("Should be TestToInterfaceSlice - slice string to slice interface success", func(t *testing.T) {
		input := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
		output := ToInterfaceSlice(input)

		if len(input) != len(output) {
			t.FailNow()
		}
	})

	action.Run("Should be TestToInterfaceSlice - slice int to slice interface failed", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		output := ToInterfaceSlice(input)

		if len(input) == (len(output) - 1) {
			t.FailNow()
		}
	})

	action.Run("Should be TestToInterfaceSlice - slice string to slice interface failed", func(t *testing.T) {
		input := []string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
		output := ToInterfaceSlice(input)

		if len(input) == (len(output) - 1) {
			t.FailNow()
		}
	})
}
