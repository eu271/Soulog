package soulObjects

import(
	"regexp"
	"log"
	"strings"
	
	"encoding/ascii85"
	
	"unicode"
	
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)


func createSlugFromTitle(title string) string {
	
	urlNotPermited := regexp.MustCompile("[^a-zA-Z0-9 -]")
	omitableInUrl := regexp.MustCompile("[ -]+")
	dash := "-"
	
	slug := strings.ToLower(title)
	isMn := func(r rune) bool {
    	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	slug, _, _ = transform.String(t, slug)
	
	//TODO: Support other alphabets.

	slug = urlNotPermited.ReplaceAllString(slug, dash)
	slug = omitableInUrl.ReplaceAllString(slug, dash)
	
	
	log.Println(slug)

	
	return slug
}