package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		return 0, nil
	}

	defer f.Close()

	return f.Write(data)
}

func Run(destination string) {
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

func RegisterHandler() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		req, err := ioutil.ReadAll(r.Body)
		if err != nil || len(req) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		write(string(req))
	})
}


func write(msg string) {
	log.Printf("%v\n", msg)
}