package internal

import "reflect"

func mapIndex(mp reflect.Value, index reflect.Value) reflect.Value {
	v := mp.MapIndex(index)
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

func Merge(v1, v2 reflect.Value) reflect.Value {
	if v1.Kind() != reflect.Map || v2.Kind() != reflect.Map || !v1.IsValid() {
		return v2
	}

	for _, key := range v2.MapKeys() {
		e1 := mapIndex(v1, key)
		e2 := mapIndex(v2, key)
		if e1.Kind() == reflect.Map && e2.Kind() == reflect.Map {
			e2 = Merge(e1, e2)
		}
		v1.SetMapIndex(key, e2)
	}
	return v1
}
