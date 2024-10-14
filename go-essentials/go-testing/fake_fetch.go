package main

import (
	"errors"
	"testing"
)

type Record struct {
	Name string
	Age  int
}

type recorder interface {
	Record(name string) (Record, error)
}

type fakeRecorder struct {
	rec Record
	err bool
}

func (f fakeRecorder) Record(name string) (Record, error) {
	if f.err {
		return Record{}, errors.New("error")
	}
	return f.rec, nil
}

func TestGreeter(t *testing.T) {
	tests := []struct {
		desc      string
		name      string
		recorder  recorder
		want      string
		expectErr bool
	}{
		{
			desc:      "Error: recorder had some server error",
			name:      "John",
			recorder:  fakeRecorder{err: true},
			expectErr: true,
		},
		{
			desc: "Error: server returned wrong name",
			name: "John",
			recorder: fakeRecorder{
				rec: Record{Name: "Bob", Age: 20},
			},
			expectErr: true,
		},
		{
			desc: "Success",
			name: "John",
			recorder: fakeRecorder{
				rec: Record{Name: "John", Age: 20},
			},
			want: "Greetings John",
		},
	}

	// Executes each test.
	for _, test := range tests {
		got, err := Greeter(test.name, test.recorder)
		switch {
		// We did not get an error, but expected one
		case err == nil && test.expectErr:
			t.Errorf("TestGreet(%s): got err == nil, want err != nil", test.desc)
			continue
		// We got an error but did not expect one
		case err != nil && !test.expectErr:
			t.Errorf("TestGreet(%s): got err == %s, want err == nil", test.desc, err)
			continue
		// We got an error we expected, so just go to the next test
		case err != nil:
			continue
		}

		// We did not get the result we expected
		if got != test.want {
			t.Errorf("TestGreet(%s): got result %q, want %q", test.desc, got, test.want)
		}
	}
}