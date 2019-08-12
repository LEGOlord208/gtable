```
Abandoned.
```

# gtable
[jTable](https://krake.one/table) rewritten in **Go**!  

### Used
Here are some examples of where gtable is used.

- [DiscordConsole](https://github.com/LEGOlord208/DiscordConsole) uses gtable to show a bunch of values nicely.
- [CmdTable](https://github.com/LEGOlord208/CmdTable) uses gtable to generate a table based on parameters.

## Usage
Check the [GoDoc](https://godoc.org/github.com/LEGOlord208/gtable)  
and an example:
```Go
package main;

import (
	"fmt"
	"github.com/jD91mZM2/gtable"
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
I think `go get github.com/jD91mZM2/gtable` should do it...  

# Have fun!
