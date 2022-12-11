package gocek

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/restuwahyu13/gocek/helpers"
)

func (h *gocekAssertion) ToMatchObject(value interface{}) {
	input := h.v

	defer h.t.Cleanup(func() {
		input = nil
		value = nil
	})

	matchData := true
	notMatchData := false

	typeValueFirst := ""
	typeValueSecond := ""

	if !reflect.DeepEqual(input, nil) && !reflect.DeepEqual(value, nil) {
		typeOfFirst := reflect.TypeOf(input).Kind().String()
		typeOfSecond := reflect.TypeOf(value).Kind().String()

		isSliceOrArray := typeOfFirst == reflect.Array.String() || typeOfFirst == reflect.Slice.String() || typeOfSecond == reflect.Array.String() || typeOfSecond == reflect.Slice.String()
		IsMapsOrStruct := reflect.TypeOf(input).Kind().String() == reflect.Map.String() && reflect.TypeOf(value).Kind().String() == reflect.Map.String() || reflect.TypeOf(input).Kind().String() == reflect.Struct.String() && reflect.TypeOf(value).Kind().String() == reflect.Struct.String()

		if isSliceOrArray {
			convertFirst := helpers.ToInterfaceSlice(input)
			convertSecond := helpers.ToInterfaceSlice(value)

			if len(convertFirst) == len(convertSecond) {
				typeValueFirst = reflect.TypeOf(convertFirst[0]).Kind().String()
				typeValueSecond = reflect.TypeOf(convertSecond[0]).Kind().String()
			}

			notMapOrSlice := typeValueFirst != reflect.Map.String() && typeValueSecond != reflect.Map.String() && typeValueFirst != reflect.Slice.String() && typeValueSecond != reflect.Slice.String()

			if h.o == "==" && notMapOrSlice && len(convertFirst) == len(convertSecond) {
				if notMapOrSlice && reflect.TypeOf(input).Kind().String() == reflect.TypeOf(value).Kind().String() {
					contain := ""

					for i := 0; i < len(convertFirst); i++ {
						compare := (convertFirst[i] != convertSecond[i])
						contain += strconv.FormatBool(compare)
					}

					if strings.Contains(contain, "true") {
						matchData = false
					}
				}
			} else if h.o == "!=" && notMapOrSlice && len(convertFirst) == len(convertSecond) {
				if reflect.TypeOf(input).Kind().String() == reflect.TypeOf(value).Kind().String() {
					contain := ""

					for i := 0; i < len(convertFirst); i++ {
						compare := (convertFirst[i] == convertSecond[i])
						contain += strconv.FormatBool(compare)
					}

					if !strings.Contains(contain, "false") {
						notMatchData = true
					}
				}
			} else if len(convertFirst) == len(convertSecond) && typeValueFirst == reflect.Slice.String() || typeValueSecond == reflect.Slice.String() || typeValueFirst == reflect.Map.String() || typeValueSecond == reflect.Map.String() {
				for i := 0; i < len(convertFirst); i++ {
					if reflect.TypeOf(convertFirst[i]).Kind().String() == reflect.Map.String() && reflect.TypeOf(convertSecond[i]).Kind().String() == reflect.Map.String() || reflect.TypeOf(input).Kind().String() == reflect.Struct.String() && reflect.TypeOf(value).Kind().String() == reflect.Struct.String() {

						convertMapFirst := helpers.ToInterfaceMap(convertFirst[i])
						convertMapSecond := helpers.ToInterfaceMap(convertSecond[i])

						if h.o == "==" {
							for i, v := range convertMapFirst {
								if !reflect.DeepEqual(convertMapFirst[i], nil) && !reflect.DeepEqual(convertMapSecond[i], nil) && v != convertMapSecond[i] {
									matchData = false
								} else if reflect.DeepEqual(convertMapFirst[i], nil) || reflect.DeepEqual(convertMapSecond[i], nil) {
									h.t.FailNow()
								}
							}
						} else if h.o == "!=" {
							for i, v := range convertMapFirst {
								if !reflect.DeepEqual(convertMapFirst[i], nil) && !reflect.DeepEqual(convertMapSecond[i], nil) && v == convertMapSecond[i] {
									notMatchData = true
								} else if reflect.DeepEqual(convertMapFirst[i], nil) && reflect.DeepEqual(convertMapSecond[i], nil) {
									notMatchData = true
								}
							}
						} else {
							h.t.FailNow()
						}
					} else if reflect.TypeOf(convertFirst).Kind().String() == reflect.Slice.String() && reflect.TypeOf(convertSecond).Kind().String() == reflect.Slice.String() {
						cvrtFirst := helpers.ToInterfaceSlice(convertFirst[i])
						cvrtSecond := helpers.ToInterfaceSlice(convertSecond[i])

						if h.o == "==" {
							for x, v := range cvrtFirst {
								convertMapFirst := helpers.ToInterfaceMap(cvrtFirst[x])
								convertMapSecond := helpers.ToInterfaceMap(cvrtSecond[x])

								for j, val := range helpers.ToInterfaceMap(v) {
									if !reflect.DeepEqual(convertMapFirst[j], nil) && !reflect.DeepEqual(convertMapSecond[j], nil) && val != convertMapSecond[j] {
										matchData = false
									} else if reflect.DeepEqual(convertMapFirst[j], nil) || reflect.DeepEqual(convertMapSecond[j], nil) {
										h.t.FailNow()
									}
								}
							}
						} else if h.o == "!=" {
							for x, v := range cvrtFirst {
								convertMapFirst := helpers.ToInterfaceMap(cvrtFirst[x])
								convertMapSecond := helpers.ToInterfaceMap(cvrtSecond[x])

								for j, val := range helpers.ToInterfaceMap(v) {
									if !reflect.DeepEqual(convertMapFirst[j], nil) && !reflect.DeepEqual(convertMapSecond[j], nil) && val == convertMapSecond[j] {
										notMatchData = true
									} else if reflect.DeepEqual(convertMapFirst[j], nil) && reflect.DeepEqual(convertMapSecond[j], nil) {
										h.t.FailNow()
									}
								}
							}
						} else {
							notMatchData = true
						}
					}
				}
			} else {
				h.t.FailNow()
			}

		} else if len(helpers.ToInterfaceMap(input)) == len(helpers.ToInterfaceMap(value)) && IsMapsOrStruct {
			convertFirst := helpers.ToInterfaceMap(h.v)
			convertSecond := helpers.ToInterfaceMap(value)

			if h.o == "==" {
				for i, v := range convertFirst {
					if !reflect.DeepEqual(convertFirst[i], nil) && !reflect.DeepEqual(convertSecond[i], nil) && v != convertSecond[i] {
						matchData = false
					} else if reflect.DeepEqual(convertFirst[i], nil) || reflect.DeepEqual(convertSecond[i], nil) {
						h.t.FailNow()
					}
				}
			} else if h.o == "!=" {
				for i, v := range convertFirst {
					if !reflect.DeepEqual(convertFirst[i], nil) && !reflect.DeepEqual(convertSecond[i], nil) && v == convertSecond[i] {
						notMatchData = true
					} else if reflect.DeepEqual(convertFirst[i], nil) && reflect.DeepEqual(convertSecond[i], nil) {
						h.t.FailNow()
					}
				}
			} else {
				h.t.FailNow()
			}
		}
	} else {
		h.t.FailNow()
	}

	if h.o == "==" && !matchData {
		h.t.FailNow()
	} else if h.o == "!=" && notMatchData {
		h.t.FailNow()
	}
}
