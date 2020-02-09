package errors

type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	ErrEventBusy     = EventError("Vremya zanyato")
	ErrEventNotFound = EventError("Sobutie ne naideno")
	ErrTimeRange     = EventError("Nepravilnyi period")
	ErrFieldEmpty    = EventError("Ne hvataet parametrov")
)
