package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Workorder represents a workorder in the system.
type Workorder struct {
	OBJNR       string    `json:"objnr,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	Description string    `json:"description,omitempty"`
	Adress      string    `json:"adress,omitempty"`
	Start       string    `json:"start,omitempty"`
	Status      string    `json:"status,omitempty"`
	Invoice     string    `json:"invoice,omitempty"`
}

var workorders []Workorder

// getWorkordersHandler fetches all workorders from db
func getWorkordersHandler(w http.ResponseWriter, r *http.Request) {
	workorderList, err := json.Marshal(workorders)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err)) // to make it print to stdout, easier for debugging during dev.
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(workorderList)
}

// createWorkorderHandler creates a workorder
func createWorkorderHandler(w http.ResponseWriter, r *http.Request) {
	order := Workorder{}

	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	order.OBJNR = r.Form.Get("objnr")
	order.CreatedAt = time.Now().Local()
	order.Description = r.Form.Get("description")
	order.Adress = r.Form.Get("adress")
	order.Start = r.Form.Get("start")
	order.Status = r.Form.Get("status")
	order.Invoice = r.Form.Get("invoice")

	workorders = append(workorders, order)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
