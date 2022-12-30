package ini

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func ReadIni(path string) {
	var config basetype
	ini.MapTo(&config, path)
	fmt.Println("Author:"+config.Author, "Title:", config.Title, config.Obj.Inner)
}
