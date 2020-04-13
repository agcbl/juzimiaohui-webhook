package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Transport struct {
	Endpoint string
	Token string
}

func (p *Transport) Get(path string, param *map[string]string) ([]byte, error) {
	baseUrl := fmt.Sprintf("%s%s", p.Endpoint, path)
	query := url.Values{}
	query.Add("token", p.Token)
	for k, v := range *param {
		query.Add(k, v)
	}

	queryParam := query.Encode()
	log.Printf("%s?%s\n", baseUrl, queryParam)
	resp, err := http.Get(fmt.Sprintf("%s?%s", baseUrl, queryParam))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *Transport) Post(path string, param map[string]interface{}) ([]byte, error) {
	baseUrl := fmt.Sprintf("%s%s", p.Endpoint, path)
	param["token"] = p.Token
	request, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	log.Printf("%s %v\n", baseUrl, param)
	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}