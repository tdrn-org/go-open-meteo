//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package openmeteo_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	openmeteo "github.com/tdrn-org/go-open-meteo"
	"github.com/tdrn-org/go-open-meteo/api"
)

const testTimezone string = "Europe/Berlin"
const testTime string = "2006-01-02T15:04"

func TestGetTimeISO8601(t *testing.T) {
	location, err := time.LoadLocation(testTimezone)
	require.NoError(t, err)
	now := time.Now().In(location).Truncate(time.Minute)
	responseTime := now.Format(testTime)
	currentTime := &api.GetV1Forecast200JSONResponseBody_Current_Time{}
	require.NoError(t, currentTime.FromGetV1Forecast200JSONResponseBodyCurrentTime0(responseTime))
	dailySunrise := &api.GetV1Forecast200JSONResponseBody_Daily_Sunrise_Item{}
	require.NoError(t, dailySunrise.FromGetV1Forecast200JSONResponseBodyDailySunrise0(responseTime))
	dailySunset := &api.GetV1Forecast200JSONResponseBody_Daily_Sunset_Item{}
	require.NoError(t, dailySunset.FromGetV1Forecast200JSONResponseBodyDailySunset0(responseTime))
	dailyTime := &api.GetV1Forecast200JSONResponseBody_Daily_Time_Item{}
	require.NoError(t, dailyTime.FromGetV1Forecast200JSONResponseBodyDailyTime0(responseTime))
	hourlyTime := &api.GetV1Forecast200JSONResponseBody_Hourly_Time_Item{}
	require.NoError(t, hourlyTime.FromGetV1Forecast200JSONResponseBodyHourlyTime0(responseTime))
	minutely15Time := &api.GetV1Forecast200JSONResponseBody_Minutely15_Time_Item{}
	require.NoError(t, minutely15Time.FromGetV1Forecast200JSONResponseBodyMinutely15Time0(responseTime))
	times := []any{
		currentTime,
		dailySunrise,
		dailySunset,
		dailyTime,
		hourlyTime,
		minutely15Time,
	}
	for _, time := range times {
		t.Run(fmt.Sprintf("%T", time), func(t *testing.T) {
			decodedTime, err := openmeteo.GetTime(time, testTimezone)
			require.NoError(t, err)
			require.Equal(t, now, decodedTime)
		})
	}
}

func TestGetTimeUnix(t *testing.T) {
	location, err := time.LoadLocation(testTimezone)
	require.NoError(t, err)
	now := time.Now().In(location).Truncate(time.Minute)
	responseTime := int(now.Unix())
	currentTime := &api.GetV1Forecast200JSONResponseBody_Current_Time{}
	require.NoError(t, currentTime.FromGetV1Forecast200JSONResponseBodyCurrentTime1(responseTime))
	dailySunrise := &api.GetV1Forecast200JSONResponseBody_Daily_Sunrise_Item{}
	require.NoError(t, dailySunrise.FromGetV1Forecast200JSONResponseBodyDailySunrise1(responseTime))
	dailySunset := &api.GetV1Forecast200JSONResponseBody_Daily_Sunset_Item{}
	require.NoError(t, dailySunset.FromGetV1Forecast200JSONResponseBodyDailySunset1(responseTime))
	dailyTime := &api.GetV1Forecast200JSONResponseBody_Daily_Time_Item{}
	require.NoError(t, dailyTime.FromGetV1Forecast200JSONResponseBodyDailyTime1(responseTime))
	hourlyTime := &api.GetV1Forecast200JSONResponseBody_Hourly_Time_Item{}
	require.NoError(t, hourlyTime.FromGetV1Forecast200JSONResponseBodyHourlyTime1(responseTime))
	minutely15Time := &api.GetV1Forecast200JSONResponseBody_Minutely15_Time_Item{}
	require.NoError(t, minutely15Time.FromGetV1Forecast200JSONResponseBodyMinutely15Time1(responseTime))
	times := []any{
		currentTime,
		dailySunrise,
		dailySunset,
		dailyTime,
		hourlyTime,
		minutely15Time,
	}
	for _, time := range times {
		t.Run(fmt.Sprintf("%T", time), func(t *testing.T) {
			decodedTime, err := openmeteo.GetTime(time, testTimezone)
			require.NoError(t, err)
			require.Equal(t, now, decodedTime)
		})
	}
}
