package img

import (
	"fmt"
	"log"
	"testing"
)

func TestGetUrl(t *testing.T) {
	i := Img{
		Path: "C://Users/Administrator/Desktop/1.jpg",
	}
	b, err := i.Upload()
	if err != nil {
		fmt.Println(err)
	}
	if len(b) < 200 {
		fmt.Print(string(b))
	}
	url, err := GetUrl(b)
	if err != nil {
		log.Error(err)
	}
	fmt.Print(url)
}
