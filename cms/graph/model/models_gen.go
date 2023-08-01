// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Categories struct {
	Count int         `json:"count"`
	Data  []*Category `json:"data,omitempty"`
}

type Category struct {
	ID          string                 `json:"id"`
	Parent      *Category              `json:"parent,omitempty"`
	Children    []*Category            `json:"children,omitempty"`
	Locale      string                 `json:"locale"`
	Taxonomy    string                 `json:"taxonomy"`
	Slug        *string                `json:"slug,omitempty"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Order       int                    `json:"order"`
	Count       int                    `json:"count"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	CreatedBy   *User                  `json:"created_by,omitempty"`
	UpdatedBy   *User                  `json:"updated_by,omitempty"`
}

func (Category) IsEntity() {}

type Comment struct {
	ID          string                 `json:"id"`
	Locale      string                 `json:"locale"`
	Content     string                 `json:"content"`
	Rating      *int                   `json:"rating,omitempty"`
	Status      string                 `json:"status"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Parent      *Comment               `json:"parent,omitempty"`
	Children    []*Comment             `json:"children,omitempty"`
	Object      map[string]interface{} `json:"object"`
	Owner       *User                  `json:"owner"`
	Reaction    []*Reaction            `json:"reaction,omitempty"`
	CreatedBy   *User                  `json:"created_by,omitempty"`
	UpdatedBy   *User                  `json:"updated_by,omitempty"`
	Attachments []*File                `json:"attachments,omitempty"`
	Reactions   []*Reaction            `json:"reactions,omitempty"`
}

func (Comment) IsEntity() {}

type Comments struct {
	Count int        `json:"count"`
	Data  []*Comment `json:"data,omitempty"`
}

type Content struct {
	ID          string                 `json:"id"`
	Locale      string                 `json:"locale"`
	Type        string                 `json:"type"`
	Title       string                 `json:"title"`
	Summary     string                 `json:"summary"`
	Body        string                 `json:"body"`
	Slug        string                 `json:"slug"`
	Status      string                 `json:"status"`
	Reviewable  *bool                  `json:"reviewable,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	PublishedAt int                    `json:"published_at"`
	CreatedBy   *User                  `json:"created_by"`
	UpdatedBy   *User                  `json:"updated_by"`
	Owner       *User                  `json:"owner"`
	Parent      *Content               `json:"parent,omitempty"`
	Followers   []*Follow              `json:"followers,omitempty"`
	Comments    []*Comment             `json:"comments,omitempty"`
	Attachments []*File                `json:"attachments,omitempty"`
}

func (Content) IsEntity() {}

type Contents struct {
	Count int        `json:"count"`
	Data  []*Content `json:"data,omitempty"`
}

type Countries struct {
	Data  []*Country `json:"data,omitempty"`
	Count int        `json:"count"`
}

