// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Boards struct {
	Data  []*Board `json:"data,omitempty"`
	Count int      `json:"count"`
}

type Comment struct {
	ID string `json:"id"`
}

func (Comment) IsEntity() {}

type Goals struct {
	Count int     `json:"count"`
	Data  []*Goal `json:"data,omitempty"`
}

type Lists struct {
	Count int     `json:"count"`
	Data  []*List `json:"data,omitempty"`
}

type Metric struct {
	Precision *int     `json:"precision,omitempty"`
	Unit      *string  `json:"unit,omitempty"`
	Initial   *float64 `json:"initial,omitempty"`
	Target    *float64 `json:"target,omitempty"`
	Current   *float64 `json:"current,omitempty"`
	Source    *string  `json:"source,omitempty"`
}

type NewBoard struct {
	Type         string                 `json:"type"`
	User         *string                `json:"user,omitempty"`
	Organization *string                `json:"organization,omitempty"`
	Title        string                 `json:"title"`
	Description  *string                `json:"description,omitempty"`
	DueDate      *string                `json:"due_date,omitempty"`
	IsTemplate   bool                   `json:"is_template"`
	Starred      bool                   `json:"starred"`
	Order        int                    `json:"order"`
	Background   *string                `json:"background,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Status       string                 `json:"status"`
}

type NewGoal struct {
	Name         string                 `json:"name"`
	Notes        *string                `json:"notes,omitempty"`
	Time         string                 `json:"time"`
	IsCompany    *bool                  `json:"is_company,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Status       *string                `json:"status,omitempty"`
	User         string                 `json:"user"`
	Parent       *string                `json:"parent,omitempty"`
	Organization *string                `json:"organization,omitempty"`
}

type NewList struct {
	Name     string                 `json:"name"`
	Board    string                 `json:"board"`
	Order    *int                   `json:"order,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type NewTask struct {
	User       *string                `json:"user,omitempty"`
	Parent     *string                `json:"parent,omitempty"`
	List       *string                `json:"list,omitempty"`
	Name       string                 `json:"name"`
	Notes      *string                `json:"notes,omitempty"`
	Priority   *string                `json:"priority,omitempty"`
	Order      *int                   `json:"order,omitempty"`
	StartDate  *string                `json:"start_date,omitempty"`
	DueDate    *string                `json:"due_date,omitempty"`
	Labels     []*string              `json:"labels,omitempty"`
	Categories []string               `json:"categories,omitempty"`
	Status     *string                `json:"status,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type Tasks struct {
	Data  []*Task `json:"data,omitempty"`
	Count int     `json:"count"`
}

type TimeInput struct {
	Parent      *string                `json:"parent,omitempty"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	StartDate   string                 `json:"start_date"`
	EndDate     int                    `json:"end_date"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type Times struct {
	Data  []*Time `json:"data,omitempty"`
	Count int     `json:"count"`
}

type UpdateBoard struct {
	User         *string                `json:"user,omitempty"`
	Organization *string                `json:"organization,omitempty"`
	Title        *string                `json:"title,omitempty"`
	Description  *string                `json:"description,omitempty"`
	DueDate      *string                `json:"due_date,omitempty"`
	IsTemplate   *bool                  `json:"is_template,omitempty"`
	Starred      *bool                  `json:"starred,omitempty"`
	Order        *int                   `json:"order,omitempty"`
	Background   *string                `json:"background,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Status       *string                `json:"status,omitempty"`
}

type UpdateGoal struct {
	Name         *string                `json:"name,omitempty"`
	Notes        *string                `json:"notes,omitempty"`
	Time         *string                `json:"time,omitempty"`
	IsCompany    *bool                  `json:"is_company,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	Status       *string                `json:"status,omitempty"`
	User         *string                `json:"user,omitempty"`
	Parent       *string                `json:"parent,omitempty"`
	Organization *string                `json:"organization,omitempty"`
}

type UpdateList struct {
	Name     *string                `json:"name,omitempty"`
	Order    *int                   `json:"order,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateTask struct {
	User       *string                `json:"user,omitempty"`
	Parent     *string                `json:"parent,omitempty"`
	List       *string                `json:"list,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Notes      *string                `json:"notes,omitempty"`
	Priority   *string                `json:"priority,omitempty"`
	Order      *int                   `json:"order,omitempty"`
	StartDate  *string                `json:"start_date,omitempty"`
	DueDate    *string                `json:"due_date,omitempty"`
	Labels     []*string              `json:"labels,omitempty"`
	Categories []string               `json:"categories,omitempty"`
	Status     *string                `json:"status,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type BoardType string

const (
	BoardTypeRequest   BoardType = "REQUEST"
	BoardTypeSales     BoardType = "SALES"
	BoardTypeProject   BoardType = "PROJECT"
	BoardTypeMarketing BoardType = "MARKETING"
)

var AllBoardType = []BoardType{
	BoardTypeRequest,
	BoardTypeSales,
	BoardTypeProject,
	BoardTypeMarketing,
}

func (e BoardType) IsValid() bool {
	switch e {
	case BoardTypeRequest, BoardTypeSales, BoardTypeProject, BoardTypeMarketing:
		return true
	}
	return false
}

func (e BoardType) String() string {
	return string(e)
}

func (e *BoardType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BoardType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BoardType", str)
	}
	return nil
}

func (e BoardType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type GoalStatus string

const (
	GoalStatusInProgress GoalStatus = "IN_PROGRESS"
	GoalStatusCompleted  GoalStatus = "COMPLETED"
	GoalStatusOnHold     GoalStatus = "ON_HOLD"
	GoalStatusCanceled   GoalStatus = "CANCELED"
)

var AllGoalStatus = []GoalStatus{
	GoalStatusInProgress,
	GoalStatusCompleted,
	GoalStatusOnHold,
	GoalStatusCanceled,
}

func (e GoalStatus) IsValid() bool {
	switch e {
	case GoalStatusInProgress, GoalStatusCompleted, GoalStatusOnHold, GoalStatusCanceled:
		return true
	}
	return false
}

func (e GoalStatus) String() string {
	return string(e)
}

func (e *GoalStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GoalStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GoalStatus", str)
	}
	return nil
}

func (e GoalStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "PENDING"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusCompleted  TaskStatus = "COMPLETED"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusPending,
	TaskStatusInProgress,
	TaskStatusCompleted,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusPending, TaskStatusInProgress, TaskStatusCompleted:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
