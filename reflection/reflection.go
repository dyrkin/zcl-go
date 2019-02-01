package reflection

import "reflect"

func Copy(n interface{}) interface{} {
	v := reflect.ValueOf(n)
	switch v.Kind() {
	case reflect.Struct:
		copy := reflect.New(v.Type()).Elem()
		return copy.Interface()
	case reflect.Ptr:
		e := v.Elem()
		copy := reflect.New(e.Type())
		return copy.Interface()
	}
	return nil
}

func ApplyArgs(n interface{}, args ...interface{}) {
	v := reflect.ValueOf(n)
	ApplyArgsToValue(v, args...)
}

func ApplyArgsToValue(v reflect.Value, args ...interface{}) {
	if v.Kind() == reflect.Ptr {
		ApplyArgsToValue(v.Elem(), args...)
		return
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		field.Set(reflect.ValueOf(args[i]))
	}
}
