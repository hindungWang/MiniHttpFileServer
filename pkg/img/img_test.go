package img

import (
	"testing"
	"fmt"
)

func TestUpload(t *testing.T)  {
	i := Img{
		Path: "C://Users/Administrator/Desktop/1.jpg",
	}
	b ,err := i.Upload()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	fmt.Print(len(string(b)))
}
