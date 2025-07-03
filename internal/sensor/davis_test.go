package sensor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	mocksensor "github.com/emiliogozo/panahon-api-go/internal/mocks"
	"github.com/emiliogozo/panahon-api-go/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNewDavis(t *testing.T) {
	testCases := []struct {
		name           string
		apiCredentials DavisAPICredentials
		checkInstance  func(sensor *Davis, err error)
	}{
		{
			name: "Default",
			apiCredentials: DavisAPICredentials{
				Type:     "v1",
				User:     "testdavisUser",
				Pass:     "testdav!sPAss",
				APIToken: "123DAV15890456ZXCARSLUY",
			},
			checkInstance: func(sensor *Davis, err error) {
				require.NotNil(t, sensor)
				require.NoError(t, err)
				require.Equal(t, "testdavisUser", sensor.apiCredentials.User)
				require.Equal(t, "testdav!sPAss", sensor.apiCredentials.Pass)
				require.Equal(t, "123DAV15890456ZXCARSLUY", sensor.apiCredentials.APIToken)
			},
		},
		{
			name: "V2",
			apiCredentials: DavisAPICredentials{
				Type:      "v2",
				APIKey:    "123DAV15890456ZXCARSLUY",
				APISecret: "xxxx123DAV15890456ZXCARSLUYxxxx",
			},
			checkInstance: func(sensor *Davis, err error) {
				require.NotNil(t, sensor)
				require.NoError(t, err)
				require.Equal(t, "123DAV15890456ZXCARSLUY", sensor.apiCredentials.APIKey)
				require.Equal(t, "xxxx123DAV15890456ZXCARSLUYxxxx", sensor.apiCredentials.APISecret)
			},
		},
		{
			name: "Dashboard",
			apiCredentials: DavisAPICredentials{
				Type:    "dashboard",
				StnUUID: "a72efc82c04d43b6801d7d4de46aa79a",
			},
			checkInstance: func(sensor *Davis, err error) {
				require.NotNil(t, sensor)
				require.NoError(t, err)
				require.Equal(t, "a72efc82c04d43b6801d7d4de46aa79a", sensor.apiCredentials.StnUUID)
			},
		},
		{
			name: "NoUser",
			apiCredentials: DavisAPICredentials{
				Type:     "v1",
				Pass:     "testdav!sPAss",
				APIToken: "123DAV15890456ZXCARSLUY",
			},
			checkInstance: func(sensor *Davis, err error) {
				require.Nil(t, sensor)
				require.Error(t, err)
			},
		},
		{
			name: "NoPass",
			apiCredentials: DavisAPICredentials{
				Type:     "v1",
				User:     "testdavisUser",
				APIToken: "123DAV15890456ZXCARSLUY",
			},
			checkInstance: func(sensor *Davis, err error) {
				require.Nil(t, sensor)
				require.Error(t, err)
			},
		},
		{
			name: "NoAPIToken",
			apiCredentials: DavisAPICredentials{
				Type: "v1",
				User: "testdavisUser",
				Pass: "testdav!sPAss",
			},
			checkInstance: func(sensor *Davis, err error) {
				require.Nil(t, sensor)
				require.Error(t, err)
			},
		},
		{
			name: "NoAPIKey",
			apiCredentials: DavisAPICredentials{
				Type:      "v2",
				APISecret: "xxxx123DAV15890456ZXCARSLUYxxxx",
			},
			checkInstance: func(sensor *Davis, err error) {
				require.Nil(t, sensor)
				require.Error(t, err)
			},
		},
		{
			name: "NoAPISecret",
			apiCredentials: DavisAPICredentials{
				Type:   "v2",
				APIKey: "123DAV15890456ZXCARSLUY",
			},
			checkInstance: func(sensor *Davis, err error) {
				require.Nil(t, sensor)
				require.Error(t, err)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			testSensor, err := NewDavis(tc.apiCredentials, 10)
			tc.checkInstance(testSensor, err)
		})
	}
}

