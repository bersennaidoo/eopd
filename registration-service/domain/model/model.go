package model

// RegistrationRequest contains data about the patient.
type RegistrationRequest struct {
	FullName string `json:"full_name,omitempty"`

	Address string `json:"address,omitempty"`

	ID int `json:"id"`

	Sex string `json:"sex,omitempty"`

	Email string `json:"email,omitempty"`

	Phone int `json:"phone,omitempty"`

	Remarks string `json:"remarks,omitempty"`

	RequestID string `json:"request_id,omitempty"`
}

// RegistrationEvent contains the details for a given registration instance.
type RegistrationEvent struct {
	ID int `json:"id"`

	Token uint64 `json:"token"`
}
