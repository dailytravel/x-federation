// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Activities struct {
	Data  []*Activity `json:"data,omitempty"`
	Count int         `json:"count"`
}

type Activity struct {
	ID        string                 `json:"id"`
	User      *User                  `json:"user"`
	Object    map[string]interface{} `json:"object"`
	Comment   *Comment               `json:"comment,omitempty"`
	Action    string                 `json:"action"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Status    string                 `json:"status"`
	CreatedAt string                 `json:"created_at"`
}

type Campaign struct {
	ID string `json:"id"`
}

func (Campaign) IsEntity() {}

type Comment struct {
	ID string `json:"id"`
}

func (Comment) IsEntity() {}

type Follow struct {
	ID string `json:"id"`
}

func (Follow) IsEntity() {}

type Log struct {
	ID        string                 `json:"id"`
	Message   *string                `json:"message,omitempty"`
	Method    string                 `json:"method"`
	Latency   int                    `json:"latency"`
	Path      string                 `json:"path"`
	Status    string                 `json:"status"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Timestamp int                    `json:"Timestamp"`
	User      *User                  `json:"user,omitempty"`
	URL       string                 `json:"url"`
	UserAgent string                 `json:"user_agent"`
	ClientIP  string                 `json:"client_ip"`
}

type Logs struct {
	Data  []*Log `json:"data,omitempty"`
	Count int    `json:"count"`
}

type NewResponse struct {
	Type      string                 `json:"type"`
	Reference string                 `json:"reference"`
	Campaign  string                 `json:"campaign"`
	Metadata  map[string]interface{} `json:"metadata"`
}

type Response struct {
	ID        string                 `json:"id"`
	Reference *string                `json:"reference,omitempty"`
	Campaign  *Campaign              `json:"campaign"`
	Type      string                 `json:"type"`
	UserAgent string                 `json:"user_agent"`
	Metadata  map[string]interface{} `json:"metadata"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

func (Response) IsEntity() {}

type Responses struct {
	Count int         `json:"count"`
	Data  []*Response `json:"data,omitempty"`
}

type UpdateResponse struct {
	Type      *string                `json:"type,omitempty"`
	Reference *string                `json:"reference,omitempty"`
	Campaign  *string                `json:"campaign,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type User struct {
	ID string `json:"id"`
}

func (User) IsEntity() {}

type ActivityAction string

const (
	ActivityActionCreated   ActivityAction = "CREATED"
	ActivityActionUpdated   ActivityAction = "UPDATED"
	ActivityActionDeleted   ActivityAction = "DELETED"
	ActivityActionCommented ActivityAction = "COMMENTED"
	ActivityActionCompleted ActivityAction = "COMPLETED"
	ActivityActionAssigned  ActivityAction = "ASSIGNED"
	ActivityActionApproved  ActivityAction = "APPROVED"
	ActivityActionRejected  ActivityAction = "REJECTED"
	ActivityActionArchived  ActivityAction = "ARCHIVED"
	ActivityActionRestored  ActivityAction = "RESTORED"
)

var AllActivityAction = []ActivityAction{
	ActivityActionCreated,
	ActivityActionUpdated,
	ActivityActionDeleted,
	ActivityActionCommented,
	ActivityActionCompleted,
	ActivityActionAssigned,
	ActivityActionApproved,
	ActivityActionRejected,
	ActivityActionArchived,
	ActivityActionRestored,
}

func (e ActivityAction) IsValid() bool {
	switch e {
	case ActivityActionCreated, ActivityActionUpdated, ActivityActionDeleted, ActivityActionCommented, ActivityActionCompleted, ActivityActionAssigned, ActivityActionApproved, ActivityActionRejected, ActivityActionArchived, ActivityActionRestored:
		return true
	}
	return false
}

func (e ActivityAction) String() string {
	return string(e)
}

func (e *ActivityAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActivityAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActivityAction", str)
	}
	return nil
}

func (e ActivityAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RequestMethod string

const (
	RequestMethodGet     RequestMethod = "GET"
	RequestMethodPost    RequestMethod = "POST"
	RequestMethodPut     RequestMethod = "PUT"
	RequestMethodDelete  RequestMethod = "DELETE"
	RequestMethodHead    RequestMethod = "HEAD"
	RequestMethodOptions RequestMethod = "OPTIONS"
	RequestMethodPatch   RequestMethod = "PATCH"
)

var AllRequestMethod = []RequestMethod{
	RequestMethodGet,
	RequestMethodPost,
	RequestMethodPut,
	RequestMethodDelete,
	RequestMethodHead,
	RequestMethodOptions,
	RequestMethodPatch,
}

func (e RequestMethod) IsValid() bool {
	switch e {
	case RequestMethodGet, RequestMethodPost, RequestMethodPut, RequestMethodDelete, RequestMethodHead, RequestMethodOptions, RequestMethodPatch:
		return true
	}
	return false
}

func (e RequestMethod) String() string {
	return string(e)
}

func (e *RequestMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RequestMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RequestMethod", str)
	}
	return nil
}

func (e RequestMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
