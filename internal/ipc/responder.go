package ipc

import (
	"encoding/json"
	"fmt"
	"io"
)

const tilixIcon = "com.gexperts.Tilix"

type Response struct {
	Append PluginResult `json:"Append,omitempty"`
}

type Icon struct {
	Name string `json:"Name"`
}

type PluginResult struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        Icon   `json:"icon"`
	Exec        string `json:"exec"`
}

type Responder struct {
	w io.Writer
}

func NewResponder(w io.Writer) *Responder {
	return &Responder{w: w}
}

func (r *Responder) Append(id int, name string, path string) error {
	description := fmt.Sprintf("Open: %s", name)
	exec := fmt.Sprintf("tilix --working-directory=%s", path)
	response := Response{Append: PluginResult{
		ID:          id,
		Name:        name,
		Description: description,
		Icon:        Icon{Name: tilixIcon},
		Exec:        exec,
	}}

	return r.send(response)
}

func (r *Responder) Finished() error {
	return r.send("Finished")
}

func (r *Responder) Close() error {
	return r.send("Close")
}

func (r *Responder) send(payload interface{}) error {
	if err := json.NewEncoder(r.w).Encode(payload); err != nil {
		return err
	}

	return nil
}
