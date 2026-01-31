package models

import (
	"time"
)

//any event must have these lables.
// we can use struct tags to enforce some of these labels (make them must be included in the 
// request body)
type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (evt Event) Save() {
	// we will save the event to a database later...
	events = append(events, evt)
}

func GetAllEvents() []Event {
	//later: might convert it to json format and return....
	// the data will come from the database.
	return events
}


