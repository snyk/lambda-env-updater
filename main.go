package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	awslambda "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/cloudskiff/lambda-env-updater/pkg"
)

func main() {

	inputs, err := pkg.ParseInputs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	err = pkg.UpdateLambdaEnv(inputs, awslambda.New(sess))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("Function configuration updated with success")
}
