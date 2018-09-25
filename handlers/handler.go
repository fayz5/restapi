package handlers

// Message represents data model for JSON api
type Message struct {
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
}

//Message handler responds with predefined JSON data
func MessageHandler(w http.ResponseWriter, r *http.Request) {
	m := Message{"Hello From Server", time.Now().Unix()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}
