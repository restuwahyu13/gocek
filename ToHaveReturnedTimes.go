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
			sliceInterface := []interface{}{}

			if reflect.TypeOf(input).Elem().Kind().String() == reflect.Interface.String() {
				sliceInterface = input.([]interface{})
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Int.String() {
				for _, v := range input.([]int) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Int8.String() {
				for _, v := range input.([]int8) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Int16.String() {
				for _, v := range input.([]int16) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Int32.String() {
				for _, v := range input.([]int32) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Int64.String() {
				for _, v := range input.([]int64) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Float32.String() {
				for _, v := range input.([]float32) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Float64.String() {
				for _, v := range input.([]float64) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.String.String() {
				for _, v := range input.([]string) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Bool.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Map.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Chan.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Slice.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Array.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Func.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Pointer.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Ptr.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Struct.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Complex64.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Complex128.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.TypeOf(input).Elem().Kind().String() == reflect.Invalid.String() {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else if reflect.DeepEqual(reflect.TypeOf(input).Elem().Kind().String(), nil) {
				for _, v := range input.([]bool) {
					sliceInterface = append(sliceInterface, v)
				}
			} else {
				h.t.FailNow()
			}

			for _, v := range sliceInterface {
				if !reflect.DeepEqual(reflect.ValueOf(v).Kind().String(), reflect.Func.String()) {
					ok = false
				} else {
					funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name(), ".")
					sliceNames = append(sliceNames, funcName[len(funcName)-1])
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
		} else if h.o == "==" && cnotEqual != 0 {
			h.t.FailNow()
		} else if h.o == "!=" && cEqual != 0 {
			h.t.FailNow()
		} else if h.o == "==" && int64(cEqual) != value {
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
