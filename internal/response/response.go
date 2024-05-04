package response

import (
	"music-libray-management/internal/helper"
	"music-libray-management/usecase/auth"
	"net/http"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResponseOK struct {
	Data   interface{} `json:"data"`
	Total  interface{} `json:"total,omitempty"`
	Status bool        `json:"status"`
}

type ResponseError struct {
	Status int         `json:"status"`
	Error  *string     `json:"error,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type ResponseBindingError struct {
	Field string  `json:"field"`
	Tag   string  `json:"tag"`
	Error *string `json:"error,omitempty"`
}

type ServiceResponse interface {
	ResponseSuccess(*gin.Context, interface{}, bool)
	AbortWithAppBindingError(*gin.Context, error)
	AbortWithAppSingleError(*gin.Context, error)
	AbortWithAppMultipleError(*gin.Context, []error)
}

type serviceResponse struct {
}

func NewServiceResponse() ServiceResponse {
	return &serviceResponse{}
}

func (s *serviceResponse) ResponseSuccess(ctx *gin.Context, data interface{}, status bool) {
	obj := &ResponseOK{
		Data:   data,
		Status: status,
	}
	if helper.IsSlice(data) || helper.IsArray(data) {
		values := reflect.ValueOf(data)
		obj.Total = values.Len()
	}
	ctx.JSON(http.StatusOK, obj)
}

func (s *serviceResponse) AbortWithAppBindingError(ctx *gin.Context, err error) {
	var structError []ResponseBindingError

	if vErrs, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range vErrs {
			fieldName := vErr.Namespace()
			validationTag := vErr.Tag()
			errorMessage := vErr.Param()

			structError = append(structError, ResponseBindingError{
				Field: helper.LowerCaseFieldName(fieldName),
				Tag:   validationTag,
				Error: &errorMessage,
			})
		}
	}

	obj := &ResponseError{
		Status: http.StatusBadRequest,
		Errors: &structError,
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, obj)
}

func (s *serviceResponse) AbortWithAppSingleError(ctx *gin.Context, err error) {
	obj := &ResponseError{
		Error: aws.String(err.Error()),
	}
	switch err {
	case auth.ErrRegisterFailed:
		obj.Status = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(obj.Status, obj)
		return
	case auth.ErrPasswordConfirmPasswordDifference:
		fallthrough
	case auth.ErrUserExisted:
		fallthrough
	case auth.ErrGenPasswordFailed:
		fallthrough
	case auth.ErrGenAccessTokenFailed:
		fallthrough
	case auth.ErrGenRefreshTokenFailed:
		fallthrough
	case auth.ErrLoginFailed:
		obj.Status = http.StatusBadRequest
		ctx.AbortWithStatusJSON(obj.Status, obj)
	case auth.ErrUserNotFound:
		obj.Status = http.StatusNotFound
		ctx.AbortWithStatusJSON(obj.Status, obj)
	default:
		obj.Status = http.StatusInternalServerError
		ctx.AbortWithStatusJSON(obj.Status, obj)
		return
	}
}

func (s *serviceResponse) AbortWithAppMultipleError(ctx *gin.Context, errs []error) {}
