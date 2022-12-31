package json

import (
	"readconfig/types"
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson(path string) types.Maintype {
	fileread, _ := os.ReadFile(path)
	var config types.Maintype
	e := json.Unmarshal(fileread, &config)
	if e != nil {
		panic("json.Unmarshal failed")
	}
	fmt.Println("Author:"+config.Author, "Title:", config.Title, config.Obj.Inner, config.Obj2["inner"])
	return config
}