package validate

import (
	"reflect"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"hello/internal/pkg/response"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate(v interface{}) error {
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ := uni.GetTranslator("zh")

	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	zh_translations.RegisterDefaultTranslations(validate, trans)

	if err := validate.Struct(v); err != nil {
		errs := err.(validator.ValidationErrors)

		for _, err := range errs.Translate(trans) {
			return errors.New(0, response.VALIDATE_ERROR, err)
		}
	}

	return nil
}
