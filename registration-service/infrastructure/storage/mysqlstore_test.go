package storage

import (
	"testing"

	"github.com/bersennaidoo/eopd/registration-service/domain/model"
	"github.com/bersennaidoo/eopd/registration-service/infrastructure/mocks"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {

	t.Run("register user", func(t *testing.T) {

		registration := model.RegistrationRequest{
			ID:       200,
			FullName: "bersen naidoo",
			Address:  "42 Ale street",
			Sex:      "Male",
			Phone:    011111111,
			Remarks:  "urgent",
		}

		store := mocks.NewStorer(t)

		store.On("Register", &registration).Return(nil).Once()

		err := store.Register(&registration)

		require.Nil(t, err)
		store.AssertExpectations(t)
	})

	t.Run("update user", func(t *testing.T) {

		registration := model.RegistrationRequest{
			FullName: "bersen naidoo",
			Address:  "42 Ale street",
			Sex:      "Male",
			Phone:    011111111,
			Remarks:  "urgent",
			ID:       200,
		}

		store := mocks.NewStorer(t)

		store.On("Update", &registration).Return(nil).Once()

		err := store.Update(&registration)

		require.Nil(t, err)
		store.AssertExpectations(t)
	})

	t.Run("view user", func(t *testing.T) {

		id := "200"

		registration := model.RegistrationRequest{
			FullName: "bersen naidoo",
			Address:  "42 Ale street",
			Sex:      "Male",
			Phone:    011111111,
			Remarks:  "urgent",
			ID:       200,
		}

		store := mocks.NewStorer(t)

		store.On("View", id).Return(&registration, nil).Once()

		reg, err := store.View(id)

		require.Nil(t, err)
		require.Equal(t, registration.ID, reg.ID)
		require.Equal(t, registration.FullName, reg.FullName)
		store.AssertExpectations(t)
	})
}
