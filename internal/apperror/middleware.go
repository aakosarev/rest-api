package apperror

import (
	"errors"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appError *AppError
		err := h(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			if errors.As(err, &appError) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(appError.Marshal())
					return
				}

				err = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(ErrNotFound.Marshal())
			}

			w.WriteHeader(http.StatusTeapot)
			w.Write(systemErr(err).Marshal())
		}
	}
}
