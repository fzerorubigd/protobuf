package typespb

import (
	"database/sql/driver"
	"encoding/json"
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
