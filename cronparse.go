package cronparse

import (
	"fmt"
	"strings"
	"time"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/winebarrel/cronparse/utils"
)

var (
	cronLexer = lexer.MustSimple([]lexer.SimpleRule{
		{`Number`, `\d+`},
		{`Month`, `(?i)(?:` + strings.Join(utils.MonthNames, "|") + `)`},
		{`Week`, `(?i)(?:` + strings.Join(utils.WeekNames, "|") + `)`},
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

func (v *CommonExp) Match(x int) bool {
	if v.Increment != nil {
		return v.Increment.Match(x)
	} else if v.NumberRange != nil {
		return v.NumberRange.Match(x)
	} else if v.Number != nil {
		return v.Number.Match(x)
	} else if v.All != nil {
		return v.All.Match(x)
	}

	return false
}

// minutes
type MinutesExp struct {
	CommonExp
}

func (v *MinutesExp) String() string {
	return v.CommonExp.String()
}

func (v *MinutesExp) Match(t time.Time) bool {
	return v.CommonExp.Match(t.Minute())
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

func (v *Minutes) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
}

// hours
type HoursExp struct {
	CommonExp
}

func (v *HoursExp) String() string {
	return v.CommonExp.String()
}

func (v *HoursExp) Match(t time.Time) bool {
	return v.CommonExp.Match(t.Hour())
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

func (v *Hours) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
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

func (v *DayOfMonthExp) Match(t time.Time) bool {
	if v.CommonExp.Present() {
		return v.CommonExp.Match(t.Day())
	} else if v.Weekday != nil {
		return v.Weekday.Match(t)
	} else if v.Any != nil {
		return v.Any.Match(t.Day())
	} else if v.Last != nil {
		return utils.LastOfMonth(t) == t.Day()
	}

	return false
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

func (v *DayOfMonth) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
}

// month
type MonthExp struct {
	CommonExp
	NameRange *MonthRange `| @@`
	Name      *MonthName  `| @@`
	Any       *Any        `| @@`
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
	}

	return ""
}

func (v *MonthExp) Match(t time.Time) bool {
	if v.CommonExp.Present() {
		return v.CommonExp.Match(int(t.Month()))
	} else if v.NameRange != nil {
		return v.NameRange.Match(t.Month())
	} else if v.Name != nil {
		return v.Name.Match(t.Month())
	} else if v.Any != nil {
		return v.Any.Match(t.Month())
	}

	return false
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

func (v *Month) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
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

func (v *DayOfWeekExp) Match(t time.Time) bool {
	if v.CommonExp.Present() {
		wday := int(t.Weekday())

		if wday == 0 {
			wday = 7
		}

		return v.CommonExp.Match(wday)
	} else if v.Instance != nil {
		return v.Instance.Match(t)
	} else if v.NameRange != nil {
		return v.NameRange.Match(t.Weekday())
	} else if v.Name != nil {
		return v.Name.Match(t.Weekday())
	} else if v.Any != nil {
		return v.Any.Match(t.Weekday())
	} else if v.Last != nil {
		return t.Weekday() == time.Saturday
	}

	return false
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

func (v *DayOfWeek) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
}

// year
type YearExp struct {
	CommonExp
}

func (v *YearExp) String() string {
	return v.CommonExp.String()
}

func (v *YearExp) Match(t time.Time) bool {
	return v.CommonExp.Match(t.Hour())
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

func (v *Year) Match(t time.Time) bool {
	for _, e := range v.Exps {
		if e.Match(t) {
			return true
		}
	}

	return false
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

func (v *Expression) Match(t time.Time) bool {
	return v.Minutes.Match(t) &&
		v.Hours.Match(t) &&
		v.DayOfMonth.Match(t) &&
		v.Month.Match(t) &&
		v.DayOfWeek.Match(t) &&
		v.Year.Match(t)
}
