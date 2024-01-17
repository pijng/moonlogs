package entities

import (
	"errors"
	"strconv"
	"strings"
)

type Tag struct {
	ID   int    `json:"id" sql:"id"`
	Name string `json:"name" sql:"name"`
}

type Tags []int

func (t *Tags) Scan(value interface{}) error {
	if value == nil {
		*t = []int{}
		return nil
	}

	switch v := value.(type) {
	case []byte:
		if len(v) == 0 {
			*t = []int{}
			return nil
		}

		strValues := strings.Split(string(v), ",")
		intValues := make([]int, len(strValues))
		for i, strValue := range strValues {
			intValues[i] = 0
			if strValue != "" {
				var err error

				intValues[i], err = strconv.Atoi(strValue)
				if err != nil {
					return err
				}
			}
		}
		*t = intValues

		return nil
	case string:
		if len(v) == 0 {
			*t = []int{}
			return nil
		}

		strValues := strings.Split(v, ",")
		intValues := make([]int, len(strValues))
		for i, strValue := range strValues {
			intValues[i] = 0

			if strValue != "" {
				var err error

				intValues[i], err = strconv.Atoi(strValue)
				if err != nil {
					return err
				}
			}
		}
		*t = intValues

		return nil
	default:
		return errors.New("unsupported type for Tags")
	}
}
