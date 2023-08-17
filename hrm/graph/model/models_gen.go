// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Attendances struct {
	Data  []*Attendance `json:"data,omitempty"`
	Count int           `json:"count"`
}

type CreateResume struct {
	Title          string `json:"title"`
	Summary        string `json:"summary"`
	Experience     string `json:"experience"`
	Education      string `json:"education"`
	Skills         string `json:"skills"`
	Certifications string `json:"certifications"`
	Languages      string `json:"languages"`
	Projects       string `json:"projects"`
	References     string `json:"references"`
}

type Employees struct {
	Data  []*Employee `json:"data,omitempty"`
	Count int         `json:"count"`
}

type Leaves struct {
	Data  []*Leave `json:"data,omitempty"`
	Count int      `json:"count"`
}

type NewAttendance struct {
	Owner    string                 `json:"owner"`
	TimeIn   string                 `json:"time_in"`
	TimeOut  string                 `json:"time_out"`
	Notes    *string                `json:"notes,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   string                 `json:"status"`
}

type NewEmployee struct {
	Reference      string                 `json:"reference"`
	FirstName      string                 `json:"first_name"`
	LastName       string                 `json:"last_name"`
	Email          string                 `json:"email"`
	Phone          string                 `json:"phone"`
	Address        string                 `json:"address"`
	Birthday       *string                `json:"birthday,omitempty"`
	Status         *string                `json:"status,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	HireDate       *string                `json:"hire_date,omitempty"`
	UserID         *string                `json:"user_id,omitempty"`
	OrganizationID string                 `json:"organization_id"`
	PositionID     *string                `json:"position_id,omitempty"`
}

type NewLeave struct {
	Employee  string                 `json:"employee"`
	Type      string                 `json:"type"`
	StartDate string                 `json:"start_date"`
	EndDate   *string                `json:"end_date,omitempty"`
	Reason    string                 `json:"reason"`
	Status    string                 `json:"status"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type NewOrganization struct {
	Employee    *string                `json:"employee,omitempty"`
	Parent      *string                `json:"parent,omitempty"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Status      *string                `json:"status,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewPayroll struct {
	Employee string                 `json:"employee"`
	PayDate  string                 `json:"pay_date"`
	Amount   float64                `json:"amount"`
	Currency string                 `json:"currency"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   string                 `json:"status"`
}

type NewPosition struct {
	Locale      string                 `json:"locale"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Salary      string                 `json:"salary"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      string                 `json:"status"`
}

type NewSalary struct {
	Employee  string                 `json:"employee"`
	Amount    float64                `json:"amount"`
	Currency  string                 `json:"currency"`
	StartDate string                 `json:"start_date"`
	EndDate   *string                `json:"end_date,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type Organizations struct {
	Data  []*Organization `json:"data,omitempty"`
	Count int             `json:"count"`
}

type Payrolls struct {
	Data  []*Payroll `json:"data,omitempty"`
	Count int        `json:"count"`
}

type Positions struct {
	Data  []*Position `json:"data,omitempty"`
	Count int         `json:"count"`
}

type Resumes struct {
	Data  []*Resume `json:"data,omitempty"`
	Count int       `json:"count"`
}

type Salaries struct {
	Data  []*Salary `json:"data,omitempty"`
	Count int       `json:"count"`
}

type UpdateAttendance struct {
	Owner    *string                `json:"owner,omitempty"`
	TimeIn   *string                `json:"time_in,omitempty"`
	TimeOut  *string                `json:"time_out,omitempty"`
	Notes    *string                `json:"notes,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   *string                `json:"status,omitempty"`
}

type UpdateEmployee struct {
	Reference      *string                `json:"reference,omitempty"`
	FirstName      *string                `json:"first_name,omitempty"`
	LastName       *string                `json:"last_name,omitempty"`
	Email          *string                `json:"email,omitempty"`
	Phone          *string                `json:"phone,omitempty"`
	Address        *string                `json:"address,omitempty"`
	Birthday       *string                `json:"birthday,omitempty"`
	Status         *string                `json:"status,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	HireDate       *string                `json:"hire_date,omitempty"`
	UserID         *string                `json:"user_id,omitempty"`
	OrganizationID *string                `json:"organization_id,omitempty"`
	PositionID     *string                `json:"position_id,omitempty"`
}

type UpdateLeave struct {
	Type      *string                `json:"type,omitempty"`
	StartDate *string                `json:"start_date,omitempty"`
	EndDate   *string                `json:"end_date,omitempty"`
	Reason    *string                `json:"reason,omitempty"`
	Status    *string                `json:"status,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateOrganization struct {
	Employee    *string                `json:"employee,omitempty"`
	Parent      *string                `json:"parent,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdatePayroll struct {
	PayDate  *string                `json:"pay_date,omitempty"`
	Amount   *float64               `json:"amount,omitempty"`
	Currency *string                `json:"currency,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status   *string                `json:"status,omitempty"`
}

