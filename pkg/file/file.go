package file

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

// ReadFile 读取文件内容
func ReadFile(filePath string) (string, error) {
	fmt.Println(os.Getwd())
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

// ReadTemplateToString 读取模板并转换为 string
func ReadTemplateToString(tplName string, tplPath string, tplData map[string]interface{}) (string, error) {
	t, err := template.New(tplName).ParseFiles(tplPath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, tplData); err != nil {
		return "", err
	}

	return buf.String(), nil
}
