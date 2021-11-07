package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r *Retriever) Get(url string) string {
	reps, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer reps.Body.Close()

	result, err := httputil.DumpResponse(reps, true)
	if err != nil {
		panic(err)
	}

	return string(result)
}
