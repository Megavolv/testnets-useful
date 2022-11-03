package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	kvlib "github.com/Megavolv/okp4kviewlib"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var loglevel string
var path string
var limit int64

func init() {
	flag.StringVar(&loglevel, "level", "debug", "Log level (error|warn|info|debug)")
	flag.StringVar(&path, "path", "db/", "Path to keys")
	flag.Int64Var(&limit, "limit", 100000, "Limit on the number of keys per request")
	flag.Parse()
}

type Server struct {
	list   *kvlib.List
	logger *logrus.Logger
}

type Q struct {
	Start int64
	Count int64
}

func (s Server) handleGet(w http.ResponseWriter, req *http.Request) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		s.logger.Error(err)
		return
	}
	start := time.Now().In(loc)

	var q Q

	err = json.NewDecoder(req.Body).Decode(&q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.logger.Error(req.Body, err)
		return
	}

	from := fmt.Sprintf("Date: %s, From: %s, Start: %d, count: %d\n", start.String(), req.RemoteAddr, q.Start, q.Count)
	s.logger.Info(from)

	if q.Count == 0 {
		q.Count = 1
		s.logger.Info(fmt.Sprintf("From: %s, New count: %d\n", req.RemoteAddr, q.Count))
	} else if q.Count > limit {
		q.Count = limit
		s.logger.Info(fmt.Sprintf("From: %s, New count: %d\n", req.RemoteAddr, q.Count))
	}

	data, err := s.list.GetKeys(q.Start, q.Count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		duration := time.Since(start)
		s.logger.Info(fmt.Sprintf("From: %s. Request completed in %s", req.RemoteAddr, duration.String()))

		return
	}

	_, err = fmt.Fprintf(w, data)
	if err != nil {
		s.logger.Error(err)
	}

	duration := time.Since(start)
	s.logger.Info(fmt.Sprintf("From: %s. Request completed in %s", req.RemoteAddr, duration.String()))
	return
}

func main() {
	logger := logrus.New()
	switch loglevel {
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)

	default:
		logger.SetLevel(logrus.DebugLevel)
	}

	list := kvlib.NewList(path, logger)
	server := Server{list: list, logger: logger}
	defer list.CloseAll()

	logger.Info("Listen at:7001")

	r := mux.NewRouter()
	r.HandleFunc("/keys", server.handleGet).Methods("GET")
	srv := &http.Server{
		Addr:    ":7001",
		Handler: r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
