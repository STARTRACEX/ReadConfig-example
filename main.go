package main

import (
	"readconfig/glob"
	"readconfig/ini"
	"readconfig/json"
	"readconfig/toml"
	"readconfig/yaml"
)

func main() {
	// yaml.ReadYamlWithMD("./yaml/.md")
	json.ReadJson("./json/.json")
	yaml.ReadYaml("./yaml/.yml")
	toml.ReadToml("./toml/.toml")
	ini.ReadIni("./ini/.ini")
	glob.ReadGlob("./json/.json")
	glob.ReadGlob("./yaml/.yml")
	glob.ReadGlob("./toml/.toml")
}
