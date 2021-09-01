package domains

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
)

func SuccessResponse(c *gin.Context, status int, data interface{}, meta interface{}) {
	c.JSON(status, gin.H{"meta": meta, "data": data})
}

func SuccessResponseMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"meta": gin.H{"message": message}, "data": nil})
}

func ErrorResponse(c *gin.Context, status int, errors interface{}, meta interface{}) {
	c.JSON(status, gin.H{"meta": meta, "errors": errors})
}

func ErrorResponseMessage(c *gin.Context, status int, errors interface{}, message string) {
	c.JSON(status, gin.H{"meta": gin.H{"message": message}, "errors": errors})
}

func InternalServerError(c *gin.Context, err string) {
	var errMsg string
	localizer := c.MustGet("localizer").(*i18n.Localizer)

	if viper.GetString("application.env") == "production" {
		errMsg = localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID: "errors.common.internal_server_error",
			},
		})
	} else {
		errMsg = err
	}

	meta := gin.H{"message": ErrInternalServerError.Error()}
	errors := []gin.H{
		gin.H{"error": errMsg},
	}

	ErrorResponse(c, http.StatusInternalServerError, errors, meta)
}

func RecordNotFound(c *gin.Context, modelName string) {
	localizer := c.MustGet("localizer").(*i18n.Localizer)
	errMsg := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "errors.common.record_not_found",
		},
		TemplateData: map[string]string{
			"modelName": modelName,
		},
	})

	meta := gin.H{"message": ErrNotFound.Error()}
	errors := []gin.H{
		gin.H{"error": errMsg},
	}

	ErrorResponse(c, http.StatusNotFound, errors, meta)
}

func BadRequest(c *gin.Context, errors interface{}) {
	meta := gin.H{"message": ErrBadRequest.Error()}
	ErrorResponse(c, http.StatusBadRequest, errors, meta)
}

func Unauthorized(c *gin.Context) {
	localizer := c.MustGet("localizer").(*i18n.Localizer)
	errMsg := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "errors.common.unauthorized",
		},
	})

	meta := gin.H{"message": ErrUnauthorized.Error()}
	errors := []gin.H{
		gin.H{"error": errMsg},
	}

	ErrorResponse(c, http.StatusUnauthorized, errors, meta)
}

func SessionExpired(c *gin.Context) {
	localizer := c.MustGet("localizer").(*i18n.Localizer)
	errMsg := localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "errors.common.session_expired",
		},
	})

	meta := gin.H{"message": ErrUnauthorized.Error()}
	errors := []gin.H{
		gin.H{"error": errMsg},
	}

	ErrorResponse(c, http.StatusUnauthorized, errors, meta)
}

func Error(c *gin.Context, err error) {
	switch e := err.(type) {
	case *InternalErrorException:
		InternalServerError(c, e.Error())
	case *UnauthorizedException:
		Unauthorized(c)
	case *BadRequestException:
		BadRequest(c, []gin.H{gin.H{"error": err.Error()}})
	case *RecordNotFoundException:
		RecordNotFound(c, e.ModelName)
	default:
		InternalServerError(c, e.Error())
	}
}
