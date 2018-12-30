package jogo

import (
	"encoding/json"
)

type ExportedJson struct {
	rawJson           interface{}
	cachedKeysContent map[string](map[string]interface{})
}

type ResultJson struct {
	rawJson interface{}
}

func Export(importedJson string) (ExportedJson, ResultJson, error) {

	var expJson ExportedJson
	expJson.cachedKeysContent = make(map[string](map[string]interface{}))
	err := json.Unmarshal([]byte(importedJson), &expJson.rawJson)
	if err != nil {
		return ExportedJson{}, ResultJson{}, err
	}

	var resJson ResultJson
	resJson.rawJson = expJson.rawJson
	if err != nil {
		return ExportedJson{}, ResultJson{}, err
	}

	return expJson, resJson, nil
}