func TestFetchLatest(t *testing.T) {
	testCases := []struct {
		name           string
		apiCredentials DavisAPICredentials
		builStubs      func(client *mocksensor.MockFetcher) []davisRawCurrentResponse
		checkResponse  func(client *mocksensor.MockFetcher, rawObsSlice []davisRawCurrentResponse, obsSlice []DavisCurrentObservation, err error)
	}{
		{
			name: "Default",
			apiCredentials: DavisAPICredentials{
				Type:     "v1",
				User:     "testuser001",
				Pass:     "secrEtp@s$",
				APIToken: "qwfparst1234ar655",
			},
			builStubs: func(client *mocksensor.MockFetcher) []davisRawCurrentResponse {
				rawObs := randomDavisRawResponseV1()
				body, _ := json.Marshal(rawObs)
				bodyReader := bytes.NewReader(body)
				client.EXPECT().Do(mock.MatchedBy(func(req *http.Request) bool { return req.Header.Get("User-Agent") != "" })).Return(&http.Response{
					Body: io.NopCloser(bodyReader),
				}, nil)
				return []davisRawCurrentResponse{rawObs}
			},
			checkResponse: func(client *mocksensor.MockFetcher, rawObsSlice []davisRawCurrentResponse, obsSlice []DavisCurrentObservation, err error) {
				client.AssertExpectations(t)
				assert.NoError(t, err)
				requireDavisEqual(t, rawObsSlice, obsSlice)
			},
		},
		{
			name: "V2",
			apiCredentials: DavisAPICredentials{
				Type:      "v2",
				APIKey:    "123DAV15890456ZXCARSLUY",
				APISecret: "xxxx123DAV15890456ZXCARSLUYxxxx",
			},
			builStubs: func(client *mocksensor.MockFetcher) []davisRawCurrentResponse {
				rawStns := davisRawStationsResponseV2{
					Stations: []davisRawStationResponseV2{
						{
							StationID: 111,
						},
						{
							StationID: 121,
						},
						{
							StationID: 511,
						},
					},
				}
				body, _ := json.Marshal(rawStns)
				bodyReader := bytes.NewReader(body)
				client.EXPECT().Do(mock.MatchedBy(func(req *http.Request) bool { return req.Header.Get("User-Agent") != "" })).Return(&http.Response{
					Body: io.NopCloser(bodyReader),
				}, nil).Once()
				rawObsSlice := make([]davisRawCurrentResponse, 0)
				for range rawStns.Stations {
					rawObs := randomDavisRawResponseV2()
					rawObsSlice = append(rawObsSlice, rawObs)
					body, _ := json.Marshal(rawObs)
					bodyReader := bytes.NewReader(body)
					client.EXPECT().Do(mock.MatchedBy(func(req *http.Request) bool { return req.Header.Get("User-Agent") != "" })).Return(&http.Response{
						Body: io.NopCloser(bodyReader),
					}, nil).Once()
				}
				return rawObsSlice
			},
			checkResponse: func(client *mocksensor.MockFetcher, rawObsSlice []davisRawCurrentResponse, obsSlice []DavisCurrentObservation, err error) {
				client.AssertExpectations(t)
				assert.NoError(t, err)
				requireDavisEqual(t, rawObsSlice, obsSlice)
			},
		},
		{
			name: "Dashboard",
			apiCredentials: DavisAPICredentials{
				Type:    "dashboard",
				StnUUID: "a72efc82c04d43b6801d7d4de46aa79a",
			},
			builStubs: func(client *mocksensor.MockFetcher) []davisRawCurrentResponse {
				rawObs := randomDavisRawWeatherDataResponseDashboard()
				body, _ := json.Marshal(rawObs)
				bodyReader := bytes.NewReader(body)
				client.EXPECT().Do(mock.MatchedBy(func(req *http.Request) bool { return req.Header.Get("User-Agent") != "" })).Return(&http.Response{
					Body: io.NopCloser(bodyReader),
				}, nil)
				return []davisRawCurrentResponse{rawObs}
			},
			checkResponse: func(client *mocksensor.MockFetcher, rawObsSlice []davisRawCurrentResponse, obsSlice []DavisCurrentObservation, err error) {
				client.AssertExpectations(t)
				assert.NoError(t, err)
				requireDavisEqual(t, rawObsSlice, obsSlice)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			testFetcher := mocksensor.NewMockFetcher(t)
			testSensor := &Davis{
				apiCredentials: tc.apiCredentials,
				client:         testFetcher,
				sleep:          0,
			}

			rawObsSlice := tc.builStubs(testFetcher)
			obs, err := testSensor.FetchLatest()

			tc.checkResponse(testFetcher, rawObsSlice, obs, err)
		})
	}
}

