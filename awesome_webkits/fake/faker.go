package fake

import (
	"awesome_webkits/fake/lib"
	"errors"
	"reflect"
)

//type Faker struct{}
//
//var fakerMap = map[string]interface{}{
//	"FemaleName": lib.RandomFemaleName,
//	"MaleName":   lib.RandomMaleName,
//}
//
//func CallDynamically(name string, args ...interface{}) string {
//	var value string
//
//	switch name {
//	case "FemaleName":
//		value = fakerMap["FemaleName"].(func() string)()
//	case "MaleName":
//		value = fakerMap["MaleName"].(func() string)()
//	}
//
//	return value
//}

type stubMapping map[string]interface{}

var StubStorage = map[string]interface{}{
	"female_name":   lib.RandomFemaleName,
	"male_name":     lib.RandomMaleName,
	"gender":        lib.RandomGender,
	"last_name":     lib.RandomSurName,
	"username":      lib.RandomUsername,
	"activity":      lib.RandomActivity,
	"avatar":        lib.RandomAvatar,
	"birth_day":     lib.RandomBirthDay,
	"color":         lib.RandomColor,
	"rgb":           lib.RandomRGBColor,
	"email":         lib.RandomEmail,
	"food":          lib.RandomFood,
	"word":          lib.Word,
	"sentence":      lib.Sentence,
	"url":           lib.Url,
	"host":          lib.Host,
	"national_code": lib.RandomNationalCode,
	"male_title":    lib.RandomMaleTitle,
	"female_title":  lib.RandomFemaleTitle,
	"uni_title":     lib.RandomUniSexTitle,
	"timestamp":     lib.RandomTimestamp,
	"date":          lib.RandomDate,
	"ipv4":          lib.RandomIpV4,
	"ipv6":          lib.RandomIpV6,
}

func Call(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(StubStorage[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	return
}
