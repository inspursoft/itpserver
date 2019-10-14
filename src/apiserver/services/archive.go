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

	"github.com/inspursoft/itpserver/src/apiserver/utils"

	"github.com/inspursoft/itpserver/src/apiserver/models"

	"github.com/astaxie/beego"
)

var pathPrefix = beego.AppConfig.String("pathprefix")
var workpath = path.Join(pathPrefix, beego.AppConfig.String("vagrant::baseworkpath"))
var outputPath = path.Join(pathPrefix, beego.AppConfig.String("vagrant::outputpath"))
var artifactsURL = beego.AppConfig.String("nexus::url")
var nexusUsername = beego.AppConfig.String("nexus::username")
var nexusPassword = beego.AppConfig.String("nexus::password")
var sshUsername = beego.AppConfig.String("ssh::username")
var sshPassword = beego.AppConfig.String("ssh::password")
var sshHost = beego.AppConfig.String("ssh::host")
var sshPort = beego.AppConfig.String("ssh::port")

func RetrieveVMFiles(vmName string) ([]string, *models.ITPError) {
	e := &models.ITPError{}
	targetPath := filepath.Join(workpath, vmName)
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		e.Notfound(targetPath, err)
		return nil, e
	}
	results := []string{}
	filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			e.InternalError(fmt.Errorf("Failed to retrieve VM files: %+v", err))
			return err
		}
		relPath, _ := filepath.Rel(filepath.FromSlash(targetPath), path)
		if !(relPath[0:1] == "." || info.IsDir()) {
			results = append(results, relPath)
		}
		return nil
	})
	return results, e
}

func ResolveBoxFilePath(vmName string) string {
	boxFilePath := filepath.Join(outputPath, vmName+".box")
	beego.Debug(fmt.Sprintf("Get VM box download file path: %s", boxFilePath))
	return boxFilePath
}

func UploadArtifacts(vmName, repoName, principle string, output io.Writer) error {
	sshClient, err := utils.NewSecureShell(output)
	if err != nil {
		return err
	}
	boxFilepath := ResolveBoxFilePath(vmName)
	scpCommand := fmt.Sprintf("scp -P %s %s@%s:%s %s", sshPort, sshUsername, sshHost, boxFilepath, outputPath)
	beego.Debug(fmt.Sprintf("SCP command: %s", scpCommand))
	err = sshClient.ExecuteCommand(scpCommand)
	if err != nil {
		return err
	}
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
	uploadURL := artifactsURL + "/" + path.Join(repoName, principle, fileName)
	req, err := http.NewRequest(http.MethodPut, uploadURL, bodyBuf)
	if err != nil {
		return err
	}
	beego.Debug(fmt.Sprintf("Upload artifacts URL: %s", uploadURL))
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
