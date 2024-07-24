package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"

	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
)

const url = "https://dog.ceo/api/breeds/image/random"

type dogy struct {
	Url string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	GetDog(w, r, url)
}

func FindDog(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var dog dogy
	if err := json.Unmarshal(data, &dog); err != nil {
		log.Fatal(err)
	}
	return dog.Url
}

func GetDog(w http.ResponseWriter, r *http.Request, url string) {
	dogurl := FindDog(url)

	resp, err := http.Get(dogurl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	errr := jpeg.Encode(w, img, nil)
	if errr != nil {
		log.Fatal(errr)
	}
}

func main() {
	// Регистрируем обработчик для пути "/"

	http.HandleFunc("/getdog", handler)

	// Запускаем веб-сервер на порту 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}

}
