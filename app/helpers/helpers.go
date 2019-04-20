package helpers

import (
	"encoding/json"
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

// func HasSession(key string) bool {

// }

// func GetSession(key string) string {

// }

// func Old(key string, defaultValue string) string {

// }
