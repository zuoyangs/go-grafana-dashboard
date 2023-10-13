package elasticsearch

/*
import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

type esResponse struct {
	Hits struct {
		Hits []struct {
			ID     string                 `json:"_id"`
			Source map[string]interface{} `json:"_source"`
			Score  float64                `json:"_score,omitempty"`
		} `json:"hits,omitempty"`
		Total struct {
			Value int
		} `json:"total,omitempty"`
	} `json:"hits,omitempty"`

	ScrollID string `_scroll_id`
}

func main1() {
	cfg := elasticsearch8.Config{
		Addresses: []string{
			"http://10.230.193.24:9200",
			"http://10.230.198.148:9200",
			"http://10.230.208.95:9200",
		},

		Username: "elastic",
		Password: "Aa123456QAZ",
	}

	client, err := elasticsearch8.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	log.Println(elasticsearch8.Version)
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	log.Println(res)

	pattern := []string{"sh-prod-c-nginx-log-2023.10.10", "sh-prod-b-nginx-log-2023.10.10", "gz-prod-all-nginx-log-2023.10.10"}

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"_source": []string{"domain"},
		"size":    1000,
		//"scroll":  2 * time.Minute,
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, _ = client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(pattern),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	fmt.Printf("HTTP BODY: %s", data)
	defer res.Body.Close()

	var r esResponse
	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		log.Fatal(err)
	}

	scrollID := r.ScrollID

	for {

		scrollRes, _ := client.Scroll(client.Scroll.WithScrollID(scrollID),
			client.Scroll.WithScroll(time.Minute))

		if scrollRes.IsError() {
			break
		}

		defer scrollRes.Body.Close()

		var sr esResponse
		if json.NewDecoder(scrollRes.Body).Decode(&sr) != nil {
			break
		}

		if len(sr.Hits.Hits) == 0 {
			break
		}

		for _, hit := range sr.Hits.Hits {
			fmt.Println(hit.Source["domain.keywords"])
		}

		scrollID = sr.ScrollID
	}
}
*/
