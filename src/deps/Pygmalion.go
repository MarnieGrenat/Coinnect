package Pygmalion

import (
	"fmt"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

var ServiceName string
var SettingsName string
var SettingsPath string

func InitConfigReader(settingsName string, settingsPath string) {
	SettingsName = settingsName
	SettingsPath = settingsPath
	ReadString("ServiceName")
}

func readYaml() map[string]interface{} {
	var path string = path.Join(SettingsPath, SettingsName)
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Pygmalion.readYaml : Failed to Read : Error=%v\n", err)
	}
	obj := make(map[string]interface{})
	err = yaml.Unmarshal(yamlFile, obj)

	if err != nil {
		fmt.Printf("Pygmalion.readYaml : Failed to decode file : Error=%v\n", err)
	}

	return obj
}

func ReadString(key string) string {
	settings := readYaml()

	if value, err := settings[key].(string); err {
		return value
	}
	fmt.Println("Pygmalion.ReadString : Key is not an String : Key=" + key)
	return ""
}

func ReadInteger(key string) int {
	settings := readYaml()
	if value, err := settings[key].(int); err {
		return value
	}
	fmt.Println("Pygmalion.ReadInteger : Key is not an integer : Key=" + key)
	return 0
}

func ReadBoolean(key string) bool {
	settings := readYaml()
	if value, err := settings[key].(bool); err {
		return value
	}
	fmt.Println("Pygmalion.ReadBoolean : Key is not an Boolean : Key=" + key)
	return false
}

func ReadList(key string) []string {
	var content string = ReadString(key)
	return strings.Split(content, ",")
}
