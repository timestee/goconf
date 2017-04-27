package goconf

import (
	"flag"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func coerceBool(v interface{}) (bool, error) {
	switch v.(type) {
	case bool:
		return v.(bool), nil
	case string:
		return strconv.ParseBool(v.(string))
	case int, int16, uint16, int32, uint32, int64, uint64:
		return reflect.ValueOf(v).Int() == 0, nil
	}
	return false, fmt.Errorf("invalid bool value type %T", v)
}

func coerceInt64(v interface{}) (int64, error) {
	switch v.(type) {
	case string:
		return strconv.ParseInt(v.(string), 10, 64)
	case int, int16, int32, int64:
		return reflect.ValueOf(v).Int(), nil
	case uint16, uint32, uint64:
		return int64(reflect.ValueOf(v).Uint()), nil
	}
	return 0, fmt.Errorf("invalid int64 value type %T", v)
}

func coerceFloat64(v interface{}) (float64, error) {
	switch v.(type) {
	case string:
		return strconv.ParseFloat(v.(string), 64)
	case float32, float64:
		return reflect.ValueOf(v).Float(), nil
	}
	return 0, fmt.Errorf("invalid float64 value type %T", v)
}

func coerceDuration(v interface{}, arg string) (time.Duration, error) {
	switch v.(type) {
	case string:
		if regexp.MustCompile(`^[0-9]+$`).MatchString(v.(string)) {
			intVal, err := strconv.Atoi(v.(string))
			if err != nil {
				return 0, err
			}
			mult, err := time.ParseDuration(arg)
			if err != nil {
				return 0, err
			}
			return time.Duration(intVal) * mult, nil
		}
		return time.ParseDuration(v.(string))
	case int, int16, uint16, int32, uint32, int64, uint64:
		return time.Duration(reflect.ValueOf(v).Int()) * time.Millisecond, nil
	case time.Duration:
		return v.(time.Duration), nil
	}
	return 0, fmt.Errorf("invalid time.Duration value type %T", v)
}

func coerceStringSlice(v interface{}) ([]string, error) {
	var tmp []string
	switch v.(type) {
	case string:
		for _, s := range strings.Split(v.(string), ",") {
			tmp = append(tmp, s)
		}
	case []interface{}:
		for _, si := range v.([]interface{}) {
			tmp = append(tmp, si.(string))
		}
	case []string:
		tmp = v.([]string)
	}
	return tmp, nil
}

func coerceFloat64Slice(v interface{}) ([]float64, error) {
	var tmp []float64
	switch v.(type) {
	case string:
		for _, s := range strings.Split(v.(string), ",") {
			f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, f)
		}
	case []interface{}:
		for _, fi := range v.([]interface{}) {
			tmp = append(tmp, fi.(float64))
		}
	case []string:
		for _, s := range v.([]string) {
			f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, f)
		}
	case []float64:
		tmp = v.([]float64)
	}
	return tmp, nil
}

func coerceString(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		return v.(string), nil
	}
	return fmt.Sprintf("%s", v), nil
}

func coerceAutoSet(v interface{}, opt interface{}, fs *flag.FlagSet, name string) error {
	_, err := _coerce(v, opt, "", fs, name)
	return err
}
func coerce(v interface{}, opt interface{}, arg string) (interface{}, error) {
	return _coerce(v, opt, arg, nil, "")
}

func _coerce(v interface{}, opt interface{}, arg string, fs *flag.FlagSet, name string) (interface{}, error) {
	switch opt.(type) {
	case bool:
		b, err := coerceBool(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Bool(name, b, "")
		}
		return b, nil
	case int:
		i, err := coerceInt64(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Int(name, int(i), "")
		}
		return int(i), nil
	case int16:
		i, err := coerceInt64(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Int(name, int(i), "")
		}
		return int16(i), nil
	case uint16:
		i, err := coerceInt64(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Uint(name, uint(i), "")
		}
		return uint16(i), nil
	case int32:
		i, err := coerceInt64(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Int(name, int(i), "")
		}
		return int32(i), nil
	case uint32:
		i, err := coerceInt64(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Uint(name, uint(i), "")
		}
		return uint32(i), nil
	case int64:
		i, err := coerceInt64(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Int64(name, i, "")
		}
		return i, nil
	case uint64:
		i, err := coerceInt64(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Uint64(name, uint64(i), "")
		}
		return uint64(i), nil
	case float64:
		i, err := coerceFloat64(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Float64(name, i, "")
		}
		return float64(i), nil
	case string:
		s, err := coerceString(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.String(name, s, "")
		}
		return s, nil
	case time.Duration:
		d, err := coerceDuration(v, arg)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			fs.Duration(name, d, "")
		}
		return d, nil
	case []string:
		cv, err := coerceStringSlice(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			return nil, fmt.Errorf("type not support []string")
		}
		return cv, nil
	case []float64:
		cv, err := coerceFloat64Slice(v)
		if err != nil {
			return nil, err
		}
		if fs != nil {
			return nil, fmt.Errorf("type not support []float64")
		}
		return cv, nil
	}
	return nil, fmt.Errorf("type not support %T", v)
}
