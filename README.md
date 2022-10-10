# cronparse

Cron expression parser for Amazon EventBridge.

## Installation

```sh
go get github.com/winebarrel/cronparse@v1.3.0
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/winebarrel/cronparse"
)

func main() {
	cron, err := cronparse.Parse("0 10 * * ? *")

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

	fmt.Println(cron.Next(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC)))
	//=> 2022-11-03 10:00:00 +0000 UTC
	fmt.Println(cron.Next(time.Date(2022, 11, 3, 11, 0, 0, 0, time.UTC)))
	//=> 2022-11-04 10:00:00 +0000 UTC
	fmt.Println(cron.NextN(time.Date(2022, 11, 3, 10, 0, 0, 0, time.UTC), 3))
	//=> [2022-11-03 10:00:00 +0000 UTC 2022-11-04 10:00:00 +0000 UTC 2022-11-05 10:00:00 +0000 UTC]
}
```

# cronnext

CLI to show next triggers.

## Installation

```
brew install winebarrel/cronnext/cronnext
```

## Usage

```
Usage: cronnext [OPTION] CRON_EXPR
  -n int
    	number of next triggers (default 10)
  -version
    	print version and exit
```

```
$ cronnext "0 10 * * ? *"
2022-10-11 10:00:00 +0900 JST
2022-10-12 10:00:00 +0900 JST
2022-10-13 10:00:00 +0900 JST
2022-10-14 10:00:00 +0900 JST
2022-10-15 10:00:00 +0900 JST
2022-10-16 10:00:00 +0900 JST
2022-10-17 10:00:00 +0900 JST
2022-10-18 10:00:00 +0900 JST
2022-10-19 10:00:00 +0900 JST
2022-10-20 10:00:00 +0900 JST
```

# Related Links

* [Schedule Expressions for Rules - Amazon CloudWatch Events](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
