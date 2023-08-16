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

type Contents struct {
	Count int        `json:"count"`
	Data  []*Content `json:"data,omitempty"`
}

type Countries struct {
	Data  []*Country `json:"data,omitempty"`
	Count int        `json:"count"`
}

type Currencies struct {
	Data  []*Currency `json:"data,omitempty"`
	Count int         `json:"count"`
}

type Files struct {
	Count int     `json:"count"`
	Data  []*File `json:"data,omitempty"`
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

type NewContent struct {
	Parent      *string                `json:"parent,omitempty"`
	User        *string                `json:"user,omitempty"`
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
	Categories  []string               `json:"categories,omitempty"`
}

type NewLocale struct {
	ID         *string                `json:"id,omitempty"`
	Name       string                 `json:"name"`
	Locale     string                 `json:"locale"`
	Code       string                 `json:"code"`
	Order      int                    `json:"order"`
	Rtl        bool                   `json:"rtl"`
	DateFormat string                 `json:"date_format"`
	TimeFormat string                 `json:"time_format"`
	WeekStart  int                    `json:"week_start"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
}

type NewTaxonomy struct {
	Category      string `json:"category"`
	Taxonomizable string `json:"Taxonomizable"`
	Type          string `json:"type"`
	Sticky        *bool  `json:"sticky,omitempty"`
}

type NewTimezone struct {
	ID       *string                `json:"id,omitempty"`
	Name     string                 `json:"name"`
	Offset   int                    `json:"offset"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
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

type UpdateContent struct {
	Parent      *string                `json:"parent,omitempty"`
	User        *string                `json:"user,omitempty"`
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
	Locale      string                 `json:"locale"`
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Starred     *bool                  `json:"starred,omitempty"`
	Status      *string                `json:"status,omitempty"`
	Categories  []string               `json:"categories,omitempty"`
}

type UpdateLocale struct {
	Name       *string                `json:"name,omitempty"`
	Locale     string                 `json:"locale"`
	Code       *string                `json:"code,omitempty"`
	Order      *int                   `json:"order,omitempty"`
	Rtl        *bool                  `json:"rtl,omitempty"`
	DateFormat *string                `json:"date_format,omitempty"`
	TimeFormat *string                `json:"time_format,omitempty"`
	WeekStart  *int                   `json:"week_start,omitempty"`
	Metadata   map[string]interface{} `json:"metadata,omitempty"`
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
