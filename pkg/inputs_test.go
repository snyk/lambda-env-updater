package pkg

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

func TestParseInputs(t *testing.T) {
	tests := []struct {
		name    string
		want    *Inputs
		args []string
		err string
	}{
		{
			name: "no flags provided",
			err: "No function name specified",
		},
		{
			name: "no env provided",
			args: []string{
				"-name", "test",
			},
			err: "No env var to update, exiting",
		},
		{
			name: "invalid env",
			args: []string{
				"-name", "test",
				"-env", "FOOBAR",
			},
			err: "invalid value \"FOOBAR\" for flag -env: Unable to parse environment variable parameter: FOOBAR",
		},
		{
			name: "valid inputs",
			args: []string{
				"-name", "test",
				"-env", "FOO=BAR",
				"-env", "BAR=FOO",
			},
			want: &Inputs{
				FunctionName: "test",
				Env: Env{
					"FOO": aws.String("BAR"),
					"BAR": aws.String("FOO"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			got, err := ParseInputs(tt.args)
			if err != nil && tt.err == "" {
				t.Fatalf("Expected no error but got: '%s'", err.Error())
			}
			if tt.err != "" {
				assert.EqualError(err, tt.err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInputs() got = %v, want %v", got, tt.want)
			}
		})
	}
}
