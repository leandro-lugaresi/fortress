package hydraidp

import (
	"context"
	"testing"

	eh "github.com/looplab/eventhorizon"
	"github.com/stretchr/testify/suite"
	"gopkg.in/hlandau/passlib.v1"
)

type UserAggregateTestSuite struct {
	suite.Suite
	UUID eh.UUID
	Hash string
	Agg  *UserAggregate
}

func (s *UserAggregateTestSuite) SetupTest() {
	s.UUID = eh.NewUUID()
	s.Agg = NewUserAggregate(s.UUID, &passlib.DefaultContext)
	var err error
	s.Hash, err = passlib.Hash("DoYouKnowME?")
	if err != nil {
		s.Error(err)
	}
	err = s.Agg.ApplyEvent(context.Background(), eh.NewEventForAggregate(
		AccountCreatedEvent,
		&AccountCreatedData{
			Email:             "jhon_doe@gmail.com",
			FirstName:         "Jhon",
			LastName:          "Doe",
			EncryptedPassword: s.Hash,
		},
		UserAggregateType,
		s.UUID,
		1,
	))
	if err != nil {
		s.Error(err, "UserAggregate.ApplyEvent() with error")
	}
}

func (s *UserAggregateTestSuite) TestUserCreateCommandShouldReturnUserCreatedEvent() {
	err := s.Agg.HandleCommand(context.Background(), &CreateAccount{
		Email:     "jhon_doe@gmail.com",
		FirstName: "Jhon",
		LastName:  "Doe",
		UserID:    s.UUID,
		Password:  "DoYouKnowME?",
	})
	if err != nil {
		s.Error(err, "UserAggregate.HandleCommand() with error")
	}
	changes := s.Agg.UncommittedEvents()
	s.EqualValues(1, len(changes), " Command CreateAccount should create just one event.")
	s.Equal(AccountCreatedEvent, changes[0].EventType())
	d := changes[0].Data().(*AccountCreatedData)
	s.Equal(d.Email, d.Email)
	s.Equal(d.FirstName, d.FirstName)
	s.Equal(d.LastName, d.LastName)
	new, err := s.Agg.pw.Verify("DoYouKnowME?", d.EncryptedPassword)
	s.NoError(err, "The hash is invalid and the passoword did not work as expected")
	s.Empty(new, "We just create the password with the default schemas, passlib should not return a new password")
}

func (s *UserAggregateTestSuite) TestApplyAccountCreatedEventShouldFillTheData() {
	s.Equal("jhon_doe@gmail.com", s.Agg.email)
	s.Equal("Jhon", s.Agg.firstName)
	s.Equal("Doe", s.Agg.lastName)
	s.Equal(s.Hash, s.Agg.encryptedPassword)
	s.Equal("", s.Agg.phone)
	s.Equal(false, s.Agg.needConfirmation)
	s.Equal("", s.Agg.confirmationCode)
}

func (s *UserAggregateTestSuite) TestChangeEmailShouldReturnUserChangeEmailEvent() {
	err := s.Agg.HandleCommand(context.Background(), &ChangeEmail{
		UserID: s.UUID,
		Email:  "jhon_doe123@gmail.com",
	})
	if err != nil {
		s.Error(err, "UserAggregate.HandleCommand() with error")
	}
	changes := s.Agg.UncommittedEvents()
	s.EqualValues(1, len(changes), " Command ChangeEmail should create just one event.")
	s.Equal(AccountEmailChangedEvent, changes[0].EventType())
	s.Equal(
		&AccountEmailChangedData{
			Email: "jhon_doe123@gmail.com",
		},
		changes[0].Data(),
	)
}

func (s *UserAggregateTestSuite) TestApplyAccountEmailChangedShouldUpdateTheEmail() {
	err := s.Agg.ApplyEvent(context.Background(), eh.NewEventForAggregate(
		AccountEmailChangedEvent,
		&AccountEmailChangedData{
			Email: "jhon_doe_fooo@gmail.com",
		},
		UserAggregateType,
		s.UUID,
		1,
	))
	if err != nil {
		s.Error(err, "UserAggregate.ApplyEvent() with error")
	}
	s.Equal("jhon_doe_fooo@gmail.com", s.Agg.email)
	s.Equal("Jhon", s.Agg.firstName)
	s.Equal("Doe", s.Agg.lastName)
	s.Equal(s.Hash, s.Agg.encryptedPassword)
	s.Equal("", s.Agg.phone)
	s.Equal(false, s.Agg.needConfirmation)
	s.Equal("", s.Agg.confirmationCode)
}

func (s *UserAggregateTestSuite) TestChangePasswordShouldWorksWhenThePassowrdIsRight() {
	err := s.Agg.HandleCommand(context.Background(), &ChangePassword{
		UserID:      s.UUID,
		Password:    "DoYouKnowME?",
		NewPassword: "NoIDoNot",
	})
	if err != nil {
		s.Error(err, "UserAggregate.HandleCommand() with error")
	}
	changes := s.Agg.UncommittedEvents()
	s.EqualValues(1, len(changes), " Command ChangePassword should create just one event.")
	s.Equal(AccountPasswordChangedEvent, changes[0].EventType())
	d := changes[0].Data().(*AccountPasswordChangedData)
	_, err = s.Agg.pw.Verify("NoIDoNot", d.EncryptedPassword)
	s.NoError(err, "The hash is invalid and the passoword did not work as expected")
}

func (s *UserAggregateTestSuite) TestApplyAccountPasswordChangedShouldUpdateThePassword() {
	h, err := s.Agg.pw.Hash("FoooBar")
	if err != nil {
		s.Error(err, "PasswordHasher with error")
	}
	err = s.Agg.ApplyEvent(context.Background(), eh.NewEventForAggregate(
		AccountPasswordChangedEvent,
		&AccountPasswordChangedData{
			EncryptedPassword: h,
		},
		UserAggregateType,
		s.UUID,
		1,
	))
	if err != nil {
		s.Error(err, "UserAggregate.ApplyEvent() with error")
	}
	s.Equal(h, s.Agg.encryptedPassword)
}
func TestUserAggregateTestSuite(t *testing.T) {
	suite.Run(t, new(UserAggregateTestSuite))
}
