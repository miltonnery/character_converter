package translator

type MorseBinaryAdapter struct {
	inputText string
}

func NewMorseBinaryAdapter(inputText string) *MorseBinaryAdapter {
	return &MorseBinaryAdapter{
		inputText: inputText,
	}
}

func (mba MorseBinaryAdapter) Translate() (output string, err error) {
	return "", nil
}

//----------------------------------------------------------------------------------------------------------------------/

type MorseTextAdapter struct {
	inputText string
}

func NewMorseTextAdapter(inputText string) *MorseTextAdapter {
	return &MorseTextAdapter{
		inputText: inputText,
	}
}

func (mta MorseTextAdapter) Translate() (output string, err error) {
	return "", nil
}

//----------------------------------------------------------------------------------------------------------------------/
