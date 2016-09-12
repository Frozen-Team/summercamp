package forms

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type FormModel struct {
	Errors []string `json:"-"`
}

func (f *FormModel) addError(e string) {
	f.Errors = append(f.Errors, e)
}

// validate performs the validation of the form whatToCheck based on
// the `valid` tag for each field of the form struct.
// It's common that whatToCheck == f, as we extend each form struct with FormModel
func (f *FormModel) validate(whatToCheck interface{}) bool {
	valid := validation.Validation{}
	ok, err := valid.Valid(whatToCheck)
	if err != nil {
		f.addError("validation-failed")
		beego.BeeLogger.Error("validation process error: %v", err)
		return ok
	}
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			e := fmt.Sprintf("%v : %v", err.Key, err.Message)
			f.addError(e)
			beego.Warning(e)
		}
	}
	return ok
}
