package jogo

import (
	"encoding/json"
	"errors"
	"reflect"
)

type ExportedJson struct {
	rawJson           interface{}
	cachedKeysContent map[string](map[string]interface{})
}

type ResultJson struct {
	rawJson interface{}
	Type    string
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
	resJson.Type, err = GetType(expJson.rawJson)
	if err != nil {
		return ExportedJson{}, ResultJson{}, err
	}

	return expJson, resJson, nil
}

func GetType(v interface{}) (string, error) {

	switch getType(v) {
	case "map[string]interface {}":
		return "OBJECT", nil
	case "[]interface {}":
		return "LIST", nil
	case "string":
		return "STRING", nil
	case "float64":
		return "NUMBER", nil
	case "bool":
		return "BOOLEAN", nil
	case "":
		return "", throwError("InvalidArg", "GetType")
	default:
		return "", nil
	}

}

func getType(v interface{}) string {

	if v == nil {
		return ""
	}
	return reflect.ValueOf(v).Type().String()

}

func throwError(code string, data string) error {

	switch code {
	case "KeyIndexError":
		return errors.New("KeyIndexError: JoGO cannot index over non-map object.")
	case "KeyError":
		return errors.New("KeyError: Key '" + data + "' does not exist.")
	case "InvalidArg":
		return errors.New("InvalidArg: Invalid argument passed in '" + data + "' method.")
	case "InvalidType":
		return errors.New(data)
	default:
		return errors.New("Error code unmatched")
	}

}
