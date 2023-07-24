package utils

import (
	"encoding/json"
)

func JsonifyStruct(jsonObj interface{}) map[string]interface{} {
	jsonBytes, _ := json.Marshal(jsonObj)
	resMap := make(map[string]interface{})
	json.Unmarshal(jsonBytes, &resMap)
	return resMap
}

func ArrayResponse(jsonObj interface{}) []interface{} {
	jsonBytes, _ := json.Marshal(jsonObj)
	resMap := make([]interface{}, 0)
	json.Unmarshal(jsonBytes, &resMap)
	return resMap
}
