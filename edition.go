package openlib_objects

import "fmt"
import "encoding/json"
import "reflect"

type EditionSchemaJson struct {
	// Authors corresponds to the JSON schema field "authors".
	Authors []Author `json:"authors,omitempty" yaml:"authors,omitempty"`

	// ByStatement corresponds to the JSON schema field "by_statement".
	ByStatement *string `json:"by_statement,omitempty" yaml:"by_statement,omitempty"`

	// Contributions corresponds to the JSON schema field "contributions".
	Contributions StringArray `json:"contributions,omitempty" yaml:"contributions,omitempty"`

	// Covers corresponds to the JSON schema field "covers".
	Covers []float64 `json:"covers,omitempty" yaml:"covers,omitempty"`

	// Created corresponds to the JSON schema field "created".
	Created *InternalDatetime `json:"created,omitempty" yaml:"created,omitempty"`

	// Description corresponds to the JSON schema field "description".
	Description *TextBlock `json:"description,omitempty" yaml:"description,omitempty"`

	// DeweyDecimalClass corresponds to the JSON schema field "dewey_decimal_class".
	DeweyDecimalClass StringArray `json:"dewey_decimal_class,omitempty" yaml:"dewey_decimal_class,omitempty"`

	// EditionName corresponds to the JSON schema field "edition_name".
	EditionName *string `json:"edition_name,omitempty" yaml:"edition_name,omitempty"`

	// FirstSentence corresponds to the JSON schema field "first_sentence".
	FirstSentence *TextBlock `json:"first_sentence,omitempty" yaml:"first_sentence,omitempty"`

	// Genres corresponds to the JSON schema field "genres".
	Genres StringArray `json:"genres,omitempty" yaml:"genres,omitempty"`

	// Identifiers corresponds to the JSON schema field "identifiers".
	Identifiers Identifiers `json:"identifiers,omitempty" yaml:"identifiers,omitempty"`

	// Isbn10 corresponds to the JSON schema field "isbn_10".
	Isbn10 []Isbn10 `json:"isbn_10,omitempty" yaml:"isbn_10,omitempty"`

	// Isbn13 corresponds to the JSON schema field "isbn_13".
	Isbn13 []Isbn13 `json:"isbn_13,omitempty" yaml:"isbn_13,omitempty"`

	// Key corresponds to the JSON schema field "key".
	Key EditionKey `json:"key" yaml:"key"`

	// Languages corresponds to the JSON schema field "languages".
	Languages []Language `json:"languages,omitempty" yaml:"languages,omitempty"`

	// LastModified corresponds to the JSON schema field "last_modified".
	LastModified InternalDatetime `json:"last_modified" yaml:"last_modified"`

	// LatestRevision corresponds to the JSON schema field "latest_revision".
	LatestRevision *float64 `json:"latest_revision,omitempty" yaml:"latest_revision,omitempty"`

	// LcClassifications corresponds to the JSON schema field "lc_classifications".
	LcClassifications []LcClassification `json:"lc_classifications,omitempty" yaml:"lc_classifications,omitempty"`

	// Library of Congress Control Numbers, linkable via https://lccn.loc.gov/<lccn>
	Lccn []string `json:"lccn,omitempty" yaml:"lccn,omitempty"`

	// Links corresponds to the JSON schema field "links".
	Links []Link `json:"links,omitempty" yaml:"links,omitempty"`

	// Notes corresponds to the JSON schema field "notes".
	Notes *TextBlock `json:"notes,omitempty" yaml:"notes,omitempty"`

	// NumberOfPages corresponds to the JSON schema field "number_of_pages".
	NumberOfPages *float64 `json:"number_of_pages,omitempty" yaml:"number_of_pages,omitempty"`

	// Links to the Internet Archive record via https://archive.org/details/<ocaid>
	Ocaid *string `json:"ocaid,omitempty" yaml:"ocaid,omitempty"`

	// OCLC Online Computer Library Center / WorldCat id, linkable via
	// https://www.worldcat.org/oclc/<oclc_number>
	OclcNumbers []string `json:"oclc_numbers,omitempty" yaml:"oclc_numbers,omitempty"`

	// OtherTitles corresponds to the JSON schema field "other_titles".
	OtherTitles StringArray `json:"other_titles,omitempty" yaml:"other_titles,omitempty"`

	// Pagination corresponds to the JSON schema field "pagination".
	Pagination *string `json:"pagination,omitempty" yaml:"pagination,omitempty"`

	// PhysicalDimensions corresponds to the JSON schema field "physical_dimensions".
	PhysicalDimensions *string `json:"physical_dimensions,omitempty" yaml:"physical_dimensions,omitempty"`

	// PhysicalFormat corresponds to the JSON schema field "physical_format".
	PhysicalFormat *string `json:"physical_format,omitempty" yaml:"physical_format,omitempty"`

	// PublishCountry corresponds to the JSON schema field "publish_country".
	PublishCountry *PublishCountry `json:"publish_country,omitempty" yaml:"publish_country,omitempty"`

	// PublishDate corresponds to the JSON schema field "publish_date".
	PublishDate *string `json:"publish_date,omitempty" yaml:"publish_date,omitempty"`

	// PublishPlaces corresponds to the JSON schema field "publish_places".
	PublishPlaces StringArray `json:"publish_places,omitempty" yaml:"publish_places,omitempty"`

	// Publishers corresponds to the JSON schema field "publishers".
	Publishers StringArray `json:"publishers,omitempty" yaml:"publishers,omitempty"`

	// Revision corresponds to the JSON schema field "revision".
	Revision float64 `json:"revision" yaml:"revision"`

	// Series corresponds to the JSON schema field "series".
	Series StringArray `json:"series,omitempty" yaml:"series,omitempty"`

	// SourceRecords corresponds to the JSON schema field "source_records".
	SourceRecords StringArray `json:"source_records,omitempty" yaml:"source_records,omitempty"`

	// Subjects corresponds to the JSON schema field "subjects".
	Subjects StringArray `json:"subjects,omitempty" yaml:"subjects,omitempty"`

	// Subtitle corresponds to the JSON schema field "subtitle".
	Subtitle *string `json:"subtitle,omitempty" yaml:"subtitle,omitempty"`

	// TableOfContents corresponds to the JSON schema field "table_of_contents".
	TableOfContents []interface{} `json:"table_of_contents,omitempty" yaml:"table_of_contents,omitempty"`

	// Title corresponds to the JSON schema field "title".
	Title string `json:"title" yaml:"title"`

	// Type corresponds to the JSON schema field "type".
	Type EditionSchemaJsonType `json:"type" yaml:"type"`

	// Weight corresponds to the JSON schema field "weight".
	Weight *string `json:"weight,omitempty" yaml:"weight,omitempty"`

	// WorkTitles corresponds to the JSON schema field "work_titles".
	WorkTitles StringArray `json:"work_titles,omitempty" yaml:"work_titles,omitempty"`

	// Works corresponds to the JSON schema field "works".
	Works Works `json:"works" yaml:"works"`
}

