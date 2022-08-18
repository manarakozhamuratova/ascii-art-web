package delivery

import (
	"log"
	"strings"
	"unicode"
)

type input struct {
	text   string
	banner string
}

func (i *input) normalize() {
	i.banner = strings.TrimFunc(i.banner, func(r rune) bool {
		return unicode.IsSpace(r)
	})
}

func (i *input) validate() bool {
	i.normalize()
	log.Printf("%#v\n", i)
	if i.text == "" || i.banner == "" {
		return false
	}
	return true
}

func (i input) checkBanner(arr ...string) bool {
	for index := range arr {
		if arr[index] == i.banner {
			return true
		}
	}
	return false
}
