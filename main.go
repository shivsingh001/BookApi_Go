package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("running a server")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Getting error while fetching data from API", err)
		return
	}

	defer res.Body.Close()
	fmt.Println("Response received from API", res)

	//read the data from res.Body

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading data from response body", err)
		return
	}

	fmt.Println("Data from API is", string(data))

}
