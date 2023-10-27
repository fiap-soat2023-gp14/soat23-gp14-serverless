# soat23-gp14-serverless
The repository to serveless application


TODO:
[] Github Actions
[] Endpoint to Confirm SignUp
[] Fix terraform
[] Unit tests

## Tools needed:
- AWS SAM CLI
- Go 1.19

## Build
`sam build`

## Deploy
`sam deploy --guided`

## Local Testing
`sam local start-api --env-vars test/env.json`