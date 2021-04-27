package test

import "github.com/aws/aws-sdk-go/service/lambda/lambdaiface"

type LambdaAPI interface {
	lambdaiface.LambdaAPI
}
