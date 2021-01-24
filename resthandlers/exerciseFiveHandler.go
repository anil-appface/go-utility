package resthandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"sync"

	"github.com/anil-appface/go-utility/util"
	"github.com/gorilla/mux"
)

type exercise5Handler struct {
	router *mux.Router
}

//NewExercise5Handler instatiates the exercise4 handler
func NewExercise5Handler(_router *mux.Router) *exercise5Handler {
	return &exercise5Handler{
		router: _router,
	}
}

func (me *exercise5Handler) GetPrimeNumbers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	n, err := strconv.Atoi(mux.Vars(r)["number"])
	if err != nil {
		fmt.Printf("error while converting the string to number %#v", err)
		return
	}
	primeNumbers := []int{}
	var wg sync.WaitGroup
	primeChan := make(chan int)

	go func() {
		for {
			select {
			case primeNumber := <-primeChan:
				primeNumbers = append(primeNumbers, primeNumber)
			}
		}
	}()

	for i := 2; i < n; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			if util.CheckPrime(j) {
				primeChan <- j
			}
		}(i)
	}

	wg.Wait()
	sort.Ints(primeNumbers)
	json.NewEncoder(w).Encode(primeNumbers)

}
