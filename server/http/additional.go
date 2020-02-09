package http

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type EventHTTP struct {
	ID          uint      `json:"ID"`
	Owner       string    `json:"owner"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartString string    `json:"-"`
	EndString   string    `json:"-"`
	Start       time.Time `json:"start"`
	End         time.Time `json:"end"`
}

type HTTPError struct {
	Error error `json:"error"`
}

func toTimeFromString(source string) (time.Time, error) {
	slice := strings.Split(source, "-")
	if len(slice) != 3 {
		return time.Now(), errors.New("Nepravilnui format(need yyyy-mm-dd")
	}
	if val, err := strconv.Atoi(slice[0]); err != nil || val < (time.Now().Year()-10) || val > (time.Now().Year()+10) {
		return time.Now(), errors.New("Nepravilnui format(need yyyy-mm-dd")
	}
	if val, err := strconv.Atoi(slice[1]); err != nil || val < 1 || val > 12 {
		return time.Now(), errors.New("Nepravilnui format(need yyyy-mm-dd")
	}
	if val, err := strconv.Atoi(slice[2]); err != nil || val < 1 || val > 31 {
		return time.Now(), errors.New("Nepravilnui format(need yyyy-mm-dd")
	}
	switch {
	case slice[1] == "1" || slice[1] == "01":
		slice[1] = "Jan"
	case slice[1] == "2" || slice[1] == "02":
		slice[1] = "Feb"
	case slice[1] == "3" || slice[1] == "03":
		slice[1] = "Mar"
	case slice[1] == "4" || slice[1] == "04":
		slice[1] = "Apr"
	case slice[1] == "5" || slice[1] == "05":
		slice[1] = "May"
	case slice[1] == "6" || slice[1] == "06":
		slice[1] = "Jun"
	case slice[1] == "7" || slice[1] == "07":
		slice[1] = "Jul"
	case slice[1] == "8" || slice[1] == "08":
		slice[1] = "Aug"
	case slice[1] == "9" || slice[1] == "09":
		slice[1] = "Sep"
	case slice[1] == "10":
		slice[1] = "Oct"
	case slice[1] == "11":
		slice[1] = "Nov"
	case slice[1] == "12":
		slice[1] = "Dec"
	}
	shortForm := "2006-Jan-02"
	return time.Parse(shortForm, strings.Join(slice, "-"))
}
