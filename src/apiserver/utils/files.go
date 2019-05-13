package utils

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
)

func CheckDir(targetPath string) error {
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		err = os.MkdirAll(targetPath, 0755)
		if err != nil {
			beego.Error(fmt.Sprintf("Failed to create target dir: %s, with err: %+v", targetPath, err))
			return err
		}
	}
	return nil
}
