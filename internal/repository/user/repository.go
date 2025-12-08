package userRepository

import "github.com/Masterminds/squirrel"

type repository struct {
	pgDB pgDBInstance

	sq squirrel.StatementBuilderType
}

func New(
	pgDBInstance pgDBInstance,
) *repository {
	return &repository{
		pgDB: pgDBInstance,

		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
