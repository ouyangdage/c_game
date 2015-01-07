package token

import (
	"crypto/md5"
	"encoding/hex"
	"libs/redis"
	"strconv"
	"time"
)

var GameToken = NewToken(redis.Redis)

func init() {
}

type adapter interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Del(key string) error
}

type Token struct {
	adapter  adapter
	isUnique bool
}

func NewToken(a adapter) *Token {
	return &Token{a, true}
}

func (this *Token) NotUnique() {
	this.isUnique = false
}

// get uid from token
func (this *Token) GetUid(token string) (int, error) {

	if str, err := this.adapter.Get(token); err != nil {
		return 0, err
	} else {
		return strconv.Atoi(str)
	}
}

// create new token
func (this *Token) AddToken(uid int) (string, error) {

	m := md5.New()
	m.Write([]byte(strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.Itoa(uid)))
	token := hex.EncodeToString(m.Sum(nil))

	if this.isUnique {
		this.setUidToken(uid, token)
	}

	return token, this.adapter.Set(token, strconv.Itoa(int(uid)))
}

func (this *Token) setUidToken(uid int, token string) error {

	key := "uid_token_" + strconv.Itoa(uid)

	if oldToken, err := this.adapter.Get(key); err == nil {
		this.adapter.Del(oldToken)
		this.adapter.Del("TOKEN_ROLE_" + oldToken)
	}

	return this.adapter.Set(key, token)
}

func (this *Token) BindingServerRoleId(token string, roleId int) error {

	key := "TOKEN_ROLE_" + token
	return this.adapter.Set(key, strconv.Itoa(roleId))

	//	key := "TOKEN_SERVER_" + token
	//
	//	return this.adapter.Set(key, strconv.Itoa(int(serverId)))
}

func (this *Token) GetRoleId(token string) (int, error) {

	if str, err := this.adapter.Get("TOKEN_ROLE_" + token); err != nil {
		return 0, err
	} else {
		return strconv.Atoi(str)
	//	strconv.ParseInt(str, 10, 0)
	}
}
