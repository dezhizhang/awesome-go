package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")

	if err != nil {
		panic(err)
	}

	response, _ := client.Do(request)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", response.StatusCode)
	}

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	return result, err
}
