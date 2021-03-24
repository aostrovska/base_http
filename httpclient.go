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
	body.Write([]byte("\"action\":\"create\",\"object\":\"Teacher\",\"data\":{\"id\":\"001\",\"subject\":\"Math\",\"salary\":2345,\"classroom\":\"CL-001\",\"person\":{\"name\":\"Ivan\",\"surname\":\"Popov\",\"personalCode\":\"123422-43235\"}"))
	
	req, err := http.NewRequest("POST", "http://localhost:8080/", &body)
	req, err = http.NewRequest("GET", "http://localhost:8080/", &body)
	resp, err := client.Do(req)
	if err != nil {panic(err) }
	
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	
	err = ioutil.WriteFile("p.html", data, 0666)
	if err != nil {panic(err) }
	
	
}
