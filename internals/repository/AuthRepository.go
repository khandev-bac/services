package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/services/db-models"
)

type DBQueries struct {
	db *db.Queries
}

func NewRepository(db *db.Queries) *DBQueries {
	return &DBQueries{
		db: db,
	}
}
func (dr *DBQueries) CreateUser(ctx context.Context, userinfo db.SingupParams) (db.SingupRow, error) {
	return dr.db.Singup(ctx, userinfo)
}
func (dr *DBQueries) CreateUserWithGoogle(ctx context.Context, googleUsers db.GoogleloginParams) (db.GoogleloginRow, error) {
	return dr.db.Googlelogin(ctx, googleUsers)
}
func (dr *DBQueries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return dr.db.DeleteUser(ctx, id)
}
func (dr *DBQueries) EditPicture(ctx context.Context, newPic db.EditPictureParams) error {
	return dr.db.EditPicture(ctx, newPic)
}

func (dr *DBQueries) EditRevoke(ctx context.Context, revokeArg db.EditRevokeParams) error {
	return dr.db.EditRevoke(ctx, revokeArg)
}
func (dr *DBQueries) EditUser(ctx context.Context, newuserinfo db.EditUserParams) error {
	return dr.db.EditUser(ctx, newuserinfo)
}

func (dr *DBQueries) FindByEmail(ctx context.Context, email string) (db.FindByEmailRow, error) {
	return dr.db.FindByEmail(ctx, email)
}
func (dr *DBQueries) FindById(ctx context.Context, id uuid.UUID) (db.FindByIdRow, error) {
	return dr.db.FindById(ctx, id)
}
func (dr *DBQueries) FindByUserName(ctx context.Context, username string) (db.FindByUserNameRow, error) {
	var nullUsername sql.NullString
	if username != "" {
		nullUsername = sql.NullString{
			String: username,
			Valid:  true,
		}
	} else {
		nullUsername = sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	return dr.db.FindByUserName(ctx, nullUsername)
}
func (dr *DBQueries) SeeRevoke(ctx context.Context, id uuid.UUID) (sql.NullBool, error) {
	return dr.db.SeeRevoke(ctx, id)
}
