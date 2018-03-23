package hub

import (
	"log"
	"time"
)

// Fields is a [key]value storage for events.
type Fields map[string]interface{}

// Event represent something that occurred with the application
// Event also contain some helper functions to convert the fields to primitive types.
type Event struct {
	Name   string
	Fields Fields
}

// Int return an int field of an Event.
func (e *Event) Int(key string) int {
	v, ok := e.Fields[key].(int)
	if !ok {
		log.Fatalf("Event %#v didn`t have the int field %s", e, key)
	}
	return v
}

// Int64 return an int64 field of an Event.
func (e *Event) Int64(key string) int64 {
	v, ok := e.Fields[key].(int64)
	if !ok {
		log.Fatalf("Event %#v didn`t have the int64 field %s", e, key)
	}
	return v
}

// Int32 return an int32 field of an Event.
func (e *Event) Int32(key string) int32 {
	v, ok := e.Fields[key].(int32)
	if !ok {
		log.Fatalf("Event %#v didn`t have the int32 field %s", e, key)
	}
	return v
}

// Float64 return a float64 field of an Event.
func (e *Event) Float64(key string) float64 {
	v, ok := e.Fields[key].(float64)
	if !ok {
		log.Fatalf("Event %#v didn`t have the float64 field %s", e, key)
	}
	return v
}

// String return a string field of an Event.
func (e *Event) String(key string) string {
	v, ok := e.Fields[key].(string)
	if !ok {
		log.Fatalf("Event %#v didn`t have the string field %s", e, key)
	}
	return v
}

// Duration return a time.Duration field of an Event.
func (e *Event) Duration(key string) time.Duration {
	v, ok := e.Fields[key].(time.Duration)
	if !ok {
		log.Fatalf("Event %#v didn`t have the time.Duration field %s", e, key)
	}
	return v
}
