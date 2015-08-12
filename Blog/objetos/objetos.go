/*
	Copyright (c) 2015 Eugenio Ochoa

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/

package soulObjects

import(
	"regexp"
	"log"
	"strings"
	

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
