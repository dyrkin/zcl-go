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
