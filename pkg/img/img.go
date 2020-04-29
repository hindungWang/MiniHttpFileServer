package img

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

type Img struct {
	Path string
}

const (
	Url = "https://sm.ms/api/upload"
)

func (i *Img) Upload() (by []byte, err error) {
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)

	f, err := os.Open(i.Path)
	if err != nil {
		log.Error("error opening file")
		return nil, err
	}
	defer f.Close()

	fw, err := w.CreateFormFile("smfile", i.Path)
	if err != nil {
		log.Error("error writing to buffer")
		return nil, err
	}

	if _, err = io.Copy(fw, f); err != nil {
		return nil, err
	}

	contenType := w.FormDataContentType()
	w.Close()

	resp, err := http.Post(Url, contenType, b)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil

	/*if fw, err = w.CreateFormField("key"); err != nil {
		return nil, err
	}
	if _, err = fw.Write([]byte("KEY")); err != nil {
		return nil, err
	}
	w.Close()
	req, err := http.NewRequest("POST", Url, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return body, nil*/
}
