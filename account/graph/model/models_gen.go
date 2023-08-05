// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Board struct {
	ID string `json:"id"`
}

func (Board) IsEntity() {}

type Clients struct {
	Data  []*Client `json:"data,omitempty"`
	Count int       `json:"count"`
}

type Contact struct {
	ID string `json:"id"`
}

func (Contact) IsEntity() {}

type Identity struct {
	ID          string                 `json:"id"`
	User        *User                  `json:"user"`
	Provider    string                 `json:"provider"`
	AccessToken string                 `json:"access_token"`
	ExpiresIn   int                    `json:"expires_in"`
	Connection  string                 `json:"connection"`
	IsSocial    bool                   `json:"is_social"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      string                 `json:"status"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

type Invitation struct {
	ID        string                 `json:"id"`
	Sender    *User                  `json:"sender"`
	Recipient string                 `json:"recipient"`
	Roles     []string               `json:"roles"`
	Status    string                 `json:"status"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

type Invitations struct {
	Data  []*Invitation `json:"data,omitempty"`
	Count int           `json:"count"`
}

type Keys struct {
	Data  []*Key `json:"data,omitempty"`
	Count int    `json:"count"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mfa struct {
	Enabled bool   `json:"enabled"`
	Code    string `json:"code"`
}

type MFAInput struct {
	Enabled bool   `json:"enabled"`
	Code    string `json:"code"`
}

type Membership struct {
	ID string `json:"id"`
}

func (Membership) IsEntity() {}

type NewClient struct {
	Type        string                 `json:"type"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Domains     []*string              `json:"domains,omitempty"`
	Redirect    string                 `json:"redirect"`
	Permissions []*string              `json:"permissions,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewInvitation struct {
	Email     string                 `json:"email"`
	FirstName *string                `json:"firstName,omitempty"`
	LastName  *string                `json:"lastName,omitempty"`
	Role      string                 `json:"role"`
	Team      string                 `json:"team"`
	Inviter   string                 `json:"inviter"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type NewKey struct {
	Name string `json:"name"`
	Kid  string `json:"kid"`
}

type NewPermission struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type NewRole struct {
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Permissions []*string `json:"permissions,omitempty"`
}

type NewUser struct {
	Locale   string                 `json:"locale"`
	Name     string                 `json:"name"`
	Email    string                 `json:"email"`
	Phone    *string                `json:"phone,omitempty"`
	Password string                 `json:"password"`
	Roles    []*string              `json:"roles,omitempty"`
	Timezone *string                `json:"timezone,omitempty"`
	Status   *string                `json:"status,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Notifications struct {
	Data  []*Notification `json:"data,omitempty"`
	Count int             `json:"count"`
}

type Organization struct {
	ID string `json:"id"`
}

func (Organization) IsEntity() {}

type PasswordInput struct {
	CurrentPassword      string `json:"current_password"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type Payload struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

type Permissions struct {
	Data  []*Permission `json:"data,omitempty"`
	Count int           `json:"count"`
}

type Point struct {
	ID string `json:"id"`
}

func (Point) IsEntity() {}

type RegisterInput struct {
	Locale   string `json:"locale"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Roles struct {
	Data  []*Role `json:"data,omitempty"`
	Count int     `json:"count"`
}

type UpdateClient struct {
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Domains     []*string              `json:"domains,omitempty"`
	Redirect    *string                `json:"redirect,omitempty"`
	Permissions []string               `json:"permissions,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	ExpiresAt   *string                `json:"expires_at,omitempty"`
}

type UpdateInvitation struct {
	Email     *string                `json:"email,omitempty"`
	FirstName *string                `json:"firstName,omitempty"`
	LastName  *string                `json:"lastName,omitempty"`
	Role      *string                `json:"role,omitempty"`
	Status    *string                `json:"status,omitempty"`
	Team      *string                `json:"team,omitempty"`
	Inviter   *string                `json:"inviter,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type UpdatePermission struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateRole struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Permissions []*string `json:"permissions,omitempty"`
}

type UpdateUser struct {
	Locale   *string                `json:"locale,omitempty"`
	Name     *string                `json:"name,omitempty"`
	Email    *string                `json:"email,omitempty"`
	Phone    *string                `json:"phone,omitempty"`
	Roles    []*string              `json:"roles,omitempty"`
	Photos   []*string              `json:"photos,omitempty"`
	Timezone *string                `json:"timezone,omitempty"`
	Status   *string                `json:"status,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Users struct {
	Count int     `json:"count"`
	Data  []*User `json:"data,omitempty"`
}

type VerifyEmailInput struct {
	Token string `json:"token"`
}

type InvitationStatus string

const (
	InvitationStatusPending  InvitationStatus = "PENDING"
	InvitationStatusAccepted InvitationStatus = "ACCEPTED"
	InvitationStatusDeclined InvitationStatus = "DECLINED"
)

var AllInvitationStatus = []InvitationStatus{
	InvitationStatusPending,
	InvitationStatusAccepted,
	InvitationStatusDeclined,
}

func (e InvitationStatus) IsValid() bool {
	switch e {
	case InvitationStatusPending, InvitationStatusAccepted, InvitationStatusDeclined:
		return true
	}
	return false
}

func (e InvitationStatus) String() string {
	return string(e)
}

func (e *InvitationStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InvitationStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InvitationStatus", str)
	}
	return nil
}

func (e InvitationStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SocialProvider string

const (
	SocialProviderFacebook SocialProvider = "FACEBOOK"
	SocialProviderGoogle   SocialProvider = "GOOGLE"
	SocialProviderTwitter  SocialProvider = "TWITTER"
	SocialProviderGithub   SocialProvider = "GITHUB"
)

var AllSocialProvider = []SocialProvider{
	SocialProviderFacebook,
	SocialProviderGoogle,
	SocialProviderTwitter,
	SocialProviderGithub,
}

func (e SocialProvider) IsValid() bool {
	switch e {
	case SocialProviderFacebook, SocialProviderGoogle, SocialProviderTwitter, SocialProviderGithub:
		return true
	}
	return false
}

func (e SocialProvider) String() string {
	return string(e)
}

func (e *SocialProvider) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SocialProvider(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SocialProvider", str)
	}
	return nil
}

func (e SocialProvider) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
