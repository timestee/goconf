package goconf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type FileLoader struct {
	log     func(string)
	data    reflect.Value
	infoLog func(string)
	errLog  func(string)
}

func genTemplate(opts interface{}, fn string) error {
	file, err := os.Create(fn)
	defer file.Close()
	if err != nil {
		return err
	}
	ext := strings.ToLower(filepath.Ext(fn))
	if encoder, ok := encodeFuncMap[ext]; ok {
		ret, err := encoder(opts)
		if err != nil {
			return err
		}
		file.WriteString(ret)
		return nil
	}
	return fmt.Errorf("file format not supported: " + filepath.Ext(fn))
}

func (c *FileLoader) Data() map[string]interface{} {
	if c.data.IsValid() {
		return c.data.Interface().(map[string]interface{})
	}
	return nil
}

func (c *FileLoader) Load(files []string) (err error) {
	c.data, err = c._load(c.data, files, true)
	return
}

func (c *FileLoader) _load(rdata reflect.Value, files []string, asc bool) (reflect.Value, error) {
	var tmp reflect.Value
	for _, file := range files {
		if data, err := c.__load(file); err != nil {
			return rdata, err
		} else {
			c.log(fmt.Sprintf("load: %s", file))
			tmp = merge(tmp, reflect.ValueOf(data))
		}
	}
	if asc {
		rdata = merge(rdata, tmp)
	} else {
		rdata = merge(tmp, rdata)
	}
	return rdata, nil
}

func (c *FileLoader) __load(file string) (interface{}, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	ext := strings.ToLower(filepath.Ext(file))
	var unmarshal decodeFunc
	var ok bool
	if unmarshal, ok = decodeFuncMap[ext]; !ok {
		return nil, fmt.Errorf("file format not supported: %s %s", file, filepath.Ext(file))
	}

	var data interface{}
	var inherit interface{}
	if err = unmarshal(bytes, &data); err == nil {
		dm := data.(map[string]interface{})
		if inherit, ok = dm["inherit_files"]; !ok {
			return data, err
		}
		var ret reflect.Value
		basePath := filepath.Dir(file) + "/"
		switch inherit.(type) {
		case string:
			name := basePath + inherit.(string)
			ret, err = c._load(reflect.ValueOf(data), []string{name}, false)
			return ret.Interface(), err
		case []interface{}:
			var files []string
			for _, fi := range inherit.([]interface{}) {
				files = append(files, basePath+fi.(string))
			}
			ret, err = c._load(reflect.ValueOf(data), files, false)
			return ret.Interface(), err
		}
		return data, err
	}

	return nil, fmt.Errorf("load %s with error %s", file, err.Error())
}

func mapIndex(mp reflect.Value, index reflect.Value) reflect.Value {
	v := mp.MapIndex(index)
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

func merge(v1, v2 reflect.Value) reflect.Value {
	if v1.Kind() != reflect.Map || v2.Kind() != reflect.Map || !v1.IsValid() {
		return v2
	}

	for _, key := range v2.MapKeys() {
		e1 := mapIndex(v1, key)
		e2 := mapIndex(v2, key)
		if e1.Kind() == reflect.Map && e2.Kind() == reflect.Map {
			e2 = merge(e1, e2)
		}
		v1.SetMapIndex(key, e2)
	}
	return v1
}
