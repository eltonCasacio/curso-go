package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Microsecond} // vai dar panic porque a resposta do google vai demorar mais que o esperado
	// c := http.Client{Timeout: time.Second}
	res, err := c.Get("http://google.com")
	ifError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	ifError(err)
	println(string(body))

}

func ifError(err error) {
	if err != nil {
		panic(err)
	}
}
