package scala

import (
	"testing"
)

func TestZipWithIndexWithSlice(t *testing.T) {
	s := []string{
		"1", "2", "3",
	}

	v := ZipWithIndex(s)

	x := v.([]interface{})

	{
		f := x[0].([]interface{})
		if f[0] != "1" || f[1] != 0 {
			t.Error("Expected result should match.")
		}
	}

	{
		f := x[1].([]interface{})
		if f[0] != "2" || f[1] != 1 {
			t.Error("Expected result should match.")
		}
	}

	{
		f := x[2].([]interface{})
		if f[0] != "3" || f[1] != 2 {
			t.Error("Expected result should match.")
		}
	}

}

func TestZipWithIndexWithArray(t *testing.T) {
	s := [3]string{
		"1", "2", "3",
	}

	v := ZipWithIndex(s)

	x := v.([]interface{})

	{
		f := x[0].([]interface{})
		if f[0] != "1" || f[1] != 0 {
			t.Error("Expected result should match.")
		}
	}

	{
		f := x[1].([]interface{})
		if f[0] != "2" || f[1] != 1 {
			t.Error("Expected result should match.")
		}
	}

	{
		f := x[2].([]interface{})
		if f[0] != "3" || f[1] != 2 {
			t.Error("Expected result should match.")
		}
	}
}

func TestZipWithIndexWithMap(t *testing.T) {
	m := map[interface{}]interface{}{
		"1": 1,
		"2": 2,
		3:   "3",
	}

	v := ZipWithIndex(m)

	checkCount := 0

	for key := range v.(map[[2]interface{}]interface{}) {
		if key[0] == "1" {
			if key[1] != 1 {
				t.Error("Expected result should match.")
			}
			checkCount++
		} else if key[0] == "2" {
			if key[1] != 2 {
				t.Error("Expected result should match.")
			}
			checkCount++
		} else if key[0] == 3 {
			if key[1] != "3" {
				t.Error("Expected result should match.")
			}
			checkCount++
		}
	}

	if checkCount != 3 {
		t.Error("Expected result should match.")
	}
}

func TestToMapWithSlice(t *testing.T) {
	s := []string{
		"1", "2", "3",
	}

	v := ToMap(ZipWithIndex(s))

	checkCount := 0

	for key, value := range v.(map[interface{}]interface{}) {
		if key == "1" {
			if value != 0 {
				t.Error("Expected result should match.")
			}
			checkCount++
		} else if key == "2" {
			if value != 1 {
				t.Error("Expected result should match.")
			}
			checkCount++
		} else if key == "3" {
			if value != 2 {
				t.Error("Expected result should match.")
			}
			checkCount++
		}
	}

	if checkCount != 3 {
		t.Error("Expected result should match.")
	}
}

func TestToMapWithMap(t *testing.T) {
	m := map[interface{}]interface{}{
		"1": 1,
		"2": 2,
		3:   "3",
	}

	v := ToMap(ZipWithIndex(m))

	checkCount := 0

	for key := range v.(map[[2]interface{}]interface{}) {
		if key[0] == "1" {
			if key[1] != 1 {
				t.Error("Expected result should match.")
			}
			checkCount++
		} else if key[0] == "2" {
			if key[1] != 2 {
				t.Error("Expected result should match.")
			}
			checkCount++
		} else if key[0] == 3 {
			if key[1] != "3" {
				t.Error("Expected result should match.")
			}
			checkCount++
		}
	}

	if checkCount != 3 {
		t.Error("Expected result should match.")
	}
}

func TestValueAccess(t *testing.T) {
	m := map[interface{}]interface{}{
		"1": 1,
		"2": 2,
		3:   "3",
	}

	v := ToMap(ZipWithIndex(m))

	if _, ok := v.(map[[2]interface{}]interface{})[[2]interface{}{"1", 1}]; !ok {
		t.Error("Expected result should match.")
	}
}

func TestMapWithSlice(t *testing.T) {
	x := []rune{'1', '2'}
	y := Map(x, func(key interface{}, value interface{}) interface{} {
		return "31"
	}).([]interface{})

	if !(y[0] == "31" && y[1] == "31") {
		t.Error("Expected result should match.")
	}
}

func TestMapWithArray(t *testing.T) {
	x := [2]rune{'1', '2'}

	y := Map(x, func(key interface{}, value interface{}) interface{} {
		return "31"
	}).([]interface{})

	if !(y[0] == "31" && y[1] == "31") {
		t.Error("Expected result should match.")
	}
}

