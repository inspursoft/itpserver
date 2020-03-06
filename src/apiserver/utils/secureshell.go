package utils

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/astaxie/beego"
	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
)

const maxSSHRetries = 10
const maxSSHDelay = 50 * time.Microsecond

type SecureShell struct {
	client *ssh.Client
	output io.Writer
}

func NewSecureShell(output io.Writer) (*SecureShell, error) {
	host := beego.AppConfig.String("ssh::host")
	port, _ := beego.AppConfig.Int("ssh::port")
	username := beego.AppConfig.String("ssh::username")
	password := beego.AppConfig.String("ssh::password")
	// Retry few times if ssh connection fails
	for i := 0; i < maxSSHRetries; i++ {
		client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		})
		if err != nil {
			time.Sleep(maxSSHDelay)
			log.Printf("Failed to dial host: %+v\n", err)
			continue
		}
		s, err := client.NewSession()
		if err != nil {
			client.Close()
			time.Sleep(maxSSHDelay)
			continue
		}
		modes := ssh.TerminalModes{
			ssh.ECHO:          0,
			ssh.TTY_OP_ISPEED: 14400,
			ssh.TTY_OP_OSPEED: 14400,
		}
		// Request pseudo terminal
		if err := s.RequestPty("xterm", 40, 80, modes); err != nil {
			return nil, fmt.Errorf("failed to get pseudo-terminal: %v", err)
		}
		return &SecureShell{client: client, output: output}, nil
	}
	return nil, nil
}

func (s *SecureShell) ExecuteCommand(cmd string) error {
	session, err := s.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	stdout, _ := session.StdoutPipe()
	if err != nil {
		return err
	}
	log.Printf("Execute command: %s\n", cmd)
	err = session.Start(cmd)
	if err != nil {
		return err
	}
	go io.Copy(s.output, stdout)
	return session.Wait()
}

func (s *SecureShell) SecureCopyData(fileName string, data []byte, destinationPath string) error {
	session, err := s.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	var buf bytes.Buffer
	length, err := buf.Write(data)
	if err != nil {
		log.Printf("Failed to load contents: %+v\n", err)
		return err
	}
	return scp.Copy(int64(length), 0755, fileName, &buf, filepath.Join(destinationPath, fileName), session)
}

func (s *SecureShell) SecureCopy(filePath string, destinationPath string) error {
	return filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		session, err := s.client.NewSession()
		if err != nil {
			return err
		}
		defer session.Close()
		if err != nil {
			return err
		}
		if info.IsDir() {
			log.Printf("From path: %s to path: %s\n", path, destinationPath)
			return nil
		}
		log.Printf("From path: %s to path: %s\n", path, filepath.Join(destinationPath, info.Name()))
		return scp.CopyPath(path, filepath.Join(destinationPath, info.Name()), session)
	})
}

func (s *SecureShell) CheckDir(dir string) error {
	return s.ExecuteCommand(fmt.Sprintf("mkdir -p %s", dir))
}