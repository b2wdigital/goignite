package gimemdb

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/hashicorp/go-memdb"
)

func NewDatabase(ctx context.Context, schema *memdb.DBSchema) (db *memdb.MemDB, err error) {

	logger := gilog.FromContext(ctx)

	db, err = memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	logger.Info("Connected to go-memdb")

	return db, err

}
