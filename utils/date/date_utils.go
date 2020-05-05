package date

import "time"

	// date time format always stays the same for the reference inside the quotes.  
	// can view format if you go to def of now.Format()
	// updating date created to user struct.

const(
	apiDateLayout = "2006-01-02-T15:04:05Z"
	apiDbLayout = "2006-01-02 15:04:05"
)

// GetNow ... setting this function to return standard time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString ...
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDbFormat ...
func GetNowDbFormat() string {
	return GetNow().Format(apiDbLayout)
}