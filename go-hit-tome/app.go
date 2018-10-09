package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HeaderCustom struct {
	Authorization string `json:"Authorization"`
	ContentType   string `json:"Content-Type"`
	Origin        string `json:"Origin"`
	Referer       string `json:"Referer"`
	XDevice       string `json:"X-Device"`
}

type Product struct {
	ProductID    int `json:"product_id"`
	ProductStock int `json:"product_stock"`
	ProductPrice int `json:"product_price"`
}

type BodyCustom struct {
	ProductList []Product `json:"product_list"`
	Sender      string    `json:"sender"`
}

func main() {
	//add client api
	client := http.Client{}

	fmt.Println("access tome API")
	method := "PATCH"
	endpointTome := "https://tome-staging.tokopedia.com/v2/batch/product/"
	h := HeaderCustom{
		Authorization: "Bearer HsuEIJ97SoKwKn2fNNsgrQ",
		ContentType:   "application/json",
		Origin:        "https://www.tokopedia.com",
		Referer:       "https://www.tokopedia.com",
		XDevice:       "IOS",
	}

	// p := Product{ProductID: 15223032, ProductStock: 250, ProductPrice: 750000}
	body := BodyCustom{
		Sender: "fulfillment_service",
		ProductList: []Product{
			Product{ProductID: 15223032, ProductStock: 250, ProductPrice: 750000},
			Product{ProductID: 15223032, ProductStock: 250, ProductPrice: 750000},
		},
	}

	tbody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("err: ", err)
	}
	req, err := http.NewRequest(method, endpointTome, bytes.NewReader([]byte(tbody)))
	if err != nil {
		fmt.Println("has error", err)
	}
	req.Header.Add("Authorization", h.Authorization)
	req.Header.Add("Content-Type", h.ContentType)
	req.Header.Add("Origin", h.Origin)
	req.Header.Add("Referer", h.Referer)
	req.Header.Add("X-Device", h.XDevice)
	fmt.Println("Header", req.Header)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf(string(data))
	}
}
