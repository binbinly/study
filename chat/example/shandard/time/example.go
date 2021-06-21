package main

import (
	"fmt"
	"time"
)

func main()  {

	time.Now()

	time.Unix(1e9, 0)

	loc, _ := time.LoadLocation("Europe/Berlin")

	time.FixedZone("Beijing Time", int((8 * time.Second).Seconds()))

	t := time.Date(2019, time.February, 7, 0, 0, 0, 0, time.UTC)

	t.Location()

	t.Local()

	t.Unix()

	t.UTC()

	t.Zone()

	t.IsZero()

	t.In(loc)

	t.UnixNano()

	t.Equal(time.Now())

	t.Before(time.Now())

	t.After(time.Now())

	t.Date()

	t.Clock()

	t.Year()

	t.YearDay()

	t.Month()

	t.Day()

	t.Weekday()

	t.ISOWeek()

	t.Hour()

	t.Minute()

	t.Second()

	t.Nanosecond()

	t.Add(time.Hour * 2)

	t.AddDate(1, 2, 3)

	t.Sub(time.Now())

	t.Round(time.Hour)

	t.Truncate(time.Hour)

	t.Format("2006-01-02 15:04:05")

	t.String()

	gobBytes, _ := t.GobEncode()

	_ = t.GobDecode(gobBytes)

	binraryBytes, _ := t.MarshalBinary()

	_ = t.UnmarshalBinary(binraryBytes)

	jsonBytes, _ := t.MarshalJSON()

	_ = t.UnmarshalJSON(jsonBytes)

	textBytes, _ := t.MarshalText()

	_ = t.UnmarshalText(textBytes)

	time.Parse("2006 Jan 02 15:04:05", "2019 Feb 07 12:15:30.918273645")

	time.ParseInLocation("2006 Jan 02 15:04:05", "2019 Feb 07 12:15:30.918273645", loc)

	d, _ := time.ParseDuration("1h15m30.918273645s")
	fmt.Println(d.Hours(), d.Minutes(), d.Nanoseconds())

	time.Since(t)

	time.Sleep(time.Second * 2)

	time.NewTimer(time.Minute * 2)

	time.After(time.Minute * 2)

	time.AfterFunc(time.Millisecond*1000, func() {
		fmt.Println("time after 1 second test")
	})

	time.NewTicker(time.Second * 10)

	time.Tick(time.Second * 10)
}