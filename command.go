package hydraidp

import (
	eh "github.com/looplab/eventhorizon"
)

func init() {
	eh.RegisterCommand(func() eh.Command { return &CreateAccount{} })
	eh.RegisterCommand(func() eh.Command { return &ChangeEmail{} })
	eh.RegisterCommand(func() eh.Command { return &ConfirmPhone{} })
	eh.RegisterCommand(func() eh.Command { return &ConfirmEmail{} })
	eh.RegisterCommand(func() eh.Command { return &ChangePassword{} })
	eh.RegisterCommand(func() eh.Command { return &AddPhone{} })
	eh.RegisterCommand(func() eh.Command { return &EditPhone{} })
}

// This consts represent all the command types used to registry commands.
const (
	CreateAccountCommand  eh.CommandType = "CreateAccount"
	ChangeEmailCommand    eh.CommandType = "ChangeEmail"
	ConfirmPhoneCommand   eh.CommandType = "ConfirmPhone"
	ConfirmEmailCommand   eh.CommandType = "ConfirmEmail"
	ChangePasswordCommand eh.CommandType = "ChangePassword"
	LoginCommand          eh.CommandType = "Login"
	AddPhoneCommand       eh.CommandType = "AddPhone"
	EditPhoneCommand      eh.CommandType = "EditPhone"
)

// CreateAccount is a command for create accounts.
type CreateAccount struct {
	UserID    eh.UUID
	Email     string
	Password  string
	FirstName string
	LastName  string
}

// AggregateID returns the ID of the aggregate.
func (c CreateAccount) AggregateID() eh.UUID { return c.UserID }

// AggregateType returns the type of the aggregate.
func (c CreateAccount) AggregateType() eh.AggregateType { return UserAggregateType }

// CommandType returns the type of the command.
func (c CreateAccount) CommandType() eh.CommandType { return CreateAccountCommand }

// ChangeEmail is a command for change the email.
type ChangeEmail struct {
	UserID eh.UUID
	Email  string
}

// AggregateID returns the ID of the aggregate.
func (c ChangeEmail) AggregateID() eh.UUID { return c.UserID }

// AggregateType returns the type of the aggregate.
func (c ChangeEmail) AggregateType() eh.AggregateType { return UserAggregateType }

// CommandType returns the type of the command.
func (c ChangeEmail) CommandType() eh.CommandType { return ChangeEmailCommand }

// ConfirmPhone is a command for confirm the account phone.
type ConfirmPhone struct {
	UserID            eh.UUID
	ConfirmationToken string
}

// AggregateID returns the ID of the aggregate.
func (c ConfirmPhone) AggregateID() eh.UUID { return c.UserID }

// AggregateType returns the type of the aggregate.
func (c ConfirmPhone) AggregateType() eh.AggregateType { return UserAggregateType }

// CommandType returns the type of the command.
func (c ConfirmPhone) CommandType() eh.CommandType { return ConfirmPhoneCommand }

// ConfirmEmail is a command for confirm the email account.
type ConfirmEmail struct {
	UserID            eh.UUID
	ConfirmationToken string
}

// AggregateID returns the ID of the aggregate.
func (c ConfirmEmail) AggregateID() eh.UUID { return c.UserID }

// AggregateType returns the type of the aggregate.
func (c ConfirmEmail) AggregateType() eh.AggregateType { return UserAggregateType }

// CommandType returns the type of the command.
func (c ConfirmEmail) CommandType() eh.CommandType { return ConfirmEmailCommand }

// ChangePassword is a command for change the password.
type ChangePassword struct {
	UserID      eh.UUID
	Password    string
	NewPassword string
}

// AggregateID returns the ID of the aggregate.
func (c ChangePassword) AggregateID() eh.UUID { return c.UserID }

// AggregateType returns the type of the aggregate.
func (c ChangePassword) AggregateType() eh.AggregateType { return UserAggregateType }

// CommandType returns the type of the command.
func (c ChangePassword) CommandType() eh.CommandType { return ChangePasswordCommand }

// Login is a command for login.
type Login struct {
	UserID eh.UUID
}

// AggregateID returns the ID of the aggregate.
func (c Login) AggregateID() eh.UUID { return c.UserID }

// AggregateType returns the type of the aggregate.
func (c Login) AggregateType() eh.AggregateType { return UserAggregateType }

// CommandType returns the type of the command.
func (c Login) CommandType() eh.CommandType { return LoginCommand }

// AddPhone is a command for add a phone number for this account.
type AddPhone struct {
	UserID eh.UUID
	Phone  string
}

// AggregateID returns the ID of the aggregate.
func (c AddPhone) AggregateID() eh.UUID { return c.UserID }

// AggregateType returns the type of the aggregate.
func (c AddPhone) AggregateType() eh.AggregateType { return UserAggregateType }

// CommandType returns the type of the command.
func (c AddPhone) CommandType() eh.CommandType { return AddPhoneCommand }

// EditPhone is a command for change the phone number.
type EditPhone struct {
	UserID eh.UUID
	Phone  string
}

// AggregateID returns the ID of the aggregate.
func (c EditPhone) AggregateID() eh.UUID { return c.UserID }

// AggregateType returns the type of the aggregate.
func (c EditPhone) AggregateType() eh.AggregateType { return UserAggregateType }

// CommandType returns the type of the command.
func (c EditPhone) CommandType() eh.CommandType { return EditPhoneCommand }
