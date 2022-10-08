package main

import (
	"fmt"
	"time"

	"github.com/winebarrel/cronparse"
)

func main() {
	cron, err := cronparse.Parser.ParseString("", "0 10 * * ? *")

	if err != nil {
		panic(err)
	}

	fmt.Println(cron.Minutes.Exps[0].Number.Value) //=> 0
	fmt.Println(cron.Hours.Exps[0].Number.Value)   //=> 10
	fmt.Println(cron.String())                     //=> "0 10 * * ? *"

	fmt.Println(cron.Match(time.Date(2022, 11, 3, 9, 0, 0, 0, time.UTC)))
	// => false

	fmt.Println(cron.Match(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC)))
	// => true
}
