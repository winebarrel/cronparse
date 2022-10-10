# cronparse

Cron expression parser for Amazon EventBridge.

## Installation

```sh
go get github.com/winebarrel/cronparse@v1.5.0
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
$ cronnext "*/10 10 ? * MON-FRI *"
Tue, 11 Oct 2022 10:00:00
Tue, 11 Oct 2022 10:10:00
Tue, 11 Oct 2022 10:20:00
Tue, 11 Oct 2022 10:30:00
Tue, 11 Oct 2022 10:40:00
Tue, 11 Oct 2022 10:50:00
Wed, 12 Oct 2022 10:00:00
Wed, 12 Oct 2022 10:10:00
Wed, 12 Oct 2022 10:20:00
Wed, 12 Oct 2022 10:30:00
```

# Related Links

* [Schedule Expressions for Rules - Amazon CloudWatch Events](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
