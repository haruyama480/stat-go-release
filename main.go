package main

import (
	"fmt"
	gover "go/version"
	"io"
	"log"
	"net/http"
	"regexp"
	"sort"
)

func main() {
	// fetch
	url := "https://go.dev/doc/devel/release"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	releaseLines := []string{}
	re := regexp.MustCompile(`go\d+(\.\d+){0,2}[\s\t]+\(released [\d\-]{10}\)`)
	matches := re.FindAllString(string(body), -1)
	for _, match := range matches {
		// Replace multiple spaces/tabs with single space
		cleanMatch := regexp.MustCompile(`[\s\t]{2,}`).ReplaceAllString(match, " ")
		releaseLines = append(releaseLines, cleanMatch)
	}

	// fmt.Println(releaseLines)
	// Output:
	// go1.24.0 (released 2025-02-11)
	// go1.23.0 (released 2024-08-13)
	// ..

	// collect releases
	type Release struct {
		Version string
		Date    string
	}
	releases := []Release{}
	for _, l := range releaseLines {
		version := regexp.MustCompile(`go\d+(\.\d+){0,2}`).FindString(l)
		date := regexp.MustCompile(`\(released [\d\-]{10}\)`).FindString(l)[10:20]
		if !gover.IsValid(version) {
			panic(fmt.Sprintf("Invalid go version. %s", version))
		}
		releases = append(releases, Release{Version: version, Date: date})
	}

	// sort
	sort.Slice(releases, func(i, j int) bool {
		return gover.Compare(releases[i].Version, releases[j].Version) > 0
	})

	// output as csv
	fmt.Println("Lang,GoVersion,Released")
	for _, r := range releases {
		lang := gover.Lang(r.Version)
		fmt.Printf("%s,%s,%s\n", lang, r.Version, r.Date)
	}
}
