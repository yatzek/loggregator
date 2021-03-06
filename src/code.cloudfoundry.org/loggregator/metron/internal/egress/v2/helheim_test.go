// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package v2_test

import v2 "code.cloudfoundry.org/loggregator/plumbing/v2"

type mockNexter struct {
	TryNextCalled chan bool
	TryNextOutput struct {
		Ret0 chan *v2.Envelope
		Ret1 chan bool
	}
}

func newMockNexter() *mockNexter {
	m := &mockNexter{}
	m.TryNextCalled = make(chan bool, 100)
	m.TryNextOutput.Ret0 = make(chan *v2.Envelope, 100)
	m.TryNextOutput.Ret1 = make(chan bool, 100)
	return m
}
func (m *mockNexter) TryNext() (*v2.Envelope, bool) {
	m.TryNextCalled <- true
	return <-m.TryNextOutput.Ret0, <-m.TryNextOutput.Ret1
}

type mockWriter struct {
	WriteCalled chan bool
	WriteInput  struct {
		Msg chan []*v2.Envelope
	}
	WriteOutput struct {
		Ret0 chan error
	}
}

func newMockWriter() *mockWriter {
	m := &mockWriter{}
	m.WriteCalled = make(chan bool, 100)
	m.WriteInput.Msg = make(chan []*v2.Envelope, 100)
	m.WriteOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockWriter) Write(msg []*v2.Envelope) error {
	m.WriteCalled <- true
	m.WriteInput.Msg <- msg
	return <-m.WriteOutput.Ret0
}
