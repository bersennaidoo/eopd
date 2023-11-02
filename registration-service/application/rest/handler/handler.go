package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/bersennaidoo/eopd/registration-service/domain/contract"
	"github.com/bersennaidoo/eopd/registration-service/domain/model"
	"github.com/bersennaidoo/eopd/registration-service/foundation/gentoken"
	"github.com/gorilla/mux"
	"github.com/nats-io/nuid"
)

type Handler struct {
	store contract.Storer
}

func New(store contract.Storer) *Handler {

	return &Handler{
		store: store,
	}
}

func (h *Handler) HandleTest(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello world")
}

// HandleTokenReset processes token reset requests.
func (h *Handler) HandleTokenReset(w http.ResponseWriter, r *http.Request) {

	resetID, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	gentoken.GenerateTokenNumber(resetID)

	json.NewEncoder(w).Encode("Token reset successful")
}

func (h *Handler) HandleToken(w http.ResponseWriter, r *http.Request) {

	token := gentoken.GenerateTokenNumber(0)
	patientID, _ := strconv.Atoi(mux.Vars(r)["id"])
	fmt.Println("Token %d generated for user %d", token, patientID)

	registration_event := model.RegistrationEvent{patientID, token}
	reg_event, err := json.Marshal(registration_event)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(reg_event)
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var registration *model.RegistrationRequest
	err = json.Unmarshal(body, &registration)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Registers patient
	h.store.Register(registration)

	registration.RequestID = nuid.Next()
	fmt.Println(registration)

	json.NewEncoder(w).Encode(registration)
}
