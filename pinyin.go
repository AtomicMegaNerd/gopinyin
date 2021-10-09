package main

import (
	"log"
	"strings"
	"unicode"
)

type pinyinConverter struct {
	debug bool
}

func NewPinyinConverter(debug bool) *pinyinConverter {
	return &pinyinConverter{debug}
}

func runeInString(ch rune, str string) bool {
	for _, str_ch := range str {
		if ch == str_ch {
			return true
		}
	}
	return false
}

/*
   This method is the meat of the program.

   This is the algorithm we are going to use:

   For each character in each line:
   1. Start a search string as a buffer.
   2. If a character is an initial consonant ignore it.
   3. If a character is a vowel or an ending consonant append it to the search
   string.
   4. If a character is a tone indicator (1,2,3, or 4) that's the terminator for a
   search.  in that case, add it as the final character to the search string. Then
   search the 4, 3, 2, or 1 letter dictionaries for a match.  If a match is found
   replace it.

   It is very important that we search the bigger maps before the smaller ones
   because the smaller ones will also match on longer sequences and we do not want
   that.

   Append modified lines to a new list called converted which the caller can use
   to output pinyin tone marks.
*/
func (p *pinyinConverter) DoConvert(text []string) []string {
	converted := make([]string, 0)
	for _, line := range text {
		if p.debug {
			log.Printf("Line [%s]", line)
		}
		search_str := ""
		for _, ch := range line {
			// Add all letters that are not initial consonants to the search string.
			// Spaces and punctuation are all ignored.
			if unicode.IsLetter(ch) {
				if !runeInString(ch, Vowels) && len(search_str) == 0 {
					continue
				}
				search_str += string(ch)

				// All of our search strings will end with the tone indicator which will be
				// 1, 2, 3, or 4.  Note that the this number may come after a vowel or a consonant
				// but we cover both cases in our 4 ConvertMaps.
			} else if runeInString(ch, Tones) {

				// Do not forget to append the number at the end as this is part of our
				// search!
				search_str += string(ch)

				if p.debug {
					log.Println("buffer = ", search_str)
				}

				// Make sure that we always search the bigger maps first otherwise
				// we will not catch the longer sequences.  The search string is
				// compared with the keys from each dict.  If it matches a key in
				// any of our maps the value from that dict will be used to
				// replace the text in the line.
				if _, ok := ConvertMap4[search_str]; ok {
					line = strings.Replace(line, search_str, ConvertMap4[search_str], 1)
				} else if _, ok := ConvertMap3[search_str]; ok {
					line = strings.Replace(line, search_str, ConvertMap3[search_str], 1)
				} else if _, ok := ConvertMap2[search_str]; ok {
					line = strings.Replace(line, search_str, ConvertMap2[search_str], 1)
				} else if _, ok := ConvertMap1[search_str]; ok {
					line = strings.Replace(line, search_str, ConvertMap1[search_str], 1)
				}

				// If we find a match reset the search_str for the next word
				search_str = ""
			} else {
				// If we do not find a match reset the search_str for the next word
				search_str = ""
			}
		}

		if p.debug {
			log.Printf("Processed line [%s]", line)
		}
		converted = append(converted, line)
	}

	return converted
}
