package openlib_objects

import "fmt"
import "encoding/json"
import "reflect"

type AuthorSchemaJson struct {
	// AlternateNames corresponds to the JSON schema field "alternate_names".
	AlternateNames StringArray `json:"alternate_names,omitempty" yaml:"alternate_names,omitempty"`

	// Bio corresponds to the JSON schema field "bio".
	Bio *TextBlock `json:"bio,omitempty" yaml:"bio,omitempty"`

	// BirthDate corresponds to the JSON schema field "birth_date".
	BirthDate *string `json:"birth_date,omitempty" yaml:"birth_date,omitempty"`

	// Created corresponds to the JSON schema field "created".
	Created *InternalDatetime `json:"created,omitempty" yaml:"created,omitempty"`

	// Date corresponds to the JSON schema field "date".
	Date *string `json:"date,omitempty" yaml:"date,omitempty"`

	// DeathDate corresponds to the JSON schema field "death_date".
	DeathDate *string `json:"death_date,omitempty" yaml:"death_date,omitempty"`

	// EntityType corresponds to the JSON schema field "entity_type".
	EntityType *AuthorSchemaJsonEntityType `json:"entity_type,omitempty" yaml:"entity_type,omitempty"`

	// FullerName corresponds to the JSON schema field "fuller_name".
	FullerName *string `json:"fuller_name,omitempty" yaml:"fuller_name,omitempty"`

	// Key corresponds to the JSON schema field "key".
	Key AuthorKey_1 `json:"key" yaml:"key"`

	// LastModified corresponds to the JSON schema field "last_modified".
	LastModified InternalDatetime `json:"last_modified" yaml:"last_modified"`

	// LatestRevision corresponds to the JSON schema field "latest_revision".
	LatestRevision *float64 `json:"latest_revision,omitempty" yaml:"latest_revision,omitempty"`

	// Links corresponds to the JSON schema field "links".
	Links []Link `json:"links,omitempty" yaml:"links,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *string `json:"location,omitempty" yaml:"location,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name" yaml:"name"`

	// PersonalName corresponds to the JSON schema field "personal_name".
	PersonalName *string `json:"personal_name,omitempty" yaml:"personal_name,omitempty"`

	// Photos corresponds to the JSON schema field "photos".
	Photos []float64 `json:"photos,omitempty" yaml:"photos,omitempty"`

	// RemoteIds corresponds to the JSON schema field "remote_ids".
	RemoteIds *AuthorSchemaJsonRemoteIds `json:"remote_ids,omitempty" yaml:"remote_ids,omitempty"`

	// Revision corresponds to the JSON schema field "revision".
	Revision float64 `json:"revision" yaml:"revision"`

	// Title corresponds to the JSON schema field "title".
	Title *string `json:"title,omitempty" yaml:"title,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type AuthorSchemaJsonType `json:"type" yaml:"type"`

	// DEPRECATED: Not many authors have this property. Should be moved to Links
	// and/or remote_ids: wikidata.
	Wikipedia *string `json:"wikipedia,omitempty" yaml:"wikipedia,omitempty"`
}

type AuthorSchemaJsonEntityType string

const AuthorSchemaJsonEntityTypeEvent AuthorSchemaJsonEntityType = "event"
const AuthorSchemaJsonEntityTypeOrg AuthorSchemaJsonEntityType = "org"
const AuthorSchemaJsonEntityTypePerson AuthorSchemaJsonEntityType = "person"

type AuthorSchemaJsonRemoteIds struct {
	// Viaf corresponds to the JSON schema field "viaf".
	Viaf *string `json:"viaf,omitempty" yaml:"viaf,omitempty"`

	// Wikidata corresponds to the JSON schema field "wikidata".
	Wikidata *string `json:"wikidata,omitempty" yaml:"wikidata,omitempty"`
}

type AuthorSchemaJsonType struct {
	// Key corresponds to the JSON schema field "key".
	Key *AuthorSchemaJsonTypeKey `json:"key,omitempty" yaml:"key,omitempty"`
}

type AuthorSchemaJsonTypeKey string

const AuthorSchemaJsonTypeKeyTypeAuthor AuthorSchemaJsonTypeKey = "/type/author"

// UnmarshalJSON implements json.Unmarshaler.

var enumValues_AuthorSchemaJsonEntityType = []interface{}{
	"person",
	"org",
	"event",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorSchemaJsonEntityType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorSchemaJsonEntityType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorSchemaJsonEntityType, v)
	}
	*j = AuthorSchemaJsonEntityType(v)
	return nil
}

var enumValues_AuthorSchemaJsonTypeKey = []interface{}{
	"/type/author",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorSchemaJsonTypeKey) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorSchemaJsonTypeKey {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorSchemaJsonTypeKey, v)
	}
	*j = AuthorSchemaJsonTypeKey(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorSchemaJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["key"]; !ok || v == nil {
		return fmt.Errorf("field key in AuthorSchemaJson: required")
	}
	if v, ok := raw["last_modified"]; !ok || v == nil {
		return fmt.Errorf("field last_modified in AuthorSchemaJson: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in AuthorSchemaJson: required")
	}
	if v, ok := raw["revision"]; !ok || v == nil {
		return fmt.Errorf("field revision in AuthorSchemaJson: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in AuthorSchemaJson: required")
	}
	type Plain AuthorSchemaJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AuthorSchemaJson(plain)
	return nil
}
