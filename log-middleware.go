package snitch

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

type NetworkLog struct {
	ServiceName        string        `json:"service_name"`
	LogType            string        `json:"log_type"`
	RequestURL         string        `json:"request_url"`
	StartTime          time.Time     `json:"start_time"`
	ResponseTime       time.Duration `json:"duration"`
	HTTPStatusReturned int           `json:"http_status"`
}

//LogMiddleware : also compatible with negroni
func LogMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	starttime := time.Now()
	next(w, r)
	res := w.(negroni.ResponseWriter)

	newEntry := &NetworkLog{
		ServiceName:        ServiceName,
		LogType:            "network_log",
		RequestURL:         r.URL.Path,
		StartTime:          starttime,
		ResponseTime:       time.Since(starttime),
		HTTPStatusReturned: res.Status(),
	}

	byt, _ := json.Marshal(newEntry)
	SendLog(string(byt))

}
