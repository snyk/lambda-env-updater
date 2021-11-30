# lambda-env-updater

## Why ?

AWS CLI allows you to update your lambda function's configuration with the following command:

```shell
aws lambda update-function-configuration --function-name <function> --environment <env-as-json>
```

Unfortunately this command replaces the whole environemnt with the newly provided, making it impossible to update just one env variable.

## Usage

```shell
lambda-env-updater -name function-name -env FOO=BAR -env BAR=FOO
```

## Contributing

Should you wish to make a contribution please open a pull request against this repository with a clear description of the change with tests demonstrating the functionality. You will also need to agree to the [Contributor Agreement](./Contributor-Agreement.md) before the code can be accepted and merged.
