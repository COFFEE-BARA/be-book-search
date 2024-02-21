package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

// var (
// 	elasticCloudID = ""
// 	elasticAPIKey  = ""
// 	elasticIndex   = "book-index"
// )

// Book 구조체 정의
type Book struct {
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Image  string `json:"image"`
	Price  string `json:"price"`
}

// 응답 메시지 구조체
type ResponseMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Lambda 핸들러 함수
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	elasticCloudID := os.Getenv("ELASTIC_CLOUD_ID")
	elasticAPIKey := os.Getenv("ELASTIC_API_KEY")
	elasticIndex := os.Getenv("ELASTIC_INDEX")

	// Set CORS headers
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*", // Allow requests from any origin
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Allow-Methods": "*", // Allow all methods
		// Add more CORS headers if needed
	}

	cfg := elasticsearch.Config{
		CloudID: elasticCloudID,
		APIKey:  elasticAPIKey,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println(err)
		response, _ := json.Marshal(ResponseMessage{500, "서버 에러", nil})
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       string(response),
		}, nil
	}

	searchWord, ok := request.QueryStringParameters["keyword"]
	if !ok || searchWord == "" {
		response, _ := json.Marshal(ResponseMessage{400, "keyword가 없습니다.", nil})
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers:    headers,
			Body:       string(response),
		}, nil
	}

	query := map[string]any{
		"query": map[string]any{
			"multi_match": map[string]any{
				"query":  searchWord,
				"fields": []string{"Title"},
			},
		},
		"_source": []string{"Title", "Author", "ImageURL", "Price"},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		fmt.Println(err)
		response, _ := json.Marshal(ResponseMessage{500, "서버 에러", nil})
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       string(response),
		}, nil
	}

	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex(elasticIndex),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		fmt.Println(err)
		response, _ := json.Marshal(ResponseMessage{500, "서버 에러", nil})
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       string(response),
		}, nil
	}
	defer res.Body.Close()

	var r map[string]any
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		fmt.Println(err)
		response, _ := json.Marshal(ResponseMessage{500, "서버 에러", nil})
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       string(response),
		}, nil
	}

	hits := r["hits"].(map[string]any)["hits"].([]any)
	bookList := make([]Book, len(hits))
	for i, hit := range hits {
		source := hit.(map[string]any)["_source"].(map[string]any)
		var price string
		if p, ok := source["Price"].(float64); ok {
			price = fmt.Sprintf("%.2f", p)
		} else if p, ok := source["Price"].(string); ok {
			price = p
		} else {
			price = "Unknown"
		}
		bookList[i] = Book{
			Isbn:   hit.(map[string]any)["_id"].(string),
			Title:  source["Title"].(string),
			Author: source["Author"].(string),
			Image:  source["ImageURL"].(string),
			Price:  price,
		}
	}

	response, _ := json.Marshal(ResponseMessage{200, "검색 결과를 가져오는데 성공했습니다.", map[string]any{"bookList": bookList}})
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers:    headers,
		Body:       string(response),
	}, nil
}

func main() {
	lambda.Start(Handler)
}

// func main() {
// 	// 테스트용 이벤트 데이터
// 	event := events.APIGatewayProxyRequest{
// 		QueryStringParameters: map[string]string{
// 			"keyword": "서버리스",
// 		},
// 	}

// 	result, err := Handler(context.Background(), event)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}

// 	fmt.Println("Status Code:", result.StatusCode)
// 	fmt.Println("Body:", result.Body)
// }
