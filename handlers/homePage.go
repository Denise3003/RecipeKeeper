package handlers

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello bla bla bla hdfjkfhkfhakdf 757584939395767839")

}
