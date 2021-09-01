package domains

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrNotFound            = errors.New("Your requested item is not found")
	ErrConflict            = errors.New("Your item already exist")
	ErrBadRequest          = errors.New("Given param is not valid")
	ErrUnprocessableEntity = errors.New("Unprocessable entity")
	ErrUnauthenticate      = errors.New("Unauthenticate")
	ErrUnauthorized        = errors.New("Unauthorized")
)

func FormatValidationErrors(localizer *i18n.Localizer, err error) ([]gin.H, error) {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return []gin.H{}, err
	}

	var errs []gin.H

	for _, err := range validationErrors {
		field := err.Field()
		tag := err.Tag()

		errMsg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID: "errors.validation." + tag,
			},
			TemplateData: map[string]string{
				"field": field,
			},
		})

		errs = append(errs, gin.H{"error": errMsg})
	}

	return errs, nil
}

type InternalErrorException struct {
	Err string
}

func NewInternalErrorException(err string) *InternalErrorException {
	return &InternalErrorException{
		Err: err,
	}
}

func (e *InternalErrorException) Error() string {
	return e.Err
}

type UnauthorizedException struct{}

func NewUnauthorizedException() *UnauthorizedException {
	return &UnauthorizedException{}
}

func (e *UnauthorizedException) Error() string {
	return ErrUnauthorized.Error()
}

type BadRequestException struct {
	Err string
}

func NewBadRequestException(err string) *BadRequestException {
	return &BadRequestException{
		Err: err,
	}
}

func (e *BadRequestException) Error() string {
	return e.Err
}

type RecordNotFoundException struct {
	ModelName string
}

func NewRecordNotFoundException(modelName string) *RecordNotFoundException {
	return &RecordNotFoundException{
		ModelName: modelName,
	}
}

func (e *RecordNotFoundException) Error() string {
	return ErrNotFound.Error()
}
