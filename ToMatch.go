package gocek

import (
	"reflect"
	"regexp"
	"strconv"
)

func (h *gocekAssertion) ToMatch(regex string) {
	input := h.v
	pattern := regex

	defer h.t.Cleanup(func() {
		input = nil
		pattern = ``
	})

	if !reflect.DeepEqual(input, nil) && !reflect.DeepEqual(pattern, nil) {
		inputType := reflect.TypeOf(input).Kind().String()

		notObjectType := inputType == reflect.Slice.String() || inputType == reflect.Array.String() || inputType == reflect.Map.String() || inputType == reflect.Func.String() || inputType == reflect.Chan.String() || inputType == reflect.Pointer.String() || inputType == reflect.Struct.String() || inputType == reflect.Complex64.String() || inputType == reflect.Complex128.String() || inputType == reflect.Ptr.String()

		if h.o == "==" && notObjectType {
			h.t.FailNow()
		}

		regexExp, err := regexp.Compile(pattern)
		if err != nil {
			h.t.FailNow()
		}

		if h.o == "==" && inputType == reflect.String.String() {
			ok := regexExp.MatchString(input.(string))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Bool.String() {
			ok := regexExp.MatchString(strconv.FormatBool(input.(bool)))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Float32.String() {
			ok := regexExp.MatchString(strconv.FormatFloat(float64(input.(float32)), 'E', -1, 32))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Float64.String() {
			ok := regexExp.MatchString(strconv.FormatFloat(float64(input.(float64)), 'E', -1, 64))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Int.String() {
			ok := regexExp.MatchString(strconv.FormatInt(int64(input.(int)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Int8.String() {
			ok := regexExp.MatchString(strconv.FormatInt(int64(input.(int8)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Int16.String() {
			ok := regexExp.MatchString(strconv.FormatInt(int64(input.(int16)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Int32.String() {
			ok := regexExp.MatchString(strconv.FormatInt(int64(input.(int32)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Int64.String() {
			ok := regexExp.MatchString(strconv.FormatInt(int64(input.(int64)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Uint.String() {
			ok := regexExp.MatchString(strconv.FormatUint(uint64(input.(uint)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Uint8.String() {
			ok := regexExp.MatchString(strconv.FormatUint(uint64(input.(uint8)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Uint16.String() {
			ok := regexExp.MatchString(strconv.FormatUint(uint64(input.(uint16)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Uint32.String() {
			ok := regexExp.MatchString(strconv.FormatUint(uint64(input.(uint32)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "==" && inputType == reflect.Uint64.String() {
			ok := regexExp.MatchString(strconv.FormatUint(uint64(input.(uint64)), 10))
			if !ok {
				h.t.FailNow()
			}
		} else if h.o == "!=" && inputType == reflect.String.String() {
			ok := regexExp.MatchString(input.(string))
			if ok {
				h.t.FailNow()
			}
		}
	} else {
		if h.o == "==" && reflect.DeepEqual(input, nil) {
			h.t.FailNow()
		}
	}
}
