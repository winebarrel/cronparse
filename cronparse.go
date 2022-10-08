package cronparse

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	cronLexer = lexer.MustSimple([]lexer.SimpleRule{
		{`Number`, `\d+`},
		{`Month`, `(?i)(?:JAN|FEB|MAR|APR|MAY|JUN|JUL|AUG|SEP|OCT|NOV|DEC)`},
		{`Week`, `(?i)(?:MON|TUE|WED|THU|FRI|SAT|SUN)`},
		{`Symbol`, `[,\-\*\?/LW#]`},
		{`SP`, `\s+`},
	})

	Parser = participle.MustBuild[Expression](
		participle.Lexer(cronLexer),
	)
)

type Number struct {
	Value int `@Number`
}

type MonthValue struct {
	Value string `@Month`
}

type WeekValue struct {
	Value string `@Week`
}

type NumberRange struct {
	From int `@Number`
	To   int `"-" @Number`
}

type WeekRange struct {
	From string `@Week`
	To   string `"-" @Week`
}

type MonthRange struct {
	From string `@Month`
	To   string `"-" @Month`
}

type All struct {
	Value struct{} `"*"`
}

type Increment struct {
	Wildcard bool `( @"*"`
	Top      int  `| @Number )`
	Buttom   int  `"/" @Number`
}

type Any struct {
	Value struct{} `"?"`
}

type Last struct {
	Value struct{} `"L"`
}

type Weekday struct {
	Value int `@Number "W"`
}

type Instance struct {
	DayOfWeek    int `@Number`
	NthDayOfWeek int `"#" @Number`
}

type CommonExp struct {
	Increment   *Increment   `@@`
	NumberRange *NumberRange `| @@`
	Number      *Number      `| @@`
	All         *All         `| @@`
}

// minutes
type MinutesExp struct {
	CommonExp
}

type Minutes struct {
	Exps []*MinutesExp `@@ ( "," @@ )*`
}

// hours
type HoursExp struct {
	CommonExp
}

type Hours struct {
	Exps []*HoursExp `@@ ( "," @@ )*`
}

// day of month
type DayOfMonthExp struct {
	Weekday *Weekday `@@ |`
	CommonExp
	Any  *Any  `| @@`
	Last *Last `| @@`
}

type DayOfMonth struct {
	Exps []*DayOfMonthExp `@@ ( "," @@ )*`
}

// month
type MonthExp struct {
	CommonExp
	StringRange *MonthRange `| @@`
	String      *MonthValue `| @@`
	Any         *Any        `| @@`
	Last        *Last       `| @@`
}

type Month struct {
	Exps []*MonthExp `@@ ( "," @@ )*`
}

// day of week
type DayOfWeekExp struct {
	Instance *Instance `@@ |`
	CommonExp
	StringRange *WeekRange `| @@`
	String      *WeekValue `| @@`
	Any         *Any       `| @@`
	Last        *Last      `| @@`
}

type DayOfWeek struct {
	Exps []*DayOfWeekExp `@@ ( "," @@ )*`
}

// year
type YearExp struct {
	CommonExp
}

type Year struct {
	Exps []*YearExp `@@ ( "," @@ )*`
}

type Expression struct {
	Minutes    *Minutes    `@@`
	Hours      *Hours      `SP @@`
	DayOfMonth *DayOfMonth `SP @@`
	Month      *Month      `SP @@`
	DayOfWeek  *DayOfWeek  `SP @@`
	Year       *Year       `SP @@`
}
