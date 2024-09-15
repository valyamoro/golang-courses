package repos

// https://fenyuk.medium.com/golang-crud-in-rest-api-in-a-generic-way-9c395a60309e
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BrandRepo struct {
	brands []entities.Brand
}

func NewBrandRepo() *BrandRepo {
	var br = BrandRepo{make([]entities.Brand, 0)}

	return &br
}

func (b *BrandRepo) Create(partial etntities.Brand) entities.Brand {
	newItem := entities.Brand{uint(len(b.brands)) + 1, partial.Name, partial.Year}
	b.brands = append(b.brands, newItem)
	return newItem
}

func (b *BrandRepo) GetList() []entities.Brand {
	return b.brands
}

func (p *BrandRepo) GetOne(id uint) (entitites.Brandh, error) {
	for _, it := range p.brands {
		if it.ID == id {
			return it, nil
		}
	}

	return entities.Brand{}, fmt.Errorf("key %d not found", id)
}

func (p *BrandRepo) Update(id uint, amended entities.Brand) (entities.Brand, error) {
	for i, it := range p.brands {
		if it.ID == id {
			amended.ID = id
			p.brands = append(p.brands[:i], p.brands[i+1:]...)
			p.brands = append(p.brands, amended)

			return amended, nil
		}
	}

	return entities.Brand{}, fmt.Errorf("key %d not found", amended.ID)
}

func (p *BrandRepo) DeleteOne(id uint) (bool, error) {
	for i, it := range p.brands {
		if it.ID == id {
			p.brands = append(p.brands[:i], p.brands[i+1:]...)
			return true, nil
		}
	}

	return false, fmt.Errorf("key %d nnot found", id)
}

type GenericRouter[TT any, T repos.GenericRepo[TT]] struct {
	muxBase string
	repo    *repos.GenericRepo[TT]
}

func (rtr *GenericRouter[TT, T]) handle(w http.ResponseWriter, r *http.Request) {
	idLong, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if r.URL.EscapedPath() != rtr.muxBase && err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		if idLong != 0 {
			item, err := (*(rtr.repo)).GetOne(uint(idLong))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode("Entity not found!")
				return
			}

			json.NewEncoder(w).Encode(&item)
		} else {
			items := (*(rtr.repo)).GetList()
			json.NewEncoder(w).Encode(&items)
		}

		w.WriteHeader(http.StatusOK)
	case http.MethodPost:
		var model TT
		json.NewDecoder(r.Body).Decode(&model)
		(*(rtr.repo)).Create(model)
		w.WriteHeader(http.StatusCreated)
	case http.MethodPut:
		var model TT
		json.NewDecoder(r.Body).Decode(&model)
		_, err := (*(rtr.repo)).Update(uint(idLong), model)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("entity not found")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	case http.MethodDelete:
		if idLong != 0 {
			_, err := (*(rtr.repo)).DeleteOne(uint(idLong))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode("Entity not found!")
				return
			}
			w.WriteHeader(http.StatusNoContent)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (rtr *GenericRouter[TT, T]) registerRoutes(mux *mux.Router) {
	mux.HandleFunc(rtr.muxBase, rtr.handle)
	mux.HandleFunc(fmt.Sprintf("%v/{id}", rtr.muxBase), rtr.handle)
}

func NewGenericRouter[TT any, T repos.GenericRepo[TT]](
	muxBase string,
	mux *mux.Router,
	repo *repos.GenericRepo[TT],
) *GenericRouter[TT, T] {
	router := GenericRouter[TT, T]{
		muxBase: muxBase,
		repo:    repo,
	}

	router.registerRoutes(mux)

	return &router
}

func RegisterBrandRoutes(router *mux.Router) {
	var brandRepo repos.GenericRepo[entitites.Brand] = repos.NewBrandRepo()

	NewGenericRouter[entities.Brand, *repos.BrandRepo]("/api/brands", router, &brandRepo)
}
