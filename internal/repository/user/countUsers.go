package userRepository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx"
)

func (t *repository) CountUsers() (int, error) {
	var result int
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query, args, err := t.sq.
		Select(fmt.Sprintf("COUNT(%v)", allFields)).
		From(tableName).
		ToSql()

	if err != nil {
		return 0, err
	}

	err = t.pgDB.QueryRow(ctx, query, args...).Scan(&result)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, nil
	}

	return result, err
}
