package jogo

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
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

func (r ResultJson) Int() int64 {

	v := r.rawJson
	t := v.(float64)
	return int64(t)

}

func (r ResultJson) IntStrict() (int64, error) {

	v := r.rawJson
	if v == nil {
		return 0, throwError("InvalidArg", "IntStrict")
	}
	tp, _ := GetType(v)
	if tp == "NUMBER" {
		return int64(v.(float64)), nil
	} else {
		return 0, throwError("InvalidType", "(Method -> IntStrict) Expected NUMBER got "+tp)
	}

}

func (r ResultJson) Float() float64 {

	v := r.rawJson
	return v.(float64)

}

func (r ResultJson) FloatStrict() (float64, error) {

	v := r.rawJson
	if v == nil {
		return 0, throwError("InvalidArg", "FloatStrict")
	}
	tp, _ := GetType(v)
	if tp == "NUMBER" {
		return v.(float64), nil
	} else {
		return 0, throwError("InvalidType", "(Method -> FloatStrict) Expected NUMBER got "+tp)
	}

}

func (r ResultJson) Bool() bool {

	v := r.rawJson
	return v.(bool)

}

func (r ResultJson) BoolStrict() (bool, error) {

	v := r.rawJson
	if v == nil {
		return false, throwError("InvalidArg", "BoolStrict")
	}
	tp, _ := GetType(v)
	if tp == "BOOLEAN" {
		return v.(bool), nil
	} else {
		return false, throwError("InvalidType", "(Method -> BoolStrict) Expected BOOLEAN got "+tp)
	}

}

func (r ResultJson) String() string {

	v := r.rawJson
	return v.(string)
}

func (r ResultJson) StringStrict() (string, error) {

	v := r.rawJson
	if v == nil {
		return "", throwError("InvalidArg", "StringStrict")
	}
	tp, _ := GetType(v)
	if tp == "STRING" {
		return v.(string), nil
	} else {
		return "", throwError("InvalidType", "(Method -> StringStrict), Expected STRING got "+tp)
	}

}

func (r ResultJson) Object() map[string]interface{} {

	v := r.rawJson
	return v.(map[string]interface{})
}

func (r ResultJson) ObjectStrict() (map[string]interface{}, error) {

	v := r.rawJson
	if v == nil {
		return nil, throwError("InvalidArg", "ObjectStrict")
	}
	tp, _ := GetType(v)
	if tp == "OBJECT" {
		return v.(map[string]interface{}), nil
	} else {
		return nil, throwError("InvalidType", "(Method -> ObjectStrict), Expected OBJECT got"+tp)
	}

}

func (r ResultJson) List() []interface{} {

	v := r.rawJson
	return v.([]interface{})

}

func (r ResultJson) ListStrict() ([]interface{}, error) {

	v := r.rawJson
	if v == nil {
		return nil, throwError("InvalidArg", "ListStrict")
	}
	tp, _ := GetType(v)
	if tp == "LIST" {
		return v.([]interface{}), nil
	} else {
		return nil, throwError("InvalidType", "(Method -> ListStrict), Expected LIST got"+tp)
	}

}

func (expJson *ExportedJson) Get(keyRef string) (ResultJson, error) {

	var keyChain []string = strings.Split(keyRef, ".")

	if getType(expJson.rawJson) != "map[string]interface {}" && len(keyChain) > 0 {
		return ResultJson{}, throwError("KeyIndexError", "")
	}

	var ResultJson ResultJson
	var err error

	flag_ := expJson.rawJson.(map[string]interface{})
	toBreak := false
	iterated := 0

	for _, key := range keyChain {

		iterated++
		v, ok := expJson.cachedKeysContent[key]
		if ok {
			flag_ = v
		}
		t := getType(flag_[key])

		switch t {
		case "map[string]interface {}":
			flag_ = flag_[key].(map[string]interface{})
			expJson.cachedKeysContent[key] = flag_
			ResultJson.rawJson = flag_
			ResultJson.Type = "OBJECT"
		case "[]interface {}":
			ResultJson.rawJson = flag_[key].([]interface{})
			ResultJson.Type = "LIST"
			toBreak = true
		case "string":
			ResultJson.rawJson = flag_[key].(string)
			ResultJson.Type = "STRING"
			toBreak = true
		case "float64":
			ResultJson.rawJson = flag_[key].(float64)
			ResultJson.Type = "NUMBER"
			toBreak = true
		case "bool":
			ResultJson.rawJson = flag_[key].(bool)
			ResultJson.Type = "BOOLEAN"
			toBreak = true
		case "":
			toBreak = true
			err = throwError("KeyError", key)
		}

		if toBreak {
			break
		}

	}

	if iterated != len(keyChain) && err == nil {
		err = throwError("KeyIndexError", "")
	}

	return ResultJson, err

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
