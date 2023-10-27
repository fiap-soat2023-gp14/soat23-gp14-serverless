package settings

import "os"

func GetUserPoolId() string {
	return os.Getenv("AWS_USER_POOL_ID")
}

func GetRegion() string {
	return os.Getenv("AWS_USER_POOL_REGION")
}

func GetClientId() string {
	return os.Getenv("AWS_CLIENT_ID")
}
