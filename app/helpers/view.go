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

// Static 生成项目静态文件地址
func Static(staticFilePath string) string {
	return "/" + config.AppConfig.PublicPath + staticFilePath
}

// Mix 根据 laravel-mix 生成静态文件 path
func Mix(staticFilePath string) string {
	result := manifests[staticFilePath]

	if result == "" {
		filename := path.Join(config.AppConfig.PublicPath, "mix-manifest.json")
		file, err := os.Open(filename)
		if err != nil {
			return Static(staticFilePath)
		}
		defer file.Close()

		dec := json.NewDecoder(file)
		if err := dec.Decode(&manifests); err != nil {
			return Static(staticFilePath)
		}

		result = manifests[staticFilePath]
	}

	if result == "" {
		return Static(staticFilePath)
	}

	return Static(result)
}
