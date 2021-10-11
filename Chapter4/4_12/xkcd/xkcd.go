package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

const XkcdURL = "https://xkcd.com/"

func (comic *Comic) Match(keyword string) bool {
	return comic.Month == keyword ||
		strconv.Itoa(comic.Num) == keyword ||
		strings.Contains(comic.Link, keyword) ||
		comic.Year == keyword ||
		strings.Contains(comic.News, keyword) ||
		strings.Contains(comic.SafeTitle, keyword) ||
		strings.Contains(comic.Transcript, keyword) ||
		strings.Contains(comic.Alt, keyword) ||
		strings.Contains(comic.Img, keyword) ||
		strings.Contains(comic.Title, keyword)
}

func FetchComic(id int) (*Comic, error) {
	reqURL := XkcdURL + strconv.Itoa(id) + "/info.0.json"
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch Comic %d Error", id)
	}
	var comic Comic
	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return nil, err
	}
	return &comic, nil
}
