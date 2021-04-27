package pkg

import (
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/pkg/errors"
)

func UpdateLambdaEnv(inputs *Inputs, lambdaApi lambdaiface.LambdaAPI) error {

	lambda := NewLambda(lambdaApi)

	functionEnv, err := lambda.ReadEnv(inputs.FunctionName)
	if err != nil {
		return errors.Wrap(err, "Unable to read function env")
	}

	functionEnv.Merge(inputs.Env)

	err = lambda.WriteEnv(inputs.FunctionName, functionEnv)
	if err != nil {
		return errors.Wrap(err, "Unable to update function env")
	}

	return nil
}