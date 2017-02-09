# gTable
[jTable](https://krake.one/table) rewritten in **Go**!  

## Examples
```Go
package main;

import (
	"fmt"
	"github.com/legolord208/gtable"
)

func main(){
	table := gtable.NewStringTable();
	table.AddStrings("This", "is");
	table.AddRow();
	table.AddStrings("a", "test", "lol");
	table.Each(func(t *gtable.TableItem){
		t.Padding(1);
		t.Center = true;
	});

	fmt.Println(table.String());
}
```

## Installing
I think `go get github.com/legolord208/gtable` should do it...  

# Have fun!
