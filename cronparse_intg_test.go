package cronparse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronparse"
)

func TestIntegration(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		exp string
		ast *cronparse.Expression
	}{
		// https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html
		{
			"0 10 * * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 10,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"15 12 * * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 15,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 12,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"0 18 ? * MON-FRI *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 18,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							NameRange: &cronparse.WeekRange{
								From: "MON",
								To:   "FRI",
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"0 8 1 * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 8,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 1,
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"0/15 * * * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Increment: &cronparse.Increment{
									Buttom: 15,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"0/10 * ? * MON-FRI *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Increment: &cronparse.Increment{
									Buttom: 10,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							NameRange: &cronparse.WeekRange{
								From: "MON",
								To:   "FRI",
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"0/5 8-17 ? * MON-FRI *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Increment: &cronparse.Increment{
									Buttom: 5,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								NumberRange: &cronparse.NumberRange{
									From: 8,
									To:   17,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							NameRange: &cronparse.WeekRange{
								From: "MON",
								To:   "FRI",
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		// https://docs.oracle.com/cd/E12058_01/doc/doc.1014/e12030/cron_expression.htm
		{
			"0 12 * * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 12,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"15 10 ? * * *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 15,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 10,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"15 10 * * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 15,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 10,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"15 10 * * ? 2005",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 15,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 10,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 2005,
								},
							},
						},
					},
				},
			},
		},
		{
			"* 14 * * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 14,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"0/5 14 * * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Increment: &cronparse.Increment{
									Buttom: 5,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 14,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"0/5 14,18 * * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Increment: &cronparse.Increment{
									Buttom: 5,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 14,
								},
							},
						},
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 18,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"0-5 14 * * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								NumberRange: &cronparse.NumberRange{
									To: 5,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 14,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"10,44 14 ? 3 WED *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 10,
								},
							},
						},
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 44,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 14,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 3,
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Name: &cronparse.WeekName{
								Value: "WED",
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"15 10 ? * MON-FRI *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 15,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 10,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							NameRange: &cronparse.WeekRange{
								From: "MON",
								To:   "FRI",
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"15 10 15 * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 15,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 10,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 15,
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"15 10 L * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 15,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 10,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{},
							Last: &cronparse.LastOfMonth{
								Value: struct{}{},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"15 10 ? * 6#3 *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 15,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 10,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							Instance: &cronparse.Instance{
								DayOfWeek:    6,
								NthDayOfWeek: 3,
							},
							CommonExp: cronparse.CommonExp{},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"0 12 1/5 * ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 12,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								Increment: &cronparse.Increment{
									Top:    1,
									Buttom: 5,
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
		{
			"11 11 11 11 ? *",
			&cronparse.Expression{
				Minutes: &cronparse.Minutes{
					Exps: []*cronparse.MinutesExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 11,
								},
							},
						},
					},
				},
				Hours: &cronparse.Hours{
					Exps: []*cronparse.HoursExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 11,
								},
							},
						},
					},
				},
				DayOfMonth: &cronparse.DayOfMonth{
					Exps: []*cronparse.DayOfMonthExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 11,
								},
							},
						},
					},
				},
				Month: &cronparse.Month{
					Exps: []*cronparse.MonthExp{
						{
							CommonExp: cronparse.CommonExp{
								Number: &cronparse.Number{
									Value: 11,
								},
							},
						},
					},
				},
				DayOfWeek: &cronparse.DayOfWeek{
					Exps: []*cronparse.DayOfWeekExp{
						{
							CommonExp: cronparse.CommonExp{},
							Any: &cronparse.Any{
								Value: struct{}{},
							},
						},
					},
				},
				Year: &cronparse.Year{
					Exps: []*cronparse.YearExp{
						{
							CommonExp: cronparse.CommonExp{
								All: &cronparse.All{
									Value: struct{}{},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, t := range tt {
		cron, err := cronparse.Parser.ParseString("", t.exp)
		assert.NoError(err)
		assert.Equal(cron, t.ast, t.exp)
		assert.Equal(t.exp, t.ast.String(), t.exp)
	}
}
