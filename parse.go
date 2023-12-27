package parse

import (
	"log"
	"os"
)

func ReadConfig(filepath string) map[string]string {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("unable to read file")
	}

	return parse(file)
}

func parse(bytes []byte) map[string]string {
	var headStart bool
	word := ""
	field := ""
	changeLine := false
	newLine := true
	configMap := make(map[string]string)
	for _, byt := range bytes {
		if newLine && byt == byte('\n') {
			newLine = true
			continue
		}
		newLine = false
		if string(byt) == "[" {
			headStart = true
			word = ""
			continue
		}

		if string(byt) == "]" {
			headStart = false
			changeLine = true
			continue
		}

		if !headStart && changeLine {
			changeLine = false
			continue
		}

		if headStart {
			word = word + string(byt)
		}

		if !headStart {
			if byt == byte('\n') {
				newLine = true
				a := split(field)
				if len(a) == 0 {
					field = ""
					continue
				}
				configMap[word+"."+a[0]] = a[1]
				field = ""
				continue
			}
			field = field + string(byt)
		}
	}

	return configMap
}

// custom splitter so we split the string on at the first occurence of the equalTo(=).
func split(str string) []string {
	if len(str) == 0 {
		return []string{}
	}

	splitStrings := make([]string, 0)
	for i, s := range str {
		if s == rune('=') {
			splitStrings = append(splitStrings, str[:i], str[i+1:])
		}
	}

	return splitStrings
}
