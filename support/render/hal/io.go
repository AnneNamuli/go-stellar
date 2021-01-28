package hal

import (
	"net/http"

	"github.com/AnneNamuli/go-stellar/support/render/httpjson"
)

// Render write data to w, after marshalling to json
func Render(w http.ResponseWriter, data interface{}) {
	httpjson.Render(w, data, httpjson.HALJSON)
}
