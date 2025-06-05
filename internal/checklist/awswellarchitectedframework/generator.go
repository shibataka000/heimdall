//go:build ignore

// Generate requirements.json.

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	urllib "net/url"
	"os"
	"slices"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func walk(url string, depth int) ([]string, error) {
	resp, err := fetch(url)
	if err != nil {
		return nil, err
	}
	contents := []string{}
	if depth == 0 {
		content, err := getMainContent(resp)
		if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	} else {
		links, err := getLinks(resp)
		if err != nil {
			return nil, err
		}
		for _, link := range links {
			childlen, err := walk(link, depth-1)
			if err != nil {
				return nil, err
			}
			contents = append(contents, childlen...)
		}
	}
	return contents, nil
}

func getMainContent(resp []byte) (string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(resp))
	if err != nil {
		return "", err
	}
	return doc.Find("#main-content").Text(), nil
}

func getLinks(resp []byte) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(resp))
	if err != nil {
		return nil, err
	}
	urls := doc.Find("#main-content").Find("a").Map(func(i int, s *goquery.Selection) string {
		href, exists := s.Attr("href")
		if !exists {
			panic(errors.New("failed to get href attribute value"))
		}
		if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
			return href
		}
		url, err := urllib.JoinPath("https://docs.aws.amazon.com/wellarchitected/latest/framework/", href)
		if err != nil {
			panic(err)
		}
		return url
	})
	ignore := []string{
		"https://docs.aws.amazon.com/wellarchitected/latest/framework/general/latest/gr/docconventions.html",
	}
	return slices.DeleteFunc(urls, func(url string) bool {
		return slices.Contains(ignore, url) || !strings.HasPrefix(url, "https://docs.aws.amazon.com/wellarchitected/latest/framework/")
	}), nil
}

func fetch(url string) ([]byte, error) {
	log.Printf("fetch %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint:errcheck
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch %s: %s", url, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	// Appendix -> Pillars -> Best practice areas -> Questions -> Best practices.
	bestpractices, err := walk("https://docs.aws.amazon.com/wellarchitected/latest/framework/appendix.html", 4)
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(bestpractices, "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("requirements.json", data, 0664)
}
