package goconf

import (
	"flag"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func HasArg(fs *flag.FlagSet, s string) bool {
	var found bool
	fs.Visit(func(flag *flag.Flag) {
		if flag.Name == s {
			found = true
		}
	})
	return found
}

func innserResolve(options interface{}, flagSet *flag.FlagSet, cfg map[string]interface{}, tomap map[string]interface{}, autoSet bool) {
	val := reflect.ValueOf(options).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Anonymous {
			var fieldPtr reflect.Value
			switch val.FieldByName(field.Name).Kind() {
			case reflect.Struct:
				fieldPtr = val.FieldByName(field.Name).Addr()
			case reflect.Ptr:
				fieldPtr = reflect.Indirect(val).FieldByName(field.Name)
			}
			if !fieldPtr.IsNil() {
				innserResolve(fieldPtr.Interface(), flagSet, cfg, tomap, autoSet)
			}
			continue
		}

		var v interface{}
		flagName := field.Tag.Get("flag")
		cfgName := field.Tag.Get("cfg")
		dvalue := field.Tag.Get("default")

		if flagName == "" {
			flagName = ToSnakeCase(field.Name)
		}

		if cfgName == "" {
			cfgName = strings.Replace(flagName, "-", "_", -1)
		}

		if autoSet {
			if flagSet.Lookup(flagName) == nil {
				if dvalue != "" {
					v = dvalue
				} else {
					v = val.Field(i).Interface()
				}
				if err := coerceAutoSet(v, val.FieldByName(field.Name).Interface(), flagSet, flagName); err != nil {
					fmt.Printf("[Config] auto flag fail, name: %s val: %v err: %s\n", flagName, v, err.Error())
				} else {
					fmt.Printf("[Config] auto flag succ, name: %s val: %v\n", flagName, v)
				}
			}
		} else {
			// resolve the flags according to priority
			if flagSet != nil && HasArg(flagSet, flagName) { // command line flag value
				flagInst := flagSet.Lookup(flagName)
				v = flagInst.Value.String()
			} else if cfgVal, ok := cfg[cfgName]; ok { // config file value
				v = cfgVal
			} else if dvalue != "" { // default value
				v = dvalue
			} else {
				v = val.Field(i).Interface()
			}
			fieldVal := val.FieldByName(field.Name)
			coerced, err := coerce(v, fieldVal.Interface(), field.Tag.Get("arg"))
			if err != nil {
				fmt.Printf("[Config] coerce fail: %v for %s (%+v) - %s\n", v, field.Name, fieldVal, err)
			}
			fieldVal.Set(reflect.ValueOf(coerced))
			if tomap != nil {
				if err == nil {
					tomap[flagName] = coerced
				} else {
					tomap[flagName] = v
				}
			}
		}
	}
}
