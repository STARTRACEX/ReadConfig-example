package glob

import (
	"fmt"
	"readconfig/json"
	"readconfig/toml"
	"readconfig/types"
	"readconfig/yaml"
)

func ReadGlob(path string) types.Maintype {
	var config types.Maintype
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
					config = toml.ReadToml(path)
					return
				}
			}()
			config = yaml.ReadYaml(path)
			return
		}
	}()
	config = json.ReadJson(path)
	return config
}
