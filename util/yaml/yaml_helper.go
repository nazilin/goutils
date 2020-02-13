package yamlhelper

import (
	"fmt"
	"gopkg.in/yaml.v2"
)


func ReadStringList(in interface{}) ([]string,bool){
	array,ok:=in.([]interface{})
	if ok {
		list := make ([]string,0)
		for _,i:= range array {
			str,ok:=i.(string)
			if ok {
				list=append(list,str)
			}else{
				return nil,false
			}
		}
		return list,true
	}
	return nil,false
}
// UnpackMap gets a yaml.MapSlice if possible.
func UnpackMap(in interface{}) (yaml.MapSlice, bool) {
	m, ok := in.(yaml.MapSlice)
	if ok {
		return m, true
	}
	// do we have an empty array?
	a, ok := in.([]interface{})
	if ok && len(a) == 0 {
		// if so, return an empty map
		return yaml.MapSlice{}, true
	}
	return nil, false
}
// MapHasKey returns true if a yaml.MapSlice contains a specified key.
func MapHasKey(m yaml.MapSlice, key string) bool {
	for _, item := range m {
		itemKey, ok := item.Key.(string)
		if ok && key == itemKey {
			return true
		}
	}
	return false
}

// MapValueForKey gets the value of a map value for a specified key.
func MapValueForKey(m yaml.MapSlice, key string) interface{} {
	for _, item := range m {
		itemKey, ok := item.Key.(string)
		if ok && key == itemKey {
			return item.Value
		}
	}
	return nil
}
// StringValue returns the string value of an item.
func StringValue(item interface{}) (value string, ok bool) {
	value, ok = item.(string)
	if ok {
		return value, ok
	}
	intValue, ok := item.(int)
	if ok {
		return fmt.Sprintf("%d",intValue), true
	}
	return "", false
}
