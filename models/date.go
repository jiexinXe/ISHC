package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

const (
	CtLayoutDate     = "2006-01-02"
	CtLayoutDateTime = "2006-01-02 15:04:05"
)

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := string(b)
	if str == `""` {
		ct.Time = time.Time{}
		return nil
	}
	t, err := time.Parse(`"`+CtLayoutDateTime+`"`, str)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + ct.Time.Format(CtLayoutDateTime) + `"`), nil
}

func (ct CustomTime) Value() (driver.Value, error) {
	if ct.Time.IsZero() {
		return nil, nil
	}
	return ct.Time.Format(CtLayoutDateTime), nil
}

func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*ct = CustomTime{Time: time.Time{}}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*ct = CustomTime{Time: v}
	case []byte:
		t, err := time.Parse(CtLayoutDateTime, string(v))
		if err != nil {
			t, err = time.Parse(CtLayoutDate, string(v))
			if err != nil {
				return err
			}
		}
		*ct = CustomTime{Time: t}
	case string:
		t, err := time.Parse(CtLayoutDateTime, v)
		if err != nil {
			t, err = time.Parse(CtLayoutDate, v)
			if err != nil {
				return err
			}
		}
		*ct = CustomTime{Time: t}
	default:
		return fmt.Errorf("cannot convert %v to CustomTime", value)
	}
	return nil
}
