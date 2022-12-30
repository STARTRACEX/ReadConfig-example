package ini

type basetype struct {
	Author string    `ini:"author"`
	Title  string    `ini:"title"`
	Obj    innertype `ini:"obj"`
}
type innertype struct {
	Inner bool `ini:"inner"`
}
