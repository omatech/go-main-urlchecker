package main

import (
	"fmt"
	"testing"
	"time"

	"omatech.com/urlchecker"
)

func TestMain(t *testing.T) {
	url := "/2010/HEM/01/09/file.pdf"
	timestamp := time.Now().Unix()
	token := urlchecker.GenerateToken(url, timestamp)
	message := urlchecker.Debug(url, timestamp, token)
	fmt.Println(message)
	is_ok, err := urlchecker.Check(url, timestamp, token)
	if !is_ok {
		t.Fatalf("Error, this test should not have error but it gets %v\n", err)
	} else {
		fmt.Println("Todo ok")
	}
}
