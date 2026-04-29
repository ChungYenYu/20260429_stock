package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type YahooFinanceResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				RegularMarketPrice float64 `json:"regularMarketPrice"`
				Symbol             string  `json:"symbol"`
			} `json:"meta"`
			Timestamp  []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Open  []float64 `json:"open"`
					High  []float64 `json:"high"`
					Low   []float64 `json:"low"`
					Close []float64 `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"chart"`
}

type HistoricalData struct {
	Time  int64   `json:"time"`
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	Close float64 `json:"close"`
}

func FetchRealTimePrice(code string) (float64, error) {
	// Yahoo Finance suffix for Taiwan stocks
	symbol := code
	if !strings.Contains(symbol, ".") {
		symbol = fmt.Sprintf("%s.TW", code)
	}

	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s", symbol)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch price: status code %d", resp.StatusCode)
	}

	var data YahooFinanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	if data.Chart.Error != nil {
		return 0, fmt.Errorf("yahoo finance error: %v", data.Chart.Error)
	}

	if len(data.Chart.Result) == 0 {
		return 0, fmt.Errorf("no result found for symbol %s", symbol)
	}

	return data.Chart.Result[0].Meta.RegularMarketPrice, nil
}

func FetchHistoricalData(code string) ([]HistoricalData, error) {
	symbol := code
	if !strings.Contains(symbol, ".") {
		symbol = fmt.Sprintf("%s.TW", code)
	}

	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?interval=1d&range=1mo", symbol)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch historical data: status code %d", resp.StatusCode)
	}

	var data YahooFinanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Chart.Error != nil {
		return nil, fmt.Errorf("yahoo finance error: %v", data.Chart.Error)
	}

	if len(data.Chart.Result) == 0 {
		return nil, fmt.Errorf("no result found for symbol %s", symbol)
	}

	result := data.Chart.Result[0]
	if len(result.Timestamp) == 0 {
		return []HistoricalData{}, nil
	}

	quote := result.Indicators.Quote[0]
	history := make([]HistoricalData, 0, len(result.Timestamp))

	for i, ts := range result.Timestamp {
		// Some values might be nil in the JSON, which Decoder might decode as 0.
		// Usually Yahoo Finance provides enough data points.
		history = append(history, HistoricalData{
			Time:  ts,
			Open:  quote.Open[i],
			High:  quote.High[i],
			Low:   quote.Low[i],
			Close: quote.Close[i],
		})
	}

	return history, nil
}
