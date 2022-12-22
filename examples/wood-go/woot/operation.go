package woot

import "fmt"

var (
	SiteID     = 0 // siteID
	LocalClock = 0 // 相对时间
)

func (wstring *WString) Find(ID string) WCharacter {
	for _, w := range wstring.Sequence {
		if w.ID == ID {
			return w
		}
	}
	return WCharacter{ID: "-1"}
}

func (wstring *WString) GenerateInsert(position int,
	alphabet string) (*WString, error) {
	LocalClock++

	// previous w
	previousW := IthVisible(*wstring, position-1)
	nextW := IthVisible(*wstring, position)

	if previousW.ID == "-1" {
		previousW = wstring.Find("start") // 最开始
	}
	if nextW.ID == "-1" {
		nextW = wstring.Find("end")
	}

	wcharacter := WCharacter{
		ID:         fmt.Sprint(SiteID) + fmt.Sprint(LocalClock),
		Visible:    true,
		Alphabet:   alphabet,
		WCPrevious: previousW.ID,
		WCNext:     nextW.ID,
	}

	return wstring.IntegrateInsert(wcharacter, previousW, nextW)
}

func (wstring *WString) IntegrateInsert(wcharacter, previousW,
	nextW WCharacter) (*WString, error) {
	sub, _ := wstring.Subseq(previousW, nextW)
	position := wstring.Position(nextW.ID)

	position--

	if len(sub) == 0 {
		return wstring.LocalInsert(wcharacter, position)
	}

	if len(sub) == 1 {
		return wstring.LocalInsert(wcharacter, position-1)
	}

	index := 1
	for index < len(sub)-1 && sub[index].ID < wcharacter.ID {
		index++
	}

	return wstring.IntegrateInsert(wcharacter, sub[index-1], sub[index])
}

func (wstring *WString) GenerateDelete(position int) *WString {
	wcharacter := IthVisible(*wstring, position)
	return wstring.IntegrateDelete(wcharacter)
}

func (wstring *WString) IntegrateDelete(wcharacter WCharacter) *WString {
	position := wstring.Position(wcharacter.ID)

	if position == -1 {
		return wstring
	}

	wstring.Sequence[position-1].Visible = false
	return wstring
}
