package oauth2

import (
	"github.com/mars-projects/mars/conf"
	"github.com/mars-projects/oauth2/v4"
	"github.com/mars-projects/oauth2/v4/manage"
	"github.com/mars-projects/oauth2/v4/models"
	"github.com/mars-projects/oauth2/v4/store"
	"time"
)

func NewManage(tokenStore oauth2.TokenStore, auth *conf.Auth) *manage.Manager {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(tokenStore, nil)
	manager.MapAccessGenerate(NewAccessGenerate())
	config := &manage.Config{
		AccessTokenExp:    time.Second * time.Duration(auth.AccessExpired) * 24,
		RefreshTokenExp:   time.Second * time.Duration(auth.RefreshExpired) * 24,
		IsGenerateRefresh: true,
	}
	manager.SetPasswordTokenCfg(config)
	// client memory store
	clientStore := store.NewClientStore()
	clientStore.Set(auth.ClientId, &models.Client{
		ID:     auth.ClientId,
		Secret: auth.ClientSecret,
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)
	return manager
}
