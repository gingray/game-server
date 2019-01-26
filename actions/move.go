package actions

import (
	"fmt"
	"net/http"
)

type Move struct {
	id string
	x int
	y int
}

func (h *Move) HandleMove(w http.ResponseWriter, r *http.Request)  {
	message := "Handle move"
	fmt.Println(message)
	w.Write([]byte(message))
}

