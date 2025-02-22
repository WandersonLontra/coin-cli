package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/WandersonLontra/coin-cli/internal/entity"
)

type Fetcher struct {
	BaseUrl 		string
	Route 			string
	AccessKey 		string
}

func NewFetcher(baseUrl, accessKey string) *Fetcher {
	return &Fetcher{
		BaseUrl: baseUrl,
		AccessKey: accessKey,
	}
}

func (f * Fetcher) GetEndpoint() string {
	return fmt.Sprintf("%s?access_key=%s", f.BaseUrl, f.AccessKey)
}
func (f *Fetcher) GetCurrencies() (*entity.Currency, error) {
	endpoint := f.GetEndpoint()
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	currency := entity.Currency{}
	err = json.NewDecoder(res.Body).Decode(&currency)
	if err != nil {
		return nil, err
	}
	currency.Timestamp = time.Now().UTC().UnixMilli()
	return &currency, nil
}