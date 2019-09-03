package services

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/astaxie/beego"
)

var pathPrefix = beego.AppConfig.String("pathprefix")
var workpath = path.Join(pathPrefix, beego.AppConfig.String("vagrant::baseworkpath"))
var outputPath = path.Join(pathPrefix, beego.AppConfig.String("vagrant::outputpath"))
var artifactsURL = beego.AppConfig.String("nexus::url")
var nexusUsername = beego.AppConfig.String("nexus:username")
var nexusPassword = beego.AppConfig.String("nexus:password")

func RetrieveVMFiles(vmName string) []string {
	results := []string{}
	filepath.Walk(filepath.Join(workpath, vmName), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			beego.Error(fmt.Sprintf("Failed to retrieve VM files: %+v", err))
			return err
		}
		if info.IsDir() {
			return nil
		}
		results = append(results, filepath.Base(path))
		return nil
	})
	return results
}

func ResolveBoxFilePath(vmName string) string {
	boxFilePath := filepath.Join(outputPath, vmName+".box")
	beego.Debug(fmt.Sprintf("Get VM box download file path: %s", boxFilePath))
	return boxFilePath
}

func UploadArtifacts(vmName, repoName, principle string) error {
	boxFilepath := ResolveBoxFilePath(vmName)
	file, err := os.Open(boxFilepath)
	if err != nil {
		return err
	}
	defer file.Close()
	fileName := filepath.Base(boxFilepath)
	artifactsURL := path.Join(artifactsURL, repoName, principle, fileName)
	beego.Debug(fmt.Sprintf("Upload artifacts URL: %s", artifactsURL))
	req, err := http.NewRequest(http.MethodPost, artifactsURL, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(nexusUsername, nexusPassword)
	client := http.Client{}
	resp, err := client.Do(req)
	if resp != nil && resp.StatusCode >= 400 {
		err = fmt.Errorf("failed to upload artifacts with status code: %d", resp.StatusCode)
		beego.Error(err)
		return err
	}
	beego.Debug(fmt.Sprintf("Successful uploaded file: %s to artifacts with URL: %s", fileName, artifactsURL))
	return nil
}
