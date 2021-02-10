package translator

import (
	"strings"
	"unicode"
)

type TextBinaryAdapter struct {
	inputText string
}

func NewTextBinaryAdapter(inputText string) *TextBinaryAdapter {
	return &TextBinaryAdapter{
		inputText: inputText,
	}
}

func (tba TextBinaryAdapter) Translate() (output string, err error) {
	return "", nil
}

//----------------------------------------------------------------------------------------------------------------------/

type TextMorseAdapter struct {
	inputText    string
	textMorseMap map[string]string
}

func NewTextMorseAdapter(inputText string) *TextMorseAdapter {
	tMMap := makeTextMorseMap()
	return &TextMorseAdapter{
		inputText:    inputText,
		textMorseMap: tMMap,
	}
}

func (tma TextMorseAdapter) Translate() (output string, err error) {

	var sb strings.Builder
	for _, character := range tma.inputText {
		character := unicode.ToUpper(character)
		morseChar := tma.textMorseMap[string(character)]
		sb.WriteString(morseChar)
	}

	output = sb.String()
	return
}

func makeTextMorseMap() map[string]string {
	m := make(map[string]string)
	m["A"] = ".- "
	m["B"] = "-... "
	m["C"] = "-.-. "
	m["D"] = "-.. "
	m["E"] = ". "
	m["F"] = "..-. "
	m["G"] = "--. "
	m["H"] = ".... "
	m["I"] = "... "
	m["J"] = ".--- "
	m["K"] = "-.- "
	m["L"] = ".-.. "
	m["M"] = "-- "
	m["N"] = "-. "
	m["O"] = "--- "
	m["P"] = ".--. "
	m["Q"] = "--.- "
	m["R"] = ".-. "
	m["S"] = "... "
	m["T"] = "- "
	m["U"] = "..- "
	m["V"] = "...- "
	m["W"] = ".-- "
	m["X"] = "-..- "
	m["Y"] = "-.-- "
	m["Z"] = "--.. "
	m[" "] = "   "

	return m
}

//----------------------------------------------------------------------------------------------------------------------/
