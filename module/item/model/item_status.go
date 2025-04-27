package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemsStatus int

const (
	ItemsStatusDoing ItemsStatus = iota
	ItemsStatusDone
	ItemsStatusDeleted
)

var allItemStatues = [3]string{"doing", "done", "deleted"}

func (item *ItemsStatus) String() string {
	return allItemStatues[*item]
}

func parseItemStatusString(s string) (ItemsStatus, error) {
	for i := range allItemStatues {
		if allItemStatues[i] == s {
			return ItemsStatus(i), nil
		}
	}
	return ItemsStatus(0), errors.New("cannot parse")
}

// parse data show list
func (item *ItemsStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("formart failed: %s ", value))

	}

	v, err := parseItemStatusString(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("failed to scan data sql %s", value))
	}
	*item = v
	return nil
}

// parse bbody -> data (status)

func (item *ItemsStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil

}

// json -> bytes
func (item *ItemsStatus) MarshalJSON() ([]byte, error) {

	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil

}

//bytes -> json

func (item *ItemsStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	itemValue, err := parseItemStatusString(str)

	if err != nil {
		return err
	}

	*item = itemValue
	return nil
}
