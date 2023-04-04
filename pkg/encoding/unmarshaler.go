package encoding

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"reflect"
)

type Unmarshaler interface {
	UnmarshalMap(map[string]interface{}) error
}

func UnmarshalKey(key string, v interface{}) error {
	return viper.UnmarshalKey(key, &v, viper.DecodeHook(UnmarshalMap()))
}

func UnmarshalMap() mapstructure.DecodeHookFuncValue {
	return func(from reflect.Value, to reflect.Value) (interface{}, error) {
		if to.CanAddr() {
			to = to.Addr()
		}

		// If the destination implements the unmarshaling interface
		u, ok := to.Interface().(Unmarshaler)
		if !ok {
			return from.Interface(), nil
		}

		// If it is nil and a pointer, create and assign the target value first
		if to.IsNil() && to.Type().Kind() == reflect.Ptr {
			to.Set(reflect.New(to.Type().Elem()))
			u = to.Interface().(Unmarshaler)
		}

		var node map[string]interface{}
		switch v := from.Interface().(type) {
		case map[string]interface{}:
			node = v
		default:
			return v, nil
		}

		if err := u.UnmarshalMap(node); err != nil {
			return to.Interface(), err
		}
		return to.Interface(), nil
	}
}
