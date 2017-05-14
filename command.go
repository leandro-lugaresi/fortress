package hydraidp

// CreateAccount is a command for create accounts.
type CreateAccount struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

// ChangeEmail is a command for change the email.
type ChangeEmail struct {
	Email string
}

// ConfirmPhone is a command for confirm the account phone.
type ConfirmPhone struct {
	ConfirmationToken []byte
}

// ConfirmEmail is a command for confirm the email account.
type ConfirmEmail struct {
	ConfirmationToken []byte
}

// ChangePassword is a command for change the password.
type ChangePassword struct {
	Password    []byte
	NewPassword []byte
}

// AddPhone is a command for add a phone number for this account.
type AddPhone struct {
	Phone string
}

// EditPhone is a command for change the phone number.
type EditPhone struct {
	OldPhone string
	NewPhone string
}
