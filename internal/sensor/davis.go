package sensor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const (
	DavisAPIV1URL     = "https://api.weatherlink.com/v1/NoaaExt.json"
	DavisAPIV2URL     = "https://api.weatherlink.com/v2"
	DavisDashboardURL = "https://www.weatherlink.com/embeddablePage/summaryData"
	HTTPUserAgent     = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"
)

type DavisSensor interface {
	FetchLatest() ([]DavisCurrentObservation, error)
}

type DavisFactory func(cred DavisAPICredentials, sleepDuration time.Duration) (DavisSensor, error)

type DavisAPICredentials struct {
	Type                 string
	User, Pass, APIToken string
	APIKey, APISecret    string
	StnUUID              string
}

type Davis struct {
	apiCredentials DavisAPICredentials
	client         Fetcher
	sleep          time.Duration
}

type DavisCurrentObservation struct {
	Rr            pgtype.Float4      `json:"rain"`
	Temp          pgtype.Float4      `json:"temp"`
	Rh            pgtype.Float4      `json:"rh"`
	Wdir          pgtype.Float4      `json:"wdir"`
	Wspd          pgtype.Float4      `json:"wspd"`
	Wspdx         pgtype.Float4      `json:"gust"`
	Srad          pgtype.Float4      `json:"srad"`
	Pres          pgtype.Float4      `json:"mslp"`
	Tn            pgtype.Float4      `json:"tn"`
	Tx            pgtype.Float4      `json:"tx"`
	Hi            pgtype.Float4      `json:"hi"`
	RainAccum     pgtype.Float4      `json:"rain_accum"`
	TnTimestamp   pgtype.Timestamptz `json:"tn_timestamp"`
	TxTimestamp   pgtype.Timestamptz `json:"tx_timestamp"`
	GustTimestamp pgtype.Timestamptz `json:"gust_timestamp"`
	Timestamp     pgtype.Timestamptz `json:"timestamp"`
}

type davisRawCurrentResponse interface {
	ToDavisCurrentObservation() *DavisCurrentObservation
}

func NewDavis(apiCredentials DavisAPICredentials, sleep time.Duration) (*Davis, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	if isValid := isCredentialValid(apiCredentials); !isValid {
		return nil, fmt.Errorf("API credentials are invalid")
	}
	return &Davis{
		apiCredentials: apiCredentials,
		client:         client,
		sleep:          sleep,
	}, nil
}

func isCredentialValid(c DavisAPICredentials) bool {
	switch c.Type {
	case "dashboard":
		return c.StnUUID != ""
	case "v2":
		return c.APIKey != "" && c.APISecret != ""
	case "v1":
		return c.User != "" && c.Pass != ""
	default:
		return false
	}
}

func (d Davis) FetchLatest() ([]DavisCurrentObservation, error) {
	var (
		apiURL     string
		params     url.Values
		rawParam   string
		finalParam string
	)

	switch d.apiCredentials.Type {
	case "v2":
		apiURL = DavisAPIV2URL + "/stations"
		params = url.Values{
			"api-key": {d.apiCredentials.APIKey},
		}
	case "v1":
		apiURL = DavisAPIV1URL
		rawParam = "pass=" + d.apiCredentials.Pass
		params = url.Values{
			"user":     {d.apiCredentials.User},
			"apiToken": {d.apiCredentials.APIToken},
		}
	case "dashboard":
		apiURL = fmt.Sprintf("%s/%s", DavisDashboardURL, d.apiCredentials.StnUUID)
	}

	baseURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}
	encParam := params.Encode()

	switch {
	case encParam != "" && rawParam != "":
		finalParam = encParam + "&" + rawParam
	case encParam != "":
		finalParam = encParam
	case rawParam != "":
		finalParam = rawParam
	default:
		finalParam = ""
	}

	baseURL.RawQuery = finalParam
	encodedURL := baseURL.String()

	req, err := http.NewRequest("GET", encodedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", HTTPUserAgent)

	if d.apiCredentials.Type == "v2" {
		req.Header.Set("X-API-Secret", d.apiCredentials.APISecret)
	}

	time.Sleep(d.sleep)

	res, err := d.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch d.apiCredentials.Type {
	case "v2":
		var rawStations davisRawStationsResponseV2
		err = json.NewDecoder(res.Body).Decode(&rawStations)
		if err != nil {
			return nil, err
		}

		obsSlice := make([]DavisCurrentObservation, 0)
		for _, rawStn := range rawStations.Stations {
			apiURL = fmt.Sprintf("%s/current/%d", DavisAPIV2URL, rawStn.StationID)
			params = url.Values{
				"api-key": {d.apiCredentials.APIKey},
			}

			baseURL, err := url.Parse(apiURL)
			if err != nil {
				return nil, err
			}
			baseURL.RawQuery = params.Encode()
			encodedURL := baseURL.String()

			req, err := http.NewRequest("GET", encodedURL, nil)
			if err != nil {
				return nil, err
			}

			req.Header.Set("User-Agent", HTTPUserAgent)
			req.Header.Set("X-API-Secret", d.apiCredentials.APISecret)

			time.Sleep(d.sleep)

			res, err := d.client.Do(req)
			if err != nil {
				return nil, err
			}
			defer res.Body.Close()

			var rawObs davisRawCurrentResponseV2
			err = json.NewDecoder(res.Body).Decode(&rawObs)
			if err != nil {
				return nil, err
			}

			obs := rawObs.ToDavisCurrentObservation()
			if obs != nil {
				obsSlice = append(obsSlice, *obs)
			}
		}
		return obsSlice, nil
	case "v1":
		var rawObs davisRawCurrentResponseV1
		err = json.NewDecoder(res.Body).Decode(&rawObs)
		if err != nil {
			return nil, err
		}

		obs := rawObs.ToDavisCurrentObservation()
		return []DavisCurrentObservation{*obs}, nil
	}
	var rawObs davisRawWeatherDataResponseDashboard
	err = json.NewDecoder(res.Body).Decode(&rawObs)
	if err != nil {
		return nil, err
	}

	obs := rawObs.ToDavisCurrentObservation()
	return []DavisCurrentObservation{*obs}, nil
}

func parseTimeStrToDateTime(timeStr string) (time.Time, error) {
	layout := "3:04pm"
	currentDate := time.Now()

	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return currentDate, err
	}

	newDateTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), t.Hour(), t.Minute(), 0, 0, currentDate.Location())
	return newDateTime, nil
}
