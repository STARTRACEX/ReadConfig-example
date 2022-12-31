package yaml

import (
	"fmt"
	"os"
	"readconfig/types"
	"github.com/goccy/go-yaml"
)

func ReadYaml(path string) types.Maintype {
	fileread, _ := os.ReadFile(path)
	var config types.Maintype
	e := yaml.Unmarshal(fileread, &config)
	if e != nil {
		panic("yaml.Unmarshal failed")
	}
	fmt.Println("Author:"+config.Author, "Title:", config.Title, config.Obj.Inner, config.Obj2["inner"])
	return config
}