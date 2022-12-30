package toml

type basetype struct {
	Author string          `toml:"author,omitempty"`
	Title  string          `toml:"title,omitempty"`
	Obj    innertype       `toml:"obj"`
	Obj2   map[string]bool `toml:"obj2"`
}
type innertype struct {
	Inner bool `toml:"inner"`
}
