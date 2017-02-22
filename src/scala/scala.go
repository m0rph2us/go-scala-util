package scala

import (
	"reflect"
)

func getInterfaceType() reflect.Type {
	return reflect.TypeOf([]interface{}{}).Elem()
}

func getSliceOfInterfaceType() reflect.Type {
	return reflect.TypeOf([]interface{}{})
}

func getMapKeyType(vo reflect.Value) reflect.Type {
	return vo.Type().Key()
}

func ZipWithIndex(value interface{}) interface{} {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	if to.Kind() == reflect.Slice || to.Kind() == reflect.Array || to.Kind() == reflect.Map {
		if to.Kind() == reflect.Map {
			m := reflect.MakeMap(reflect.MapOf(reflect.TypeOf([2]interface{}{}),
				getInterfaceType()))

			for i, k := range vo.MapKeys() {
				s := reflect.New(reflect.ArrayOf(2, getInterfaceType()))

				s.Elem().Index(0).Set(k)
				s.Elem().Index(1).Set(vo.MapIndex(k))

				m.SetMapIndex(s.Elem(), reflect.ValueOf(i))
			}

			return m.Interface()
		} else if to.Kind() == reflect.Array || to.Kind() == reflect.Slice {
			len := vo.Len()

			s := reflect.MakeSlice(getSliceOfInterfaceType(), len, len)

			for i := 0; len > i; i++ {
				s1 := reflect.MakeSlice(getSliceOfInterfaceType(), 2, 2)
				s1.Index(0).Set(vo.Index(i))
				s1.Index(1).Set(reflect.ValueOf(i))

				s.Index(i).Set(s1)
			}

			return s.Interface()
		}
	}

	panic("Unsupported kind.")
}

func ToMap(value interface{}) interface{} {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	if to.Kind() == reflect.Map {
		return value
	} else if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		m := reflect.MakeMap(reflect.MapOf(vo.Index(0).Elem().Index(0).Type(),
			vo.Index(0).Elem().Index(1).Type()))

		len := vo.Len()

		for i := 0; len > i; i++ {
			m.SetMapIndex(
				vo.Index(i).Elem().Index(0),
				vo.Index(i).Elem().Index(1))
		}

		return m.Interface()
	}

	panic("Unsupported kind.")
}

func Map(value interface{}, f func(key interface{}, value interface{}) interface{}) interface{} {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	if to.Kind() == reflect.Map {
		m := reflect.MakeMap(reflect.MapOf(getMapKeyType(vo), getInterfaceType()))

		for _, k := range vo.MapKeys() {
			v := f(k.Interface(), vo.MapIndex(k).Interface())
			m.SetMapIndex(k, reflect.ValueOf(v))
		}

		return m.Interface()
	} else if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()

		s := reflect.MakeSlice(getSliceOfInterfaceType(), len, len)

		for i := 0; len > i; i++ {
			v := f(i, vo.Index(i).Interface())
			s.Index(i).Set(reflect.ValueOf(v))
		}

		return s.Interface()
	}

	panic("Unsupported kind.")
}

func Filter(value interface{}, f func(value interface{}) bool) interface{} {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	if to.Kind() == reflect.Map {
		m := reflect.MakeMap(reflect.MapOf(getMapKeyType(vo), getInterfaceType()))

		for _, k := range vo.MapKeys() {
			v := f(vo.MapIndex(k).Interface())
			if v {
				m.SetMapIndex(k, vo.MapIndex(k))
			}
		}

		return m.Interface()
	} else if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()

		s := reflect.MakeSlice(getSliceOfInterfaceType(), 0, len)

		for i := 0; len > i; i++ {
			v := f(vo.Index(i).Interface())
			if v {
				s = reflect.Append(s, vo.Index(i))
			}
		}

		return s.Interface()
	}

	panic("Unsupported kind.")
}

