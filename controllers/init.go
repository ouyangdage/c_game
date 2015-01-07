package controllers

import (
	"github.com/fhbzyc/c_game/libs/log"
	"reflect"
)

var FuncMap = make(map[string]reflect.Value)

func init() {

	value := reflect.ValueOf(Controller{})

	numMethod := value.NumMethod()

	for i := 0; i < numMethod; i++ {

		FuncMap[value.Type().Method(i).Name] = value.Method(i)
	}

	log.Logger.Info("Program Start")
}
