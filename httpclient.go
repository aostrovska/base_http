package main

import (
	"net/http"
	//"io"
	"io/ioutil"
	"bytes"
)

func main() {
	client := &http.Client{}
	
	var body bytes.Buffer
	body.Write([]byte("Hello server!"))
	
	req, err := http.NewRequest("GET", "http://localhost:8080/", &body)
	
	resp, err := client.Do(req)
	if err != nil {panic(err) }
	
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	
	err = ioutil.WriteFile("p.html", data, 0666)
	if err != nil {panic(err) }
	
	
}
