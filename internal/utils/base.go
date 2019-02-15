package utils

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"reflect"
	"strconv"
)

func ExtractFlag(flags []*kingpin.FlagModel, target interface{}) {

	tvt := reflect.TypeOf(target)
	tv := reflect.ValueOf(target)
	if tvt.Kind() == reflect.Ptr && !tv.IsNil() {
		flagMap := make(map[string]*kingpin.FlagModel, 0)
		for _, flagValue := range flags {
			flagMap[flagValue.Name] = flagValue
		}
		tv := reflect.ValueOf(target).Elem()
		tvt := tv.Type()
		fn := tv.NumField()
		for i := 0; i < fn; i++ {
			ft := tv.Field(i)
			if fieldName := tvt.Field(i).Tag.Get("flag"); fieldName != "" {
				setFieldValue(flagMap[fieldName].Value, &ft)
			}

		}
	} else {
		log.Fatal("target must ptr and not nil")
	}

}

func setFieldValue(val kingpin.Value, field *reflect.Value) {
	switch field.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		{
			if v, err := strconv.ParseInt(val.String(), 10, 64); err == nil {
				field.SetInt(v)
			}
		}
	case reflect.Float32, reflect.Float64:
		{
			if v, err := strconv.ParseFloat(val.String(), 64); err == nil {
				field.SetFloat(v)
			}
		}
	case reflect.Bool:
		{
			if v, err := strconv.ParseBool(val.String()); err == nil {
				field.SetBool(v)
			}
		}
	case reflect.String:
		{
			field.SetString(val.String())
		}
	}
}
