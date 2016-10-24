package thrift

// Ping is implementing a basic ping function in the Flipd service.
func (f *Flipd) Ping() (string, error) { return "pong", nil }
