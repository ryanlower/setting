package setting

import (
	"log"
	"os"
	"reflect"
)

// Load ...
func Load(config interface{}) {
	t := reflect.TypeOf(config).Elem()
	v := reflect.ValueOf(config).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.FieldByName(field.Name)

		set(field, value)
	}
}

func set(field reflect.StructField, value reflect.Value) {
	// TODO, ensure value is set-able

	s := getString(field)

	// Set value based on type
	// TODO, handle more than strings!
	switch value.Kind() {
	case reflect.String:
		value.SetString(s)
	default:
		// TODO, error instead of logging
		log.Printf("%v field of type '%v' cannot be set. Defaulting to zero value",
			field.Name, field.Type)
	}
}

func getString(field reflect.StructField) string {
	if envKey := field.Tag.Get("env"); envKey != "" {
		if env := os.Getenv(envKey); env != "" {
			return env
		}
	}

	if d := field.Tag.Get("default"); d != "" {
		return d
	}

	return ""
}
