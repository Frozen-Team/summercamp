package forms

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// validate performs the validation of the form fields based on
// the `valid` tag for each struct field
func validate(whatToCheck interface{}) (errs []error, ok bool) {
	valid := validation.Validation{}
	ok, err := valid.Valid(whatToCheck)
	if err != nil {
		errs = append(errs, errors.New("validation-failed"))
		beego.BeeLogger.Error("validation process error: %v", err)
		return
	}
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			e := fmt.Errorf("%v : %v", err.Key, err.Message)
			errs = append(errs, e)
			beego.Warning(e)
		}
	}
	return
}
