package toml

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

func ReadToml(path string) {
	fileread, _ := os.ReadFile(path)
	var config basetype
	toml.Unmarshal(fileread, &config)
	// toml.Decode(string(fileread), &config)
	fmt.Println("Author:"+config.Author, "Title:", config.Title, config.Obj.Inner, config.Obj2["inner"])

}
