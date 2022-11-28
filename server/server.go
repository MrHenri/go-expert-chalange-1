package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/MrHenri/go-expert-chalange-1/database"
	"github.com/MrHenri/go-expert-chalange-1/model"
)

func Quotation(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	quote, err := quotePayloader(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = quoteResponse(quote, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = database.InsertDolarQuotation(&quote.Usdbrl)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func quoteResponse(quote *model.Quote, w http.ResponseWriter) error {
	quoteClient := model.QuoteDTO{
		Bid: quote.Usdbrl.Bid,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(quoteClient)
}

func quotePayloader(resp *http.Response) (*model.Quote, error) {

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var quote model.Quote
	err = json.Unmarshal(body, &quote)
	return &quote, err
}
