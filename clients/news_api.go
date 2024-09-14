package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joshua468/news-aggregator-api/config"
)

type NewsAPIClient struct {
	apiKey string
}

func NewNewsAPIClient() *NewsAPIClient {
	return &NewsAPIClient{
		apiKey: config.GetEnv("NEWS_API_KEY"),
	}
}

type Article struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

func (client *NewsAPIClient) FetchNews(country string) ([]Article, error) {
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s", country, client.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch news: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	var response struct {
		Articles []Article `json:"articles"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode news API response: %w", err)
	}

	return response.Articles, nil
}
