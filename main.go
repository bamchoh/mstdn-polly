package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bamchoh/pollydent"

	"gopkg.in/yaml.v2"
)

type AwsCredential struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}

func Load(filename string) (*AwsCredential, error) {
	var err error
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data []byte
	data, err = ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var aws AwsCredential
	err = yaml.Unmarshal(data, &aws)
	if err != nil {
		return nil, err
	}

	return &aws, err
}

type message struct {
	ID       uint
	UserName string `json:"username"`
	Acct     string `json:"acct"`
	Content  string `json:"content"`
}

func setLog(prefix string) *log.Logger {
	basedir := filepath.Dir(os.Args[0])
	logPath := filepath.Join(basedir, prefix+".log")
	f, _ := os.Create(logPath)
	return log.New(f, prefix+":", 0)
}

func main() {
	logger := setLog("mstdn")

	ac, err := Load("mstdn-polly.yml")
	if err != nil {
		logger.Println("Load error")
		logger.Println(err)
	}

	p := pollydent.NewPollydent(
		ac.AccessKey,
		ac.SecretKey,
		nil,
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(txt)
		dec := json.NewDecoder(strings.NewReader(txt))
		msg := new(message)
		err := dec.Decode(msg)
		if err != nil {
			logger.Println(err)
		}
		reg := regexp.MustCompile("https?://.*")
		readContent := reg.ReplaceAllString(msg.Content, "")
		go p.ReadAloud(readContent)
	}

	if err := scanner.Err(); err != nil {
		logger.Println("reading standard input:", err)
	}
}
