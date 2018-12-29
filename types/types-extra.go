package typespb

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
)

// Null is this a null?
func (m *Timestamp) Null() bool {
	if m == nil {
		return true
	}
	return !m.Valid
}

// Timestamp try to return the timestamp
func (m *Timestamp) Timestamp() time.Time {
	return time.Unix(m.Unix, 0)
}

// MarshalJSON  try to marshal it in json
func (m *Timestamp) MarshalJSON() ([]byte, error) {
	if m == nil || !m.Valid {
		return []byte("null"), nil
	}

	t := time.Unix(m.GetUnix(), 0)
	return json.Marshal(t)
}

// UnmarshalJSON try to unmarshal this from a json string
func (m *Timestamp) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		m.Valid, m.Unix = false, 0
		return nil
	}

	var ts time.Time
	err := json.Unmarshal(src, &ts)
	if err != nil {
		return err
	}

	m.Valid, m.Unix = true, ts.Unix()
	return nil
}

// Scan convert the json array into string slice
func (m *Timestamp) Scan(src interface{}) error {
	var p pq.NullTime
	if err := p.Scan(src); err != nil {
		return err
	}

	m.Valid, m.Unix = p.Valid, p.Time.Unix()
	return nil
}

// Value try to get the string slice representation in database
func (m *Timestamp) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	if !m.Valid {
		return nil, nil
	}
	return time.Unix(m.Unix, 0), nil
}

func TimestampProto(ts time.Time) *Timestamp {
	return &Timestamp{
		Unix:  ts.Unix(),
		Valid: true,
	}
}

// MarshalJSON  try to marshal it in json
func (m *JSONArray) MarshalJSON() ([]byte, error) {
	if m == nil || m.Data == nil {
		return []byte("null"), nil
	}

	return json.Marshal(m.Data)
}

// UnmarshalJSON try to unmarshal this from a json string
func (m *JSONArray) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		m.Data = nil
		return nil
	}

	return json.Unmarshal(src, &m.Data)
}

// Scan convert the json array into string slice
func (m *JSONArray) Scan(src interface{}) error {
	var b []byte
	switch t := src.(type) {
	case []byte:
		b = t
	case string:
		b = []byte(t)
	case nil:
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}

	return m.UnmarshalJSON(b)
}

// Value try to get the string slice representation in database
func (m *JSONArray) Value() (driver.Value, error) {
	if m == nil || m.Data == nil {
		return nil, nil
	}

	return m.MarshalJSON()
}

// Null return true if the field is null
func (m *JSONArray) Null() bool {
	return m == nil || m.Data == nil
}

// MarshalJSON  try to marshal it in json
func (m *JSONMap) MarshalJSON() ([]byte, error) {
	if m == nil || m.Data == nil {
		return []byte("null"), nil
	}

	return json.Marshal(m.Data)
}

// UnmarshalJSON try to unmarshal this from a json string
func (m *JSONMap) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		m.Data = nil
		return nil
	}

	m.Data = make(map[string]string)
	return json.Unmarshal(src, &m.Data)
}

// Scan convert the json array into string slice
func (m *JSONMap) Scan(src interface{}) error {
	var b []byte
	switch t := src.(type) {
	case []byte:
		b = t
	case string:
		b = []byte(t)
	case nil:
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}

	return m.UnmarshalJSON(b)
}

// Value try to get the string slice representation in database
func (m *JSONMap) Value() (driver.Value, error) {
	if m == nil || m.Data == nil {
		return nil, nil
	}

	return m.MarshalJSON()
}

// Null return true if the field is null
func (m *JSONMap) Null() bool {
	return m == nil || m.Data == nil
}

// MarshalJSON  try to marshal it in json
func (m *NullString) MarshalJSON() ([]byte, error) {
	if m == nil || !m.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(m.String_)
}

// UnmarshalJSON try to unmarshal this from a json string
func (m *NullString) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		m.Valid = false
		return nil
	}

	m.Valid = true
	return json.Unmarshal(src, &m.String_)
}

// Scan convert the json array into string slice
func (m *NullString) Scan(src interface{}) error {
	var b sql.NullString
	if err := b.Scan(src); err != nil {
		return err
	}
	m.Valid, m.String_ = b.Valid, b.String
	return nil
}

// Value try to get the string slice representation in database
func (m *NullString) Value() (driver.Value, error) {
	if !m.Valid {
		return nil, nil
	}
	return m.String_, nil
}

// Null return true if the field is null
func (m *NullString) Null() bool {
	return m == nil || !m.Valid
}

// MarshalJSON  try to marshal it in json
func (m *NullInt64) MarshalJSON() ([]byte, error) {
	if m == nil || !m.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(m.Int64)
}

// UnmarshalJSON try to unmarshal this from a json string
func (m *NullInt64) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		m.Valid = false
		return nil
	}

	m.Valid = true
	return json.Unmarshal(src, &m.Int64)
}

// Scan convert the json array into string slice
func (m *NullInt64) Scan(src interface{}) error {
	var b sql.NullInt64
	if err := b.Scan(src); err != nil {
		return err
	}
	m.Valid, m.Int64 = b.Valid, b.Int64
	return nil
}

// Value try to get the string slice representation in database
func (m *NullInt64) Value() (driver.Value, error) {
	if !m.Valid {
		return nil, nil
	}
	return m.Int64, nil
}

// Null return true if the field is null
func (m *NullInt64) Null() bool {
	return m == nil || !m.Valid
}
