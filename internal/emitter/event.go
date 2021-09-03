package emitter

const EMIT_CONNECTED = "Connected"
const EMIT_LINK = "Link"
const EMIT_INVITE = "Invite"
const EMIT_REPLY = "Reply"
const EMIT_CONFIRMED = "Confirmed"
const EMIT_COMPLETED = "Completed"
const EMIT_ERROR = "Error"
const EMIT_MAIL_EVENT = "MailEvent"
const EMIT_PROGRESS_EVENT = "ProgressEvent"
const EMIT_ROOM_EVENT = "RoomEvent"
const EMIT_SYNC_EVENT = "SyncEvent"

// Event is a structure to send events contains
// some helpers to cast primitive types easily.
type Event struct {
	Topic, OriginalTopic string
	Flags                Flag
	Args                 []interface{}
}

// Int returns casted into int type argument by index.
// `dflt` argument is an optional default value returned
// either in case of casting error or in case of index error.
func (e Event) Int(index uint, dflt ...int) int {
	var d int
	for _, first := range dflt {
		d = first
		break
	}
	if len(e.Args) > int(index) {
		if casted, okey := e.Args[index].(int); okey {
			d = casted
		}
	}
	return d
}

// String returns casted into string type argument by index.
// `dflt` argument is an optional default value returned
// either in case of casting error or in case of index error.
func (e Event) String(index uint, dflt ...string) string {
	var d string
	for _, first := range dflt {
		d = first
		break
	}
	if len(e.Args) > int(index) {
		if casted, okey := e.Args[index].(string); okey {
			d = casted
		}
	}
	return d
}

// Float returns casted into float64 type argument by index.
// `dflt` argument is an optional default value returned
// either in case of casting error or in case of index error.
func (e Event) Float(index uint, dflt ...float64) float64 {
	var d float64
	for _, first := range dflt {
		d = first
		break
	}
	if len(e.Args) > int(index) {
		if casted, okey := e.Args[index].(float64); okey {
			d = casted
		}
	}
	return d
}

// Bool returns casted into bool type argument by index.
// `dflt` argument is an optional default value returned
// either in case of casting error or in case of index error.
func (e Event) Bool(index uint, dflt ...bool) bool {
	var d bool
	for _, first := range dflt {
		d = first
		break
	}
	if len(e.Args) > int(index) {
		if casted, okey := e.Args[index].(bool); okey {
			d = casted
		}
	}
	return d
}