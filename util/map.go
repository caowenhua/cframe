package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//type define to map[string]interface{}
type Map map[string]interface{}

func GetInt(v interface{}) int64 {
	val, err := GetIntV(v)
	if err == nil {
		return val
	} else {
		return 0
	}
}

func GetIntV(value interface{}) (int64, error) {
	if value == nil {
		return 0, errors.New("value is null")
	}
	k := reflect.TypeOf(value)
	if k.Name() == "Time" {
		t := value.(time.Time)
		return Timestamp(t), nil
	}
	switch k.Kind() {
	case reflect.Int:
		return int64(value.(int)), nil
	case reflect.Int8:
		return int64(value.(int8)), nil
	case reflect.Int16:
		return int64(value.(int16)), nil
	case reflect.Int32:
		return int64(value.(int32)), nil
	case reflect.Int64:
		return value.(int64), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(GetUint(value)), nil
	case reflect.Float32, reflect.Float64:
		return int64(GetFloat(value)), nil
	case reflect.String:
		if fv, err := strconv.ParseInt(value.(string), 10, 64); err == nil {
			return fv, nil
		} else {
			return 0, err
		}
	case reflect.Struct:
		if k.Name() == "Time" {
			return Timestamp(value.(time.Time)), nil
		} else {
			return 0, fmt.Errorf("incompactable kind(%v)", k.Kind())
		}
	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind())
	}
}

func GetUint(value interface{}) uint64 {
	val, err := GetUintV(value)
	if err == nil {
		return val
	} else {
		return 0
	}
}

func GetUintV(value interface{}) (uint64, error) {
	if value == nil {
		return 0, errors.New("value is null")
	}
	k := reflect.TypeOf(value)
	switch k.Kind() {
	case reflect.Uint:
		return uint64(value.(uint)), nil
	case reflect.Uint8:
		return uint64(value.(uint8)), nil
	case reflect.Uint16:
		return uint64(value.(uint16)), nil
	case reflect.Uint32:
		return uint64(value.(uint32)), nil
	case reflect.Uint64:
		return value.(uint64), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(GetInt(value)), nil
	case reflect.Float32, reflect.Float64:
		return uint64(GetFloat(value)), nil
	case reflect.String:
		if fv, err := strconv.ParseUint(value.(string), 10, 64); err == nil {
			return fv, nil
		} else {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind().String())
	}
}

func GetFloat(value interface{}) float64 {
	val, err := GetFloatV(value)
	if err == nil {
		return val
	} else {
		return 0
	}
}

