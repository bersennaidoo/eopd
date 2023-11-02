package contract

import "github.com/bersennaidoo/eopd/registration-service/domain/model"

// A Contract for a patient store.
type Storer interface {
	Register(registration *model.RegistrationRequest)

	Update(registration *model.RegistrationRequest)

	View(patientID string) *model.RegistrationRequest
}
