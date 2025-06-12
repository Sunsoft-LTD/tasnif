package tasnif

import (
	"encoding/json"

	"resty.dev/v3"
)

func New(lan Language) *Tasnif {
	return &Tasnif{
		BaseURL:  API_URL,
		Language: lan,
	}
}

func (t *Tasnif) Search(search string) ([]*Product, error) {
	c := resty.New()
	defer c.Close()
	c.SetBaseURL(t.BaseURL)
	c.SetHeader("Content-Type", "application/json")
	c.SetHeader("Accept", "application/json")
	resp, err := c.R().SetQueryParams(map[string]string{
		"lang":   string(t.Language),
		"search": search,
		"page":   "0",
		"size":   "20",
	}).Get("/elasticsearch/search")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var result *Response
	if err = json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}
