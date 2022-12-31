package yaml

import (
	"fmt"
	"os"
	"readconfig/types"
	"regexp"
	"strings"
	"github.com/goccy/go-yaml"
)

func ReadYamlWithMD(path string) {
	fileread, _ := os.ReadFile(path)
	str := string(fileread)
	data, l := getMeta(str)
	fmt.Println("Meta date:", data)
	fmt.Println("Meta date Length:", l)
	fmt.Println("File content:", string(fileread[l:]))
	var config types.Maintype
	yaml.Unmarshal([]byte(data), &config)
	if config.Title == "" {
		config.Title = getTitie(string(fileread[l:]))
	}
	fmt.Println("config.Author:"+config.Author, "config.Title:", config.Title)
}

func getMeta(data string) (string, int) {
	re := regexp.MustCompile(`---([\s\S]*?\n)---`)
	return re.FindString(data), len(re.FindString(data))
}

func getTitie(data string) string {
	re := regexp.MustCompile(`(#)[^\n]*?\n`)
	return strings.ReplaceAll(re.FindString(data), "# ", "")
}
