package woot

import "strings"

type (
	WString struct {
		Sequence []WCharacter
	}

	WCharacter struct {
		ID         string
		Visible    bool
		Alphabet   string
		WCPrevious string
		WCNext     string
	}
)

var (
	WCharacterStart = WCharacter{
		ID:         "start",
		Visible:    false,
		Alphabet:   "",
		WCPrevious: "",
		WCNext:     "end",
	}

	WCharacterEnd = WCharacter{
		ID:         "end",
		Visible:    false,
		Alphabet:   "",
		WCPrevious: "start",
		WCNext:     "",
	}
)

func Initialize() WString {
	return WString{Sequence: []WCharacter{WCharacterStart, WCharacterEnd}}
}

func (wstring *WString) Length() int {
	return len(wstring.Sequence)
}

func (wstring *WString) ElementAt(position int) (WCharacter, error) {
	if position < 0 || position >= wstring.Length() {
		return WCharacter{}, ErrPositionOutOfBounds
	}

	return wstring.Sequence[position], nil
}

func (wstring *WString) Position(wcharacterID string) int {
	for i, w := range wstring.Sequence {
		if w.ID == wcharacterID {
			return i + 1
		}
	}
	return -1 // not found
}

func (wstring *WString) LocalInsert(wcharacter WCharacter,
	position int) (*WString, error) {
	if position <= 0 || position >= wstring.Length() {
		return wstring, ErrPositionOutOfBounds
	}

	if wcharacter.ID == "" {
		return wstring, ErrEmptyWCharacter
	}

	wstring.Sequence = append(wstring.Sequence[:position],
		append([]WCharacter{wcharacter}, wstring.Sequence[position:]...)...,
	)
	// 更新指针
	wstring.Sequence[position-1].WCNext = wcharacter.ID
	wstring.Sequence[position+1].WCPrevious = wcharacter.ID

	return wstring, nil
}

func (wstring *WString) Subseq(wcharacterStart, wcharacterEnd WCharacter) ([]WCharacter, error) {
	startPosition := wstring.Position(wcharacterStart.ID)
	endPosition := wstring.Position(wcharacterEnd.ID)

	if startPosition == -1 || endPosition == -1 {
		return wstring.Sequence, ErrBoundNotPresent
	}

	if startPosition == endPosition {
		return []WCharacter{}, nil
	}

	return wstring.Sequence[startPosition : endPosition-1], nil
}

func (wstring *WString) Contains(wcharacterID string) bool {
	position := wstring.Position(wcharacterID)
	return position != -1
}

func Value(wstring WString) string {
	builder := strings.Builder{}
	for _, w := range wstring.Sequence {
		if w.Visible {
			builder.WriteString(w.Alphabet)
		}
	}
	return builder.String()
}

func IthVisible(wstring WString, position int) WCharacter {
	count := 0

	for _, w := range wstring.Sequence {
		if w.Visible {
			if count == position-1 {
				return w
			}
			count++
		}
	}
	// not found
	return WCharacter{ID: "-1"}
}
