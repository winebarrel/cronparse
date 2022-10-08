package cronparse

import (
	"fmt"
	"strings"

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

type CommonExp struct {
	Increment   *Increment   `@@`
	NumberRange *NumberRange `| @@`
	Number      *Number      `| @@`
	All         *All         `| @@`
}

func (v *CommonExp) String() string {
	if v.Increment != nil {
		return v.Increment.String()
	} else if v.NumberRange != nil {
		return v.NumberRange.String()
	} else if v.Number != nil {
		return v.Number.String()
	} else if v.All != nil {
		return v.All.String()
	}

	return ""
}

func (v *CommonExp) Present() bool {
	return v.Increment != nil || v.NumberRange != nil || v.Number != nil || v.All != nil
}

// minutes
type MinutesExp struct {
	CommonExp
}

func (v *MinutesExp) String() string {
	return v.CommonExp.String()
}

type Minutes struct {
	Exps []*MinutesExp `@@ ( "," @@ )*`
}

func (v *Minutes) String() string {
	strs := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		strs = append(strs, e.String())
	}

	return strings.Join(strs, ",")
}

// hours
type HoursExp struct {
	CommonExp
}

func (v *HoursExp) String() string {
	return v.CommonExp.String()
}

type Hours struct {
	Exps []*HoursExp `@@ ( "," @@ )*`
}

func (v *Hours) String() string {
	strs := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		strs = append(strs, e.String())
	}

	return strings.Join(strs, ",")
}

// day of month
type DayOfMonthExp struct {
	Weekday *Weekday `@@ |`
	CommonExp
	Any  *Any  `| @@`
	Last *Last `| @@`
}

func (v *DayOfMonthExp) String() string {
	if v.CommonExp.Present() {
		return v.CommonExp.String()
	} else if v.Weekday != nil {
		return v.Weekday.String()
	} else if v.Any != nil {
		return v.Any.String()
	} else if v.Last != nil {
		return v.Last.String()
	}

	return ""
}

type DayOfMonth struct {
	Exps []*DayOfMonthExp `@@ ( "," @@ )*`
}

func (v *DayOfMonth) String() string {
	strs := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		strs = append(strs, e.String())
	}

	return strings.Join(strs, ",")
}

// month
type MonthExp struct {
	CommonExp
	NameRange *MonthRange `| @@`
	Name      *MonthName  `| @@`
	Any       *Any        `| @@`
	Last      *Last       `| @@`
}

func (v *MonthExp) String() string {
	if v.CommonExp.Present() {
		return v.CommonExp.String()
	} else if v.NameRange != nil {
		return v.NameRange.String()
	} else if v.Name != nil {
		return v.Name.String()
	} else if v.Any != nil {
		return v.Any.String()
	} else if v.Last != nil {
		return v.Last.String()
	}

	return ""
}

type Month struct {
	Exps []*MonthExp `@@ ( "," @@ )*`
}

func (v *Month) String() string {
	strs := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		strs = append(strs, e.String())
	}

	return strings.Join(strs, ",")
}

// day of week
type DayOfWeekExp struct {
	Instance *Instance `@@ |`
	CommonExp
	NameRange *WeekRange `| @@`
	Name      *WeekName  `| @@`
	Any       *Any       `| @@`
	Last      *Last      `| @@`
}

func (v *DayOfWeekExp) String() string {
	if v.CommonExp.Present() {
		return v.CommonExp.String()
	} else if v.Instance != nil {
		return v.Instance.String()
	} else if v.NameRange != nil {
		return v.NameRange.String()
	} else if v.Name != nil {
		return v.Name.String()
	} else if v.Any != nil {
		return v.Any.String()
	} else if v.Last != nil {
		return v.Last.String()
	}

	return ""
}

type DayOfWeek struct {
	Exps []*DayOfWeekExp `@@ ( "," @@ )*`
}

func (v *DayOfWeek) String() string {
	strs := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		strs = append(strs, e.String())
	}

	return strings.Join(strs, ",")
}

// year
type YearExp struct {
	CommonExp
}

func (v *YearExp) String() string {
	return v.CommonExp.String()
}

type Year struct {
	Exps []*YearExp `@@ ( "," @@ )*`
}

func (v *Year) String() string {
	strs := make([]string, 0, len(v.Exps))

	for _, e := range v.Exps {
		strs = append(strs, e.String())
	}

	return strings.Join(strs, ",")
}

type Expression struct {
	Minutes    *Minutes    `@@`
	Hours      *Hours      `SP @@`
	DayOfMonth *DayOfMonth `SP @@`
	Month      *Month      `SP @@`
	DayOfWeek  *DayOfWeek  `SP @@`
	Year       *Year       `SP @@`
}

func (v *Expression) String() string {
	return fmt.Sprintf("%s %s %s %s %s %s",
		v.Minutes.String(),
		v.Hours.String(),
		v.DayOfMonth.String(),
		v.Month.String(),
		v.DayOfWeek.String(),
		v.Year.String(),
	)
}
