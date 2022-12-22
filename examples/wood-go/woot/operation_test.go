package woot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IntegrateInsert(t *testing.T) {
	wstring = Initialize()

	wcharacter := WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID}

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: wcharacter.ID},
			{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: wcharacter.ID, WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.IntegrateInsert(wcharacter, WCharacterStart, WCharacterEnd)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

func Test_IntegrateInsert_Middle(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "a"},
			{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "b"},
			{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "a", WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "b", WCNext: ""},
		},
	}

	wcharacter := WCharacter{ID: "x", Visible: false, Alphabet: "x", WCPrevious: "a", WCNext: "b"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "a"},
			{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "x"},
			{ID: "x", Visible: false, Alphabet: "x", WCPrevious: "a", WCNext: "b"},
			{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "x", WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "b", WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.IntegrateInsert(
		wcharacter,
		WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "b"},
		WCharacter{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "a", WCNext: WCharacterEnd.ID},
	)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

func Test_IntegrateInsert_Transpose(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "a"},
			{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "b"},
			{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "a", WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "b", WCNext: ""},
		},
	}

	wcharacter := WCharacter{ID: "x", Visible: false, Alphabet: "x", WCPrevious: WCharacterStart.ID, WCNext: "a"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "x"},
			{ID: "x", Visible: false, Alphabet: "x", WCPrevious: WCharacterStart.ID, WCNext: "a"},
			{ID: "a", Visible: false, Alphabet: "a", WCPrevious: "x", WCNext: "b"},
			{ID: "b", Visible: false, Alphabet: "b", WCPrevious: "a", WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "b", WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.IntegrateInsert(
		wcharacter,
		WCharacterStart,
		WCharacter{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: "b"},
	)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
}

func Test_IntegrateDelete(t *testing.T) {
	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID}

	wstring = WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: wcharacter.ID},
			{ID: "a", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: wcharacter.ID, WCNext: ""},
		},
	}

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: wcharacter.ID},
			{ID: "a", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: wcharacter.ID, WCNext: ""},
		},
	}

	actualWString := wstring.IntegrateDelete(wcharacter)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
}

func Test_IntegrateDelete_WCharacterNotPresent(t *testing.T) {
	wcharacter := WCharacter{ID: "a", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID}

	wstring = WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: WCharacterStart.ID, WCNext: ""},
		},
	}

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: WCharacterStart.ID, WCNext: ""},
		},
	}

	actualWString := wstring.IntegrateDelete(wcharacter)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
}

