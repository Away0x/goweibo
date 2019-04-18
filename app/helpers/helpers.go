package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"gin_weibo/config"
)

// Mix 根据 laravel-mix 生成静态文件 path
func Mix(staticFilePath string) string {
	manifests := make(map[string]string)

	filename := path.Join(config.ProjectConfig.PublicPath, "mix-manifest.json")
	file, err := os.Open(filename)
	if err != nil {
		// log.Fatalf("mix-manifest.json load fail: %v", err)
		return staticFilePath
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if err := dec.Decode(&manifests); err != nil {
		// log.Fatalf("mix-manifest.json decode fail: %v", err)
		return staticFilePath
	}

	// for k, v := range manifests {
	// 	log.Printf("%#v %#v", k, v)
	// }

	result := manifests[staticFilePath]
	if result == "" {
		return staticFilePath
	}

	return "/" + config.ProjectConfig.PublicPath + result
}

// FormatAsDate 格式化日期
func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}
