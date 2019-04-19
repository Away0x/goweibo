package helpers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path"

	"gin_weibo/config"
)

var (
	// 存储 mix-manifest.json 解析出来的 path map
	manifests = make(map[string]string)
)

// Mix 根据 laravel-mix 生成静态文件 path
func Mix(staticFilePath string) string {
	result := manifests[staticFilePath]

	if result == "" {
		filename := path.Join(config.ProjectConfig.PublicPath, "mix-manifest.json")
		file, err := os.Open(filename)
		if err != nil {
			return staticFilePath
		}
		defer file.Close()

		dec := json.NewDecoder(file)
		if err := dec.Decode(&manifests); err != nil {
			return staticFilePath
		}

		result = manifests[staticFilePath]
	}

	if result == "" {
		return staticFilePath
	}

	return "/" + config.ProjectConfig.PublicPath + result
}

// CsrfField csrf input
func CsrfField() template.HTML {
	return template.HTML(fmt.Sprintf(`<input type="hidden" name="_token" value="%s">`, "asd"))
}

// func HasSession(key string) bool {

// }

// func GetSession(key string) string {

// }

// func Old(key string, defaultValue string) string {

// }
