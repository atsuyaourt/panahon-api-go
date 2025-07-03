package util

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

type JSONFloat4 struct {
	Value float32
	Valid bool
}

func (f *JSONFloat4) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	s := string(b)
	if len(s) > 0 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	fv := float32(v)
	f.Value = fv
	f.Valid = true
	return nil
}

func (f JSONFloat4) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatFloat(float64(f.Value), 'f', -1, 32)), nil
}

func (f JSONFloat4) ToFloat4() pgtype.Float4 {
	return pgtype.Float4{Float32: f.Value, Valid: f.Valid}
}

// RandomJSONFloat generates a random float between min and max:w http.ResponseWriter, r *http.Request
func RandomJSONFloat4(min, max float32) JSONFloat4 {
	return JSONFloat4{
		Value: RandomFloat(min, max),
		Valid: true,
	}
}
