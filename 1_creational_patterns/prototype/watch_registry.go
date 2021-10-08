package prototype

// All types of watch.
const (
	WATCH_ANALOG = iota
	WATCH_DIGITAL
)

// WatchRegistry stores prototype instances.
type WatchRegistry struct {
	Registry map[int]Watch
}

// InitWatchRegistry initilizes watch registry.
func InitWatchRegistry() *WatchRegistry {
	aw := &AnalogWatch{}
	aw.SetModel("ANALOG")

	dw := &DigitalWatch{}
	dw.SetModel("DIGITAL")
	dw.SetAlarm(true)

	return &WatchRegistry{
		Registry: map[int]Watch{
			WATCH_ANALOG:  aw,
			WATCH_DIGITAL: dw,
		},
	}
}

// GetWatch retrieves registered watch instance by type.
func (r *WatchRegistry) GetWatch(typ int) (Watch, error) {
	return r.Registry[typ].Clone()
}
