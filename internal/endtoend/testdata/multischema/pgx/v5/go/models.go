// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bar struct {
	ID int32
}

type Foo struct {
	ID  int32
	Bar pgtype.Int4
}