type UpdatePosition struct {
	Locale      string                 `json:"locale"`
	Title       *string                `json:"title,omitempty"`
	Description *string                `json:"description,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Salary      *string                `json:"salary,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
}

type UpdateResume struct {
	Title          *string `json:"title,omitempty"`
	Summary        *string `json:"summary,omitempty"`
	Experience     *string `json:"experience,omitempty"`
	Education      *string `json:"education,omitempty"`
	Skills         *string `json:"skills,omitempty"`
	Certifications *string `json:"certifications,omitempty"`
	Languages      *string `json:"languages,omitempty"`
	Projects       *string `json:"projects,omitempty"`
	References     *string `json:"references,omitempty"`
}

type UpdateSalary struct {
	Amount    *float64               `json:"amount,omitempty"`
	Currency  *string                `json:"currency,omitempty"`
	StartDate *string                `json:"start_date,omitempty"`
	EndDate   *string                `json:"end_date,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type AttendanceStatus string

const (
	AttendanceStatusPresent AttendanceStatus = "PRESENT"
	AttendanceStatusAbsent  AttendanceStatus = "ABSENT"
	AttendanceStatusLate    AttendanceStatus = "LATE"
	AttendanceStatusExcused AttendanceStatus = "EXCUSED"
)

var AllAttendanceStatus = []AttendanceStatus{
	AttendanceStatusPresent,
	AttendanceStatusAbsent,
	AttendanceStatusLate,
	AttendanceStatusExcused,
}

func (e AttendanceStatus) IsValid() bool {
	switch e {
	case AttendanceStatusPresent, AttendanceStatusAbsent, AttendanceStatusLate, AttendanceStatusExcused:
		return true
	}
	return false
}

func (e AttendanceStatus) String() string {
	return string(e)
}

func (e *AttendanceStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AttendanceStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AttendanceStatus", str)
	}
	return nil
}

func (e AttendanceStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LeaveStatus string

const (
	LeaveStatusPending  LeaveStatus = "pending"
	LeaveStatusApproved LeaveStatus = "approved"
	LeaveStatusRejected LeaveStatus = "rejected"
)

var AllLeaveStatus = []LeaveStatus{
	LeaveStatusPending,
	LeaveStatusApproved,
	LeaveStatusRejected,
}

func (e LeaveStatus) IsValid() bool {
	switch e {
	case LeaveStatusPending, LeaveStatusApproved, LeaveStatusRejected:
		return true
	}
	return false
}

func (e LeaveStatus) String() string {
	return string(e)
}

func (e *LeaveStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeaveStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeaveStatus", str)
	}
	return nil
}

func (e LeaveStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LeaveType string

const (
	LeaveTypeParental      LeaveType = "Parental"
	LeaveTypeSick          LeaveType = "Sick"
	LeaveTypeBereavement   LeaveType = "Bereavement"
	LeaveTypeSabbatical    LeaveType = "Sabbatical"
	LeaveTypeHoliday       LeaveType = "holiday"
	LeaveTypeLeave         LeaveType = "Leave"
	LeaveTypeAnnual        LeaveType = "Annual"
	LeaveTypeUnpaid        LeaveType = "Unpaid"
	LeaveTypeMilitary      LeaveType = "Military"
	LeaveTypeMarriage      LeaveType = "Marriage"
	LeaveTypeCompensatory  LeaveType = "Compensatory"
	LeaveTypeReligion      LeaveType = "Religion"
	LeaveTypeEarned        LeaveType = "Earned"
	LeaveTypeCompassionate LeaveType = "compassionate"
)

var AllLeaveType = []LeaveType{
	LeaveTypeParental,
	LeaveTypeSick,
	LeaveTypeBereavement,
	LeaveTypeSabbatical,
	LeaveTypeHoliday,
	LeaveTypeLeave,
	LeaveTypeAnnual,
	LeaveTypeUnpaid,
	LeaveTypeMilitary,
	LeaveTypeMarriage,
	LeaveTypeCompensatory,
	LeaveTypeReligion,
	LeaveTypeEarned,
	LeaveTypeCompassionate,
}

func (e LeaveType) IsValid() bool {
	switch e {
	case LeaveTypeParental, LeaveTypeSick, LeaveTypeBereavement, LeaveTypeSabbatical, LeaveTypeHoliday, LeaveTypeLeave, LeaveTypeAnnual, LeaveTypeUnpaid, LeaveTypeMilitary, LeaveTypeMarriage, LeaveTypeCompensatory, LeaveTypeReligion, LeaveTypeEarned, LeaveTypeCompassionate:
		return true
	}
	return false
}

func (e LeaveType) String() string {
	return string(e)
}

func (e *LeaveType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeaveType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeaveType", str)
	}
	return nil
}

func (e LeaveType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
