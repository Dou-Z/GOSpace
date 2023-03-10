package ini

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

func MarshalFile(fileName string, data interface{}) (err error) {
	result, err := Marshal(data)
	if err != nil {
		return
	}
	return ioutil.WriteFile(fileName, result, 0755)
}

func Marshal(data interface{}) (result []byte, err error) {
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() != reflect.Struct {
		err = errors.New("Please pass address")
		return
	}
	var conf []string
	valueInfo := reflect.ValueOf(data)
	for i := 0; i < typeInfo.NumField(); i++ {
		sectionField := typeInfo.Field(i)
		sectionVal := valueInfo.Field(i)

		fieldType := sectionField.Type
		if fieldType.Kind() != reflect.Struct {
			continue
		}
		tagVal := sectionField.Tag.Get("ini")
		if len(tagVal) == 0 {
			tagVal = sectionField.Name
		}
		section := fmt.Sprintf("\n[%s]\n", tagVal)
		conf = append(conf, section)

		for j := 0; j < fieldType.NumField(); j++ {
			keyField := fieldType.Field(j)
			fieldTafVal := keyField.Tag.Get("ini")
			if len(fieldTafVal) == 0 {
				fieldTafVal = keyField.Name
			}
			valField := sectionVal.Field(j)
			item := fmt.Sprintf("%s=%v\n", fieldTafVal, valField.Interface())
			// fmt.Println(item)
			conf = append(conf, item)
		}
	}
	for _, val := range conf {
		// fmt.Println(val)
		byteVal := []byte(val)
		result = append(result, byteVal...)

	}
	return
}

func UnMarshalFile(fileName string, data interface{}) (err error) {
	result, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	return UnMarshal(result, result)
}

func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")

	typeInfo := reflect.TypeOf(result)
	if typeInfo.Kind() != reflect.Ptr {
		err = errors.New("Please pass address")
		return
	}
	typeStruct := typeInfo.Elem()
	if typeStruct.Kind() != reflect.Struct {
		err = errors.New("Please pass struct")
		return
	}
	var lastFieldName string
	for index, line := range lineArr {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		//  ??????????????????????????????
		if line[0] == ';' || line[0] == '#' {
			continue
		}

		if line[0] == '[' {
			lastFieldName, err = parseSection(line, typeStruct)
			if err != nil {
				err = fmt.Errorf("%v lineno:%d", err, index+1)
				return
			}
			continue
		}
		err = parseItem(lastFieldName, line, result)
		if err != nil {
			err = fmt.Errorf("%v lineno :%d", err, index+1)
			return
		}
	}

	return
}
func parseItem(lastFieldName string, line string, result interface{}) (err error) {
	index := strings.Index(line, "=")
	if index == -1 {
		err = fmt.Errorf("sytax error ,line :%s", line)
		return
	}
	key := strings.TrimSpace(line[0:index])
	val := strings.TrimSpace(line[index+1:])
	_ = val
	if len(key) == 0 {
		err = fmt.Errorf("sytax error ,line:%s", line)
		return
	}
	resultValue := reflect.ValueOf(result)
	sectionValue := resultValue.Elem().FieldByName(lastFieldName)
	sectionType := sectionValue.Type()
	if sectionType.Kind() != reflect.Struct {
		err = fmt.Errorf("field :%s must be struct", lastFieldName)
		return
	}
	keyFieldName := ""
	for i := 0; i < sectionType.NumField(); i++ {
		field := sectionType.Field(i)
		tagVal := field.Tag.Get("ini")
		if tagVal == key {
			keyFieldName = field.Name
			break
		}
	}
	if len(keyFieldName) == 0 {
		return
	}
	fmt.Println(keyFieldName)
	fieldValue := sectionValue.FieldByName(keyFieldName)
	if fieldValue == reflect.ValueOf(nil) {
		return
	}
	switch fieldValue.Type().Kind() {
	case reflect.String:
		fieldValue.SetString(val)
	case reflect.Int8, reflect.Int16, reflect.Int, reflect.Int64, reflect.Int32:
		intVal, errRet := strconv.ParseInt(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetInt(intVal)
	case reflect.Uint8, reflect.Uint16, reflect.Uint, reflect.Uint64, reflect.Uint32:
		intVal, errRet := strconv.ParseUint(val, 10, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetUint(intVal)
	case reflect.Float32, reflect.Float64:
		intVal, errRet := strconv.ParseFloat(val, 64)
		if errRet != nil {
			err = errRet
			return
		}
		fieldValue.SetFloat(intVal)
	default:
		err = fmt.Errorf("unsupport type:%v", fieldValue.Type().Kind())
	}
	return

}

func parseSection(line string, typeInfo reflect.Type) (filedName string, err error) {
	if line[0] == '[' && len(line) <= 2 {
		err = fmt.Errorf("syntax erroe ,invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] != ']' {
		err = fmt.Errorf("syntax erroe ,invalid section:%s", line)
		return
	}
	if line[0] == '[' && line[len(line)-1] == ']' {
		sectionName := strings.TrimSpace(line[1 : len(line)-1])
		if len(sectionName) == 0 {
			err = fmt.Errorf("syntax erroe ,invalid section:%s", line)
			return
		}
		for i := 0; i < typeInfo.NumField(); i++ {
			field := typeInfo.Field(i)
			tagValue := field.Tag.Get("ini")
			if tagValue == sectionName {
				filedName = field.Name
				// fmt.Println("field name:", lastSectionFieldName)
				break
			}
		}
	}
	return
}
