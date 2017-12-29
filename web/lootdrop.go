package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListLootDrop(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      Site
		LootDrops []*model.LootDrop
	}

	site := a.NewSite(r)
	site.Page = "lootDrop"
	site.Title = "LootDrop"
	site.Section = "lootDrop"

	lootDrops, err := a.lootDropRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:      site,
		LootDrops: lootDrops,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdrop/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("lootDrop", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) GetLootDrop(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     Site
		LootDrop *model.LootDrop
	}

	id, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	lootDrop, err := a.lootDropRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
	site.Page = "lootDrop"
	site.Title = "LootDrop"
	site.Section = "lootDrop"

	content := Content{
		Site:     site,
		LootDrop: lootDrop,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdrop/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("lootdrop", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
