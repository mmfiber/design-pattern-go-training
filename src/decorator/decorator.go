package decorator

import (
	"strings"
)

type Border struct {
	display Display
}

type SideBorder struct {
	*Border
	borderChar string
}

func NewSideBorder(display Display, borderChar string) *SideBorder {
	return &SideBorder{&Border{display}, borderChar}
}

func (b *SideBorder) GetColumns() int {
	return 1 + b.display.GetColumns() + 1
}

func (b *SideBorder) GetRows() int {
	return b.display.GetRows()
}

func (b *SideBorder) GetRowText(row int) (string, bool) {
	text, ok := b.display.GetRowText(row)
	if !ok {
		return "", false
	}

	char := b.borderChar
	return char + text + char, true
}

type FullBorder struct {
	*Border
}

func NewFullBorder(display Display) *FullBorder {
	return &FullBorder{&Border{display}}
}

func (b *FullBorder) GetColumns() int {
	return 1 + b.display.GetColumns() + 1
}

func (b *FullBorder) GetRows() int {
	return 1 + b.display.GetRows() + 1
}

func (b *FullBorder) GetRowText(row int) (string, bool) {
	if row == 0 || row == b.display.GetRows()+1 {
		return "+" + b.MakeLine("-", b.display.GetColumns()) + "+", true
	}

	text, ok := b.display.GetRowText(row - 1)
	if !ok {
		return "", false
	}
	return "|" + text + "|", true
}

func (b *FullBorder) MakeLine(ch string, count int) string {
	return strings.Repeat(ch, count)
}