func TestMapWithMap(t *testing.T) {
	x := map[string]string{"1": "3", "2": "4"}

	y := Map(x, func(key interface{}, value interface{}) interface{} {
		return 1
	}).(map[string]interface{})

	if !(y["1"] == 1 && y["2"] == 1) {
		t.Error("Expected result should match.")
	}
}

func TestFilterWithSlice(t *testing.T) {
	x := []int{1, 2}
	y := Filter(x, func(value interface{}) bool {
		return value.(int) > 1
	}).([]interface{})

	if !(y[0] == 2 && len(y) == 1) {
		t.Error("Expected result should match.")
	}
}

func TestFilterWithArray(t *testing.T) {
	x := []int{1, 2}
	y := Filter(x, func(value interface{}) bool {
		return value.(int) > 1
	}).([]interface{})

	if !(y[0] == 2 && len(y) == 1) {
		t.Error("Expected result should match.")
	}
}

func TestFilterWithMap(t *testing.T) {
	x := map[int]int{1: 2, 3: 4}
	y := Filter(x, func(value interface{}) bool {
		return value.(int) > 2
	}).(map[int]interface{})

	if !(y[3] == 4 && len(y) == 1) {
		t.Error("Expected result should match.")
	}
}

func TestFoldLeftWithSlice(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	v := FoldLeft(x, 0, func(folded interface{}, key interface{}, value interface{}) interface{} {
		return folded.(int) + value.(int)
	}).(interface{})

	if v != 15 {
		t.Error("Expected result should match.")
	}
}

func TestFoldRightWithSlice(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	v := FoldRight(x, 0, func(key interface{}, value interface{}, folded interface{}) interface{} {
		return value.(int) - folded.(int)
	}).(interface{})

	if v != 3 {
		t.Error("Expected result should match.")
	}
}

func TestFoldLeftWithMap(t *testing.T) {
	x := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	v := FoldLeft(x, 0, func(folded interface{}, key interface{}, value interface{}) interface{} {
		return folded.(int) + value.(int)
	}).(interface{})

	if v != 15 {
		t.Error("Expected result should match.")
	}
}

func TestFoldRightWithMap(t *testing.T) {
	x := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	v := FoldRight(x, 0, func(key interface{}, value interface{}, folded interface{}) interface{} {
		return value.(int) + folded.(int)
	}).(interface{})

	if v != 15 {
		t.Error("Expected result should match.")
	}
}

func TestGroupByWithSlice(t *testing.T) {
	v := []string{"Dog", "Bear", "Fox", "Dolphin", "Cat", "Snake", "Bird"}
	x := GroupBy(v, func(key interface{}, value interface{}) interface{} {
		return []rune(value.(string))[0]
	}).(map[interface{}][]string)

	if !(len(x) > 0) {
		t.Error("Expected result should match.")
	}
}

func TestGroupByWithMap(t *testing.T) {
	v := map[string]string{
		"Dog": "Hey Dog", "Bear": "Hey Bear", "Fox": "Hey Fox", "Dolphin": "Hey Dolphin",
		"Cat": "Hey Cat", "Snake": "Hey Snake", "Bird": "Hey Bird",
	}
	x := GroupBy(v, func(key interface{}, value interface{}) interface{} {
		return []rune(key.(string))[0]
	}).(map[interface{}]map[string]string)

	if !(len(x) > 0) {
		t.Error("Expected result should match.")
	}
}

func TestCount(t *testing.T) {
	s := []string{"1", "2", "3"}

	v := Count(s, func(value interface{}) bool {
		if value.(string) == "2" {
			return true
		}
		return false
	})

	if v != 1 {
		t.Error("Expected result should match.")
	}
}

func TestExists(t *testing.T) {
	s := []string{"1", "2", "3"}

	v := Exists(s, func(value interface{}) bool {
		if value.(string) == "2" {
			return true
		}
		return false
	})

	if v != true {
		t.Error("Expected result should match.")
	}
}

func TestDropWhile(t *testing.T) {
	s := []string{"1", "2", "1", "3"}

	v := DropWhile(s, func(value interface{}) bool {
		return value.(string) == "1"
	}).([]interface{})

	if !(v[0].(string) == "2" && v[1].(string) == "1" && v[2].(string) == "3") {
		t.Error("Expected result should match.")
	}
}
