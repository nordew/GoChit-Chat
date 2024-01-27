package main

import (
	"context"
	"user/internal/app"
)

func init() {

}

func main() {
	app.MustRun(context.Background())
}
