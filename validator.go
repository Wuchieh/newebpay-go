package ebpay

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"log"
)

var (
	validate = validator.New(validator.WithRequiredStructEnabled())
)

func validateItemType(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(ItemType)
	if !ok {
		return false
	}

	switch value {
	case ItemTypeProduct, ItemTypeTickets, ItemTypeReserve:
		return true
	default:
		return false
	}
}

func validateResponseType(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(RespondType)
	if !ok {
		return false
	}

	switch value {
	case RespondTypeJSON, RespondTypeString:
		return true
	default:
		return false
	}
}

func validateLangType(fl validator.FieldLevel) bool {
	value, ok := fl.Field().Interface().(LangType)
	if !ok {
		return false
	}
	switch value {
	case LangTypeEN, LangTypeTW, LangTypeJP:
		return true
	default:
		return false
	}
}

func isValidatorErr(err error) (validator.ValidationErrors, bool) {
	var vErr validator.ValidationErrors
	if errors.As(err, &vErr) {
		return vErr, true
	}
	return nil, false
}

func init() {
	if err := validate.RegisterValidation("item_type", validateItemType); err != nil {
		log.Println("RegisterValidation \"itemType\" Fail")
	}

	if err := validate.RegisterValidation("response_type", validateResponseType); err != nil {
		log.Println("RegisterValidation \"response_type\" Fail")
	}

	if err := validate.RegisterValidation("lang_type", validateLangType); err != nil {
		log.Println("RegisterValidation \"lang_type\" Fail")
	}

}
