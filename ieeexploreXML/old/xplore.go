package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

var punctuation = map[string]bool{
	"?": true,
	",": true,
	".": true,
	"!": true,
}

type word struct {
	str   string
	score float64
	count int
}

type Index []*word

func (idx Index) Swap(i, j int) { idx[i], idx[j] = idx[j], idx[i] }
func (idx Index) Less(i, j int) bool {
	if idx[i].score == idx[j].score {
		return idx[i].str < idx[j].str
	}
	return idx[i].score > idx[j].score
}
func (i Index) Len() int { return len(i) }

func main() {

	f, err := os.Open("in2.txt")
	if err != nil {
		f = os.Stdin
	}

	// Calculate SHIT I DONT WANT
	var s string
	fmt.Fscanln(f, &s)
	stoppers := strings.Split(s, ";")

	var replacements []string
	for v := range punctuation {
		replacements = append(replacements, v, " ")
	}
	for _, v := range stoppers {
		if v != "" {
			replacements = append(replacements, v, " ")
		}
	}
	replacer := strings.NewReplacer(replacements...)

	// CALCULATE SHIT I DO WANT
	fmt.Fscanln(f, &s)
	aux := strings.Split(s, ";")
	var idxers = make(map[string]int)
	for _, v := range aux {
		idxers[v] = 0
	}
	EOF := false
	var xmlBuild strings.Builder
	for !EOF {
		b := make([]byte, 1024)
		n, err := f.Read(b)
		if err != nil || n == 0 {
			EOF = true
		}
		xmlBuild.Write(b[:n]) // suiper fast string builder
	}
	re := regexp.MustCompile(`\w{}[.,?!]{1}\w`)
	re.FindAllIndex()
	xml := replacer.Replace(strings.ToLower(xmlBuild.String()))
	title, abstract, body := innerHTML(xml, "title"), innerHTML(xml, "abstract"), innerHTML(xml, "body")
	title, abstract, body = deleteXMLTags(title), deleteXMLTags(abstract), deleteXMLTags(body)
	// body = deleteXMLTags(body)
	fieldsTitle, fieldsAbstract, fieldsBody := strings.Fields(title), strings.Fields(abstract), strings.Fields(body)
	fieldsTitle, fieldsAbstract, fieldsBody = discardShort(fieldsTitle), discardShort(fieldsAbstract), discardShort(fieldsBody)
	L := len(fieldsTitle) + len(fieldsAbstract) + len(fieldsBody)
	for _, v := range fieldsTitle {
		_, ok := idxers[v]
		if ok {
			idxers[v] += 5
		}
	}
	for _, v := range fieldsAbstract {
		_, ok := idxers[v]
		if ok {
			idxers[v] += 3
		}
	}
	for _, v := range fieldsBody {
		_, ok := idxers[v]
		if ok {
			idxers[v] += 1
		}
	}
	var index Index
	for s, n := range idxers {
		index = append(index, &word{str: s, score: 100. * float64(n) / float64(L), count: n})
	}
	sort.Sort(index)
	// print(L, "\n")
	i := 0
	done := false
	for !done {
		w := index[i]
		if i > 2 && w.score == index[2].score {
			fmt.Printf("%s: %.9f\n", w.str, w.score)
		} else if i > 2 {
			done = true
		} else {
			fmt.Printf("%s: %.9f\n", w.str, w.score)
		}

		// fmt.Printf("%s: %.9f %d\n", w.str, w.score, w.count) // debug line
		i++
	}

}
func discardShort(a []string) []string {
	n := 0
	for _, x := range a {
		if len(x) >= 4 {
			a[n] = x
			n++
		}
	}
	return a[:n]
}

func deleteXMLTags(s string) string {
	var b strings.Builder
	done := false
	idx := 0
	for !done { // idx0 y idx1 en coordenadas relativas a idx
		idx0 := strings.Index(s[idx:], "<")
		idx1 := strings.Index(s[idx:], ">")
		if idx0 < 0 || idx1 < 0 {
			if idx == 0 {
				return s
			}
			break
		}
		if idx0 != 0 {
			ess := s[idx : idx+idx0]
			b.WriteString(ess)
			idx += idx1 + 1
		}

	}
	return b.String()
}

// first html instance
func innerHTML(s, tag string) string {
	idx0 := strings.Index(s, fmt.Sprintf("<%s", tag))
	idx1 := strings.Index(s[idx0:], ">")
	idx2 := strings.Index(s[idx1:], fmt.Sprintf("</%s", tag))
	return s[idx0+idx1+1 : idx1+idx2]
}
