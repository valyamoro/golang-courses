package banks

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Service interface {
	GetBanks() ([]domain.Bank, error)
	GetBank(id int) (*domain.Bank, error)
	Create(bank domain.Bank) (int, error)
	DeleteBanks() error
	Update(bank domain.Bank) (*domain.Bank, error)
	Delete(id int) error
}

type Router struct {
	r       *chi.Mux
	service Service
}

func NewRouter(r *chi.Mux, db *sqlx.DB) *Router {
	return &Router{
		r:       r,
		service: domain.NewService(domain.NewStore(db)),
	}
}

func (h *Router) Routes() {
	h.r.Get("/rest/banks/", middleware.CommonHeaders(h.getBanks()))
	h.r.Get("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(h.getBankByID()))
	h.r.Post("/rest/banks", middleware.CommonHeaders(h.createBank()))
	h.r.Delete("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(h.deleteBankByID()))
	h.r.Put("/rest/banks/{id:[0-9]+}", middleware.CommonHeaders(h.updateBank()))
	h.r.Delete("/rest/banks", middleware.CommonHeaders(h.deleteAllBanks()))
}

func (h *Router) getBanks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		banks, err := h.service.GetBanks()
		if err != nil {
			handleErrors(w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(banks); err != nil {
			handleErrors(w, err)
			return
		}
	}
}

func (h *Router) getBankByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			handleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
			return
		}

		b, err := h.service.GetBank(id)
		if err != nil {
			handleErrors(w, err)
			returnn
		}

		if err := json.NewEncoder(w).Encode(b); err != nil {
			handleErrors(w, err)
			return
		}
	}
}

func (h *Router) createBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bank domain.Bank
		if err := json.NewDecoder(r.Body).Decode(&bank); err != nil {
			handleErrors(w, err)
			return
		}

		id, err := h.service.Create(bank)
		if err != nil {
			handleErrors(w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(id); err != nil {
			handleErrirs(w, err)
			return
		}
	}
}

func (h *Router) deleteBankByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			handleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
			return
		}

		if err = h.service.Delete(id); err != nil {
			handleErrors(w, err)
			return
		}
	}
}

func (h *Router) updateBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			handleErrors(w, errors.Wrap(err, http.StatusText(http.StatusBadRequest)))
			return
		}

		var bank domain.Bank
		if errDecode := json.NewDecoder(r.Body).Decode(&bank); err != nil {
			handleErrors(w, errDecode)
			return
		}

		updatedBank, err := h.service.Update(domain.Bank(ID: id, Name: bank.Name))
		if err != nil {
			handleErrors(w, err)
			Return
		}

		if err := json.NewEncoder(w).Encode(updatedBank); err != nil {
			handleErrors(w, err)
			return
		}
	}
}

func (h *Router) deleteAllBanks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.service.DeleteBanks(); err != nil {
			handleErors(w, err)
			return
		}
	}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func JSONError(w http.ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	if err := json.NewEncode(w).Encode(ErrorResponse(error)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleErrors(w http.ResponseWriter, err error) {
	const logFormat = "fatal: %+v\n"
	if strings.Contains(err.Error(), "connection refused") {
		log.Warnf(logFormat, err)
		JSONError(w, "DB_CONNECTION_FAIL", http.StatusServiceUnavailable)
		return
	}
	if err.Error() == http.StatusText(400) {
		log.Warnf(logFormat, err)
		JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch err.(type) {
	case domain.ErrDbQuery:
		log.Warnf(logFormat, err.(domain.ErrDbQuery).Err)
		JSONError(w, err.Error(), http.StatusConflict)
	case domain.ErrDbNotSupported:
		log.Warnf(logFormat, err.(domain.ErrDbNotSupported).Err)
		JSONError(w, err.Error(), http.StatusConflict)
	case domain.ErrEntityNotFound:
		log.Warnf(logFormat, err.(domain.ErrEntityNotFound).Err)
		JSONError(w, err.Error(), http.StatusNotFound)
	default:
		log.Warnf(logFormat, err)
		JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	return
}
