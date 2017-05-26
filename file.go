package goconf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

type FileLoader struct {
	data reflect.Value
}

func genTemplate(opts interface{}, fname string) error {
	file, err := os.Create(fname)
	defer file.Close()
	if err != nil {
		return err
	}
	ext := strings.ToLower(filepath.Ext(fname))
	if encoder, ok := EncodeFuncMap[ext]; ok {
		ret, err := encoder(opts)
		if err != nil {
			return err
		}
		file.WriteString(ret)
		return nil
	}
	return fmt.Errorf("File format not supported: " + filepath.Ext(fname))
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
			fmt.Printf("[Config] load: %v\n", file)
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
	var unmarshal DecodeFunc
	var ok bool
	if unmarshal, ok = DecodeFuncMap[ext]; !ok {
		return nil, fmt.Errorf("File format not supported: %s %s", file, filepath.Ext(file))
	}

	var data interface{}
	var inherit interface{}
	if err = unmarshal(bytes, &data); err == nil {
		data_map := data.(map[string]interface{})
		if inherit, ok = data_map["inherit_files"]; !ok {
			return data, err
		}
		var ret reflect.Value
		basepath := filepath.Dir(file) + "/"
		switch inherit.(type) {
		case string:
			name := basepath + inherit.(string)
			ret, err = c._load(reflect.ValueOf(data), []string{name}, false)
			return ret.Interface(), err
		case []interface{}:
			var files []string
			for _, fi := range inherit.([]interface{}) {
				files = append(files, basepath+fi.(string))
			}
			ret, err = c._load(reflect.ValueOf(data), files, false)
			return ret.Interface(), err
		}
		return data, err
	}

	return nil, fmt.Errorf("Load %s with error %s", file, err.Error())
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
