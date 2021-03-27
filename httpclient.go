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
	body.Write([]byte("{\n\"action\":\"create\",\n\"object\":\"Teacher\",\n\"data\":{\n\"id\":\"001\",\n\"subject\":\"Math\",\n\"salary\":2345,\n\"classroom\":\"CL-001\",\n\"person\":{\n\"name\":\"Ivan\",\n\"surname\":\"Popov\",\n\"personalCode\":\"123422-43235\"\n}\n}\n}"))
	
	req, err := http.NewRequest("POST", "http://localhost:8080/", &body)
	resp, err := client.Do(req)
	if err != nil {panic(err) }
	data, err := ioutil.ReadAll(resp.Body)
	req, err = http.NewRequest("GET", "http://localhost:8080/", &body)
	resp, err = client.Do(req)
	if err != nil {panic(err) }
	data, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	
	err = ioutil.WriteFile("p.html", data, 0666)
	if err != nil {panic(err) }
	
	
}
