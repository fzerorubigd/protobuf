package typespb

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/lib/pq"
	"google.golang.org/protobuf/types/known/structpb"
)

// Null is this a null?
func (x *Timestamp) Null() bool {
	if x == nil {
		return true
	}
	return !x.Valid
}

// Timestamp try to return the timestamp
func (x *Timestamp) Timestamp() time.Time {
	return time.Unix(x.Unix, 0)
}

// MarshalJSON  try to marshal it in json
func (x *Timestamp) MarshalJSON() ([]byte, error) {
	if x == nil || !x.Valid {
		return []byte("null"), nil
	}

	t := time.Unix(x.GetUnix(), 0)
	return json.Marshal(t)
}

// UnmarshalJSON try to unmarshal this from a json string
func (x *Timestamp) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		x.Valid, x.Unix = false, 0
		return nil
	}

	var ts time.Time
	err := json.Unmarshal(src, &ts)
	if err != nil {
		return err
	}

	x.Valid, x.Unix = true, ts.Unix()
	return nil
}

// Scan convert the json array into string slice
func (x *Timestamp) Scan(src interface{}) error {
	var p pq.NullTime
	if err := p.Scan(src); err != nil {
		return err
	}

	x.Valid, x.Unix = p.Valid, p.Time.Unix()
	return nil
}

// Value try to get the string slice representation in database
func (x *Timestamp) Value() (driver.Value, error) {
	if x == nil {
		return nil, nil
	}
	if !x.Valid {
		return nil, nil
	}
	return time.Unix(x.Unix, 0), nil
}

func TimestampProto(ts time.Time) *Timestamp {
	return &Timestamp{
		Unix:  ts.Unix(),
		Valid: true,
	}
}

// MarshalJSON  try to marshal it in json
func (x *JSONArray) MarshalJSON() ([]byte, error) {
	if x == nil || x.Data == nil {
		return []byte("null"), nil
	}

	return json.Marshal(x.Data)
}

// UnmarshalJSON try to unmarshal this from a json string
func (x *JSONArray) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		x.Data = nil
		return nil
	}

	return json.Unmarshal(src, &x.Data)
}

// Scan convert the json array into string slice
func (x *JSONArray) Scan(src interface{}) error {
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
	if b == nil || string(b) == "null" {
		x.Data = nil
		return nil
	}

	return (*pq.StringArray)(&x.Data).Scan(src)
}

// Value try to get the string slice representation in database
func (x *JSONArray) Value() (driver.Value, error) {
	if x == nil || x.Data == nil {
		return nil, nil
	}

	return (*pq.StringArray)(&x.Data).Value()
}

// Null return true if the field is null
func (x *JSONArray) Null() bool {
	return x == nil || x.Data == nil
}

// MarshalJSON  try to marshal it in json
func (x *JSONMap) MarshalJSON() ([]byte, error) {
	if x == nil || x.Data == nil {
		return []byte("null"), nil
	}

	return json.Marshal(x.Data)
}

// UnmarshalJSON try to unmarshal this from a json string
func (x *JSONMap) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		x.Data = nil
		return nil
	}

	x.Data = make(map[string]string)
	return json.Unmarshal(src, &x.Data)
}

// Scan convert the json array into string slice
func (x *JSONMap) Scan(src interface{}) error {
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

	return x.UnmarshalJSON(b)
}

// Value try to get the string slice representation in database
func (x *JSONMap) Value() (driver.Value, error) {
	if x == nil || x.Data == nil {
		return nil, nil
	}

	return x.MarshalJSON()
}

// Null return true if the field is null
func (x *JSONMap) Null() bool {
	return x == nil || x.Data == nil
}

// MarshalJSON  try to marshal it in json
func (x *NullString) MarshalJSON() ([]byte, error) {
	if x == nil || !x.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(x.String_)
}

// UnmarshalJSON try to unmarshal this from a json string
func (x *NullString) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		x.Valid = false
		return nil
	}

	x.Valid = true
	return json.Unmarshal(src, &x.String_)
}

// Scan convert the json array into string slice
func (x *NullString) Scan(src interface{}) error {
	var b sql.NullString
	if err := b.Scan(src); err != nil {
		return err
	}
	x.Valid, x.String_ = b.Valid, b.String
	return nil
}

// Value try to get the string slice representation in database
func (x *NullString) Value() (driver.Value, error) {
	if x == nil || !x.Valid {
		return nil, nil
	}
	return x.String_, nil
}

// Null return true if the field is null
func (x *NullString) Null() bool {
	return x == nil || !x.Valid
}

// MarshalJSON  try to marshal it in json
func (x *NullInt64) MarshalJSON() ([]byte, error) {
	if x == nil || !x.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(x.Int64)
}

// UnmarshalJSON try to unmarshal this from a json string
func (x *NullInt64) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		x.Valid = false
		return nil
	}

	x.Valid = true
	return json.Unmarshal(src, &x.Int64)
}

// Scan convert the json array into string slice
func (x *NullInt64) Scan(src interface{}) error {
	var b sql.NullInt64
	if err := b.Scan(src); err != nil {
		return err
	}
	x.Valid, x.Int64 = b.Valid, b.Int64
	return nil
}

// Value try to get the string slice representation in database
func (x *NullInt64) Value() (driver.Value, error) {
	if x == nil || !x.Valid {
		return nil, nil
	}
	return x.Int64, nil
}

// Null return true if the field is null
func (x *NullInt64) Null() bool {
	return x == nil || !x.Valid
}

// MarshalJSON  try to marshal it in json
func (x *JSONObject) MarshalJSON() ([]byte, error) {
	if x == nil || x.Data == nil {
		return []byte("null"), nil
	}

	return json.Marshal(x.Data.AsMap())
}

// UnmarshalJSON try to unmarshal this from a json string
func (x *JSONObject) UnmarshalJSON(src []byte) error {
	if string(src) == "null" {
		x.Data = nil
		return nil
	}

	var iface map[string]interface{}
	if err := json.Unmarshal(src, &iface); err != nil {
		return err
	}

	var err error
	x.Data, err = structpb.NewStruct(iface)
	if err != nil {
		return err
	}

	return nil
}

// Scan convert the json array into string slice
func (x *JSONObject) Scan(src interface{}) error {
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

	return x.UnmarshalJSON(b)
}

// Value try to get the string slice representation in database
func (x *JSONObject) Value() (driver.Value, error) {
	if x == nil || x.Data == nil {
		return nil, nil
	}

	return x.MarshalJSON()
}

// Null return true if the field is null
func (x *JSONObject) Null() bool {
	return x == nil || x.Data == nil
}

// NewJSONObject create a new JSONObject from a map
func NewJSONObject(data map[string]interface{}) (*JSONObject, error) {
	d, err := structpb.NewStruct(data)
	if err != nil {
		return nil, err
	}
	return &JSONObject{Data: d}, nil
}
