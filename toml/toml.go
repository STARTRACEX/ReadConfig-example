package toml

import (
	"fmt"
	"os"
	"readconfig/types"
	"github.com/BurntSushi/toml"
)

func ReadToml(path string) types.Maintype{
	fileread, _ := os.ReadFile(path)
	var config types.Maintype
	e:=toml.Unmarshal(fileread, &config)
	// toml.Decode(string(fileread), &config)
	if e!=nil{
		panic("toml.Unmarshal failed")
	}
	fmt.Println("Author:"+config.Author, "Title:", config.Title, config.Obj.Inner, config.Obj2["inner"])
	return config
}
