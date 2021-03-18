package errors

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
	"github.com/pkg/errors"
)

func TestWrapf(t *testing.T) {
	type args struct {
		err          error
		formatString string
		args         []interface{}
	}

	tests := []struct {
		name      string
		args      args
		expected  error
		wantError bool
	}{
		{
			name: "Test no variadic args",
			args: args{
				err:          fmt.Errorf("test"),
				formatString: "test",
				args:         nil,
			},
			expected:  errors.Wrap(fmt.Errorf("test"), "test"),
			wantError: false,
		},
		{
			name: "Test no variadic args - failure",
			args: args{
				err:          fmt.Errorf("test"),
				formatString: "test",
				args:         nil,
			},
			expected:  errors.Wrap(fmt.Errorf("test"), "test2"),
			wantError: true,
		},
		{
			name: "Test int variadic args",
			args: args{
				err:          fmt.Errorf("test"),
				formatString: "test %d %d",
				args:         []interface{}{1, 2},
			},
			expected:  errors.Wrap(fmt.Errorf("test"), "test 1 2"),
			wantError: false,
		},
		{
			name: "Test int variadic args - failure",
			args: args{
				err:          fmt.Errorf("test"),
				formatString: "test %d %d",
				args:         []interface{}{1, 2},
			},
			expected:  errors.Wrap(fmt.Errorf("test"), "test 1 3"),
			wantError: true,
		},
		{
			name: "Test string variadic args",
			args: args{
				err:          fmt.Errorf("test"),
				formatString: "test %s %s",
				args:         []interface{}{"1", "2"},
			},
			expected:  errors.Wrap(fmt.Errorf("test"), "test 1 2"),
			wantError: false,
		},
		{
			name: "Test string variadic args - failure",
			args: args{
				err:          fmt.Errorf("test"),
				formatString: "test %s %s",
				args:         []interface{}{"1", "2"},
			},
			expected:  errors.Wrap(fmt.Errorf("test"), "test 1 3"),
			wantError: true,
		},
		{
			name: "Test mixed variadic args",
			args: args{
				err:          fmt.Errorf("test"),
				formatString: "test %s %d",
				args:         []interface{}{"1", 2},
			},
			expected:  errors.Wrap(fmt.Errorf("test"), "test 1 2"),
			wantError: false,
		},
		{
			name: "Test mixed variadic args - failure",
			args: args{
				err:          fmt.Errorf("test"),
				formatString: "test %s %d",
				args:         []interface{}{"1", 2},
			},
			expected:  errors.Wrap(fmt.Errorf("test"), "test 1 3"),
			wantError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrapf(tt.args.err, tt.args.formatString, tt.args.args...)
			if diff := deep.Equal(err.Error(), tt.expected.Error()); diff != nil {
				if tt.wantError == false {
					t.Error(diff)
				}
			}
		})
	}
}
