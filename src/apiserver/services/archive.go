package services

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
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
var nexusUsername = beego.AppConfig.String("nexus::username")
var nexusPassword = beego.AppConfig.String("nexus::password")

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
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("upload-file", fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return err
	}
	defer bodyWriter.Close()
	artifactsURL += "/" + path.Join(repoName, principle, fileName)
	req, err := http.NewRequest(http.MethodPut, artifactsURL, bodyBuf)
	if err != nil {
		return err
	}
	beego.Debug(fmt.Sprintf("Upload artifacts URL: %s", artifactsURL))
	req.SetBasicAuth(nexusUsername, nexusPassword)
	client := &http.Client{}
	resp, err := client.Do(req)
	if resp != nil && resp.StatusCode >= 400 {
		err = fmt.Errorf("failed to request URL: %s with status code: %d", artifactsURL, resp.StatusCode)
		beego.Debug(err.Error())
		return err
	}
	beego.Debug(fmt.Sprintf("Successful uploaded file: %s to artifacts with URL: %s", fileName, artifactsURL))
	return nil
}
