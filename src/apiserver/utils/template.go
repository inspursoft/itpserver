package utils

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/astaxie/beego"
)

func ExecuteTemplate(data interface{}, templateName string) ([]byte, error) {
	templatePath := beego.AppConfig.String("templatepath")
	t, err := template.ParseFiles(filepath.Join(templatePath, templateName))
	if err != nil {
		beego.Error(fmt.Sprintf("Failed to parse template file: %s, with err: %+v", filepath.Join(templatePath, templateName), err))
		return nil, err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	return buf.Bytes(), err
}
