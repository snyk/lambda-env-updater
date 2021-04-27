package pkg

import (
	"flag"

	"github.com/pkg/errors"
)

type Inputs struct {
	FunctionName string
	Env Env
}

func ParseInputs(args []string) (*Inputs, error) {
	var functionName string
	env := Env{}
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.StringVar(&functionName, "name", "", "Lambda function name")
	flags.Var(&env, "env", "Environment variable to set: -env \"MY_VAR=myvalue\"")
	err := flags.Parse(args)
	if err != nil {
		return nil, err
	}

	if functionName == "" {
		return nil, errors.New("No function name specified")
	}

	if len(env) == 0 {
		return nil, errors.New("No env var to update, exiting")
	}

	return &Inputs{Env: env, FunctionName: functionName}, nil
}
