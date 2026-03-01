package main

import (
	"net/http"

	"github.com/paccolamano/plugincmd"
	"github.com/pocketbase/pocketbase/core"
)

type KeycloakAuthPlugin struct{}

func (p *KeycloakAuthPlugin) Name() string {
	return "keycloakauth"
}

func (p *KeycloakAuthPlugin) Version() string {
	return "1.2.36"
}

func (p *KeycloakAuthPlugin) Register(app core.App) error {
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/keycloakauth/version", func(e *core.RequestEvent) error {
			return e.JSON(http.StatusOK, map[string]string{
				"version": "1.2.36",
			})
		})
		return se.Next()
	})

	return nil
}

var Plugin plugincmd.PBPlugin = &KeycloakAuthPlugin{}
