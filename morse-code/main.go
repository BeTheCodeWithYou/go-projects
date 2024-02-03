package main

import (
	"fmt"
	//"io"
	"net/http"
	"strings"
	"text/template"
)

var morsemap = map[string]string{
	".-":     "A",
	"-...":   "B",
	"-.-.":   "C",
	"-..":    "D",
	".":      "E",
	"..-.":   "F",
	"--.":    "G",
	"....":   "H",
	"..":     "I",
	".---":   "J",
	"-.-":    "K",
	".-..":   "L",
	"--":     "M",
	"-.":     "N",
	"---":    "O",
	".--.":   "P",
	"--.-":   "Q",
	".-.":    "R",
	"...":    "S",
	"-":      "T",
	"..-":    "U",
	"...-":   "V",
	".--":    "W",
	"-..-":   "X",
	"-.--":   "Y",
	"--..":   "Z",
	".----":  "1",
	"..---":  "2",
	"...--":  "3",
	"....-":  "4",
	".....":  "5",
	"-....":  "6",
	"--...":  "7",
	"---..":  "8",
	"----.":  "9",
	"-----":  "0",
	".-.-.-": ".",
	"--..--": ",",
	"..--..": "?",
	"-.-.--": "!",
	"-....-": "-",
	"-..-.":  "/",
	".--.-.": "@",
	"-.--.":  "(",
	"-.--.-": ")",
}

var morseValMap = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
	".": ".-.-.-",
	",": "--..--",
	"?": "..--..",
	"!": "-.-.--",
	"-": "-....-",
	"/": "-..-.",
	"@": ".--.-.",
	"(": "-.--.",
	")": "-.--.-",
	" ": "   ",
}

// -.. --- .. -. --.    .-- --- .-. -.-   ..-. --- .-.   -- .   .- -. -..   -- ..- --   - .... .- -   ..-
//When the message is written in Morse code, a single space is used to separate the character codes and 
//3 spaces are used to separate words
//The Morse code is case-insensitive, traditionally capital letters are used

func main() {

	str := decodeMorse(".... . -.--   .--- ..- -.. .")
	fmt.Println(str)
	fmt.Println("###->",encodeToMorseCode("hello hi"))

	http.HandleFunc("/", pageLoadHandler)
	http.HandleFunc("/morsecode", morseCodeHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func pageLoadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func morseCodeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles("index.html"))
	inputTxt := r.Form["morse-code"][0]
	if r.Form["actionVal"]!=nil{
	tmpl.ExecuteTemplate(w, "decoded-morese-code-here", decodeMorse(inputTxt))
	} else {
	 tmpl.ExecuteTemplate(w, "decoded-morese-code-here", encodeToMorseCode(inputTxt))
	}
	
}

func decodeMorse(morseCode string) string {
	morseCode = strings.TrimSpace(morseCode)
	wrds := strings.Split(morseCode, "   ")
	for i, word := range wrds {
		chrs := strings.Split(word, " ")
		for i, chr := range chrs {
			chrs[i] = morsemap[chr]
		}
		letter := strings.Join(chrs, "")
		wrds[i] = letter
	}

	return strings.Join(wrds, " ")
}

func encodeToMorseCode(plainText string) string {

	s1 := strings.ToUpper(strings.TrimSpace(plainText))
	s := strings.Split(s1, "")
	var cd []string
	var mcode string
	for _,v := range s{
		mcode = morseValMap[v]
		cd = append(cd, mcode)
		cd = append(cd, " ")
		if v ==" "{
		  cd = append(cd, "   ")
		}
	}
	return strings.Join(cd, " ")
}
