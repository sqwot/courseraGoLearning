package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type User struct {
	ID       int
	RealName string `unpack:"-"`
	Login    string
	Flags    int
}

func PrintReflect(u interface{}) error {
	val := reflect.ValueOf(u).Elem()

	fmt.Printf("%T have %d fields:\n", u, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		fmt.Printf("\tname=%v, type=%v, value=%v, tag=`%v`\n",
			typeField.Name,
			typeField.Type.Kind(),
			valueField,
			typeField.Tag,
		)
	}
	return nil
}

func UnpackReflect(u interface{}, data []byte) error {
	r := bytes.NewReader(data)

	val := reflect.ValueOf(u).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		if typeField.Tag.Get("Unpack") == "-" {
			continue
		}

		switch typeField.Type.Kind() {
		case reflect.Int:
			var value uint32
			binary.Read(r, binary.LittleEndian, &value)
			valueField.Set(reflect.ValueOf(int(value)))
		case reflect.String:
			var lenRaw uint32
			binary.Read(r, binary.LittleEndian, &lenRaw)

			dataRaw := make([]byte, lenRaw)
			binary.Read(r, binary.LittleEndian, &dataRaw)

			valueField.SetString(string(dataRaw))
		default:
			return fmt.Errorf("bad type^ %v for field %v", typeField.Type.Kind(), typeField.Name)

		}
	}
	return nil
}

func main() {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "***************Reflect***************")
	fmt.Println(string(colorReset), "")

	u := &User{
		ID:       42,
		RealName: "sqwot",
		Flags:    32,
	}
	err := PrintReflect(u)
	if err != nil {
		panic(err)
	}

	data := []byte{
		128, 36, 17, 0,

		9, 0, 0, 0,
		118, 46, 114, 111, 109, 97, 110, 111, 118,

		16, 0, 0, 0,
	}
	u2 := new(User)
	err2 := UnpackReflect(u2, data)
	if err2 != nil {
		panic(err)
	}

	fmt.Printf("%#v", u2)
}
