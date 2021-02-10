package translator

type BinaryMorseAdapter struct {
	inputText string
}

func NewBinaryMorseAdapter(inputText string) *BinaryMorseAdapter {
	return &BinaryMorseAdapter{
		inputText: inputText,
	}
}

func (bma BinaryMorseAdapter) Translate() (output string, err error) {
	return "", nil
}

//----------------------------------------------------------------------------------------------------------------------/

type BinaryTextAdapter struct {
	inputText string
}

func NewBinaryTextAdapter(inputText string) *BinaryTextAdapter {
	return &BinaryTextAdapter{
		inputText: inputText,
	}
}

func (mta BinaryTextAdapter) Translate() (output string, err error) {
	return "", nil
}

//----------------------------------------------------------------------------------------------------------------------/
