package gocek

import (
	"reflect"
)

func (h *gocekAssertion) ToHaveLength(value int64) {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
		value = 0
	})

	if !reflect.DeepEqual(input, nil) && !reflect.DeepEqual(value, nil) {
		inputType := reflect.TypeOf(input).Kind().String()

		notObjectType := inputType == reflect.Slice.String() || inputType == reflect.Array.String() || inputType == reflect.Map.String() || inputType == reflect.Bool.String() || inputType == reflect.Func.String() || inputType == reflect.Chan.String() || inputType == reflect.Pointer.String() || inputType == reflect.Struct.String() || inputType == reflect.Complex64.String() || inputType == reflect.Complex128.String() || inputType == reflect.String.String() || inputType == reflect.Ptr.String() || inputType == reflect.Interface.String() || inputType == reflect.Invalid.String() || inputType == reflect.Uintptr.String() || inputType == reflect.Float32.String() || inputType == reflect.Float64.String()

		if h.o == "==" && notObjectType && (input.(int64) < 0 || value < 0) {
			h.t.FailNow()
		}

		if h.o == "==" && int64(input.(int)) != value {
			h.t.FailNow()
		} else if h.o == "!=" && input == value {
			h.t.FailNow()
		}
	} else {
		if h.o == "==" && reflect.DeepEqual(input, nil) || reflect.DeepEqual(value, nil) {
			h.t.FailNow()
		}
	}
}
