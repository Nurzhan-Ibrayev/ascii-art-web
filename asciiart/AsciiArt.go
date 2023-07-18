package asciiart

import (
	"errors"
	"strings"
)

func AsciiArt(text, style string) (string, error) {
	TextFromFile, _ := CheckHash(style)
	Symbs := strings.Split(strings.ReplaceAll(TextFromFile, "\r", ""), "\n\n")
	ModStr := strings.ReplaceAll(text, "\n", "\\n")
	Text := strings.Split(ModStr, "\\n")
	res, err := FilterAndPrint(Text, ModStr, Symbs)
	if err != nil {
		return "", err
	}
	return res, nil
}

func FilterAndPrint(Text []string, ModStr string, Symbs []string) (string, error) {
	// if len(Text) == 0 || len(Text) <= 1 && Text[0] == "" {
	// 	return "", errors.New("Error: Incorrect input! " + ModStr)
	// }
	for i := 0; i < len(Text); i++ {
		if len(Text) == len(ModStr)+1 && Text[i] == "" {
			Text = append(Text[:i], Text[i+1:]...)
		}
	}
	onlyNewLines := checkNewLines(Text)
	res := ""
	// if len(Text) == 0 {
	// 	log.Panicln("Empty Input!")
	// }
	for i, word := range Text {
		if onlyNewLines && i == len(Text)-1 {
			break
		}
		if word == "" {
			res += "\n"
		} else {
			for line := 0; line < 8; line++ {
				for _, rune := range word {
					if rune-32 == -19 {
						continue
					}
					if rune >= 32 && rune <= 127 {
						res += strings.Split(Symbs[rune-32], "\n")[line]
					} else {
						res = ""
						return res, errors.New("only 32-127 asccii allert, bitch")
					}
				}
				res += "\n"
			}
		}
	}
	return res, nil
}

func checkNewLines(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return false
		}
	}
	return true
}
