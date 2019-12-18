package quadracompta

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/ianlopshire/go-fixedwidth"
)

type UpdateExistingAccountFlag int

// C/F/G
type AccountType string

type Amount float64

func (a Amount) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	if a == 0.0 {
		return []byte{}, nil
	}

	sign := "+"
	if a < 0.0 {
		sign = "-"
	}

	length := spec.EndPos + 1 - spec.StartPos
	if length < 2 {
		length = 13
	}

	i := math.RoundToEven(float64(a) * 100)

	f := fmt.Sprintf("%%0+%dv", length-1)
	s := sign + fmt.Sprintf(f, i)
	return []byte(s), nil
}

type Flag bool

func (f *Flag) UnmarshalFixedWidth(data []byte) error {
	// log.Println(string(data))
	return nil
}

func (f Flag) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	if f == true {
		return []byte("1"), nil
	} else {
		return []byte("0"), nil
	}
}

type Date struct {
	time.Time
}

func (d *Date) UnmarshalFixedWidth(data []byte) error {
	s := string(data)
	if s == "0" || s == "" || s == "000000" {
		d.Time = time.Time{}
		return nil
	}

	layout := "020106"
	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	d.Time = t
	return nil
}

func (d Date) MarshalFixedWidth(spec fixedwidth.FieldSpec) ([]byte, error) {
	if d.Time.IsZero() {
		return []byte{}, nil
	}

	layout := "020106"
	s := d.Time.Format(layout)
	return []byte(s), nil
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try iso8601 date format
	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	// try quadracompta date format
	d.Time, err = time.Parse("020106", value)
	return err
}
