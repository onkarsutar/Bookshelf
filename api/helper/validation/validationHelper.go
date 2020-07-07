package validationhelper

import (
	"github.com/go-playground/locales/en"

	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	en_translations "gopkg.in/go-playground/validator.v9/translations/en"
)

//Validate method
func Validate(s interface{}) map[string]string {
	var validate *validator.Validate
	var uni *ut.UniversalTranslator

	validate = validator.New()

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		customErrs := make(map[string]string, len(errs))

		for _, e := range errs {
			customErrs[e.Namespace()] = e.Translate(trans)
		}
		return customErrs
	}
	return nil
}
