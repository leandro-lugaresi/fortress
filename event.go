package hydraidp

import "time"

// AccountCreated is the event data for when an account has been created.
type AccountCreated struct {
	occurredOn        time.Time
	Email             string
	EncryptedPassword []byte
	FirstName         string
	LastName          string
}

// OccurredOn return the time when the event occurred. Its used to implement the goengine.DomainEvent
func (e *AccountCreated) OccurredOn() time.Time {
	return e.occurredOn
}

// AccountEmailChanged is the event data for when an email change has completed.
type AccountEmailChanged struct {
	occurredOn time.Time
	NewEmail   string
	OldEmail   string
}

// OccurredOn return the time when the event occurred. Its used to implement the goengine.DomainEvent
func (e *AccountEmailChanged) OccurredOn() time.Time {
	return e.occurredOn
}

// AccountPasswordChanged is the event data for when an password has been changed.
type AccountPasswordChanged struct {
	occurredOn        time.Time
	EncryptedPassword []byte
}

// OccurredOn return the time when the event occurred. Its used to implement the goengine.DomainEvent
func (e *AccountPasswordChanged) OccurredOn() time.Time {
	return e.occurredOn
}

// AccountPhoneAdded is the event data for when an phone has been added.
type AccountPhoneAdded struct {
	occurredOn time.Time
	Phone      string
}

// OccurredOn return the time when the event occurred. Its used to implement the goengine.DomainEvent
func (e *AccountPhoneAdded) OccurredOn() time.Time {
	return e.occurredOn
}

// AccountPhoneEdited is the event data for when an phone has been edited.
type AccountPhoneEdited struct {
	occurredOn time.Time
	Phone      string
}

// OccurredOn return the time when the event occurred. Its used to implement the goengine.DomainEvent
func (e *AccountPhoneEdited) OccurredOn() time.Time {
	return e.occurredOn
}

// AccountConfirmationRequested is the event for when a email needs a confirmation request.
type AccountConfirmationRequested struct {
	occurredOn       time.Time
	ConfirmationHash []byte
}

// OccurredOn return the time when the event occurred. Its used to implement the goengine.DomainEvent
func (e *AccountConfirmationRequested) OccurredOn() time.Time {
	return e.occurredOn
}
