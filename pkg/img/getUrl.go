package img

import "encoding/json"

type data struct {
	Url string `json:"url"`
}
type result struct {
	Code string `json:"code"`
	Data data `json:"data"`
}

func GetUrl(b []byte) (url string, err error) {
	res := &result{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return "", err
	}
	return res.Data.Url, nil
}
