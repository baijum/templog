package templog

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dir string

// Println log the message (msg.txt)
func Println(msg ...string) {
	tmpfn := filepath.Join(dir, "msg.txt")
	content := strings.Join(msg, " ") + "\n"
	f, err := os.OpenFile(tmpfn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		log.Fatal(err)
	}
}

// Println2 log the message (msg2.txt)
func Println2(msg ...string) {
	tmpfn := filepath.Join(dir, "msg2.txt")
	content := strings.Join(msg, " ") + "\n"
	f, err := os.OpenFile(tmpfn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		log.Fatal(err)
	}
}

// Println3 log the message (msg3.txt)
func Println3(msg ...string) {
	tmpfn := filepath.Join(dir, "msg3.txt")
	content := strings.Join(msg, " ") + "\n"
	f, err := os.OpenFile(tmpfn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		log.Fatal(err)
	}
}

func init() {
	var err error
	dir, err = ioutil.TempDir("", "tmplog")
	if err != nil {
		log.Fatal(err)
	}
}
