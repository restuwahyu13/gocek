package helpers

import "encoding/json"

func ToInterfaceMap(v interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	stringify, _ := json.Marshal(&v)
	json.Unmarshal(stringify, &data)

	return data
}
