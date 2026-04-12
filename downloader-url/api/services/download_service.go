package services

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

type UrlItem struct {
	Url string
	Id  *uuid.UUID
}

type UrlItemResult struct {
	item *UrlItem
	err  error
}

func downloadByUrl(urlValidated string, id uuid.UUID, ch chan *UrlItemResult) {
	dirName := id.String()
	pathDir := fmt.Sprintf("./storage/%s", dirName)

	err := os.MkdirAll(pathDir, os.ModePerm)
	if err != nil {
		ch <- &UrlItemResult{err: err}
		return
	}

	idItemCreated := uuid.New()
	path := fmt.Sprintf("%s/%s", pathDir, idItemCreated.String())

	cmd := exec.Command("wget", urlValidated, "-O", path)
	_, err = cmd.CombinedOutput()
	if err != nil {
		ch <- &UrlItemResult{err: err}
		return
	}

	ch <- &UrlItemResult{
		item: &UrlItem{
			Url: urlValidated,
			Id:  &idItemCreated,
		},
	}
}
func ExecDownload(urls []*url.URL, id uuid.UUID) {
	idItems := []*UrlItem{}
	ch := make(chan *UrlItemResult, len(urls))
	for _, urlItem := range urls {
		go downloadByUrl(urlItem.String(), id, ch)
	}

	for i := 0; i < len(urls); i++ {
		res := <-ch

		if res.err != nil {
			fmt.Println(res.err.Error())
			continue
		} else {
			idItems = append(idItems, res.item)
		}
	}
}
