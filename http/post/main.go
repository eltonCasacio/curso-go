package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"
)

type Content struct {
	name     string
	duration int16
	finished bool
}

func main() {
	content := Content{
		name:     "Go",
		duration: 120,
		finished: false,
	}

	jsonContent, err := json.Marshal(content)
	ifError(err)

	bufferedContent := bytes.NewBuffer([]byte(jsonContent))

	c := http.Client{Timeout: time.Second}
	res, err := c.Post("http://google.com", "application/json", bufferedContent)
	ifError(err)
	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}

func ifError(err error) {
	if err != nil {
		panic(err)
	}
}
