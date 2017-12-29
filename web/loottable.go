package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListLootTable(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       Site
		LootTables []*model.LootTable
	}

	site := a.NewSite(r)
	site.Page = "lootTable"
	site.Title = "LootTable"
	site.Section = "lootTable"

	lootTables, err := a.lootTableRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:       site,
		LootTables: lootTables,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "loottable/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("lootTable", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) GetLootTable(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      Site
		LootTable *model.LootTable
	}

	id, err := getIntVar(r, "lootTableId")
	if err != nil {
		err = errors.Wrap(err, "lootTableId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	lootTable, err := a.lootTableRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
	site.Page = "lootTable"
	site.Title = "LootTable"
	site.Section = "lootTable"

	content := Content{
		Site:      site,
		LootTable: lootTable,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "loottable/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("loottable", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}