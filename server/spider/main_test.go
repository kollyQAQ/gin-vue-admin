package spider

import (
	"fmt"
	"log"
	"testing"

	"github.com/levigross/grequests"
)

func get() {
	resp, err := grequests.Get("https://api.github.com/events", nil)
	// You can modify the request by passing an optional RequestOptions struct

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	fmt.Println(resp.String())
}

func TestGet(t *testing.T) {
	get()
}
