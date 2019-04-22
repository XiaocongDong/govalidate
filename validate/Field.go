package validate

type Validate struct {
}

type FieldError struct {
	tag string
	param string
	field string
}

func (f *FieldError) Tag () string {
	return f.tag
}

func (f *FieldError) Param () string {
	return f.param
}

func (f *FieldError) Field () string {
	return f.field
}
