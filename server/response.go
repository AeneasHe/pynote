package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func ResponseOutput(c echo.Context, code int, message string, result interface{}) error {
	return c.JSON(http.StatusOK, ApiResponse{
		Code:    code,
		Message: message,
		Result:  result,
	})
}
func ResponseSuccess(ctx echo.Context, result interface{}) error {
	c := ctx.(*CustomContext)
	switch c.Lang() {
	case "ja":
		return ResponseOutput(c, 0, "成功した操作", result)
	default:
		return ResponseOutput(c, 0, "操作成功", result)
	}
}

func ResponseError(c echo.Context, code int, message string) error {
	return ResponseOutput(c, code, message, nil)
}
