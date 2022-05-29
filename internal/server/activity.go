package server

import (
	"fmt"
	//"sync"
	"time"
)


var ErrIDNotFound = fmt.Errorf("ID not found")



type Activity struct {
	Time        time.Time `json:"time"`
	Description string    `json:"description"`
	ID          uint64    `json:"id"`
}



type Activities struct {
	activities []Activity
}



func (activities *Activities) Insert(activity Activity) uint64 {

	activity.ID = uint64(len(activities.activities))
	activities.activities = append(activities.activities, activity)

	return activity.ID
	
}



func (activities *Activities) Retrieve(id uint64) (Activity, error) {

	if id >= uint64(len(activities.activities)) {
		return Activity{}, ErrIDNotFound
	}
	
	return activities.activities[id], nil
	
}