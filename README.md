# cronparse

Cron expression parser.

## Usage

```go
package main

import (
	"fmt"

	"github.com/winebarrel/cronparse"
)

func main() {
	cron, err := cronparse.Parser.ParseString("", "0 10 * * ? *")

	if err != nil {
		panic(err)
	}

	fmt.Println(cron.Minutes.Exps[0].Number.Value) //=> 0
	fmt.Println(cron.Hours.Exps[0].Number.Value)   //=> 10
}
```
