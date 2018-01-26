package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) spawnRoutes() (routes []*route) {
	routes = []*route{
		//Spawn
		{
			"ListSpawn",
			"GET",
			"/spawn",
			a.listSpawn,
		},
		{
			"GetSpawn",
			"GET",
			"/spawn/{spawnID}/details",
			a.getSpawn,
		},
	}
	return
}

func (a *Web) listSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Spawns []*model.Spawn
	}

	site := a.newSite(r)
	site.Page = "spawn"
	site.Title = "Spawn"
	site.Section = "spawn"

	spawns, err := a.spawnRepo.List(user)
	if err != nil {
		return
	}

	content = Content{
		Site:   site,
		Spawns: spawns,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawn/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("spawn", tmp)
	}

	return
}

func (a *Web) getSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Spawn *model.Spawn
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}
	spawn := &model.Spawn{
		ID: spawnID,
	}
	err = a.spawnRepo.Get(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "spawn"
	site.Title = "Spawn"
	site.Section = "spawn"

	content = Content{
		Site:  site,
		Spawn: spawn,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawn/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("spawn", tmp)
	}

	return
}
