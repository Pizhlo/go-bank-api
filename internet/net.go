package internet

import (
	"io"
	"log"
	"net/http"
)

// MakeRequest отправляет GET-запрос на сервер по указанному url
func MakeRequest(url string) io.ReadCloser {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("не удалось создать запрос:", err)
		return nil
	}

	req.Header = http.Header{
		"Host":       {"www.cbr.com"},
		"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 OPR/97.0.0.0"},
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("не удалось совершить запрос:", err)
	}

	return resp.Body
}
