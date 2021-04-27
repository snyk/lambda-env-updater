package pkg

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
)

type Lambda struct {
	client lambdaiface.LambdaAPI
}

func NewLambda(client lambdaiface.LambdaAPI) *Lambda {
	return &Lambda{client: client}
}

func (l *Lambda) ReadEnv(functionName string) (Env, error) {
	config, err := l.client.GetFunctionConfiguration(&lambda.GetFunctionConfigurationInput{
		FunctionName: &functionName,
	})
	if err != nil {
		return nil, err
	}

	return config.Environment.Variables, nil
}

func (l *Lambda) WriteEnv(functionName string, env Env) error {
	_, err := l.client.UpdateFunctionConfiguration(&lambda.UpdateFunctionConfigurationInput{
		FunctionName: &functionName,
		Environment: &lambda.Environment{
			Variables: env,
		},
	})

	if err != nil {
		return err
	}

	return nil
}