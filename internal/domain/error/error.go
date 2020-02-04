package error

type EventError string

func (ee EventError) Error() string {
	return string(ee)
}

var (
	ErrBusy = EventError("Vremya zanyato")
)
