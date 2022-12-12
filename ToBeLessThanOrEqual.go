package gocek

import (
	"reflect"
)

func (h *gocekAssertion) ToBeLessThanOrEqual(value interface{}) {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
		value = nil
	})

	if !reflect.DeepEqual(input, nil) && !reflect.DeepEqual(value, nil) {
		inputType := reflect.TypeOf(input).Kind().String()
		valueType := reflect.TypeOf(value).Kind().String()

		notObjectType := inputType == reflect.Slice.String() || inputType == reflect.Array.String() || inputType == reflect.Map.String() || inputType == reflect.Bool.String() || inputType == reflect.Func.String() || inputType == reflect.Chan.String() || inputType == reflect.Pointer.String() || inputType == reflect.Struct.String() || inputType == reflect.Complex64.String() || inputType == reflect.Complex128.String() || inputType == reflect.String.String() || inputType == reflect.Ptr.String() || valueType == reflect.Slice.String() || valueType == reflect.Array.String() || valueType == reflect.Map.String() || valueType == reflect.Bool.String() || valueType == reflect.Func.String() || valueType == reflect.Chan.String() || valueType == reflect.Pointer.String() || valueType == reflect.Struct.String() || valueType == reflect.Complex64.String() || valueType == reflect.Complex128.String() || valueType == reflect.String.String() || valueType == reflect.Ptr.String()

		if h.o == "==" && notObjectType {
			h.t.FailNow()
		}

		if h.o == "==" && inputType == reflect.Int.String() && valueType == reflect.Int.String() && !(input.(int) <= value.(int)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Int8.String() && valueType == reflect.Int8.String() && !(input.(int8) <= value.(int8)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Int16.String() && valueType == reflect.Int16.String() && !(input.(int16) <= value.(int16)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Int32.String() && valueType == reflect.Int32.String() && !(input.(int32) <= value.(int32)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Int64.String() && valueType == reflect.Int64.String() && !(input.(int64) <= value.(int64)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Float32.String() && valueType == reflect.Float32.String() && !(input.(float32) <= value.(float32)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Float64.String() && valueType == reflect.Float64.String() && !(input.(float64) <= value.(float64)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Uint.String() && valueType == reflect.Uint.String() && !(input.(uint) <= value.(uint)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Uint8.String() && valueType == reflect.Uint8.String() && !(input.(uint8) <= value.(uint8)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Uint16.String() && valueType == reflect.Uint16.String() && !(input.(uint16) <= value.(uint16)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Uint32.String() && valueType == reflect.Uint32.String() && !(input.(uint32) <= value.(uint32)) {
			h.t.FailNow()
		} else if h.o == "==" && inputType == reflect.Uint64.String() && valueType == reflect.Uint64.String() && !(input.(uint64) <= value.(uint64)) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Int.String() && valueType == reflect.Int.String() && (input.(int) <= value.(int)) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Int8.String() && valueType == reflect.Int8.String() && input.(int8) <= value.(int8) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Int16.String() && valueType == reflect.Int16.String() && input.(int16) <= value.(int16) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Int32.String() && valueType == reflect.Int32.String() && input.(int32) <= value.(int32) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Int64.String() && valueType == reflect.Int64.String() && input.(int64) <= value.(int64) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Float32.String() && valueType == reflect.Float32.String() && input.(float32) <= value.(float32) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Float64.String() && valueType == reflect.Float64.String() && input.(float64) <= value.(float64) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Uint.String() && valueType == reflect.Uint.String() && input.(uint) <= value.(uint) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Uint8.String() && valueType == reflect.Uint8.String() && input.(uint8) <= value.(uint8) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Uint16.String() && valueType == reflect.Uint16.String() && input.(uint16) <= value.(uint16) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Uint32.String() && valueType == reflect.Uint32.String() && input.(uint32) <= value.(uint32) {
			h.t.FailNow()
		} else if h.o == "!=" && inputType == reflect.Uint64.String() && valueType == reflect.Uint64.String() && input.(uint64) <= value.(uint64) {
			h.t.FailNow()
		}
	} else {
		if h.o == "==" && reflect.DeepEqual(input, nil) || reflect.DeepEqual(value, nil) {
			h.t.FailNow()
		}
	}
}
