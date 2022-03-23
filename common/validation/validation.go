package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

var validate = validator.New()

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
	Json        string
}

func ValidateStruct(u interface{}) *string {
	var errors []*ErrorResponse
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			field := err.Field()
			FieldInfo, ok := reflect.TypeOf(u).Elem().FieldByName(field)
			if !ok {
				panic("Field not found")
			}
			jsonName := FieldInfo.Tag.Get("json")
			fmt.Printf("Tag: %v\n", err.Tag())
			fmt.Printf("ActualTag: %v\n", err.ActualTag())
			fmt.Printf("Namespace: %v\n", err.Namespace())
			fmt.Printf("StructNamespace: %v\n", err.StructNamespace())
			fmt.Printf("Field: %v\n", err.Field())
			fmt.Printf("StructField: %v\n", err.StructField())
			fmt.Printf("Value: %v\n", err.Value())
			fmt.Printf("Param: %v\n", err.Param())
			fmt.Printf("Kind: %v\n", err.Kind())
			fmt.Printf("Type: %v\n", err.Type())
			fmt.Printf("Error: %v\n", err.Error())
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Value()
			element.Json = jsonName
			errors = append(errors, &element)
		}
	}
	if len(errors) > 0 {
		str := ""
		firstErr := errors[0]
		switch firstErr.Tag {
		case "email":
			str = fmt.Sprintf("参数%v的值%v不是合法邮箱", firstErr.Json, firstErr.Value)
			break
		case "phone":
			str = fmt.Sprintf("参数%v的值%v不是合法手机号", firstErr.Json, firstErr.Value)
			break
		default:
			str = fmt.Sprintf("参数%v的值%v发生%v错误", firstErr.Json, firstErr.Value, firstErr.Tag)
			break
		}
		return &str
	}
	return nil
}

// InitValidation 初始化自定义校验
func InitValidation() {
	// 校验手机号
	err := validate.RegisterValidation("phone", isPhone)
	if err != nil {
		fmt.Println("===注册phone失败===")
		return
	}
}
