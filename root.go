package main

import (
	"fmt"
	"time"
)

type root struct {
	time time.Time
}

type Location struct {
	Name     string
	TimeZone string `sentinel:"time_zone"`
}

// framework.Root impl.
func (m *root) Configure(raw map[string]interface{}) error {
	if _, ok := raw["timestamp"]; !ok {
		raw["timestamp"] = time.Now().Unix()
	}

	v := raw["timestamp"]
	timestamp, ok := v.(int64)
	if !ok {
		return fmt.Errorf("invalid timestamp type %T", v)
	}

	m.time = time.Unix(timestamp, 0).UTC()
	return nil
}

// framework.Namespace impl.
func (m *root) Get(key string) (interface{}, error) {
	switch key {
	case "minute":
		return time.Now().Minute(), nil
	case "month":
		return &namespaceMonth{Month: m.time.Month()}, nil
	case "location":
		return &Location{Name: "somewhere", TimeZone: "some zone"}, nil
	}
	return nil, nil
}

type namespaceMonth struct{ Month time.Month }

func (m *namespaceMonth) Get(key string) (interface{}, error) {
	switch key {
	case "string":
		return m.Month.String(), nil

	case "index":
		return int(m.Month), nil
	}

	return nil, nil
}

func (m *root) Func(key string) interface{} {
	switch key {
	case "add_month":
		return m.addMonth
	}

	return nil
}

func (m *root) addMonth(n int) *namespaceMonth {
	return &namespaceMonth{Month: m.time.AddDate(0, n, 0).Month()}
}
