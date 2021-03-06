package tables

import (
	"back-src/model/existence"
	"github.com/go-pg/pg"
	"time"
)

type AuthTokenTable struct {
	conn *pg.DB
}

func NewAuthTokenTable(db *pg.DB) *AuthTokenTable {
	return &AuthTokenTable{db}
}

func (table *AuthTokenTable) MakeNewAuth(username, token string, isFreelancer bool) (string, error) {
	auth := existence.AuthToken{token, username, time.Now(), isFreelancer}
	if _, err := table.conn.Model(&auth).Insert(); err != nil {
		return "", err
	}
	return token, nil
}

func (table *AuthTokenTable) IsThereAuthWithToken(token string) (bool, error) {
	var resultSet []existence.AuthToken
	error := table.conn.Model(&resultSet).Where("token = ?", token).Select()
	return len(resultSet) != 0, error
}

func (table *AuthTokenTable) ExpireAuth(token string) error {
	_, err := table.conn.Model(&existence.AuthToken{}).Where("token = ?", token).Delete()
	return err
}

func (table *AuthTokenTable) GetAuthByToken(token string) (auth existence.AuthToken, e error) {
	e = table.conn.Model(&auth).Where("token = ?", token).Select()
	return
}

func (table *AuthTokenTable) GetUsernameByToken(token string) (username string, e error) {
	auth := existence.AuthToken{}
	e = table.conn.Model(&auth).Column("username").Where("token = ?", token).Select()
	username = auth.Username
	return
}

func (table *AuthTokenTable) GetAllTokens() ([]existence.AuthToken, error) {
	var auths []existence.AuthToken
	err := table.conn.Model(&auths).Column("token").Select()
	return auths, err
}
