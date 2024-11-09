package helpers

import (
	"encoding/json"
	"fmt"
)

func StructToMap(obj interface{}) map[string]interface{} {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(obj)
	json.Unmarshal(inrec, &inInterface)
	return inInterface
}

func MergeMaps(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m2 {
		m1[k] = v
	}

	return m1
}

func DefualtMap() map[string]interface{} {
	return map[string]interface{}{}
}

func MapToStruct(js interface{}, d any) error {
	jsonbody, err := json.Marshal(js)
	if err != nil {
		// do error check
		fmt.Println(err)
		return err
	}

	if err := json.Unmarshal(jsonbody, d); err != nil {
		return err
	}

	return nil
}
