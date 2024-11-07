package utils

import "time"

func ParseStringToTime (dateString string) (time.Time, error) {
	var format string
	var date time.Time

	// Commonly the Pub dates are just the year
	if len(dateString) == 4 {
		format = "2006"
	} else {
		format = "January 02, 2006"
	}

	date, err := time.Parse(format, dateString)

	if err != nil {
		format = "January 2, 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "January 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "Jan 2, 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "Jan 02, 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2006-01-02"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2006-Jan-02"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2006-January-02"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "02 Jan 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	if err != nil {
		format = "2 Jan 2006"
		date, err = time.Parse(format, dateString)
		if err == nil {return date, nil}
	}

	return date, err
}
