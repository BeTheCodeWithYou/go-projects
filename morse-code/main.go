package main

import (
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

func main() {

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
	tmpl.ExecuteTemplate(w, "decoded-morse-code-here", decodeMorse(inputTxt))
	} else {
	 tmpl.ExecuteTemplate(w, "decoded-morse-code-here", encodeToMorseCode(inputTxt))
	}
	
}

func decodeMorse(morseCode string) string {
	wrds := strings.Split(strings.TrimSpace(morseCode), "   ")
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

	s := strings.Split(strings.ToUpper(strings.TrimSpace(plainText)), "")
	var cd []string
	var mcode string
	for _,v := range s{	
		if v == " "{
			cd = append(cd, " ")			
		} else {
			mcode = morseValMap[v]
			cd = append(cd, mcode)
		}
	}
	return strings.Join(cd, " ")
	
}