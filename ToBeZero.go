package gocek

import "reflect"

func (h *gocekAssertion) ToBeZero() {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
	})

	if !reflect.DeepEqual(input, nil) {
		inputType := reflect.TypeOf(input).Kind().String()

		notObjectType := inputType == reflect.Slice.String() || inputType == reflect.Array.String() || inputType == reflect.Map.String() || inputType == reflect.Bool.String() || inputType == reflect.Func.String() || inputType == reflect.Chan.String() || inputType == reflect.Pointer.String() || inputType == reflect.Struct.String() || inputType == reflect.Complex64.String() || inputType == reflect.Complex128.String() || inputType == reflect.String.String() || inputType == reflect.Ptr.String() || inputType == reflect.Interface.String() || inputType == reflect.Invalid.String() || inputType == reflect.Uintptr.String()

		if h.o == "==" && notObjectType {
			h.t.FailNow()
		}

		if inputType == reflect.Float32.String() {
			if h.o == "==" && input.(float32) > -1 && input.(float32) != 0 {
				h.t.FailNow()
			} else if h.o == "!=" && input.(float32) > -1 && input.(float32) == 0 {
				h.t.FailNow()
			}
		} else if inputType == reflect.Float64.String() {
			if h.o == "==" && input.(float64) > -1 && input.(float64) != 0 {
				h.t.FailNow()
			} else if h.o == "!=" && input.(float64) > -1 && input.(float64) == 0 {
				h.t.FailNow()
			}
		} else if inputType == reflect.Int.String() {
			if h.o == "==" && input.(int) > -1 && input.(int) != 0 {
				h.t.FailNow()
			} else if h.o == "!=" && input.(int) > -1 && input.(int) == 0 {
				h.t.FailNow()
			}
		} else if inputType == reflect.Int8.String() {
			if h.o == "==" && input.(int8) > -1 && input.(int8) != 0 {
				h.t.FailNow()
			} else if h.o == "!=" && input.(int8) > -1 && input.(int8) == 0 {
				h.t.FailNow()
			}
		} else if inputType == reflect.Int16.String() {
			if h.o == "==" && input.(int16) > -1 && input.(int16) != 0 {
				h.t.FailNow()
			} else if h.o == "!=" && input.(int16) > -1 && input.(int16) == 0 {
				h.t.FailNow()
			}
		} else if inputType == reflect.Int32.String() {
			if h.o == "==" && input.(int32) > -1 && input.(int32) != 0 {
				h.t.FailNow()
			} else if h.o == "!=" && input.(int32) > -1 && input.(int32) == 0 {
				h.t.FailNow()
			}
		} else if inputType == reflect.Int64.String() {
			if h.o == "==" && input.(int64) > -1 && input.(int64) != 0 {
				h.t.FailNow()
			} else if h.o == "!=" && input.(int64) > -1 && input.(int64) == 0 {
				h.t.FailNow()
			}
		}
	} else {
		if h.o == "==" && reflect.DeepEqual(input, nil) {
			h.t.FailNow()
		}
	}
}
