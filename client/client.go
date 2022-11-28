package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MrHenri/go-expert-chalange-1/model"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusRequestTimeout {
		log.Fatal("Server timeout")
	}

	defer resp.Body.Close()

	var quote model.QuoteDTO
	if err := json.NewDecoder(resp.Body).Decode(&quote); err != nil {
		panic(err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write([]byte("Dólar: " + quote.Bid))
}
