package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type NpcLootRepository struct {
	stor storage.Storage
}

func (g *NpcLootRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *NpcLootRepository) Get(npcID int64, itemID int64) (npcLoot *model.NpcLoot, err error) {
	npcLoot, err = g.stor.GetNpcLoot(npcID, itemID)
	return
}

func (g *NpcLootRepository) Truncate() (err error) {
	err = g.stor.TruncateNpcLoot()
	return
}

func (g *NpcLootRepository) Create(npcLoot *model.NpcLoot) (err error) {
	if npcLoot == nil {
		err = fmt.Errorf("Empty npcLoot")
		return
	}
	schema, err := npcLoot.NewSchema([]string{}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(npcLoot))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	err = g.stor.CreateNpcLoot(npcLoot)
	if err != nil {
		return
	}
	return
}

func (g *NpcLootRepository) Edit(npcID int64, itemID int64, npcLoot *model.NpcLoot) (err error) {
	schema, err := npcLoot.NewSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(npcLoot))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}

	err = g.stor.EditNpcLoot(npcID, itemID, npcLoot)
	if err != nil {
		return
	}
	return
}

func (g *NpcLootRepository) Delete(npcID int64, itemID int64) (err error) {
	err = g.stor.DeleteNpcLoot(npcID, itemID)
	if err != nil {
		return
	}
	return
}

func (g *NpcLootRepository) List(npcID int64) (npcLoots []*model.NpcLoot, err error) {
	npcLoots, err = g.stor.ListNpcLoot(npcID)
	if err != nil {
		return
	}
	return
}

func (g *NpcLootRepository) ListByZone(zoneID int64) (npcLoots []*model.NpcLoot, err error) {
	npcLoots, err = g.stor.ListNpcLootByZone(zoneID)
	if err != nil {
		return
	}
	return
}
