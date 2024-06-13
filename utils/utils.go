package utils

import (
	"crypto/rand"
	"encoding/hex"
	"reflect"
	"strings"
)

func GenerateId () (string, error) {
	bytes := make([]byte, 5)
	_, err := rand.Read(bytes)
	if(err != nil) {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
} 

func UpdateStructFields(target interface{}, updates map[string]interface{}) error {
	v := reflect.ValueOf(target).Elem()
	
	for key, value := range updates {

		field := v.FieldByNameFunc(func(name string) bool { return strings.EqualFold(name, key)})

		if !field.IsValid() || !field.CanSet() {
			continue
		}

		val := reflect.ValueOf(value)
		if val.Type().ConvertibleTo(field.Type()) {
			field.Set(val.Convert(field.Type()))
		}
	}
	return nil
}