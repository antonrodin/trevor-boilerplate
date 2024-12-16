package forms

type errors map[string][]string

// Add to error slice a message to given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Returns a first error message from error list related to some field
func (e errors) Get(field string) string {
	fieldErrors := e[field]

	if len(fieldErrors) == 0 {
		return ""
	}

	return fieldErrors[0]
}
