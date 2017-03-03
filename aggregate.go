package hydraidp

import (
	"context"
	"fmt"

	eh "github.com/looplab/eventhorizon"
	"gopkg.in/hlandau/passlib.v1"
)

func init() {
	eh.RegisterAggregate(func(id eh.UUID) eh.Aggregate {
		return NewUserAggregate(id, &passlib.DefaultContext)
	})
}

// UserAggregateType is the type name of the user aggregate.
const UserAggregateType eh.AggregateType = "User"

// UserAggregate represent one user in the identity provider.
type UserAggregate struct {
	// AggregateBase implements most of the eventhorizon.Aggregate interface.
	*eh.AggregateBase
	pw                PasswordHasher
	email             string
	phone             string
	encryptedPassword string
	firstName         string
	lastName          string
	confirmationCode  string
	needConfirmation  bool
}

// NewUserAggregate creates a new UserAggregate with an ID.
func NewUserAggregate(id eh.UUID, pw PasswordHasher) *UserAggregate {
	return &UserAggregate{
		AggregateBase: eh.NewAggregateBase(UserAggregateType, id),
		pw:            pw,
	}
}

// HandleCommand implements the HandleCommand method of the Aggregate interface.
func (a *UserAggregate) HandleCommand(ctx context.Context, command eh.Command) error {
	switch command := command.(type) {
	case *CreateAccount:
		h, err := a.pw.Hash(command.Password)
		if err != nil {
			return err
		}
		a.StoreEvent(AccountCreatedEvent, &AccountCreatedData{
			Email:             command.Email,
			FirstName:         command.FirstName,
			LastName:          command.LastName,
			EncryptedPassword: h,
		})
		return nil
	case *ChangeEmail:
		a.StoreEvent(AccountEmailChangedEvent, &AccountEmailChangedData{
			Email: command.Email,
		})
	case *ChangePassword:
		if _, err := a.pw.Verify(command.Password, a.encryptedPassword); err != nil {
			return err
		}
		h, err := a.pw.Hash(command.NewPassword)
		if err != nil {
			return err
		}
		a.StoreEvent(AccountPasswordChangedEvent, &AccountPasswordChangedData{
			EncryptedPassword: h,
		})
	}
	return fmt.Errorf("Invalid command, type received: %T", command)
}

// ApplyEvent implements the ApplyEvent method of the Aggregate interface.
func (a *UserAggregate) ApplyEvent(ctx context.Context, event eh.Event) error {
	switch event.EventType() {
	case AccountCreatedEvent:
		if data, ok := event.Data().(*AccountCreatedData); ok {
			a.email = data.Email
			a.firstName = data.FirstName
			a.lastName = data.LastName
			a.phone = ""
			a.encryptedPassword = data.EncryptedPassword
			a.needConfirmation = false
			a.confirmationCode = ""
			return nil
		}
		return fmt.Errorf("invalid event data type: %T", event.Data())
	case AccountEmailChangedEvent:
		if data, ok := event.Data().(*AccountEmailChangedData); ok {
			a.email = data.Email
			return nil
		}
		return fmt.Errorf("invalid event data type: %T", event.Data())
	case AccountPasswordChangedEvent:
		if data, ok := event.Data().(*AccountPasswordChangedData); ok {
			a.encryptedPassword = data.EncryptedPassword
		}
		return fmt.Errorf("invalid event data type: %T", event.Data())
	}
	return nil
}
