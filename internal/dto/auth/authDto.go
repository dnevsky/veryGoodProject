package auth

import (
	"github.com/dnevsky/veryGoodProject/internal/dto"
	"github.com/go-playground/validator/v10"
	"github.com/leebenson/conform"
)

type AuthDTO struct {
	dto.ServiceDTO `swaggerignore:"true"`
	Login          string `form:"login" json:"login" conform:"trim" validate:"required,min=3,max=255"`
	Password       string `form:"password" json:"password" conform:"trim" validate:"required,min=3,max=255"`
	Ip             string `swaggerignore:"true"`
}

func (dto *AuthDTO) Validate() error {
	if err := conform.Strings(dto); err != nil {
		return err
	}
	if err := validator.New().Struct(dto); err != nil {
		return err
	}
	return nil
}
