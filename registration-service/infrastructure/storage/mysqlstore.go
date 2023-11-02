package storage

import (
	"database/sql"

	"github.com/bersennaidoo/eopd/registration-service/domain/model"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {

	return &Store{
		db: db,
	}
}

func (s *Store) Register(registration *model.RegistrationRequest) {

	insForm, err := s.db.Prepare("INSERT INTO patient_details(id, full_name, address, sex, phone, remarks) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(registration.ID, registration.FullName, registration.Address, registration.Sex, registration.Phone, registration.Remarks)
}

func (s *Store) Update(registration *model.RegistrationRequest) {

	insForm, err := s.db.Prepare("UPDATE patient_details SET full_name=?, address=?, sex=?, phone=?, remarks=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(registration.FullName, registration.Address, registration.Sex, registration.Phone, registration.Remarks, registration.ID)
}

func (s *Store) View(patientID string) *model.RegistrationRequest {

	selDB, err := s.db.Query("SELECT * FROM patient_details WHERE ID=?", patientID)
	if err != nil {
		panic(err.Error())
	}

	registration := &model.RegistrationRequest{}
	for selDB.Next() {
		var id, phone int
		var full_name, address, sex, remarks string
		err = selDB.Scan(&id, &full_name, &address, &sex, &phone, &remarks)
		if err != nil {
			panic(err.Error())
		}
		registration.ID = id
		registration.FullName = full_name
		registration.Address = address
		registration.Sex = sex
		registration.Phone = phone
		registration.Remarks = remarks
	}

	return registration
}
