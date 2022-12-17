package gocek

import (
	"reflect"
)

func (h *gocekAssertion) ToBe(value interface{}) {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
		value = nil
	})

	if !reflect.DeepEqual(input, nil) && !reflect.DeepEqual(value, nil) {
		inputType := reflect.TypeOf(input).Kind().String()
		valueType := reflect.TypeOf(value).Kind().String()

		notObjectType := inputType != reflect.Slice.String() || inputType != reflect.Array.String() || inputType != reflect.Map.String() || inputType != reflect.Func.String() || inputType != reflect.Chan.String() || inputType != reflect.Pointer.String() || inputType != reflect.Struct.String() || valueType != reflect.Slice.String() || valueType != reflect.Array.String() || valueType != reflect.Map.String() || valueType != reflect.Func.String() || valueType != reflect.Chan.String() || valueType != reflect.Pointer.String() || valueType != reflect.Struct.String()

		if h.o == "==" && !notObjectType {
			h.t.FailNow()
		}

		if h.o == "==" && !reflect.DeepEqual(input, nil) && !reflect.DeepEqual(value, nil) && input != value {
			h.t.FailNow()
		} else if h.o == "!=" && !reflect.DeepEqual(input, nil) && !reflect.DeepEqual(value, nil) && input == value {
			h.t.FailNow()
		}
	} else {
		if h.o == "==" && reflect.DeepEqual(input, nil) != reflect.DeepEqual(value, nil) {
			h.t.FailNow()
		} else if h.o == "!=" && reflect.DeepEqual(input, nil) == reflect.DeepEqual(value, nil) {
			h.t.FailNow()
		}
	}
}
