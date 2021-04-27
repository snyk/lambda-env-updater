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