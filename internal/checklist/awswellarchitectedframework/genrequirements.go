//go:build ignore

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	wa "github.com/shibataka000/heimdall/internal/checklist/awswellarchitectedframework"
)

func scrape(base *url.URL, depth int) ([]wa.Requirement, error) {
	resp, err := fetch(base)
	if err != nil {
		return nil, err
	}
	requirements := []wa.Requirement{}
	if depth == 0 {
		requirement, err := parseBestPracticePage(resp)
		if err != nil {
			return nil, err
		}
		requirements = append(requirements, requirement)
	} else {
		links, err := parseParentPage(resp)
		if err != nil {
			return nil, err
		}
		for _, link := range links {
			childlen, err := scrape(link, depth-1)
			if err != nil {
				return nil, err
			}
			requirements = append(requirements, childlen...)
		}
	}
	return requirements, nil
}

func parseBestPracticePage(resp []byte) (wa.Requirement, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(resp))
	if err != nil {
		return wa.Requirement{}, err
	}

	title := doc.Find("title").First().Text()
	title = strings.TrimSuffix(title, " - AWS Well-Architected Framework")

	id := strings.Split(title, " ")[0]

	body := doc.Find(".awsdocs-page-header-container").NextUntil("#resources").Text()
	body = regexp.MustCompile(`\s+`).ReplaceAllString(body, " ")
	body = strings.TrimSpace(body)

	return wa.Requirement{ID: id, Title: title, Body: body}, nil
}

func parseParentPage(resp []byte) ([]*url.URL, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(resp))
	if err != nil {
		return nil, err
	}
	hrefs := doc.Find("#main-content").Find("a").Map(func(i int, s *goquery.Selection) string {
		href, exists := s.Attr("href")
		if !exists {
			panic(errors.New("failed to get href attribute value"))
		}
		return href
	})
	base, err := url.Parse("https://docs.aws.amazon.com/wellarchitected/latest/framework/")
	if err != nil {
		return nil, err
	}
	urls := []*url.URL{}
	for _, href := range hrefs {
		u, err := url.Parse(href)
		if err != nil {
			return nil, err
		}
		if !u.IsAbs() {
			u = base.ResolveReference(u)
		}
		if !strings.HasPrefix(u.String(), "https://docs.aws.amazon.com/wellarchitected/latest/framework/") {
			continue
		}
		if u.String() == "https://docs.aws.amazon.com/wellarchitected/latest/framework/general/latest/gr/docconventions.html" {
			continue
		}
		urls = append(urls, u)
	}
	return urls, nil
}

func fetch(u *url.URL) ([]byte, error) {
	log.Printf("fetch %s", u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint:errcheck
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch %s: %s", u.String(), resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	// Scrape the pages by the following orders.
	// 1. Appendix -> 2. Pillar -> 3. Best practice area -> 4. Question -> 5. Best practice.
	base, err := url.Parse("https://docs.aws.amazon.com/wellarchitected/latest/framework/appendix.html")
	if err != nil {
		panic(err)
	}
	requirements, err := scrape(base, 4)
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(requirements, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("requirements.json", data, 0664)
	if err != nil {
		panic(err)
	}
}
