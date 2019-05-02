package checker

import (
	"log"
	"net/http"

	"fmt"
	"time"
)

// CheckURL check something url,
func CheckURL(url string) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)

}
