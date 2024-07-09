package helpers

import (
	"github.com/dnevsky/veryGoodProject/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
)

const (
	authUserIdField = "XUserID"
)

type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

func (m *Manager) BindData(c *gin.Context, req interface{}) error {
	err := c.Request.ParseForm()
	if err != nil {
		return err
	}

	if err := c.Bind(req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			return models.ErrInvalidRequestParams
		}

		return err
	}

	// тут можно обогатить дтошку данными. можно добавить модельку юзера и заполнить её данными, к примеру.
	// эту структуру можно описать в dto/service_dto.go. потом эту структуру встраивать в каждую дтошку и мы всегда будем
	// обогащать эту дтошку данными. Кроме юзера можно передавать какой-нибудь трейс, версию фронта и тд
	// для этой задачи это не нужно, но такое возможно

	err = m.BindAuthUser(c, &req)

	return nil
}

func (m *Manager) BindAuthUser(c *gin.Context, req *interface{}) error {
	uid, err := m.GetUserIdAuthorization(c)
	if err != nil {
		return err
	}

	v := reflect.ValueOf(*req).Elem()

	if f := v.FieldByName(authUserIdField); f.IsValid() {
		f.Set(reflect.ValueOf(uid))
	}

	return nil
}

func (m *Manager) GetIdFromPath(c *gin.Context, key string) (uint, error) {
	param := c.Param(key)
	intParam, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return uint(intParam), nil
}
