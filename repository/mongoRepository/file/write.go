package file

import (
	"context"
	"mediaStorer/param/storage"
)

func (d *DB) Insert(ctx context.Context, file storage.File) {
	d.conn.Collection.InsertOne(ctx, file)
}

func (d *DB) Delete() {
}
