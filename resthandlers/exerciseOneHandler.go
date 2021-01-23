package resthandlers

import (
	"net/http"
)

func GetTextAndCountForWeb(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
