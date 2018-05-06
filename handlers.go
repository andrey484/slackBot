package main

import (
	"net/http"
	"io/ioutil"
)

func slackWebHook(response http.ResponseWriter, request *http.Request){
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil{
		panic(err)
	}
	response.Write(body)
}
