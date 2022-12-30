package yaml

type basetype struct {
	Author string          `yaml:"author,omitempty"`
	Title  string          `yaml:"title,omitempty"`
	Obj    innertype       `toml:"obj"`
	Obj2   map[string]bool `toml:"obj2"`
}
type innertype struct {
	Inner bool `toml:"inner"`
}
