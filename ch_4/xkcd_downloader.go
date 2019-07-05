package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	err := ioutil.WriteFile(fileName, data, 0655)
	if err != nil {
		fmt.Println(err)
	}

}

func ReadFile(fileName string) []byte {

	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

func main() {
	start := 1
	end := 25

	comics := GetComicRange(start, end)

	data, err := json.Marshal(comics)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	WriteFile("comics_index.json", data)

	data = ReadFile("comics_index.json")

	if err := json.Unmarshal(data, &comics); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s\n", err)
	}

	fmt.Printf("comics length: %d\n", len(comics))

	PrintComics(comics)
}
