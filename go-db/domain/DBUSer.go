package domain

import (
	"database/sql/driver"
	"strings"
)

type PqStringArray []string

func (psa *PqStringArray) Scan(value interface{}) error {
	if value == nil {
		*psa = PqStringArray([]string{})
		return nil
	}
	stringVal := value.(string)
	stringVal = strings.TrimPrefix(stringVal, "{")
	stringVal = strings.TrimSuffix(stringVal, "}")
	stringArrayValue := strings.Split(stringVal, ",")
	*psa = PqStringArray(stringArrayValue)
	return nil
}

func (psa PqStringArray) Value() (driver.Value, error) {
	value := "{" + strings.Join(psa, ",") + "}"
	return value, nil
}

type DBUser struct {
	Id int `db:"id"`
	Name string `db:"id"`
	Contacts PqStringArray `db:"contacts"`
}
