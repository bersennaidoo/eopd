package contract

import "github.com/bersennaidoo/eopd/registration-service/domain/model"

// A Contract for a patient store.
type Storer interface {
	Register(registration *model.RegistrationRequest) error

	Update(registration *model.RegistrationRequest) error

	View(patientID string) (*model.RegistrationRequest, error)
}
