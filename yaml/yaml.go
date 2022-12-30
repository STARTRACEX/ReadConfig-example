package yaml

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"os"
)

func ReadYaml(path string) {
	fileread, _ := os.ReadFile(path)
	var config basetype
	yaml.Unmarshal(fileread, &config)
	fmt.Println("Author:"+config.Author, "Title:", config.Title, config.Obj.Inner, config.Obj2["inner"])

}
