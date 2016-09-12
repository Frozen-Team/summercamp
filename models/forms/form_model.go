package forms

type FormModel struct {
	Errors []string `json:"-"`
}

func (f *FormModel) addError(e string) {
	f.Errors = append(f.Errors, e)
}
