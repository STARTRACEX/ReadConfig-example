package types

type Maintype struct {
	Title  string          `json:"title" yaml:"title" toml:"title"`
	Author string          `json:"author" yaml:"author" toml:"author"`
	Obj    innertype       `json:"obj" yaml:"obj" toml:"obj"`
	Obj2   map[string]bool `json:"obj2" yaml:"obj2" toml:"obj2"`
}
type innertype struct {
	Inner bool `json:"inner" yaml:"inner" toml:"inner"`
}
