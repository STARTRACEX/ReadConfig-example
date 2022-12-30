package main

import (
	"readconfig/ini"
	"readconfig/json"
	"readconfig/toml"
	"readconfig/yaml"
)

func main() {
	// yaml.ReadYamlWithMD("./yaml/.md")
	yaml.ReadYaml("./yaml/.yml")
	toml.ReadToml("./toml/.toml")
	json.ReadJson("./json/.json")
	ini.ReadIni("./ini/.ini")
}
