package server

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)



var LastActivities *Activities


type httpServer struct {
	Activities *Activities
}



type IDDocument struct {
	ID uint64 `json:"id"`
}



type ActivityDocument struct {
	Activity Activity `json:"activity"`
}



func NewHTTPServer(addr string) *http.Server {
	server := &httpServer{
		Activities: &Activities{},
	}
	
	r := mux.NewRouter()
	r.HandleFunc("/", server.handlePost).Methods("POST")
	//r.HandleFunc("/", server.handleGet).Methods("GET")
	r.HandleFunc("/", server.handleGet).Methods("GET")
	
	LastActivities = server.Activities
	
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}





//func handleGet(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "get\n")
//}
func (s *httpServer) handleGet(w http.ResponseWriter, r *http.Request) {
	var req IDDocument
	
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	//...
	fmt.Fprintf(w, "get\n")
	
	activity, err := s.Activities.Retrieve(req.ID)
	if err == ErrIDNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res := ActivityDocument{Activity: activity}
	json.NewEncoder(w).Encode(res)
}



func (s *httpServer) handlePost(w http.ResponseWriter, r *http.Request) {
	var req ActivityDocument
	
	err := json.NewDecoder(r.Body).Decode(&req)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	//...
	fmt.Fprintf(w, "post\n")
	
	id := s.Activities.Insert(req.Activity)
	res := IDDocument{ID: id}
	json.NewEncoder(w).Encode(res)
}