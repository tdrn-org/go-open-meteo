//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package openmeteo

import (
	"fmt"
	"time"

	"github.com/tdrn-org/go-open-meteo/api"
)

func GetTime(t any, tzName string) (time.Time, error) {
	location, err := time.LoadLocation(tzName)
	if err != nil {
		return time.Time{}, fmt.Errorf("unable to resolve timezone %q (cause: %w)", tzName, err)
	}
	switch t := t.(type) {
	case *api.GetV1Forecast200JSONResponseBody_Current_Time:
		return getTime(t.AsGetV1Forecast200JSONResponseBodyCurrentTime0, t.AsGetV1Forecast200JSONResponseBodyCurrentTime1, location)
	case *api.GetV1Forecast200JSONResponseBody_Daily_Sunrise_Item:
		return getTime(t.AsGetV1Forecast200JSONResponseBodyDailySunrise0, t.AsGetV1Forecast200JSONResponseBodyDailySunrise1, location)
	case *api.GetV1Forecast200JSONResponseBody_Daily_Sunset_Item:
		return getTime(t.AsGetV1Forecast200JSONResponseBodyDailySunset0, t.AsGetV1Forecast200JSONResponseBodyDailySunset1, location)
	case *api.GetV1Forecast200JSONResponseBody_Daily_Time_Item:
		return getTime(t.AsGetV1Forecast200JSONResponseBodyDailyTime0, t.AsGetV1Forecast200JSONResponseBodyDailyTime1, location)
	case *api.GetV1Forecast200JSONResponseBody_Hourly_Time_Item:
		return getTime(t.AsGetV1Forecast200JSONResponseBodyHourlyTime0, t.AsGetV1Forecast200JSONResponseBodyHourlyTime1, location)
	case *api.GetV1Forecast200JSONResponseBody_Minutely15_Time_Item:
		return getTime(t.AsGetV1Forecast200JSONResponseBodyMinutely15Time0, t.AsGetV1Forecast200JSONResponseBodyMinutely15Time1, location)
	default:
		return time.Time{}, fmt.Errorf("unrecognized time type %T", t)
	}
}

func getTime(f0 func() (string, error), f1 func() (int, error), location *time.Location) (time.Time, error) {
	t0, err0 := f0()
	if err0 == nil {
		t, err := time.Parse("2006-01-02T15:04", t0)
		if err != nil {
			return time.Time{}, err
		}
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), location), nil
	}
	t1, err1 := f1()
	if err1 != nil {
		return time.Time{}, err1
	}
	return time.Unix(int64(t1), 0).In(location), nil
}
