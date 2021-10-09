package main

import "unicode"

func runeInString(ch rune, str string) {
    for _, str_ch := range str {
        if ch == str_ch {
            return true
        }
    }
    return false
}
func do_convert(text []string) []string {
	converted := make([]string, len(text))
	for _, line := range text {
		search_str := ""
		for _, ch := range line {
			if unicode.IsLetter(ch) {
                if !isVowel(ch) && len(search_str) == 0 {
					continue
				}
				search_str += string(ch)
			} else if 
		}
	}
	return converted
}
