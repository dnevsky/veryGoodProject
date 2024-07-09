package helpers

import (
	"errors"
	"github.com/dnevsky/veryGoodProject/internal/models"
	repoErrors "github.com/dnevsky/veryGoodProject/internal/repository/errors"
	"github.com/dnevsky/veryGoodProject/transport/rest/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (m *Manager) getErrCodeType(err error) (int, models.ErrType) {
	if strings.Contains(err.Error(), models.ErrBrokenPipe.Error()) ||
		strings.Contains(err.Error(), models.ErrConnectionResetByPeer.Error()) {
		return http.StatusBadRequest, models.ErrTypeDefault
	}

	if errors.Is(err, models.ErrUnauthorized) {
		return http.StatusUnauthorized, models.ErrTypeDefault
	}

	if errors.Is(err, repoErrors.ErrNotFound) {
		return http.StatusNotFound, models.ErrTypeDefault
	}

	if errors.Is(err, models.ErrInvalidRequestParams) ||
		errors.Is(err, models.ErrBadLoginOrPassword) ||
		errors.Is(err, models.ErrEmptyAuthHeader) ||
		errors.Is(err, models.ErrInvalidAuthHeader) ||
		errors.Is(err, models.ErrTokenIsEmpty) ||
		errors.Is(err, models.ErrTokenExpired) ||
		errors.Is(err, models.ErrBadAssetName) ||
		errors.Is(err, models.ErrEmptyBody) ||
		errors.Is(err, repoErrors.ErrAlreadyExists) {

		return http.StatusBadRequest, models.ErrTypeDefault
	}

	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		return http.StatusBadRequest, models.ErrTypeValidation
	} else {
		return http.StatusInternalServerError, models.ErrTypeDefault
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "min":
		return "Minimum string length is " + fe.Param()
	case "max":
		return "Maximum string length is " + fe.Param()
	case "oneof":
		return "Field can be one of: " + fe.Param()
	}
	return "Unknown error"
}

func (m *Manager) canLogError(code int) bool {
	if code >= 500 {
		return true
	}
	return false
}

func (m *Manager) defaultValidationErrorsHandle(c *gin.Context, err error, errCode int) {
	response.JsonResponse(c.Writer, response.Data{
		Code: errCode,
		Text: err.Error(),
	})
}

func (m *Manager) validationErrorsHandle(c *gin.Context, err error, errCode int) {
	var validationErr validator.ValidationErrors

	if errors.As(err, &validationErr) {
		out := make([]ErrorMsg, len(validationErr))
		for i, vErr := range err.(validator.ValidationErrors) {
			out[i] = ErrorMsg{vErr.Field(), getErrorMsg(vErr)}
		}
		response.JsonResponse(c.Writer, response.Data{
			Code:         http.StatusBadRequest,
			ClientErrors: out,
		})
	} else {
		m.defaultValidationErrorsHandle(c, err, errCode)
	}
}

func (m *Manager) LogError(err error) {
	errCode, _ := m.getErrCodeType(err)
	if m.canLogError(errCode) {
		m.Logger.Error(err)
	}
}

func (m *Manager) ErrorsHandle(c *gin.Context, err error) {
	m.LogError(err)

	errCode, errType := m.getErrCodeType(err)
	switch errType {
	case models.ErrTypeDefault:
		m.defaultValidationErrorsHandle(c, err, errCode)
	case models.ErrTypeValidation:
		m.validationErrorsHandle(c, err, errCode)
	default:
	}
}
