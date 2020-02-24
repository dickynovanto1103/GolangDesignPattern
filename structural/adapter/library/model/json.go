package model

type JSON struct {
	field string
}

func NewJSONRequest(field string) *JSON {
	return &JSON{field: field}
}
