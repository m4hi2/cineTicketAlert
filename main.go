package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type LoginPayload struct {
	UserId string `json:"user_id"`
}
type LoginResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func main() {

	login := LoginPayload{UserId: "gg"}

	lPld, err := json.Marshal(login)

	if err != nil {
		log.Println("Error: Can not marshal login payload: ", err)
		return
	}
	request, err := http.NewRequest(http.MethodPost, "https://cineplex-web-api.cineplexbd.com/api/v1/login", bytes.NewBuffer(lPld))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println("Error: Can not create request: ", err)
		return
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("Error: Can not perform HTTP request: ", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error: ", err)
		}
	}(response.Body)

	rb, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	lRes := &LoginResponse{}

	err = json.Unmarshal(rb, lRes)
	if err != nil {
		log.Println("Error: in unmarshalling login response: ", err)
		return
	}
	if lRes.Code != 200 {
		log.Println("Error: in logging in.")
		return
	}

	authToken := lRes.Data

	location := bytes.NewBuffer([]byte(`{"location":3}`))

	mRequest, err := http.NewRequest(http.MethodPost, "https://cineplex-web-api.cineplexbd.com/api/v1/movie-list", location)
	if err != nil {
		return
	}
	mRequest.Header.Set("Content-Type", "application/json")
	mRequest.Header.Set("Authorization", "Bearer "+authToken)
	mRes, err := http.DefaultClient.Do(mRequest)
	if err != nil {
		return
	}
	movies, err := io.ReadAll(mRes.Body)
	if err != nil {
		return
	}
	log.Println(string(movies))

}
