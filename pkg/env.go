package pkg

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Env map[string]*string

func (e *Env) Merge(newEnv Env) {
	for k,v := range newEnv {
		(*e)[k] = v
	}
}

func (e *Env) String() string {
	return fmt.Sprintf("%+v", *e)
}

func (e *Env) Set(s string) error {
	splitedEnv := strings.SplitN(s, "=", 2)
	if len(splitedEnv) != 2 {
		return errors.Errorf("Unable to parse environment variable parameter: %s", s)
	}
	(*e)[splitedEnv[0]] = &splitedEnv[1]
	return nil
}
