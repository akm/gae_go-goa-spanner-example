package controller

import (
	"context"
	"os"

	"cloud.google.com/go/spanner"
	// database "cloud.google.com/go/spanner/admin/database/apiv1"
	// "google.golang.org/api/iterator"

	// adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"

	"google.golang.org/appengine/log"
)

func createClient(ctx context.Context, f func(*spanner.Client) error) error {
	db := os.Getenv("SPANNER_DATABASE")
	c, err := spanner.NewClient(ctx, db)
	if err != nil {
		log.Errorf(ctx, "Failed to spanner.NewClient for %s\n", db)
		return err
	}
	return f(c)
}
