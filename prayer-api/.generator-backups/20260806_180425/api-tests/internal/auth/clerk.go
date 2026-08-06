package auth

import "github.com/clerk/clerk-sdk-go/v2"

func Configure(secretKey string) {
	clerk.SetKey(secretKey)
}
