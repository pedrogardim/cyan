package handler

import (
	"net/http"
	"time"

	"github.com/go-cyan/cyan/repository"
)

type HandlerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
type Handler struct {
	repo repository.RepositoryInterface
}

func NewHandler(name interface{}, repo repository.RepositoryInterface) HandlerInterface {
	return Handler{
		repo: repo,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format("HH:MM")
	w.Write([]byte("The time is: " + tm))
}
