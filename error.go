package error

import "github.com/labstack/echo"

// BaseError : base of the error
type BaseError struct {
	Code    int
	Message string
}

// SendNoControlledError : send this message if the error is not controlled
func SendNoControlledError(c echo.Context) error {
	errorResult := &BaseError{Code: 500, Message: "no controlled error ocurred"}
	return c.JSON(errorResult.Code, errorResult)
}

// SendValidationError : send validation erro
func SendValidationError(c echo.Context, message string) error {
	if !(len(message) > 0) {
		message = "Validation not successful"
	}
	errorResult := &BaseError{Code: 422, Message: message}
	return c.JSON(errorResult.Code, errorResult)
}

//SendSpecificError : send an especific error
func SendSpecificError(c echo.Context, message string) error {
	if !(len(message) > 0) {
		message = "an error occurred"
	}
	errorResult := &BaseError{Code: 422, Message: message}
	return c.JSON(errorResult.Code, errorResult)
}
