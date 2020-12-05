package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var punctuation = map[string]bool{
	"?": true,
	",": true,
	".": true,
	"!": true,
}

type Title struct {
	XMLName  xml.Name `xml:"title"`
	InnerXML string   `xml:",innerxml"`
}

type Body struct {
	XMLName  xml.Name `xml:"body"`
	InnerXML string   `xml:",innerxml"`
}

type Abstract struct {
	XMLName  xml.Name `xml:"abstract"`
	InnerXML string   `xml:",innerxml"`
}

type Article struct {
	Title    Title    `xml:""`
	Body     Body     `xml:""`
	Abstract Abstract `xml:""`
}

func main() {

	f, err := os.Open("in1.txt")
	if err != nil {
		f = os.Stdin
	}
	defer f.Close()
	r := bufio.NewReader(f)
	// Calculate SHIT I DONT WANT
	var s string
	fmt.Fscanln(r, &s)
	stoppers := strings.Split(s, ";")

	var replacements []string
	for v := range punctuation {
		replacements = append(replacements, v, "")
	}
	for _, v := range stoppers {
		if v != "" {
			replacements = append(replacements, v, " ")
		}
	}
	replacer := strings.NewReplacer(replacements...)
	// CALCULATE SHIT I DO WANT
	fmt.Fscanln(r, &s)
	aux := strings.Split(s, ";")
	var idxers = make(map[string]int)
	for _, v := range aux {
		idxers[v] = 0
	}

	b, _ := ioutil.ReadAll(r)
	b = append(b, []byte("</g>")...)
	b = append([]byte("<g>"), b...)
	art := new(Article)
	_ = xml.Unmarshal(b, art)
	art.Title.InnerXML = strings.ToLower(art.Title.InnerXML)
	art.Body.InnerXML = strings.ToLower(art.Body.InnerXML)
	art.Abstract.InnerXML = strings.ToLower(art.Abstract.InnerXML)
	fmt.Printf("%s", art.Body.InnerXML)
	art.Body.InnerXML = replacer.Replace(art.Body.InnerXML)
}
