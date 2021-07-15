package main

import (
	"context"

	"github.com/rekhin/repository/example/object/mongodb"
)

func main() {
	repo := mongodb.NewRepository()
	ctx := context.Background()
	repo.CreateObject(ctx, mongodb.Object{ID: 1})
}
