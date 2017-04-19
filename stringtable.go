package gtable

// These values specify what corner we're talking
// about in StringTable#Corner.
const (
	CornerTopLeft = iota
	CornerTopRight
	CornerBottomLeft
	CornerBottomRight
	CornerMiddleHorizontal
	CornerMiddleVertical
)

// CornerPlus is a corner function
// which uses the + symbol as corner
var CornerPlus = func(i int) rune {
	return '+'
}

// CornerRound is a corner function
// which makes the generated table look round
var CornerRound = func(i int) rune {
	switch i {
	default:
		return '.'
	case CornerMiddleHorizontal:
		fallthrough
	case CornerMiddleVertical:
		return '+'
	case CornerBottomLeft:
		fallthrough
	case CornerBottomRight:
		return '\''
	}
}

// StringTable is the table object.
// Please initiate using the NewStringTable() function
type StringTable struct {
	Corner func(i int) rune
	Header bool
	rows   [][]*TableItem
}

// NewStringTable creates a new string table with all default values.
func NewStringTable() StringTable {
	return StringTable{
		Corner: CornerPlus,
		Header: true,
		rows:   [][]*TableItem{make([]*TableItem, 0)},
	}
}

// AddItems adds items to the table.
func (st *StringTable) AddItems(items ...*TableItem) {
	index := len(st.rows) - 1
	col := st.rows[index]

	for _, item := range items {
		col = append(col, item)
	}
	st.rows[index] = col
}

// AddStrings creates items by label and adds them to the table.
func (st *StringTable) AddStrings(items ...string) {
	tItems := make([]*TableItem, len(items))
	for i, item := range items {
		tItem := NewItem(item)
		tItems[i] = tItem
	}
	st.AddItems(tItems...)
}

// AddRow breaks the table and continues to the next row.
func (st *StringTable) AddRow() {
	st.rows = append(st.rows, make([]*TableItem, 0))
}

// Get gets a column from the table.
func (st *StringTable) Get(row, col int) *TableItem {
	return st.rows[row][col]
}

// Columns counts the maximum table column length.
func (st *StringTable) Columns() int {
	columns := 0

	for _, row := range st.rows {
		columns = max(len(row), columns)
	}
	return columns
}

func max(i1, i2 int) int {
	if i1 >= i2 {
		return i1
	}
	return i2
}

// Rows returns a copy of all rows in the table
func (st *StringTable) Rows() [][]*TableItem {
	var arr = make([][]*TableItem, len(st.rows))
	for i := range st.rows {
		arr[i] = make([]*TableItem, len(st.rows[i]))
		copy(arr[i], st.rows[i])
	}

	return arr
}

// Each executes 'handler' for every table item.
// Items may be modified.
// Useful for setting global properties.
func (st *StringTable) Each(handler func(*TableItem)) {
	for _, row := range st.rows {
		for _, col := range row {
			handler(col)
		}
	}
}

// String generates the ASCII table!
func (st *StringTable) String() string {
	s := ""
	n := "\n"

	rows := st.Rows()

	for len(rows) > 0 && len(rows[len(rows)-1]) <= 0 {
		last := len(rows) - 1
		rows = append(rows[:last], rows[last+1:]...)
	}

	if len(rows) <= 0 {
		return ""
	}

	columns := 0
	for _, row := range rows {
		columns = max(columns, len(row))
	}

	if columns <= 0 {
		return ""
	}

	lengths := make([]int, 0)
	for c := 0; c < columns; c++ {
		length := 0
		for r := range rows {
			for len(rows[r]) < columns {
				item := NewItem("")
				rows[r] = append(rows[r], item)
			}
			col := rows[r][c]

			length = max(length, col.Size())
		}
		lengths = append(lengths, length)
	}

	frame := ""

	first := true
	for _, length := range lengths {
		if first {
			first = false
		} else {
			frame += string(st.Corner(CornerMiddleHorizontal))
		}

		for i := 0; i < length; i++ {
			frame += "-"
		}
	}

	s += string(st.Corner(CornerTopLeft)) + frame + string(st.Corner(CornerTopRight)) + n

	first = true
	for _, row := range rows {
		for i, col := range row {
			col.Width = lengths[i] - (col.PaddingLeft + col.PaddingRight)

			s += "|"
			s += col.String()
		}
		s += "|" + n

		if first {
			first = false

			if st.Header {
				s += string(st.Corner(CornerMiddleVertical)) + frame + string(st.Corner(CornerMiddleVertical)) + n
			}
		}
	}

	s += string(st.Corner(CornerBottomLeft)) + frame + string(st.Corner(CornerBottomRight))

	return s
}
