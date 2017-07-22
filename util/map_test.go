package util

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	type test struct {
		I  int      `json:"a"`
		F  float64  `json:"f"`
		As []string `json:"as"`
		Ms []Map    `json:"ms"`
	}
	var m = Map{
		"int8":            int8(5),
		"int":             int(5),
		"int32":           int32(5),
		"int64":           int64(5),
		"float32":         float32(99.99),
		"float64":         float64(99.99),
		"uint8":           uint8(5),
		"uint":            uint(5),
		"uint32":          uint32(5),
		"uint64":          uint64(5),
		"string":          "string",
		"array_int":       []int{1, 2},
		"array_float":     []float64{1.1, 2.2},
		"array_string":    []string{"a", "b"},
		"array_map":       []Map{{"a": 1}, {"b": 1}},
		"array_interface": []interface{}{1, "a", Map{"a": 1}},
		"map":             Map{"a": 1},
		"map2":            map[string]interface{}{"a": 1},
		"map3":            map[string]map[string]int{"a": {"a": 1}},
		"map4": map[string]test{
			"a": {
				I: 1, F: 2.2, As: []string{"a", "b"}, Ms: []Map{{"a": 1}, {"b": 1}},
			}},
	}
	fmt.Println(ToJson(m))

	//================================================
	//test get value
	{
		//test get int
		if m.GetInt("int") != 5 {
			t.Error("int != 5", m.GetInt("int"), m["int"])
		}
		if m.GetInt("int8") != 5 {
			t.Error("int8 != 5", m.GetInt("int8"), m["int8"])
		}
		if m.GetInt("int32") != 5 {
			t.Error("int32 != 5", m.GetInt("int32"), m["int32"])
		}
		if m.GetInt("int64") != 5 {
			t.Error("int64 != 5", m.GetInt("int64"), m["int64"])
		}

		//test get uint
		if m.GetUint("uint") != 5 {
			t.Error("uint != 5", m.GetUint("uint"), m["uint"])
		}
		if m.GetUint("uint8") != 5 {
			t.Error("uint8 != 5", m.GetUint("uint8"), m["uint8"])
		}
		if m.GetUint("uint32") != 5 {
			t.Error("uint32 != 5", m.GetUint("uint32"), m["uint32"])
		}
		if m.GetUint("uint64") != 5 {
			t.Error("uint64 != 5", m.GetUint("uint64"), m["uint64"])
		}

		//test get float
		if m.GetFloat("float32") != float64(float32(99.99)) {
			t.Error("float32 != 5", m.GetFloat("float32"), m["float32"])
		}
		if m.GetFloat("float64") != 99.99 {
			t.Error("float64 != 5", m.GetFloat("float64"), m["float64"])
		}

		//test get string
		if m.GetString("string") != "string" {
			t.Error("string err")
		}

		//test get default value
		if m.GetInt("err") != 0 {
			t.Error("get int err")
		}
		if m.GetIntV("int", 520) != 5 {
			t.Error("get int err")
		}
		if m.GetIntV("err", 520) != 520 {
			t.Error("get int err")
		}
		if m.GetFloat("err") != 0 {
			t.Error("get float err")
		}
		if m.GetUint("err") != 0 {
			t.Error("get uint err")
		}
		if m.GetString("err") != "" {
			t.Error("get string err")
		}
		if m.GetStringV("string", "hahhh") != "string" {
			t.Error("get string err")
		}
		if m.GetStringV("err", "hahhh") != "hahhh" {
			t.Error("get string err")
		}
		if m.GetMap("err") != nil {
			t.Error("get string err")
		}
	}

	//================================================
	//test get array
	{
		//normal
		is := m.GetIntArray("array_int")
		if len(is) != 2 || is[0] != 1 || is[1] != 2 {
			t.Error("get int array err")
		}

		is64 := m.GetInt64Array("array_int")
		if len(is64) != 2 || is64[0] != 1 || is64[1] != 2 {
			t.Error("get int array err")
		}

		is = m.GetIntArray("array_float")
		if len(is) != 2 || is[0] != 1 || is[1] != 2 {
			t.Error("get int array err")
		}

		fs := m.GetFloatArray("array_float")
		if len(fs) != 2 || fs[0] != 1.1 || fs[1] != 2.2 {
			t.Error("get int array err")
		}

		ss := m.GetStringArray("array_string")
		if len(ss) != 2 || ss[0] != "a" || ss[1] != "b" {
			t.Error("get int array err")
		}

		ms := m.GetMapArray("array_map")
		if len(ms) != 2 || ms[0]["a"] != 1 || ms[1]["b"] != 1 {
			t.Error("get map array err")
		}

		ifs := m.GetArray("array_interface")
		if len(ifs) != 3 || ifs[0] != 1 || ifs[1] != "a" || ifs[2].(Map)["a"] != 1 {
			t.Error("get interface array err")
		}
	}

	{
		//error
		ms := m.GetMapArray("array_interface")
		if ms != nil {
			t.Error("get map array err")
		}

		is := m.GetIntArray("array_interface")
		if is != nil {
			t.Error("get int array err")
		}

		is64 := m.GetInt64Array("array_interface")
		if is64 != nil {
			t.Error("get int64 array err")
		}

		fs := m.GetFloatArray("array_interface")
		if fs != nil {
			t.Error("get float64 array err")
		}

		ss := m.GetStringArray("array_interface")
		if len(ss) != 3 {
			t.Error("get string array err")
		}

		//test get default value
		if m.GetIntArray("err") != nil {
			t.Error("get int array err")
		}
		if m.GetFloatArray("err") != nil {
			t.Error("get float array err")
		}
		if m.GetInt64Array("err") != nil {
			t.Error("get int64 array err")
		}
		if m.GetStringArray("err") != nil {
			t.Error("get string array err")
		}
		if m.GetMapArray("err") != nil {
			t.Error("get map array err")
		}
		if m.GetArray("err") != nil {
			t.Error("get interface array err")
		}
	}

	//=================================
	//test get map
	{
		if m.GetMap("map")["a"] != 1 {
			t.Error("get map err")
		}
		if m.GetMap("map2")["a"] != 1 {
			t.Error("get map err")
		}
		if m.GetMap("map3").GetMap("a").GetInt("a") != 1 {
			t.Error("get map err")
		}
		if m.GetMap("map4").GetMap("a").GetInt("a") != 1 {
			t.Error("get map err")
		}
	}

	//=================================
	//test set path and get path
}
