package unsplash

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UnsplashPhoto struct {
	Urls struct {
		Regular string
	}
}

func GetURL(accessKey string) string {
	query := "nature"
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?query=%s&client_id=%s", query, accessKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var photo UnsplashPhoto
	err = json.NewDecoder(resp.Body).Decode(&photo)
	if err != nil {
		panic(err)
	}
	imageUrl := photo.Urls.Regular
	return imageUrl

}
func GetPhoto(accessKey string) ([]byte, error) {
	imageUrl := GetURL(accessKey)
	resp, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	imageData, err := ioutil.ReadAll(resp.Body)
	return imageData, err
}