func requireDavisEqual(t *testing.T, rawObsSlice []davisRawCurrentResponse, obsSlice []DavisCurrentObservation) {
	for i, rawObs := range rawObsSlice {
		switch v := rawObs.(type) {
		case davisRawCurrentResponseV1:
			requireDavisEqualV1(t, v, obsSlice[i])
		case davisRawCurrentResponseV2:
			requireDavisEqualV2(t, v, obsSlice[i])
		case davisRawWeatherDataResponseDashboard:
			requireDavisEqualDashboard(t, v, obsSlice[i])
		}
	}
}

func requireDavisEqualV1(t *testing.T, rawObs davisRawCurrentResponseV1, obs DavisCurrentObservation) {
	if rawObs.PressureMb.Valid {
		require.InDelta(t, rawObs.PressureMb.Value, obs.Pres.Float32, 0.001)
	}
	if rawObs.Rh.Valid {
		require.InDelta(t, rawObs.Rh.Value, obs.Rh.Float32, 0.001)
	}
	if rawObs.TempC.Valid {
		require.InDelta(t, rawObs.TempC.Value, obs.Temp.Float32, 0.001)
	}
	if rawObs.WindDeg.Valid {
		require.InDelta(t, rawObs.WindDeg.Value, obs.Wdir.Float32, 0.001)
	}
	if rawObs.WindMPH.Valid {
		require.InDelta(t, rawObs.WindMPH.Value*0.44704, obs.Wspd.Float32, 0.001)
	}
	if rawObs.Obs.WindDayHighMPH.Valid {
		require.InDelta(t, rawObs.Obs.WindDayHighMPH.Value*0.44704, obs.Wspdx.Float32, 0.001)
	}
	require.Equal(t, rawObs.Obs.TempDayHighTime, datetimeToTimeStr(obs.TxTimestamp.Time))
	require.Equal(t, rawObs.Obs.TempDayLowTime, datetimeToTimeStr(obs.TnTimestamp.Time))
	require.Equal(t, rawObs.Obs.WindDayHighTime, datetimeToTimeStr(obs.GustTimestamp.Time))
	require.Equal(t, rawObs.Time, obs.Timestamp.Time.Format("Mon, 02 Jan 2006 15:04:05 -0700"))
}

func requireDavisEqualV2(t *testing.T, rawObs davisRawCurrentResponseV2, obs DavisCurrentObservation) {
	rawObsData := rawObs.Sensors[0].Data[0]
	if rawObsData.Bar != nil {
		require.InDelta(t, util.InHgToMbar(*rawObsData.Bar), obs.Pres.Float32, 0.001)
	}
	if rawObsData.HumOut != nil {
		require.InDelta(t, *rawObsData.HumOut, obs.Rh.Float32, 0.001)
	}
	if rawObsData.TempOut != nil {
		require.InDelta(t, util.FahrenheitToCelsius(*rawObsData.TempOut), obs.Temp.Float32, 0.001)
	}
}

func requireDavisEqualDashboard(t *testing.T, rawObs davisRawWeatherDataResponseDashboard, obs DavisCurrentObservation) {
	for _, cur := range rawObs.CurrConditionValues {
		if cur.SensorDataTypeID == nil {
			continue
		}

		switch *cur.SensorDataTypeID {
		case 7:
			require.InDelta(t, *cur.Value, obs.Temp.Float32, 0.001)

		case 12:
			require.InDelta(t, *cur.Value, obs.Hi.Float32, 0.001)

		case 22:
			require.InDelta(t, *cur.Value, obs.Rr.Float32, 0.001)

		default:
			continue
		}
	}

	for _, hl := range rawObs.HighLowValues {
		if hl.SensorDataTypeID == nil {
			continue
		}
		switch *hl.SensorDataTypeID {
		case 57:
			require.InDelta(t, *hl.Value, obs.Tx.Float32, 0.001)

		default:
			continue
		}
	}
}

