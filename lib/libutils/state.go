package libutils

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

type SliceItem struct {
	Value interface{} `json:"value"`
	Desc  interface{} `json:"desc"`
}

func InitStateX(obj interface{}) error {
	state := reflect.TypeOf(obj).Elem()
	numFiled := state.NumField()

	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < numFiled; i++ {
		field := state.Field(i)

		statusStr := field.Tag.Get("v")

		switch field.Type.String() {
		case "int":
			status := cast.ToInt(statusStr)
			v.Field(i).Set(reflect.ValueOf(status))
		case "int8":
			status := cast.ToInt8(statusStr)
			v.Field(i).Set(reflect.ValueOf(status))
		case "int16":
			status := cast.ToInt16(statusStr)
			v.Field(i).Set(reflect.ValueOf(status))
		case "int32":
			status := cast.ToInt32(statusStr)
			v.Field(i).Set(reflect.ValueOf(status))
		case "int64":
			status := cast.ToInt64(statusStr)
			v.Field(i).SetInt(status)
		case "uint":
			status := cast.ToUint(statusStr)
			v.Field(i).Set(reflect.ValueOf(status))
		case "uint8":
			status := cast.ToUint8(statusStr)
			v.Field(i).Set(reflect.ValueOf(status))
		case "uint16":
			status := cast.ToUint16(statusStr)
			v.Field(i).Set(reflect.ValueOf(status))
		case "uint32":
			status := cast.ToUint32(statusStr)
			v.Field(i).Set(reflect.ValueOf(status))
		case "uint64":
			status := cast.ToUint64(statusStr)
			v.Field(i).SetUint(status)
		case "string":
			// Set the field.
			v.Field(i).SetString(statusStr)
		default:
			panic(fmt.Sprintf("type not support:%s, tag=%s", state.Field(i).Type.String(), field.Tag))
		}
	}

	return nil
}

func InitMap(obj interface{}, StateMap interface{}) error {
	m := reflect.ValueOf(StateMap)
	state := reflect.TypeOf(obj).Elem()
	numFiled := state.NumField()

	for i := 0; i < numFiled; i++ {
		field := state.Field(i)

		statusStr := field.Tag.Get("v")
		desc := field.Tag.Get("d")

		switch field.Type.String() {
		case "int":
			status := cast.ToInt(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "int8":
			status := cast.ToInt8(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "int16":
			status := cast.ToInt16(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "int32":
			status := cast.ToInt32(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "int64":
			status := cast.ToInt64(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "uint":
			status := cast.ToUint(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "uint8":
			status := cast.ToUint8(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "uint16":
			status := cast.ToUint16(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "uint32":
			status := cast.ToUint32(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "uint64":
			status := cast.ToUint64(statusStr)
			m.SetMapIndex(reflect.ValueOf(status), reflect.ValueOf(desc))
		case "string":
			// Set the field.
			m.SetMapIndex(reflect.ValueOf(statusStr), reflect.ValueOf(desc))
		default:
			panic(fmt.Sprintf("type not support:%s, tag=%s", state.Field(i).Type.String(), field.Tag))
		}
	}

	return nil
}

func InitInverseMap(obj interface{}, RevStateMap interface{}) error {
	revM := reflect.ValueOf(RevStateMap)
	state := reflect.TypeOf(obj).Elem()
	numFiled := state.NumField()

	for i := 0; i < numFiled; i++ {
		field := state.Field(i)

		statusStr := field.Tag.Get("v")
		desc := field.Tag.Get("d")

		switch field.Type.String() {
		case "int":
			status := cast.ToInt(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "int8":
			status := cast.ToInt8(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "int16":
			status := cast.ToInt16(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "int32":
			status := cast.ToInt32(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "int64":
			status := cast.ToInt64(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "uint":
			status := cast.ToUint(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "uint8":
			status := cast.ToUint8(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "uint16":
			status := cast.ToUint16(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "uint32":
			status := cast.ToUint32(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "uint64":
			status := cast.ToUint64(statusStr)
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(status))
		case "string":
			// Set the field.
			revM.SetMapIndex(reflect.ValueOf(desc), reflect.ValueOf(statusStr))
		default:
			panic(fmt.Sprintf("type not support:%s, tag=%s", state.Field(i).Type.String(), field.Tag))
		}
	}

	return nil
}

func InitSlice(obj interface{}, slicePtr interface{}, itemType interface{}) error {
	state := reflect.TypeOf(obj).Elem()
	numFiled := state.NumField()

	// 获取切片的反射值
	sliceValue := reflect.ValueOf(slicePtr).Elem()
	// 确保传入的是切片
	if sliceValue.Kind() != reflect.Slice {
		return errors.New("传入的不是切片类型")
	}

	// 创建一个新的切片
	newSlice := reflect.MakeSlice(sliceValue.Type(), numFiled, numFiled)

	for i := 0; i < numFiled; i++ {
		field := state.Field(i)

		statusStr := field.Tag.Get("v")
		desc := field.Tag.Get("d")

		v := reflect.ValueOf(itemType).Elem()
		v.FieldByName("Desc").Set(reflect.ValueOf(desc))

		switch field.Type.String() {
		case "int":
			status := cast.ToInt(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "int8":
			status := cast.ToInt8(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "int16":
			status := cast.ToInt16(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "int32":
			status := cast.ToInt32(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "int64":
			status := cast.ToInt64(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "uint":
			status := cast.ToUint(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "uint8":
			status := cast.ToUint8(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "uint16":
			status := cast.ToUint16(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "uint32":
			status := cast.ToUint32(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "uint64":
			status := cast.ToUint64(statusStr)
			v.FieldByName("Value").Set(reflect.ValueOf(status))
		case "string":
			// Set the field.
			v.FieldByName("Value").Set(reflect.ValueOf(statusStr))
		default:
			panic(fmt.Sprintf("type not support:%s, tag=%s", state.Field(i).Type.String(), field.Tag))
		}

		newSlice.Index(i).Set(v)
	}

	// 将新切片的值设置回原始切片
	sliceValue.Set(newSlice)
	return nil
}

func GetStateList(obj interface{}) ([]SliceItem, error) {
	state := reflect.TypeOf(obj).Elem()
	numFiled := state.NumField()

	s := make([]SliceItem, numFiled, numFiled)
	rs := reflect.ValueOf(s)

	for i := 0; i < numFiled; i++ {
		field := state.Field(i)

		statusStr := field.Tag.Get("v")
		desc := field.Tag.Get("d")

		switch field.Type.String() {
		case "int":
			status := cast.ToInt(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "int8":
			status := cast.ToInt8(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "int16":
			status := cast.ToInt16(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "int32":
			status := cast.ToInt32(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "int64":
			status := cast.ToInt64(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "uint":
			status := cast.ToUint(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "uint8":
			status := cast.ToUint8(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "uint16":
			status := cast.ToUint16(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "uint32":
			status := cast.ToUint32(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "uint64":
			status := cast.ToUint64(statusStr)
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: status, Desc: desc}))
		case "string":
			// Set the field.
			rs.Index(i).Set(reflect.ValueOf(SliceItem{Value: statusStr, Desc: desc}))
		default:
			panic(fmt.Sprintf("type not support:%s, tag=%s", state.Field(i).Type.String(), field.Tag))
		}
	}

	return s, nil
}
