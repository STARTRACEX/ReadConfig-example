package json

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson(path string) {
	fileread, _ := os.ReadFile(path)
	var config basetype
	json.Unmarshal(fileread, &config)
	fmt.Println("Author:"+config.Author, "Title:", config.Title, config.Obj.Inner, config.Obj2["inner"])
}
