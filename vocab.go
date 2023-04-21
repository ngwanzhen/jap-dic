package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// SAVE AS TYPE SLICE OF STRING
type stringSlice []string

func convertToStringSlice(i string) stringSlice {
	return strings.Split(i, ",")
}

func (v stringSlice) convertToString() string {
	return strings.Join(v, ",")
}

// SAVE AS TYPE STRUCT
type vocabStruct struct {
	jap string
	eng string
}

func convertToStruct(i string) vocabStruct {
	s := strings.Split(i, ",")
	a := vocabStruct{}
	for _, v := range s {
		r := []rune(v)
		if unicode.In(r[0], unicode.Hiragana, unicode.Han) {
			fmt.Println(string(v), "is a Jap character")
			a.jap = v
		} else {
			fmt.Println(string(v), "is not a Jap character")
			a.eng = v
		}
	}
	return a
}

func (v vocabStruct) convertToString() string {
	return v.jap + "," + v.eng
}

// SAVE AS TYPE MAPS
type vocabMap map[string]string

func convertToMap(i string) vocabMap {
	s := strings.Split(i, ",")
	var ja, en string
	for _, v := range s {
		if isJap(v) {
			// fmt.Println(string(v), "is a Jap character")
			ja = v
		} else {
			// fmt.Println(string(v), "is not a Jap character")
			en = v
		}
	}
	return vocabMap{ja: en}
}

func (v vocabMap) convertToString() string {
	var s string
	for ja, en := range v {
		s = ja + "," + en
	}
	return s
}

// DETECT LANGUAGE
func isJap(s string) bool {
	r := []rune(s)
	if unicode.In(r[0], unicode.Hiragana, unicode.Han) {
		return true
	} else {
		return false
	}
}

// INTERFACE TO VOCAB TYPE
type vocab interface {
	convertToString() string
}

// FUNCTIONS TAKING IN VOCAB TYPE
func print(v vocab) {
	fmt.Println(v.convertToString())
}

func saveToFile(v vocab) {
	// 	// SAVE FIRST LINE IN NEW FILE
	// 	// b := []byte(convertToString(v))
	// 	// err := ioutil.WriteFile("hello", b, 0644)
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }

	// APPEND SECOND LINE IN EXISTING FILE
	file, err := os.OpenFile("dic.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	if _, err := file.WriteString(v.convertToString() + "\n"); err != nil {
		// if _, err := file.WriteString(convertToString(v) + "\n"); err != nil {
		log.Fatal(err)
	}
}


func readFromFile() []string {
	file, _ := os.Open("dic.txt")
	scanner := bufio.NewScanner(file)

	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return lines
}

func randomVocab(a []string) (string, string) {
	// a := strings.Split(s, "\n")

	// Random index
	s1 := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s1)
	r1 := r.Intn(len(a)) // random index in list of vocab
	r2 := r.Intn(2)      // random 0 or 1

	v := strings.Split(a[r1], ",")
	if r2 == 0 {
		return v[0], v[1]
	} else {
		return v[1], v[0]
	}

}

func find(s string) string {
	file, _ := os.Open("dic.txt")
	scanner := bufio.NewScanner(file)

	ans := ""
	for {
		scanner.Scan()
		line := scanner.Text()
		ans = line
		if strings.Split(line, ",")[0] == s {
			break
		} else if len(line) == 0 {
			break
		}
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return ans

}
