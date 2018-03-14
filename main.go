package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/fax/ring", ringHandler)
	http.HandleFunc("/fax/received", receivedHandler)
	http.ListenAndServe(":2342", nil)
}

func ringHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("INCOMING: %s @ %s\n", r.FormValue("FaxSid"), r.FormValue("From"))
	w.Header().Set("Content-type", "text/xml")
	response := `<Response>
<Receive action="/fax/received"/>
</Response>`
	fmt.Fprintf(w, "%s", response)
}

func receivedHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("RECEIVED: %s\n", r.FormValue("MediaUrl"))
	w.Header().Set("Content-type", "text/xml")
	fmt.Fprintf(w, "%s", "<Response />")
}
