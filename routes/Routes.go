package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/magneto/Controller"
	"github.com/magneto/Repository"
	"github.com/magneto/core"
	"net/http"
	"strings"
)

type Routes struct {
}

func (p *Routes) SetupRoutes(router *mux.Router) {
		enableCORS(router)
		router.HandleFunc("/mutant", func(w http.ResponseWriter, r *http.Request) {
			var request core.Request
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				respondWithError(err, w)
			}else{
				var chain = request.Dna
				var magneto Controller.MagnetoController
				chain = strings.ReplaceAll(chain,"}","")
				chain = strings.ReplaceAll(chain,"{","")
				var repo = Repository.MongoMutantRepository{}
				if magneto.IsMutant(chain) {
					magneto.Save(chain, true, &repo)
					respondWithOk("Your welcome Mutant",w)
				}else{
					magneto.Save(chain, false, &repo)
					respondWithForbidden("You are not a Mutant",w)
				}
			}
		}).Methods(http.MethodPost)

	router.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		var magneto Controller.MagnetoController
		var repo = Repository.MongoMutantRepository{}
		response := magneto.FindAll(&repo)
		json.NewEncoder(w).Encode(response)
	}).Methods(http.MethodGet)

}

func enableCORS(router *mux.Router)  {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Accept", "application/json")
			next.ServeHTTP(w, req)
		})
}

func respondWithError(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(data)
}

func respondWithForbidden(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(data)
}

func respondWithOk(data interface{}, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

