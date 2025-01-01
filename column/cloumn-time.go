package column

type YearType struct{}

func (t *YearType) ID() Type       { return YEAR }
func (t *YearType) Name() string   { return "year" }
func (t *YearType) String() string { return "year" }

type DateType struct{}

func (t *DateType) ID() Type       { return DATE }
func (t *DateType) Name() string   { return "date" }
func (t *DateType) String() string { return "date" }

type DateTimeType struct{}

func (t *DateTimeType) ID() Type       { return DATETIME }
func (t *DateTimeType) Name() string   { return "datetime" }
func (t *DateTimeType) String() string { return "datetime" }

type TimeType struct{ WithZone bool }

func (t *TimeType) ID() Type       { return TIME }
func (t *TimeType) Name() string   { return "time" }
func (t *TimeType) String() string { return "time" }

type TimestampType struct {
	Len      int
	WithZone bool
	Local    bool
}

func (t *TimestampType) ID() Type       { return TIMESTAMP }
func (t *TimestampType) Name() string   { return "timestamp" }
func (t *TimestampType) String() string { return "timestamp" }

func (t *TimestampType) Length() int { return t.Len }

func (t *TimestampType) Zone() bool { return t.WithZone }

type IntervalYearToMonth struct {
	YearPrecision  int
	MonthPrecision int
}

func (t *IntervalYearToMonth) ID() Type       { return INTERVAL_YEAR_MONTHS }
func (t *IntervalYearToMonth) Name() string   { return "year_to_month" }
func (t *IntervalYearToMonth) String() string { return "year_to_month" }

type IntervalDayToSecond struct {
	DayPrecision    int
	SecondPrecision int
}

func (t *IntervalDayToSecond) ID() Type       { return INTERVAL_DAY_TIME }
func (t *IntervalDayToSecond) Name() string   { return "day_to_time" }
func (t *IntervalDayToSecond) String() string { return "day_to_time" }
