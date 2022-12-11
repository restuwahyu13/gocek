package helpers

import (
	"encoding/json"
)

func ToInterfaceSlice(v interface{}) []interface{} {
	data := []interface{}{}
	stringify, _ := json.Marshal(&v)
	json.Unmarshal(stringify, &data)

	return data

}
