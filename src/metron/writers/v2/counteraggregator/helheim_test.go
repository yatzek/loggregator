// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package counteraggregator_test

import v2 "plumbing/v2"

type mockWriter struct {
	WriteCalled chan bool
	WriteInput  struct {
		Msg chan *v2.Envelope
	}
	WriteOutput struct {
		Err chan error
	}
}

func newMockWriter() *mockWriter {
	m := &mockWriter{}
	m.WriteCalled = make(chan bool, 100)
	m.WriteInput.Msg = make(chan *v2.Envelope, 100)
	m.WriteOutput.Err = make(chan error, 100)
	return m
}
func (m *mockWriter) Write(msg *v2.Envelope) (err error) {
	m.WriteCalled <- true
	m.WriteInput.Msg <- msg
	return <-m.WriteOutput.Err
}
