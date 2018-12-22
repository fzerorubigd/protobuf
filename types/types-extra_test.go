package typespb

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	F *Timestamp
}

func TestTimestamp(t *testing.T) {
	s := testStruct{
		F: &Timestamp{},
	}

	b, err := json.Marshal(s)
	require.NoError(t, err)
	require.Equal(t, `{"F":null}`, string(b))

	s2 := testStruct{}
	err = json.Unmarshal(b, &s2)
	require.NoError(t, err)
	require.Equal(t, true, s2.F.Null())

	ts := time.Now()
	s3 := testStruct{
		F: &Timestamp{
			Unix:  ts.Unix(),
			Valid: true,
		},
	}
	b, err = json.Marshal(s3)
	require.NoError(t, err)
	var s4 testStruct
	err = json.Unmarshal(b, &s4)
	require.NoError(t, err)
	require.Equal(t, s4.F.Unix, ts.Unix())
}
