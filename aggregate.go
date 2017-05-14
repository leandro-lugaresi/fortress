package hydraidp

import (
	"time"

	"github.com/hellofresh/goengine"
)

// UserAggregate represent one user in the identity provider.
type UserAggregate struct {
	*goengine.AggregateRootBased
	Email             string
	Phone             string
	EncryptedPassword []byte
	FirstName         string
	LastName          string
	ConfirmationCode  []byte
	NeedConfirmation  bool
}

// RegisterUserWithData will create and init the User.
func RegisterUserWithData(email string, encrypedPassword []byte, firstName string, lastName string) *UserAggregate {
	user := new(UserAggregate)
	user.AggregateRootBased = goengine.NewAggregateRootBased(user)
	user.RecordThat(&AccountCreated{
		Email:             email,
		EncryptedPassword: encrypedPassword,
		FirstName:         firstName,
		LastName:          lastName,
		occurredOn:        time.Now(),
	})
	return user
}

//WhenAccountCreated process the AccountCreated event.
func (a *UserAggregate) WhenAccountCreated(event *AccountCreated) {
	a.Email = event.Email
	a.FirstName = event.FirstName
	a.LastName = event.LastName
	a.EncryptedPassword = event.EncryptedPassword
	a.ConfirmationCode = nil
	a.NeedConfirmation = false
	a.Phone = ""
}