type Country struct {
	ID        string                 `json:"id"`
	Code      string                 `json:"code"`
	Locale    string                 `json:"locale"`
	Name      string                 `json:"name"`
	Continent string                 `json:"continent"`
	Currency  *string                `json:"currency,omitempty"`
	Languages []*string              `json:"languages,omitempty"`
	Capital   *string                `json:"capital,omitempty"`
	Flag      *string                `json:"flag,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

type Currencies struct {
	Data  []*Currency `json:"data,omitempty"`
	Count int         `json:"count"`
}

type Currency struct {
	ID        string                 `json:"id"`
	Locale    string                 `json:"locale"`
	Code      string                 `json:"code"`
	Name      string                 `json:"name"`
	Rate      float64                `json:"rate"`
	Symbol    string                 `json:"symbol"`
	Precision int                    `json:"precision"`
	Decimal   string                 `json:"decimal"`
	Thousand  string                 `json:"thousand"`
	Order     int                    `json:"order"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

type File struct {
	ID          string                 `json:"id"`
	Locale      string                 `json:"locale"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Size        int                    `json:"size"`
	Provider    string                 `json:"provider"`
	URL         string                 `json:"url"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Starred     bool                   `json:"starred"`
	Status      string                 `json:"status"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	CreatedBy   *User                  `json:"created_by,omitempty"`
	UpdatedBy   *User                  `json:"updated_by,omitempty"`
	Owner       *User                  `json:"owner,omitempty"`
	Followers   []*Follow              `json:"followers,omitempty"`
}

func (File) IsEntity() {}

type Files struct {
	Count int     `json:"count"`
	Data  []*File `json:"data,omitempty"`
}

type Follow struct {
	ID        string                 `json:"id"`
	User      *User                  `json:"user"`
	Object    map[string]interface{} `json:"object"`
	Role      *string                `json:"role,omitempty"`
	Status    string                 `json:"status"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
	CreatedBy *User                  `json:"created_by"`
	UpdatedBy *User                  `json:"updated_by"`
}

func (Follow) IsEntity() {}

type Follows struct {
	Count int       `json:"count"`
	Data  []*Follow `json:"data,omitempty"`
}

type Locale struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Locale       string                 `json:"locale"`
	Code         string                 `json:"code"`
	Order        int                    `json:"order"`
	Rtl          bool                   `json:"rtl"`
	DateFormat   string                 `json:"date_format"`
	StringFormat string                 `json:"String_format"`
	WeekStart    int                    `json:"week_start"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
}

type Locales struct {
	Count int       `json:"count"`
	Data  []*Locale `json:"data,omitempty"`
}

type NewCategory struct {
	Locale      string                 `json:"locale"`
	Name        string                 `json:"name"`
	Slug        *string                `json:"slug,omitempty"`
	Description *string                `json:"description,omitempty"`
	Parent      *string                `json:"parent,omitempty"`
	Taxonomy    string                 `json:"taxonomy"`
	Order       *int                   `json:"order,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewComment struct {
	Parent      *string                `json:"parent,omitempty"`
	Object      string                 `json:"object"`
	Type        string                 `json:"type"`
	Locale      string                 `json:"locale"`
	Content     *string                `json:"content,omitempty"`
	Rating      *int                   `json:"rating,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Attachments []*string              `json:"attachments,omitempty"`
}

type NewContent struct {
	Parent      *string                `json:"parent,omitempty"`
	Owner       *string                `json:"owner,omitempty"`
	Locale      string                 `json:"locale"`
	Type        string                 `json:"type"`
	Title       *string                `json:"title,omitempty"`
	Summary     *string                `json:"summary,omitempty"`
	Body        *string                `json:"body,omitempty"`
	Slug        *string                `json:"slug,omitempty"`
	Categories  []string               `json:"categories,omitempty"`
	Images      []*string              `json:"images,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Commentable *bool                  `json:"commentable,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type NewCountry struct {
	Code      string                 `json:"code"`
	Locale    string                 `json:"locale"`
	Name      string                 `json:"name"`
	Native    string                 `json:"native"`
	Continent string                 `json:"continent"`
	Currency  *string                `json:"currency,omitempty"`
	Languages []*string              `json:"languages,omitempty"`
	Capital   *string                `json:"capital,omitempty"`
	Flag      *string                `json:"flag,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type NewCurrency struct {
	ID        *string                `json:"id,omitempty"`
	Locale    string                 `json:"locale"`
	Name      string                 `json:"name"`
	Code      string                 `json:"code"`
	Rate      float64                `json:"rate"`
	Symbol    string                 `json:"symbol"`
	Precision int                    `json:"precision"`
	Decimal   string                 `json:"decimal"`
	Thousand  string                 `json:"thousand"`
	Order     int                    `json:"order"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type NewFile struct {
	Locale      string                 `json:"locale"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Type        string                 `json:"type"`
	Size        int                    `json:"size"`
	Provider    string                 `json:"provider"`
	URL         string                 `json:"url"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Starred     *bool                  `json:"starred,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Owner       *string                `json:"owner,omitempty"`
	Categories  []string               `json:"categories,omitempty"`
}

type NewFollow struct {
	User   string `json:"user"`
	Object string `json:"object"`
	Type   string `json:"type"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

type NewLocale struct {
	ID           *string                `json:"id,omitempty"`
	Name         string                 `json:"name"`
	Locale       string                 `json:"locale"`
	Code         string                 `json:"code"`
	Order        int                    `json:"order"`
	Rtl          bool                   `json:"rtl"`
	DateFormat   string                 `json:"date_format"`
	StringFormat string                 `json:"String_format"`
	WeekStart    int                    `json:"week_start"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type NewReaction struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Action string `json:"action"`
}

type NewTaxonomy struct {
	Category string `json:"category"`
	Object   string `json:"object"`
	Type     string `json:"type"`
	Sticky   *bool  `json:"sticky,omitempty"`
}

type NewTimezone struct {
	ID       *string                `json:"id,omitempty"`
	Name     string                 `json:"name"`
	Offset   int                    `json:"offset"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type Object struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Reaction struct {
	ID        string                 `json:"id"`
	User      *User                  `json:"user"`
	Object    map[string]interface{} `json:"object"`
	Action    string                 `json:"action"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

func (Reaction) IsEntity() {}

type Reactions struct {
	Data  []*Reaction `json:"data,omitempty"`
	Count int         `json:"count"`
}

type Taxonomy struct {
	ID        string    `json:"id"`
	Category  *Category `json:"category"`
	Object    *Object   `json:"object"`
	Sticky    bool      `json:"sticky"`
	CreatedAt string    `json:"created_at"`
}

type Timezone struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Offset      int                    `json:"offset"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
}

type Timezones struct {
	Count int         `json:"count"`
	Data  []*Timezone `json:"data,omitempty"`
}

type UpdateCategory struct {
	ID          *string                `json:"id,omitempty"`
	Locale      string                 `json:"locale"`
	Name        *string                `json:"name,omitempty"`
	Slug        *string                `json:"slug,omitempty"`
	Description *string                `json:"description,omitempty"`
	Parent      *string                `json:"parent,omitempty"`
	Taxonomy    *string                `json:"taxonomy,omitempty"`
	Order       *int                   `json:"order,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateComment struct {
	Parent      *string                `json:"parent,omitempty"`
	Locale      string                 `json:"locale"`
	Content     *string                `json:"content,omitempty"`
	Rating      *int                   `json:"rating,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Attachments []*string              `json:"attachments,omitempty"`
}

type UpdateContent struct {
	Parent      *string                `json:"parent,omitempty"`
	Owner       *string                `json:"owner,omitempty"`
	Locale      string                 `json:"locale"`
	Type        *string                `json:"type,omitempty"`
	Title       *string                `json:"title,omitempty"`
	Summary     *string                `json:"summary,omitempty"`
	Body        *string                `json:"body,omitempty"`
	Slug        *string                `json:"slug,omitempty"`
	Categories  []string               `json:"categories,omitempty"`
	Images      []*string              `json:"images,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Commentable *bool                  `json:"commentable,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateCountry struct {
	Locale    *string                `json:"locale,omitempty"`
	Code      *string                `json:"code,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Native    *string                `json:"native,omitempty"`
	Continent *string                `json:"continent,omitempty"`
	Currency  *string                `json:"currency,omitempty"`
	Languages []*string              `json:"languages,omitempty"`
	Capital   *string                `json:"capital,omitempty"`
	Flag      *string                `json:"flag,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateCurrency struct {
	Locale    *string                `json:"locale,omitempty"`
	Name      *string                `json:"name,omitempty"`
	Code      *string                `json:"code,omitempty"`
	Rate      *float64               `json:"rate,omitempty"`
	Symbol    *string                `json:"symbol,omitempty"`
	Precision *int                   `json:"precision,omitempty"`
	Decimal   *string                `json:"decimal,omitempty"`
	Thousand  *string                `json:"thousand,omitempty"`
	Order     *int                   `json:"order,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Status    *string                `json:"status,omitempty"`
}

type UpdateFile struct {
	Locale      *string                `json:"locale,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Size        *int                   `json:"size,omitempty"`
	Provider    *string                `json:"provider,omitempty"`
	URL         *string                `json:"url,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Starred     *bool                  `json:"starred,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Owner       *string                `json:"owner,omitempty"`
	Categories  []string               `json:"categories,omitempty"`
}

type UpdateFollow struct {
	Role   *string `json:"role,omitempty"`
	Status *string `json:"status,omitempty"`
}

type UpdateLocale struct {
	Name         *string                `json:"name,omitempty"`
	Locale       string                 `json:"locale"`
	Code         *string                `json:"code,omitempty"`
	Order        *int                   `json:"order,omitempty"`
	Rtl          *bool                  `json:"rtl,omitempty"`
	DateFormat   *string                `json:"date_format,omitempty"`
	StringFormat *string                `json:"String_format,omitempty"`
	WeekStart    *int                   `json:"week_start,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type UpdateReaction struct {
	Type string `json:"type"`
}

type UpdateTaxonomy struct {
	Sticky *bool `json:"sticky,omitempty"`
}

type UpdateTimezone struct {
	Name        *string                `json:"name,omitempty"`
	Offset      *int                   `json:"offset,omitempty"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type User struct {
	ID   string  `json:"id"`
	Name *string `json:"name,omitempty"`
}

func (User) IsEntity() {}

type ContentStatus string

const (
	ContentStatusDraft     ContentStatus = "DRAFT"
	ContentStatusPublished ContentStatus = "PUBLISHED"
	ContentStatusArchived  ContentStatus = "ARCHIVED"
)

var AllContentStatus = []ContentStatus{
	ContentStatusDraft,
	ContentStatusPublished,
	ContentStatusArchived,
}

func (e ContentStatus) IsValid() bool {
	switch e {
	case ContentStatusDraft, ContentStatusPublished, ContentStatusArchived:
		return true
	}
	return false
}

func (e ContentStatus) String() string {
	return string(e)
}

func (e *ContentStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContentStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContentStatus", str)
	}
	return nil
}

func (e ContentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ContentType string

const (
	ContentTypePost       ContentType = "POST"
	ContentTypePage       ContentType = "PAGE"
	ContentTypeProduct    ContentType = "PRODUCT"
	ContentTypeHotel      ContentType = "HOTEL"
	ContentTypeRestaurant ContentType = "RESTAURANT"
	ContentTypePackage    ContentType = "PACKAGE"
)

var AllContentType = []ContentType{
	ContentTypePost,
	ContentTypePage,
	ContentTypeProduct,
	ContentTypeHotel,
	ContentTypeRestaurant,
	ContentTypePackage,
}

func (e ContentType) IsValid() bool {
	switch e {
	case ContentTypePost, ContentTypePage, ContentTypeProduct, ContentTypeHotel, ContentTypeRestaurant, ContentTypePackage:
		return true
	}
	return false
}

func (e ContentType) String() string {
	return string(e)
}

func (e *ContentType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContentType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContentType", str)
	}
	return nil
}

func (e ContentType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CurrencyStatus string

const (
	CurrencyStatusActive   CurrencyStatus = "ACTIVE"
	CurrencyStatusInactive CurrencyStatus = "INACTIVE"
)

var AllCurrencyStatus = []CurrencyStatus{
	CurrencyStatusActive,
	CurrencyStatusInactive,
}

func (e CurrencyStatus) IsValid() bool {
	switch e {
	case CurrencyStatusActive, CurrencyStatusInactive:
		return true
	}
	return false
}

func (e CurrencyStatus) String() string {
	return string(e)
}

func (e *CurrencyStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CurrencyStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CurrencyStatus", str)
	}
	return nil
}

func (e CurrencyStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FollowRole string

const (
	FollowRoleViewer    FollowRole = "VIEWER"
	FollowRoleCommenter FollowRole = "COMMENTER"
	FollowRoleEditor    FollowRole = "EDITOR"
)

var AllFollowRole = []FollowRole{
	FollowRoleViewer,
	FollowRoleCommenter,
	FollowRoleEditor,
}

func (e FollowRole) IsValid() bool {
	switch e {
	case FollowRoleViewer, FollowRoleCommenter, FollowRoleEditor:
		return true
	}
	return false
}

func (e FollowRole) String() string {
	return string(e)
}

func (e *FollowRole) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FollowRole(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FollowRole", str)
	}
	return nil
}

func (e FollowRole) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FollowStatus string

const (
	FollowStatusPending  FollowStatus = "PENDING"
	FollowStatusAccepted FollowStatus = "ACCEPTED"
	FollowStatusRejected FollowStatus = "REJECTED"
)

var AllFollowStatus = []FollowStatus{
	FollowStatusPending,
	FollowStatusAccepted,
	FollowStatusRejected,
}

func (e FollowStatus) IsValid() bool {
	switch e {
	case FollowStatusPending, FollowStatusAccepted, FollowStatusRejected:
		return true
	}
	return false
}

func (e FollowStatus) String() string {
	return string(e)
}

func (e *FollowStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FollowStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FollowStatus", str)
	}
	return nil
}

func (e FollowStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ReactionType string

const (
	ReactionTypeNone  ReactionType = "NONE"
	ReactionTypeLike  ReactionType = "LIKE"
	ReactionTypeLove  ReactionType = "LOVE"
	ReactionTypeWow   ReactionType = "WOW"
	ReactionTypeHaha  ReactionType = "HAHA"
	ReactionTypeSorry ReactionType = "SORRY"
	ReactionTypeAngry ReactionType = "ANGRY"
)

var AllReactionType = []ReactionType{
	ReactionTypeNone,
	ReactionTypeLike,
	ReactionTypeLove,
	ReactionTypeWow,
	ReactionTypeHaha,
	ReactionTypeSorry,
	ReactionTypeAngry,
}

func (e ReactionType) IsValid() bool {
	switch e {
	case ReactionTypeNone, ReactionTypeLike, ReactionTypeLove, ReactionTypeWow, ReactionTypeHaha, ReactionTypeSorry, ReactionTypeAngry:
		return true
	}
	return false
}

func (e ReactionType) String() string {
	return string(e)
}

func (e *ReactionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReactionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReactionType", str)
	}
	return nil
}

func (e ReactionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