func GetFloatV(value interface{}) (float64, error) {
	if value == nil {
		return 0, errors.New("arg value is null")
	}
	k := reflect.TypeOf(value)
	switch k.Kind() {
	case reflect.Float32:
		return float64(value.(float32)), nil
	case reflect.Float64:
		return float64(value.(float64)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(GetUint(value)), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(GetInt(value)), nil
	case reflect.String:
		if fv, err := strconv.ParseFloat(value.(string), 64); err == nil {
			return fv, nil
		} else {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("incompactable kind(%v)", k.Kind().String())
	}
}

func GetString(value interface{}) string {
	if value == nil {
		return ""
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return value.(string)
	case reflect.Slice:
		vals := reflect.ValueOf(value)
		var vs = []string{}
		for i := 0; i < vals.Len(); i++ {
			vs = append(vs, fmt.Sprintf("%v", vals.Index(i).Interface()))
		}
		return strings.Join(vs, ",")
	default:
		return fmt.Sprintf("%v", value)
	}
}

func GetMap(value interface{}) Map {
	if map_, ok := value.(Map); ok {
		return map_
	} else if map_, ok := value.(map[string]interface{}); ok {
		return Map(map_)
	} else {
		bys, err := json.Marshal(value)
		if err != nil {
			return nil
		}
		res := Map{}
		err = json.Unmarshal(bys, &res)
		if err != nil {
			return nil
		}
		return res
	}
}

func GetArray(value interface{}) []interface{} {
	if vals, ok := value.([]interface{}); ok {
		return vals
	}
	vals := reflect.ValueOf(value)
	if vals.Kind() != reflect.Slice {
		return nil
	}
	var vs = []interface{}{}
	for i := 0; i < vals.Len(); i++ {
		vs = append(vs, vals.Index(i).Interface())
	}
	return vs
}

func GetMapArray(value interface{}) []Map {
	var vals = GetArray(value)
	if vals == nil {
		return nil
	}
	var ms = []Map{}
	for _, val := range vals {
		var mv = GetMap(val)
		if mv == nil {
			return nil
		} else {
			ms = append(ms, mv)
		}
	}
	return ms
}

func GetStringArray(value interface{}) []string {
	var vals = GetArray(value)
	if vals == nil {
		return nil
	}
	var strs = []string{}
	for _, val := range vals {
		strs = append(strs, GetString(val))
	}
	return strs
}

func GetIntArray(value interface{}) []int {
	as := GetArray(value)
	if as == nil {
		return nil
	}
	is := []int{}
	for _, v := range as {
		iv, err := GetIntV(v)
		if err != nil {
			return nil
		}
		is = append(is, int(iv))
	}
	return is
}

func GetInt64Array(value interface{}) []int64 {
	as := GetArray(value)
	if as == nil {
		return nil
	}
	is := []int64{}
	for _, v := range as {
		iv, err := GetIntV(v)
		if err != nil {
			return nil
		}
		is = append(is, iv)
	}
	return is
}

func GetFloatArray(value interface{}) []float64 {
	as := GetArray(value)
	if as == nil {
		return nil
	}
	is := []float64{}
	for _, v := range as {
		iv, err := GetFloatV(v)
		if err != nil {
			return nil
		}
		is = append(is, iv)
	}
	return is
}

func (m Map) GetUint(key string) uint64 {
	if v, ok := m[key]; ok {
		return GetUint(v)
	} else {
		return 0
	}
}

func (m Map) GetInt(key string) int64 {
	if v, ok := m[key]; ok {
		return GetInt(v)
	} else {
		return 0
	}
}

func (m Map) GetIntV(key string, defaultValue int64) int64 {
	if v, ok := m[key]; ok {
		val, err := GetIntV(v)
		if err == nil {
			return val
		} else {
			return defaultValue
		}
	} else {
		return defaultValue
	}
}

func (m Map) GetFloat(key string) float64 {
	if v, ok := m[key]; ok {
		return GetFloat(v)
	} else {
		return 0
	}
}

func (m Map) GetString(key string) string {
	if v, ok := m[key]; ok {
		return GetString(v)
	} else {
		return ""
	}
}

func (m Map) GetStringV(key, defaultValue string) string {
	if v, ok := m[key]; ok {
		val := GetString(v)
		if len(val) < 0 {
			return defaultValue
		} else {
			return val
		}
	} else {
		return defaultValue
	}
}

func (m Map) GetMap(key string) Map {
	if v, ok := m[key]; ok {
		return GetMap(v)
	} else {
		return nil
	}
}

func (m Map) GetArray(key string) []interface{} {
	if v, ok := m[key]; ok {
		return GetArray(v)
	} else {
		return nil
	}
}

func (m Map) GetMapArray(key string) []Map {
	if v, ok := m[key]; ok {
		return GetMapArray(v)
	} else {
		return nil
	}
}

func (m Map) GetStringArray(key string) []string {
	if v, ok := m[key]; ok {
		return GetStringArray(v)
	} else {
		return nil
	}
}

func (m Map) GetIntArray(key string) []int {
	if v, ok := m[key]; ok {
		return GetIntArray(v)
	} else {
		return nil
	}
}

func (m Map) GetFloatArray(key string) []float64 {
	if v, ok := m[key]; ok {
		return GetFloatArray(v)
	} else {
		return nil
	}
}

func (m Map) GetInt64Array(key string) []int64 {
	if v, ok := m[key]; ok {
		return GetInt64Array(v)
	} else {
		return nil
	}
}

func (m Map) Value(key string) interface{} {
	if v, ok := m[key]; ok {
		return v
	} else {
		return nil
	}
}

func (m Map) SetValue(key string, val interface{}) {
	if val == nil {
		delete(m, key)
	} else {
		m[key] = val
	}
}

func (m Map) GetUintByPath(path string) uint64 {
	v, _ := m.GetValueByPath(path)
	return GetUint(v)
}

func (m Map) GetIntByPath(path string) int64 {
	v, _ := m.GetValueByPath(path)
	return GetInt(v)
}

func (m Map) GetIntByPathV(path string, defaultValue int64) int64 {
	v, _ := m.GetValueByPath(path)
	val, err := GetIntV(v)
	if err == nil {
		return val
	} else {
		return defaultValue
	}
}

func (m Map) GetFloatByPath(path string) float64 {
	v, _ := m.GetValueByPath(path)
	return GetFloat(v)
}

func (m Map) GetStringByPath(path string) string {
	v, _ := m.GetValueByPath(path)
	return GetString(v)
}

func (m Map) GetStringByPathV(path, defaultValue string) string {
	v, _ := m.GetValueByPath(path)
	val := GetString(v)
	if len(val) < 1 {
		return defaultValue
	} else {
		return val
	}
}

func (m Map) GetMapByPath(path string) Map {
	v, _ := m.GetValueByPath(path)
	return GetMap(v)
}

func (m Map) GetArrayByPath(path string) []interface{} {
	v, _ := m.GetValueByPath(path)
	return GetArray(v)
}

func (m Map) GetMapArrayByPath(path string) []Map {
	v, _ := m.GetValueByPath(path)
	return GetMapArray(v)
}

func (m Map) GetStringArrayByPath(path string) []string {
	v, _ := m.GetValueByPath(path)
	return GetStringArray(v)
}

func (m Map) GetIntArrayByPath(path string) []int {
	v, _ := m.GetValueByPath(path)
	return GetIntArray(v)
}

func (m Map) GetValueByPath(path string) (interface{}, error) {
	path = strings.TrimPrefix(path, "/")
	keys := strings.Split(path, "/")
	return m.getValueByPath(keys)
}

func (m Map) getValueByPath(keys []string) (interface{}, error) {
	count := len(keys)
	var tv interface{} = m
	for i := 0; i < count; i++ {
		if tv == nil {
			break
		}
		switch reflect.TypeOf(tv).Kind() {
		case reflect.Slice: //if array
			ary, ok := tv.([]interface{}) //check if valid array
			if !ok {
				return nil, errors.New(fmt.Sprintf(
					"invalid array(%v) in path(/%v),expected []interface{}",
					reflect.TypeOf(tv).String(), strings.Join(keys[:i+1], "/"),
				))
			}
			if keys[i] == "@len" { //check if have @len
				return len(ary), nil //return the array length
			}
			idx, err := strconv.Atoi(keys[i]) //get the target index.
			if err != nil {
				return nil, errors.New(fmt.Sprintf(
					"invalid array index(/%v)", strings.Join(keys[:i+1], "/"),
				))
			}
			if idx >= len(ary) || idx < 0 { //check index valid
				return nil, errors.New(fmt.Sprintf(
					"array out of index in path(/%v)", strings.Join(keys[:i+1], "/"),
				))
			}
			tv = ary[idx]
			continue
		case reflect.Map: //if map
			tm := GetMap(tv) //check map covert
			if tm == nil {
				return nil, errors.New(fmt.Sprintf(
					"invalid map in path(/%v)", strings.Join(keys[:i], "/"),
				))
			}
			tv = tm.Value(keys[i])
			continue
		default: //unknow type
			return nil, errors.New(fmt.Sprintf(
				"invalid type(%v) in path(/%v)",
				reflect.TypeOf(tv).Kind(), strings.Join(keys[:i], "/"),
			))
		}
	}
	if tv == nil { //if valud not found
		return nil, errors.New(fmt.Sprintf(
			"value not found in path(/%v)", strings.Join(keys, "/"),
		))
	} else {
		return tv, nil
	}
}

func (m Map) SetValueByPath(path string, val interface{}) error {
	if len(path) < 1 {
		return errors.New("path is empty")
	}
	path = strings.TrimPrefix(path, "/")
	keys := strings.Split(path, "/")
	//
	i := len(keys) - 1
	pv, err := m.getValueByPath(keys[:i])
	if err != nil {
		return err
	}
	switch reflect.TypeOf(pv).Kind() {
	case reflect.Slice:
		ary, ok := pv.([]interface{}) //check if valid array
		if !ok {
			return errors.New(fmt.Sprintf(
				"invalid array(%v) in path(/%v),expected []interface{}",
				reflect.TypeOf(pv).String(), strings.Join(keys[:i+1], "/"),
			))
		}
		idx, err := strconv.Atoi(keys[i]) //get the target index.
		if err != nil {
			return errors.New(fmt.Sprintf(
				"invalid array index(/%v)", strings.Join(keys[:i+1], "/"),
			))
		}
		if idx >= len(ary) || idx < 0 { //check index valid
			return errors.New(fmt.Sprintf(
				"array out of index in path(/%v)", strings.Join(keys[:i+1], "/"),
			))
		}
		ary[idx] = val
	case reflect.Map:
		tm := GetMap(pv) //check map covert
		if tm == nil {
			return errors.New(fmt.Sprintf(
				"invalid map in path(/%v)", strings.Join(keys[:i], "/"),
			))
		}
		tm.SetValue(keys[i], val)
	default: //unknow type
		return errors.New(fmt.Sprintf(
			"not map type(%v) in path(/%v)",
			reflect.TypeOf(pv).Kind(), strings.Join(keys[:i], "/"),
		))
	}
	return nil
}

func (m Map) Exist(key string) bool {
	_, ok := m[key]
	return ok
}

func NewMapByFile(f string) (Map, error) {
	bys, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var kvs Map = Map{}
	err = json.Unmarshal(bys, &kvs)
	return kvs, err
}

func NewMapsByFile(f string) ([]Map, error) {
	bys, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var kvs []Map = []Map{}
	err = json.Unmarshal(bys, &kvs)
	return kvs, err
}
