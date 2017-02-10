package thrift

// Ping is implementing a basic ping function in the Flipd service.
func (f *Handler) Ping() (string, error) {
	return "pong", nil
}