func FoldLeft(value interface{}, initValue interface{},
f func(folded interface{}, key interface{}, value interface{}) interface{}) interface{} {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	v := initValue

	if to.Kind() == reflect.Map {
		for _, k := range vo.MapKeys() {
			v = f(v, k.Interface(), vo.MapIndex(k).Interface())
		}

		return v
	} else if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()
		for i := 0; len > i; i++ {
			v = f(v, i, vo.Index(i).Interface())
		}

		return v
	}

	panic("Unsupported kind.")
}

func FoldRight(value interface{}, initValue interface{},
f func(key interface{}, value interface{}, folded interface{}) interface{}) interface{} {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	v := initValue

	if to.Kind() == reflect.Map {
		for _, k := range vo.MapKeys() {
			v = f(k, vo.MapIndex(k).Interface(), v)
		}

		return v
	} else if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()
		for i := len - 1; 0 <= i; i-- {
			v = f(i, vo.Index(i).Interface(), v)
		}

		return v
	}

	panic("Unsupported kind.")
}

func GroupBy(value interface{}, f func(key interface{}, value interface{}) interface{}) interface{} {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	if to.Kind() == reflect.Map {
		m := reflect.MakeMap(reflect.MapOf(getInterfaceType(), to))

		for _, k := range vo.MapKeys() {
			v := f(k.Interface(), vo.MapIndex(k).Interface())

			if m.MapIndex(reflect.ValueOf(v)).Kind() == reflect.Invalid {
				s := reflect.MakeMap(reflect.MapOf(to.Key(), to.Elem()))
				s.SetMapIndex(k, vo.MapIndex(k))
				m.SetMapIndex(reflect.ValueOf(v), s)
			} else {
				m.MapIndex(reflect.ValueOf(v)).SetMapIndex(k, vo.MapIndex(k))
			}
		}

		return m.Interface()
	} else if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()

		m := reflect.MakeMap(reflect.MapOf(getInterfaceType(), to))

		for i := 0; len > i; i++ {
			v := f(i, vo.Index(i).Interface())

			if m.MapIndex(reflect.ValueOf(v)).Kind() == reflect.Invalid {
				s := reflect.MakeSlice(to, 1, len)
				s.Index(0).Set(vo.Index(i))
				m.SetMapIndex(reflect.ValueOf(v), s)
			} else {
				m.SetMapIndex(reflect.ValueOf(v),
					reflect.Append(m.MapIndex(reflect.ValueOf(v)), vo.Index(i)))
			}

		}

		return m.Interface()
	}

	panic("Unsupported kind.")
}

func ForAll(value interface{}, f func(value interface{}) bool) bool {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()
		for i := 0; len > i; i++ {
			v := f(vo.Index(i).Interface())
			if v == false {
				return false
			}
		}
		return true
	}

	panic("Unsupported kind.")
}

func Reverse(value interface{}) interface{} {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()
		s := reflect.MakeSlice(getSliceOfInterfaceType(), 0, len)
		for i := len - 1; i >= 0; i-- {
			s = reflect.Append(s, vo.Index(i))
		}
		return s.Interface()
	}

	panic("Unsupported kind.")
}

func Count(value interface{}, f func(value interface{}) bool) int {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	count := 0

	if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()
		for i := 0; len > i; i++ {
			if f(vo.Index(i).Interface()) {
				count++
			}
		}
		return count
	}

	panic("Unsupported kind.")
}

func Exists(value interface{}, f func(value interface{}) bool) bool {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()
		for i := 0; len > i; i++ {
			if f(vo.Index(i).Interface()) {
				return true
			}
		}
		return false
	}

	panic("Unsupported kind.")
}

func DropWhile(value interface{}, f func(value interface{}) bool) interface{} {
	to := reflect.TypeOf(value)
	vo := reflect.ValueOf(value)

	if to.Kind() == reflect.Slice || to.Kind() == reflect.Array {
		len := vo.Len()

		s := reflect.MakeSlice(getSliceOfInterfaceType(), 0, len)

		cond := true
		for i := 0; len > i; i++ {
			if !(cond && f(vo.Index(i).Interface())) {
				cond = false
				s = reflect.Append(s, vo.Index(i))
			}
		}
		return s.Interface()
	}

	panic("Unsupported kind.")
}
