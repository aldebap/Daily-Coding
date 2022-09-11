////////////////////////////////////////////////////////////////////////////////
//	intervalOverlap.go  -  Sep-9-2022  -  aldebap
//
//	Problem of Date invervals overlap
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"time"
)

type dateInterval struct {
	start string
	end   string
}

func CheckOverlap(incoming dateInterval, existing []dateInterval) (bool, error) {

	//	to make the interval dates comparisons, convert dates to time
	incomingStart, err := time.Parse("2006-01-02", incoming.start)
	if err != nil {
		return false, err
	}
	incomingEnd, err := time.Parse("2006-01-02", incoming.end)
	if err != nil {
		return false, err
	}

	for _, aux := range existing {
		existingStart, err := time.Parse("2006-01-02", aux.start)
		if err != nil {
			return false, err
		}
		existingEnd, err := time.Parse("2006-01-02", aux.end)
		if err != nil {
			return false, err
		}

		//	check if the start date of incoming is between one of the existing intervals
		if existingStart.Before(incomingStart) && existingEnd.After(incomingStart) {
			return true, nil
		}

		//	check if the end date of incoming is between one of the existing intervals
		if existingStart.Before(incomingEnd) && existingEnd.After(incomingEnd) {
			return true, nil
		}

		//	check if any of the existing intervals is between the incoming interval
		if incomingStart.Before(existingStart) && incomingEnd.After(existingStart) {
			return true, nil
		}
		if incomingStart.Before(existingEnd) && incomingEnd.After(existingEnd) {
			return true, nil
		}
	}

	return false, nil
}
