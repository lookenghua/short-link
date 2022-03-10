package util

import (
	"github.com/gofiber/fiber/v2"
)
import . "short-link/common/validation"

// 错误代码
const (
	ParseError    = iota + 1001 // 解析错误
	ValidateError               // 校验错误
)

// ApiUtil API工具
type ApiUtil struct {
	Ctx *fiber.Ctx
}

// Response 相应格式
type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage *string     `json:"errorMessage,omitempty"`
	ErrorCode    *int        `json:"errorCode,omitempty"`
}

// ValidateBody 校验body
func (app *ApiUtil) ValidateBody(data interface{}) error {
	if err := app.Ctx.BodyParser(data); err != nil {
		if err != nil {
			return fiber.NewError(ParseError, "解析失败")
		}

	}
	errors := ValidateStruct(data)
	if errors != nil {
		return fiber.NewError(ValidateError, *errors)

	}
	return nil
}

// Success 成功返回
func (app *ApiUtil) Success(data interface{}) error {
	return app.Ctx.JSON(&Response{
		Success: true,
		Data:    data,
	})
}

// Fail 失败返回
func (app *ApiUtil) Fail(errorCode int, errorMessage string) error {
	return app.Ctx.JSON(&Response{
		Success:      false,
		ErrorCode:    &errorCode,
		ErrorMessage: &errorMessage,
	})
}
