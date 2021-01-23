package resthandlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	strip "github.com/grokify/html-strip-tags-go"
)

type exercise1Handler struct {
	router *mux.Router
}

//NewExercise1Handler instatiates the exercise1 handler
func NewExercise1Handler(_router *mux.Router) *exercise1Handler {
	return &exercise1Handler{
		router: _router,
	}
}

func (me *exercise1Handler) GetTextAndCountForWeb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("URL cannot be empty"))
		return
	}
	fmt.Println(url)
	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error while getting data through http: %#v", err)
		return
	}
	defer resp.Body.Close()
	// Read the response body
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error while reading the response: %#v", err)
		return
	}

	stripped := strip.StripTags(string(bytes))

	allWords := strings.Split(stripped, " ")
	wordsCount := map[string]int{}
	for _, word := range allWords {
		if val, ok := wordsCount[word]; ok {
			wordsCount[word] = val + 1
		} else {
			wordsCount[word] = 1
		}
	}

	json.NewEncoder(w).Encode(wordsCount)
}
