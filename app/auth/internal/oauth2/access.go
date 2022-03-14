package oauth2

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"github.com/mars-projects/oauth2/v4"
	"strconv"
	"strings"
)

// NewAccessGenerate create to generate the access token instance
func NewAccessGenerate() *AccessGenerate {
	return &AccessGenerate{}
}

// AccessGenerate generate the access token
type AccessGenerate struct{}

// Token based on the UUID generated token
func (ag *AccessGenerate) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (string, string, error) {
	buf := bytes.NewBufferString(data.Client.GetID())
	buf.WriteString(data.UserID)
	buf.WriteString(strconv.FormatInt(data.CreateAt.UnixNano(), 10))
	access := uuid.NewMD5(uuid.Must(uuid.NewRandom()), buf.Bytes()).String()
	access = strings.TrimRight(access, "=")
	refresh := ""
	if isGenRefresh {
		refresh = uuid.NewSHA1(uuid.Must(uuid.NewRandom()), buf.Bytes()).String()
		refresh = strings.TrimRight(refresh, "=")
	}
	return access, refresh, nil
}
