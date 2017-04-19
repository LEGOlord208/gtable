package gtable

import (
	"math"
)

// TableItem is an item in the table.
// Please use NewItem to instansiate.
type TableItem struct {
	Text         string
	Width        int
	PaddingLeft  int
	PaddingRight int
	Center       bool
}

// NewItem creates a new table item.
func NewItem(text string) *TableItem {
	return &TableItem{Text: text}
}

// Padding sets both paddings (left and right)
func (ti *TableItem) Padding(pad int) {
	ti.PaddingLeft = pad
	ti.PaddingRight = pad
}

// TextSize retrieves the size of the text.
func (ti TableItem) TextSize() int {
	width := ti.Width
	if width <= 0 {
		width = len(ti.Text)
	}
	return width
}

// Size retrieves the total item size.
func (ti TableItem) Size() int {
	return ti.PaddingLeft + ti.TextSize() + ti.PaddingRight
}

// String prints out the table item.
// Used in StringTable#String.
func (ti TableItem) String() string {
	s := ""
	for i := 0; i < ti.PaddingLeft; i++ {
		s += " "
	}

	width := ti.Width
	text := ti.Text

	if width > 0 && width < len(text) {
		text = text[:width]
	}
	width -= len(text)

	if width > 0 && ti.Center {
		half := int(math.Floor(float64(width) / 2.0))
		width -= half

		for i := 0; i < half; i++ {
			s += " "
		}
	}
	s += text
	for width > 0 {
		s += " "
		width--
	}

	for i := 0; i < ti.PaddingRight; i++ {
		s += " "
	}
	return s
}
