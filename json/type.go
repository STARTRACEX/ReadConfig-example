package json

type basetype struct {
	Title  string          `json:"title"`
	Author string          `json:"author"`
	Obj    innertype       `json:"obj"`
	Obj2   map[string]bool `json:"obj2"`
}
type innertype struct {
	Inner bool `json:"inner"`
}
