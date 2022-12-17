package gocek

import (
	"reflect"
	"runtime"
	"strings"
)

func (h *gocekAssertion) ToHaveReturnedTimes(value int64) {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
		value = 0
	})

	ok := true
	cnotEqual := 0
	cEqual := 0
	sliceNames := []string{}

	if !reflect.DeepEqual(input, nil) {
		inputType := reflect.TypeOf(input).Kind().String()

		notObjectType := inputType == reflect.Map.String() || inputType == reflect.Bool.String() || inputType == reflect.Func.String() || inputType == reflect.Chan.String() || inputType == reflect.Pointer.String() || inputType == reflect.Struct.String() || inputType == reflect.Complex64.String() || inputType == reflect.Complex128.String() || inputType == reflect.String.String() || inputType == reflect.Ptr.String() || inputType == reflect.Interface.String() || inputType == reflect.Invalid.String() || inputType == reflect.Uintptr.String() || inputType == reflect.Int.String() || inputType == reflect.Int8.String() || inputType == reflect.Int16.String() || inputType == reflect.Int32.String() || inputType == reflect.Int64.String() || inputType == reflect.Uint.String() || inputType == reflect.Uint8.String() || inputType == reflect.Uint16.String() || inputType == reflect.Uint32.String() || inputType == reflect.Uint64.String() || inputType == reflect.Float32.String() || inputType == reflect.Float64.String()

		if h.o == "==" && notObjectType {
			h.t.FailNow()
		}

		if reflect.DeepEqual(inputType, reflect.Slice.String()) {
			for _, v := range input.([]interface{}) {
				if !reflect.DeepEqual(reflect.ValueOf(v).Kind().String(), reflect.Func.String()) {
					ok = false
				} else {
					sliceNames = append(sliceNames, strings.Split(runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name(), ".")[1])
				}
			}

			if h.o == "==" && !ok {
				h.t.FailNow()
			}

			if ok {
				firstIndex := sliceNames[0]
				for _, v := range sliceNames {
					if v != firstIndex {
						cnotEqual++
					} else {
						cEqual++
					}
				}
			}
		}

		if cEqual == 0 && cnotEqual != len(sliceNames) {
			h.t.FailNow()
		}

		if h.o == "==" && int64(cEqual) != value {
			h.t.FailNow()
		} else if h.o == "!=" && int64(cEqual) == value {
			h.t.FailNow()
		}
	} else {
		if h.o == "==" && reflect.DeepEqual(input, nil) {
			h.t.FailNow()
		}
	}
}
