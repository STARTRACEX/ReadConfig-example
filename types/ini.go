package types

type Initype struct {
	Author string    `ini:"author"`
	Title  string    `ini:"title"`
	Obj    innertype `ini:"obj"`
}
type Iniinnertype struct {
	Inner bool `ini:"inner"`
}
