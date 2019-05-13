package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
)

const templatePath = "../../templates"

func ExecuteTemplate(data interface{}, templateName string, targetPath string) error {
	t, err := template.ParseFiles(filepath.Join(templatePath, templateName))
	if err != nil {
		beego.Error(fmt.Sprintf("Failed to parse template file: %s, with err: %+v", filepath.Join(templatePath, templateName), err))
		return err
	}
	f, err := os.OpenFile(filepath.Join(targetPath, templateName), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, data)
}
