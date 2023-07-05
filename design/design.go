package design

import (
	"time"

	. "goa.design/goa/v3/dsl"
)

var _ = API("api", func() {
	Title("Scheduler service")
	Description("HTTP service for scheduler events")
	Server("scheduler", func() {
		Host("localhost", func() { URI("http://localhost:8088") })
	})
})

var ThingWithTimeslice = Type("ThingWithTimeslice", func() {
	Attribute("timeslice", func() {
		Default([]time.Time{})
		Meta("struct:field:type", "[]time.Time", "time")
		Meta("struct:tag:json", "-")
	})
})

// Service describes a service
var _ = Service("service", func() {

	Method("get-thing", func() {
		Payload(ThingWithTimeslice)
		Result(ThingWithTimeslice)
		HTTP(func() {
			POST("/")
			Response(StatusOK)
		})
	})
})