func datetimeToTimeStr(dt time.Time) string {
	return fmt.Sprintf("%s:%02d%s", dt.Format("3"), dt.Minute(), dt.Format("pm"))
}

func randomDavisRawResponseV1() davisRawCurrentResponseV1 {
	return davisRawCurrentResponseV1{
		Location:   util.RandomString(24),
		Lat:        util.RandomJSONFloat4(4.0, 22.0),
		Lon:        util.RandomJSONFloat4(114.0, 121.0),
		PressureMb: util.RandomJSONFloat4(990.0, 1100.),
		Rh:         util.RandomJSONFloat4(0.0, 100.0),
		TempC:      util.RandomJSONFloat4(25.0, 33.0),
		TdC:        util.RandomJSONFloat4(25.0, 33.0),
		WindDeg:    util.RandomJSONFloat4(0, 360),
		WindMPH:    util.RandomJSONFloat4(0.0, 10.0),
		HeatIndexC: util.RandomJSONFloat4(30.0, 50.0),
		Obs: davisRawCurrentObservationV1{
			RRInPerHr:       util.RandomJSONFloat4(0.0, 5.0),
			RainDayIn:       util.RandomJSONFloat4(0.0, 100.0),
			Srad:            util.RandomJSONFloat4(0, 400),
			UVIndex:         util.RandomJSONFloat4(0.0, 1.0),
			TempDayHighF:    util.RandomJSONFloat4(77.0, 104.0),
			TempDayLowF:     util.RandomJSONFloat4(60.0, 104.0),
			WindDayHighMPH:  util.RandomJSONFloat4(0.0, 20.0),
			TempDayHighTime: randomTimeString(),
			TempDayLowTime:  randomTimeString(),
			WindDayHighTime: randomTimeString(),
		},
		Time: time.Now().Format("Mon, 02 Jan 2006 15:04:05 -0700"),
	}
}

func randomDavisRawResponseV2() davisRawCurrentResponseV2 {
	pres := util.MbarToInHg(util.RandomFloat[float32](990.0, 1100.0))
	return davisRawCurrentResponseV2{
		StationID: util.RandomInt(100000, 999999),
		Sensors: []davisRawCurrentSensorResponseV2{
			{
				LSID: util.RandomInt(100000, 999999),
				Data: []davisRawCurrentDataResponseV2{
					{
						Bar:     &pres,
						TempOut: util.RandomFloatPtr[float32](25.0, 33.0),
						HumOut:  util.RandomFloatPtr[float32](0.0, 100.0),
					},
				},
			},
		},
	}
}

func randomDavisRawWeatherDataResponseDashboard() davisRawWeatherDataResponseDashboard {
	tempID, rhID, rainID := 7, 12, 22
	txID := 57
	highTempTime := float32(randomTimeInt())
	return davisRawWeatherDataResponseDashboard{
		CurrConditionValues: []davisRawSensorDataResponseDashboard{
			{
				SensorDataTypeID: &tempID,
				Value:            util.RandomFloatPtr[float32](24.6, 33.7),
			},
			{
				SensorDataTypeID: &rhID,
				Value:            util.RandomFloatPtr[float32](30.0, 100.0),
			},
			{
				SensorDataTypeID: &rainID,
				Value:            util.RandomFloatPtr[float32](0.0, 12.5),
			},
		},
		HighLowValues: []davisRawSensorDataResponseDashboard{
			{
				SensorDataTypeID: &txID,
				Value:            util.RandomFloatPtr[float32](28.5, 39.9),
			},
			{
				SensorDataName: "High Temp Time",
				Value:          &highTempTime,
			},
		},
	}
}

func randomTimeString() string {
	h := util.RandomInt(1, 12)
	m := util.RandomInt(0, 59)
	i := util.RandomInt(0, 100)
	x := "am"
	if i%2 == 0 {
		x = "pm"
	}

	return fmt.Sprintf("%d:%02d%s", h, m, x)
}

func randomTimeInt() int32 {
	h := util.RandomInt(0, 23)
	m := util.RandomInt(0, 59)

	return int32(h*100 + m)
}