func Test_GenerateInsert(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	position, alphabet := 1, "a"

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			{ID: "01", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "01", WCNext: ""},
		},
	}

	actualWString, actualErr := wstring.GenerateInsert(position, alphabet)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateInsert_ReplaceStart(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	position := 1
	alphabet1, alphabet2 := "a", "b"

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "02"},
			{ID: "02", Visible: true, Alphabet: "b", WCPrevious: "start", WCNext: "01"},
			{ID: "01", Visible: true, Alphabet: "a", WCPrevious: "02", WCNext: "end"},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "01", WCNext: ""},
		},
	}

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(position, alphabet1)
	actualWString, actualErr := WStringPtr.GenerateInsert(position, alphabet2)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateInsert_ReplaceEnd(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	alphabet1, alphabet2 := "a", "b"

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			{ID: "01", Visible: true, Alphabet: "a", WCPrevious: "start", WCNext: "02"},
			{ID: "02", Visible: true, Alphabet: "b", WCPrevious: "01", WCNext: "end"},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "02", WCNext: ""},
		},
	}

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(1, alphabet1)
	actualWString, actualErr := WStringPtr.GenerateInsert(2, alphabet2)

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateInsert_Word(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	alphabets := []string{"a", "b", "c"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			{ID: "01", Visible: true, Alphabet: "a", WCPrevious: "start", WCNext: "02"},
			{ID: "02", Visible: true, Alphabet: "b", WCPrevious: "01", WCNext: "03"},
			{ID: "03", Visible: true, Alphabet: "c", WCPrevious: "02", WCNext: "end"},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "03", WCNext: ""},
		},
	}

	var actualErr error
	count := 1

	actualWString, _ := wstring.GenerateInsert(count, alphabets[0])

	for _, alphabet := range alphabets[1:] {
		count++
		actualWString, actualErr = wstring.GenerateInsert(count, alphabet)
	}

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	expectedText := "abc"
	actualText := Value(*actualWString)

	assert.Equal(t, expectedText, actualText)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateInsert_Sentence(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	alphabets := []string{"I", " ", "l", "i", "k", "e", " ", "d", "o", "g", "s"}

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			{ID: "01", Visible: true, Alphabet: "I", WCPrevious: "start", WCNext: "02"},
			{ID: "02", Visible: true, Alphabet: " ", WCPrevious: "01", WCNext: "03"},
			{ID: "03", Visible: true, Alphabet: "l", WCPrevious: "02", WCNext: "04"},
			{ID: "04", Visible: true, Alphabet: "i", WCPrevious: "03", WCNext: "05"},
			{ID: "05", Visible: true, Alphabet: "k", WCPrevious: "04", WCNext: "06"},
			{ID: "06", Visible: true, Alphabet: "e", WCPrevious: "05", WCNext: "07"},
			{ID: "07", Visible: true, Alphabet: " ", WCPrevious: "06", WCNext: "08"},
			{ID: "08", Visible: true, Alphabet: "d", WCPrevious: "07", WCNext: "09"},
			{ID: "09", Visible: true, Alphabet: "o", WCPrevious: "08", WCNext: "010"},
			{ID: "010", Visible: true, Alphabet: "g", WCPrevious: "09", WCNext: "011"},
			{ID: "011", Visible: true, Alphabet: "s", WCPrevious: "010", WCNext: "end"},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "011", WCNext: ""},
		},
	}

	var actualErr error
	count := 1

	actualWString, _ := wstring.GenerateInsert(count, alphabets[0])

	for _, alphabet := range alphabets[1:] {
		count++
		actualWString, actualErr = wstring.GenerateInsert(count, alphabet)
	}

	assert.Equal(t, expectedWString, actualWString)
	assert.Nil(t, actualErr)

	expectedText := "I like dogs"
	actualText := Value(*actualWString)

	assert.Equal(t, expectedText, actualText)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateDelete(t *testing.T) {
	wstring = WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			{ID: "01", Visible: true, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "01", WCNext: ""},
		},
	}
	LocalClock = 0

	position := 1

	expectedWString := &WString{
		Sequence: []WCharacter{
			{ID: "start", Visible: false, Alphabet: "", WCPrevious: "", WCNext: "01"},
			{ID: "01", Visible: false, Alphabet: "a", WCPrevious: WCharacterStart.ID, WCNext: WCharacterEnd.ID},
			{ID: "end", Visible: false, Alphabet: "", WCPrevious: "01", WCNext: ""},
		},
	}

	actualWString := wstring.GenerateDelete(position)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateDelete_NoValue(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	position := 1

	expectedWString := &wstring

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(position, "a")
	WStringPtr = WStringPtr.GenerateDelete(position)
	actualWString := WStringPtr.GenerateDelete(position)

	assert.Equal(t, expectedWString, actualWString)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateDelete_WordReplaceInPlace(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(1, "1")
	WStringPtr, _ = WStringPtr.GenerateInsert(2, "2")
	WStringPtr, _ = WStringPtr.GenerateInsert(3, "3")
	WStringPtr = WStringPtr.GenerateDelete(2)
	WStringPtr, _ = WStringPtr.GenerateInsert(2, "4")

	expectedText := "143"
	actualText := Value(*WStringPtr)

	assert.Equal(t, expectedText, actualText)

	wstring = Clear()
	LocalClock = 0
}

func Test_GenerateDelete_SentenceReplaceInPlace(t *testing.T) {
	wstring = Initialize()
	LocalClock = 0

	var WStringPtr *WString

	WStringPtr, _ = wstring.GenerateInsert(1, "1")
	WStringPtr, _ = WStringPtr.GenerateInsert(2, "2")
	WStringPtr, _ = WStringPtr.GenerateInsert(3, "3")
	WStringPtr, _ = WStringPtr.GenerateInsert(4, " ")

	WStringPtr, _ = wstring.GenerateInsert(5, "4")
	WStringPtr, _ = WStringPtr.GenerateInsert(6, "5")
	WStringPtr, _ = WStringPtr.GenerateInsert(7, "6")
	WStringPtr, _ = WStringPtr.GenerateInsert(8, " ")

	WStringPtr, _ = wstring.GenerateInsert(9, "7")
	WStringPtr, _ = WStringPtr.GenerateInsert(10, "8")
	WStringPtr, _ = WStringPtr.GenerateInsert(11, "9")
	WStringPtr, _ = WStringPtr.GenerateInsert(12, " ")

	WStringPtr = wstring.GenerateDelete(1)
	WStringPtr = WStringPtr.GenerateDelete(1)
	WStringPtr = WStringPtr.GenerateDelete(1)

	WStringPtr, _ = wstring.GenerateInsert(1, "7")
	WStringPtr, _ = WStringPtr.GenerateInsert(2, "8")
	WStringPtr, _ = WStringPtr.GenerateInsert(3, "9")

	WStringPtr = wstring.GenerateDelete(9)
	WStringPtr = WStringPtr.GenerateDelete(9)
	WStringPtr = WStringPtr.GenerateDelete(9)

	WStringPtr, _ = wstring.GenerateInsert(9, "1")
	WStringPtr, _ = wstring.GenerateInsert(10, "2")
	WStringPtr, _ = wstring.GenerateInsert(11, "3")

	expectedText := "789 456 123 "
	actualText := Value(*WStringPtr)

	assert.Equal(t, expectedText, actualText)

	wstring = Clear()
	LocalClock = 0
}
