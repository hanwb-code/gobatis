package gobatis

import (
	"reflect"
)

func GetResultType(res interface{}) ResultType {
	resType := reflect.TypeOf(res)

	// 检查res是否为nil
	if resType == nil {
		return "" // 假设有一个专门表示未知类型的常量
	}

	// 如果res是指针，获取其指向的元素的类型
	if resType.Kind() == reflect.Ptr {
		resType = resType.Elem()
	}

	switch resType.Kind() {
	case reflect.Map:
		return resultTypeMap
	case reflect.Slice, reflect.Array:
		// 进一步检查切片或数组的元素类型
		elemType := resType.Elem()
		if elemType.Kind() == reflect.Ptr {
			// 如果元素是指针，获取指针指向的类型
			elemType = elemType.Elem()
		}
		if elemType.Kind() == reflect.Struct {
			// 如果元素是结构体（或指向结构体的指针），返回resultTypeStructs
			return resultTypeStructs
		}
		// 这里可以根据需要添加更多的判断逻辑
		return resultTypeSlices
	case reflect.Struct:
		return resultTypeStruct
	default:
		return resultTypeValue
	}
}
