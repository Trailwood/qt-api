package qtapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Questrade struct {
	apiserver string
	token     string
}

func (q Questrade) request(url string) ([]byte, error) {
	client := &http.Client{}
	full_url := fmt.Sprintf("%s%s", q.apiserver, url)

	req, err := http.NewRequest("GET", full_url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", q.token))
	req.Header.Add("Content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		return body, nil
	}
	req.Header.Del("Authorization")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", q.token))
	resp, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
