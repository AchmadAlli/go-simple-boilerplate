package validator

import (
	"github.com/achmadAlli/go-simple-boilerplate/adapters/validator"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	go_validator "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni *ut.UniversalTranslator
)

type playgroundValidator struct {
	validator *go_validator.Validate
	trans     ut.Translator
}

func NewPlaygroundValidator() validator.Validator {
	playgroundValidator := new(playgroundValidator)

	validator := go_validator.New()
	language := en.New()
	uni = ut.New(language, language)
	trans, _ := uni.GetTranslator("language")
	en_translations.RegisterDefaultTranslations(validator, trans)

	playgroundValidator.validator = validator
	playgroundValidator.trans = trans

	return playgroundValidator
}

func (v *playgroundValidator) Validate(s interface{}) error {
	if err := v.validator.Struct(s); err != nil {
		result := make(map[string]string)
		errs := err.(go_validator.ValidationErrors)

		for _, e := range errs {
			switch e.Tag() {
			default:
				result[e.Field()] = e.Translate(v.trans) //translate each error one at a time.
			}
		}

		return &validator.ValidationError{
			Err:     validator.ErrValidation,
			ErrData: result,
		}
	}

	return nil
}
