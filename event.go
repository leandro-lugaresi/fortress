package hydraidp

import (
	eh "github.com/looplab/eventhorizon"
)

// Constants represeting system events.
const (
	AccountCreatedEvent               eh.EventType = "AccountCreated"
	AccountConfirmationRequestedEvent eh.EventType = "AccountConfirmationRequested"
	AccountConfirmedEvent             eh.EventType = "AccountConfirmed"
	AccountEmailChangedEvent          eh.EventType = "AccountEmailChanged"
	AccountPasswordChangedEvent       eh.EventType = "AccountPasswordChanged"
	AccountLoggedInEvent              eh.EventType = "AccountLoggedIn"
	AccountAddedOTPAuthEvent          eh.EventType = "AccountAddedOTPAuth"
	AccountScratchCodeUsedEvent       eh.EventType = "AccountScratchCodeUsed"
	AccountOTPAuthRemovedEvent        eh.EventType = "AccountOTPAuthRemoved"
	AccountPhoneAddedEvent            eh.EventType = "AccountPhoneAdded"
	AccountPhoneEditedEvent           eh.EventType = "AccountPhoneEdited"
)

func init() {
	eh.RegisterEventData(AccountCreatedEvent, func() eh.EventData {
		return &AccountCreatedData{}
	})
	eh.RegisterEventData(AccountConfirmationRequestedEvent, func() eh.EventData {
		return &AccountConfirmationRequestedData{}
	})
	eh.RegisterEventData(AccountEmailChangedEvent, func() eh.EventData {
		return &AccountEmailChangedData{}
	})
	eh.RegisterEventData(AccountPasswordChangedEvent, func() eh.EventData {
		return &AccountPasswordChangedData{}
	})
	eh.RegisterEventData(AccountPhoneAddedEvent, func() eh.EventData {
		return &AccountPhoneAddedData{}
	})
	eh.RegisterEventData(AccountPhoneEditedEvent, func() eh.EventData {
		return &AccountPhoneEditedData{}
	})
}

// AccountCreatedData is the event data for when an account has been created.
type AccountCreatedData struct {
	Email             string
	EncryptedPassword string
	FirstName         string
	LastName          string
}

// AccountEmailChangedData is the event data for when an email change has completed.
type AccountEmailChangedData struct {
	Email string
}

// AccountPasswordChangedData is the event data for when an password has been changed.
type AccountPasswordChangedData struct {
	EncryptedPassword string
}

// AccountPhoneAddedData is the event data for when an phone has been added.
type AccountPhoneAddedData struct {
	Phone string
}

// AccountPhoneEditedData is the event data for when an phone has been edited.
type AccountPhoneEditedData struct {
	Phone string
}

// AccountConfirmationRequestedData is the event for when a email needs a confirmation request.
type AccountConfirmationRequestedData struct {
	ConfirmationHash string
}
