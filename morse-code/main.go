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
}

// -.. --- .. -. --.    .-- --- .-. -.-   ..-. --- .-.   -- .   .- -. -..   -- ..- --   - .... .- -   ..-

func main() {

	str := DecodeMorse(".... . -.--   .--- ..- -.. .")
	fmt.Println(str)

	http.HandleFunc("/", pageLoadHandler)
	http.HandleFunc("/morsecode", morseCodeHandler)
	http.ListenAndServe(":8080", nil)
}

func pageLoadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func morseCodeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "decoded-morese-code-here", DecodeMorse(r.PostFormValue("morse-code")))
}

func DecodeMorse(morseCode string) string {
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
