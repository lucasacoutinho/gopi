package chi

import (
	"github.com/go-chi/chi/v5"
	"github.com/lucasacoutinho/gopi/user"
)

func Handlers(uService user.UseCase) *chi.Mux {
	r := chi.NewRouter()

	return r
}
