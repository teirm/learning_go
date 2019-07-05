package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const XkcdURL = "https://xkcd.com"

type Comic struct {
	Month      string
	Number     int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Title      string
	Day        string
}

func GetComic(index int) (*Comic, error) {
	comicUrl := fmt.Sprintf("%s/%d/info.0.json", XkcdURL, index)
	resp, err := http.Get(comicUrl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}

	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func GetComicRange(start, end int) []*Comic {

	var comics []*Comic

	for i := start; i < end; i++ {
		comic, err := GetComic(i)
		if err != nil {
			panic(err)
		}
		comics = append(comics, comic)
	}

	return comics
}

func PrintComics(comics []*Comic) {

	for index, comic := range comics {
		fmt.Printf("Index: %d\nTitle: %s\n", index, comic.Title)
	}
}

func WriteFile(fileName string, data []byte) {

}

func main() {
	start := 10
	end := 13

	comics := GetComicRange(start, end)

	data, err := json.Marshal(comics)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	if err := json.Unmarshal(data, &comics); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s\n", err)
	}

	fmt.Printf("comics length: %d\n", len(comics))

	PrintComics(comics)
}
