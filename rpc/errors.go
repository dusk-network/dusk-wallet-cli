package rpc

// NetworkError is returned from an RPC call whenever it fails due to
// timeouts, bad requests or similar problems.
type NetworkError struct {
	err error
}

func (n *NetworkError) Error() string {
	return n.err.Error()
}

// MethodError is returned from an RPC call whenever the method call
// returned an error in the JSON-RPC response.
type MethodError struct {
	err error
}

func (m *MethodError) Error() string {
	return m.err.Error()
}