type EditionSchemaJsonType struct {
	// Key corresponds to the JSON schema field "key".
	Key *EditionSchemaJsonTypeKey `json:"key,omitempty" yaml:"key,omitempty"`
}

type EditionSchemaJsonTypeKey string

const EditionSchemaJsonTypeKeyTypeEdition EditionSchemaJsonTypeKey = "/type/edition"

type Identifiers map[string]interface{}

type Isbn10 string

type Isbn13 string

// A type based on the list of MARC21 language codes. See
// https://www.loc.gov/marc/languages/
type Language struct {
	// Key corresponds to the JSON schema field "key".
	Key string `json:"key" yaml:"key"`
}

type Works []struct {
	// Key corresponds to the JSON schema field "key".
	Key WorkKey `json:"key" yaml:"key"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Language) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["key"]; !ok || v == nil {
		return fmt.Errorf("field key in Language: required")
	}
	type Plain Language
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Language(plain)
	return nil
}

var enumValues_EditionSchemaJsonTypeKey = []interface{}{
	"/type/edition",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EditionSchemaJsonTypeKey) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EditionSchemaJsonTypeKey {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EditionSchemaJsonTypeKey, v)
	}
	*j = EditionSchemaJsonTypeKey(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EditionSchemaJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["key"]; !ok || v == nil {
		return fmt.Errorf("field key in EditionSchemaJson: required")
	}
	if v, ok := raw["last_modified"]; !ok || v == nil {
		return fmt.Errorf("field last_modified in EditionSchemaJson: required")
	}
	if v, ok := raw["revision"]; !ok || v == nil {
		return fmt.Errorf("field revision in EditionSchemaJson: required")
	}
	if v, ok := raw["title"]; !ok || v == nil {
		return fmt.Errorf("field title in EditionSchemaJson: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in EditionSchemaJson: required")
	}
	if v, ok := raw["works"]; !ok || v == nil {
		return fmt.Errorf("field works in EditionSchemaJson: required")
	}
	type Plain EditionSchemaJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EditionSchemaJson(plain)
	return nil
}
