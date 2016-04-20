package soul

type TypeError struct {
	object       string
	expectedType string
	field        string
	value        interface{}
}

func (e *TypeError) Error() string {
	return e.object + " type error, expected: " + e.expectedType + " in field " + e.field
}

type NilField struct {
	object string
	field  string
}

func (e *NilField) Error() string {
	return "In object " + e.object + " field " + e.field + " could not be nil."
}
