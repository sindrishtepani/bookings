package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initalizes a form struct
func New(data url.Values) *Form {

	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if form has a specific field
func (f *Form) Has(field string) bool {
	x := f.Values.Get(field)
	if x == "" {
		f.Errors.Add(field, "Does not exist in form")
		return false
	}

	return true
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Required checks for required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks for string minium length
func (f *Form) MinLength(field string, length int) bool {
	x := f.Values.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d chars long", length))
		return false
	}

	return true
}

// IsEmail adds an error to Error object if field is not a valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
