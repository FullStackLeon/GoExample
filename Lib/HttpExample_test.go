package Lib

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestHttpClient(t *testing.T) {
	client := http.Client{}
	// 1.Get example
	getUrl := "https://images.pexels.com/photos/28874283/pexels-photo-28874283.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=2"
	res, err := client.Get(getUrl)
	if err != nil {
		log.Fatalf("get err:%v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("get response status code: %d,status: %s", res.StatusCode, res.Status)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("read response body error:%v", err)
	}
	imageFile, err := os.Create("download_image.jpg")
	if err != nil {
		log.Fatalf("create image file error:%v", err)
	}
	defer imageFile.Close()

	_, err = imageFile.Write(body)
	if err != nil {
		log.Fatalf("write response body error:%v", err)
	}
	log.Printf("image save as download_image.jpg")

	// 2. Head example
	res, err = client.Head(getUrl)
	if err != nil {
		log.Fatalf("head err:%v", err)
	}
	defer res.Body.Close()

	log.Printf("head response status code: %d", res.StatusCode)
	log.Printf("head response header: %#v", res.Header)

	// 3. Post example
	postUrl := "https://b2bc31c0cca34e119a999521872da175.api.mockbin.io/"
	res, err = client.Post(postUrl, "application/json", strings.NewReader(""))
	if err != nil {
		log.Fatalf("post err:%v", err)
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("read post response body error:%v", err)
	}
	log.Printf("post response status code: %d", res.StatusCode)
	log.Printf("post response: %s", string(body))

	// 4. PostForm example
	postFormUrl := "https://postman-echo.com/post"
	form := url.Values{
		"foo": {"bar"},
		"baz": {"qux"},
	}
	res, err = client.PostForm(postFormUrl, form)
	if err != nil {
		log.Fatalf("post form err: %v", err)
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("read post form response body error:%v", err)
	}
	log.Printf("post form status code: %d", res.StatusCode)
	log.Printf("post form response: %s", string(body))
}
