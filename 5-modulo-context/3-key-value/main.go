package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "password-secret")
	handler(ctx, "name")
}

func handler(ctx context.Context, name string) {
	token := ctx.Value("token")
	if token != nil {
		println()
		fmt.Println("Token:", token)
		fmt.Println("Name:", name)
	}
}
