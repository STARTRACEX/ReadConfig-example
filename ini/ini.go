package ini

import (
	"fmt"
	"readconfig/types"
	"gopkg.in/ini.v1"
)

func ReadIni(path string) types.Initype {
	var config types.Initype
	ini.MapTo(&config, path)
	fmt.Println("Author:"+config.Author, "Title:", config.Title, config.Obj.Inner)
	return config
}
