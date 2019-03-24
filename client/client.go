package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Request() []byte {

	client := &http.Client{}
	url := "http://www.metro.sp.gov.br/Sistemas/direto-do-metro-via4/diretoDoMetroMobile.aspx"
	void := []byte("olamundo")

	request, err := http.NewRequest("GET", url, bytes.NewBuffer(void))
	if err != nil {
		log.Println("Error to access the site, got ", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", "Firefox")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error to get the body, got: ", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body
}
