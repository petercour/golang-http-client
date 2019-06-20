package main
 
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
 
func main() {
 
	// HTTP GET request
	resp, err := http.Get("https://golangr.com")
	if err != nil {
		log.Fatal("Error getting response. ", err)
	}
	defer resp.Body.Close()
 
	// Read body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
 
	fmt.Printf("%s\n", body)
}