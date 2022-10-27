package openlib_objects

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Author struct {
	// Key corresponds to the JSON schema field "key".
	Key AuthorKey_1 `json:"key" yaml:"key"`
}

type AuthorKey_1 string

type InternalDatetimeType string

const InternalDatetimeTypeTypeDatetime InternalDatetimeType = "/type/datetime"

// The MARC21 language code. See
// https://www.loc.gov/marc/languages/language_code.html
type LanguageCode string

// The Library of Congress Classification number. See
// https://www.loc.gov/catdir/cpso/lcc.html We include the imprint date as the last
// four digits.
type LcClassification string

type Link struct {
	// Title corresponds to the JSON schema field "title".
	Title string `json:"title" yaml:"title"`

	// Type corresponds to the JSON schema field "type".
	Type *LinkType `json:"type,omitempty" yaml:"type,omitempty"`

	// Url corresponds to the JSON schema field "url".
	Url string `json:"url" yaml:"url"`
}

type LinkType struct {
	// Key corresponds to the JSON schema field "key".
	Key LinkTypeKey `json:"key" yaml:"key"`
}

type LinkTypeKey string

const LinkTypeKeyTypeLink LinkTypeKey = "/type/link"

// The MARC21 country code. See https://www.loc.gov/marc/countries/cou_home.html
type PublishCountry string

var enumValues_TextBlockType = []interface{}{
	"/type/text",
}

const TextBlockTypeTypeText TextBlockType = "/type/text"

type TextBlock struct {
	// Type corresponds to the JSON schema field "type".
	Type TextBlockType `json:"type" yaml:"type"`

	// Value corresponds to the JSON schema field "value".
	Value string `json:"value" yaml:"value"`
}

type StringArray []string

type WorkKey string

type TextBlockType string

var enumValues_LinkTypeKey = []interface{}{
	"/type/link",
}
var enumValues_InternalDatetimeType = []interface{}{
	"/type/datetime",
}

type EditionKey string

type InternalDatetime struct {
	// Type corresponds to the JSON schema field "type".
	Type InternalDatetimeType `json:"type" yaml:"type"`

	// Value corresponds to the JSON schema field "value".
	Value string `json:"value" yaml:"value"`
}

func (j *Link) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["title"]; !ok || v == nil {
		return fmt.Errorf("field title in Link: required")
	}
	if v, ok := raw["url"]; !ok || v == nil {
		return fmt.Errorf("field url in Link: required")
	}
	type Plain Link
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Link(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TextBlockType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TextBlockType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TextBlockType, v)
	}
	*j = TextBlockType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TextBlock) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in TextBlock: required")
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value in TextBlock: required")
	}
	type Plain TextBlock
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TextBlock(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LinkType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["key"]; !ok || v == nil {
		return fmt.Errorf("field key in LinkType: required")
	}
	type Plain LinkType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = LinkType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LinkTypeKey) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LinkTypeKey {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LinkTypeKey, v)
	}
	*j = LinkTypeKey(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *InternalDatetime) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in InternalDatetime: required")
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value in InternalDatetime: required")
	}
	type Plain InternalDatetime
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = InternalDatetime(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *InternalDatetimeType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_InternalDatetimeType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_InternalDatetimeType, v)
	}
	*j = InternalDatetimeType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Author) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["key"]; !ok || v == nil {
		return fmt.Errorf("field key in Author: required")
	}
	type Plain Author
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Author(plain)
	return nil
}
