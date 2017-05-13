package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/bamchoh/pollydent"
)

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

	p, err := pollydent.NewPolly(logger, "mstdn-polly.yml")
	if err != nil {
		logger.Println(err)
	}

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
