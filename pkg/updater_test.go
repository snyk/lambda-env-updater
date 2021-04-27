package pkg

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/cloudskiff/lambda-env-updater/test"
)

func TestUpdateLambdaEnv(t *testing.T) {
	tests := []struct {
		name    string
		inputs *Inputs
		mocks func(api *test.MockLambdaAPI)
		err error
	}{
		{
			name: "error reading lambda env",
			inputs: &Inputs{
				FunctionName: "hello-world",
			},
			mocks: func(api *test.MockLambdaAPI) {
				api.On("GetFunctionConfiguration",
					&lambda.GetFunctionConfigurationInput{
						FunctionName: aws.String("hello-world"),
					},
				).Once().Return(nil, errors.New("404 function not found"))
			},
			err: errors.New("Unable to read function env: 404 function not found"),
		},
		{
			name: "error updating lambda env",
			inputs: &Inputs{
				FunctionName: "hello-world",
			},
			mocks: func(api *test.MockLambdaAPI) {
				api.On("GetFunctionConfiguration",
					&lambda.GetFunctionConfigurationInput{
						FunctionName: aws.String("hello-world"),
					},
				).Once().Return(&lambda.FunctionConfiguration{
					Environment: &lambda.EnvironmentResponse{
						Variables: Env{
							"FOO": aws.String("FOO"),
						},
					},
				}, nil)

				api.On("UpdateFunctionConfiguration",
					&lambda.UpdateFunctionConfigurationInput{
						FunctionName: aws.String("hello-world"),
						Environment: &lambda.Environment{
							Variables: Env{
								"FOO": aws.String("FOO"),
							},
						},
					},
				).Return(nil, errors.New("403 access denied"))
			},
			err: errors.New("Unable to update function env: 403 access denied"),
		},
		{
			name: "update existing value",
			inputs: &Inputs{
				FunctionName: "hello-world",
				Env: map[string]*string{
					"FOO": aws.String("BAR"),
				},
			},
			mocks: func(api *test.MockLambdaAPI) {
				api.On("GetFunctionConfiguration",
					&lambda.GetFunctionConfigurationInput{
						FunctionName: aws.String("hello-world"),
					},
				).Once().Return(&lambda.FunctionConfiguration{
					Environment: &lambda.EnvironmentResponse{
						Variables: Env{
							"FOO": aws.String("FOO"),
						},
					},
				}, nil)

				api.On("UpdateFunctionConfiguration",
					&lambda.UpdateFunctionConfigurationInput{
						FunctionName: aws.String("hello-world"),
						Environment: &lambda.Environment{
							Variables: Env{
								"FOO": aws.String("BAR"),
							},
						},
					},
				).Return(nil, nil)
			},
		},
		{
			name: "add new value and keep existing ones",
			inputs: &Inputs{
				FunctionName: "hello-world",
				Env: map[string]*string{
					"NEW": aws.String("VALUE"),
				},
			},
			mocks: func(api *test.MockLambdaAPI) {
				api.On("GetFunctionConfiguration",
					&lambda.GetFunctionConfigurationInput{
						FunctionName: aws.String("hello-world"),
					},
				).Once().Return(&lambda.FunctionConfiguration{
					Environment: &lambda.EnvironmentResponse{
						Variables: Env{
							"FOO": aws.String("BAR"),
							"BAR": aws.String("FOO"),
						},
					},
				}, nil)

				api.On("UpdateFunctionConfiguration",
					&lambda.UpdateFunctionConfigurationInput{
						FunctionName: aws.String("hello-world"),
						Environment: &lambda.Environment{
							Variables: Env{
								"FOO": aws.String("BAR"),
								"BAR": aws.String("FOO"),
								"NEW": aws.String("VALUE"),
							},
						},
					},
				).Return(nil, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			fakeApi := test.MockLambdaAPI{}
			if tt.mocks != nil {
				tt.mocks(&fakeApi)
			}
			err := UpdateLambdaEnv(tt.inputs, &fakeApi)
			if err != nil {
				assert.EqualError(tt.err, err.Error())
			}
		})
	}
}
