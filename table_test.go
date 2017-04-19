package gtable_test

import (
	"fmt"

	"github.com/legolord208/gtable"
)

func Example() {
	table := gtable.NewStringTable()
	table.AddStrings("Hello", "World")
	table.AddRow()
	table.AddStrings("Testing", "123")
	fmt.Println(table.String())
	// Output:
	//+-------+-----+
	//|Hello  |World|
	//+-------+-----+
	//|Testing|123  |
	//+-------+-----+
}
