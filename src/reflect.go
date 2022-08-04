package nestpay

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func QueryString(v interface{}) (url.Values, error) {
	values := make(url.Values)
	value := reflect.ValueOf(v)
	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return values, nil
		}
		value = value.Elem()
	}
	if v == nil {
		return values, nil
	}
	err := reflector(values, value)
	return values, err
}

func String(v reflect.Value) string {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	return fmt.Sprint(v.Interface())
}

func reflector(values url.Values, val reflect.Value) error {
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		sf := typ.Field(i)
		sv := val.Field(i)
		for sv.Kind() == reflect.Ptr {
			if sv.IsNil() {
				break
			}
			sv = sv.Elem()
		}
		if sv.Kind() == reflect.Struct {
			if err := reflector(values, sv); err != nil {
				return err
			}
			continue
		}
		if n, ok := sf.Tag.Lookup("form"); ok {
			ts := strings.Split(n, ",")
			name := ts[0]
			value := String(sv)
			if len(ts) > 1 && ts[1] == "omitempty" && value != "" {
				values.Add(name, value)
			} else if len(ts) > 1 && ts[1] != "omitempty" {
				values.Add(name, value)
			} else if len(ts) == 1 {
				values.Add(name, value)
			}
		}
	}
	return nil
}
