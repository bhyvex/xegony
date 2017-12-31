package web

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//ApplyRoutes applies routes to given mux router
func (a *Web) ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("WEB_ROOT")
	if len(rootPath) == 0 {
		rootPath = ""
	}

	type Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}

	routes := []Route{
		{
			"Index",
			"GET",
			"/",
			a.listForum,
		},

		{
			"GetDashboard",
			"GET",
			"/dashboard",
			a.getDashboard,
		},

		{
			"Login",
			"GET",
			"/login",
			a.getLogin,
		},

		{
			"Logout",
			"GET",
			"/logout",
			a.getLogout,
		},

		{
			"SearchCharacter",
			"GET",
			"/character/search",
			a.searchCharacter,
		},

		{
			"SearchCharacter",
			"GET",
			"/character/search/{search}",
			a.searchCharacter,
		},

		{
			"GetCharacter",
			"GET",
			"/character/{characterID}",
			a.getCharacter,
		},

		{
			"ListCharacter",
			"GET",
			"/character",
			a.listCharacter,
		},

		{
			"ListCharacter",
			"GET",
			"/character/{characterID}/inventory",
			a.listItemByCharacter,
		},

		{
			"ListCharacterByRanking",
			"GET",
			"/character/ranking",
			a.listCharacterByRanking,
		},

		{
			"ListCharacterByOnline",
			"GET",
			"/character/byonline",
			a.listCharacterByOnline,
		},

		{
			"ListCharacterByAccount",
			"GET",
			"/character/byaccount/{accountID}",
			a.listCharacterByAccount,
		},

		{
			"ListCharacterByRanking",
			"GET",
			"/ranking",
			a.listCharacterByRanking,
		},

		{
			"GetNpc",
			"GET",
			"/npc/{npcID}",
			a.getNpc,
		},

		{
			"ListNpc",
			"GET",
			"/npc",
			a.listNpc,
		},

		{
			"ListNpcByZone",
			"GET",
			"/npc/byzone",
			a.listNpcByZone,
		},

		{
			"GetNpcByZone",
			"GET",
			"/npc/byzone/{zoneID}",
			a.getNpcByZone,
		},

		{
			"ListNpcByFaction",
			"GET",
			"/npc/byfaction",
			a.listNpcByFaction,
		},

		{
			"GetNpcByFaction",
			"GET",
			"/npc/byfaction/{factionID}",
			a.getNpcByFaction,
		},

		{
			"ListTopic",
			"GET",
			"/forum/{forumID}",
			a.listTopic,
		},

		{
			"GetTopic",
			"GET",
			"/topic/{topicID}/details",
			a.getTopic,
		},

		{
			"ListActivity",
			"GET",
			"/task/{taskID}",
			a.listActivity,
		},

		{
			"GetActivity",
			"GET",
			"/task/{taskID}/activity/{activityID}",
			a.getActivity,
		},

		{
			"GetLootTable",
			"GET",
			"/loottable/{lootTableID}",
			a.getLootTable,
		},

		{
			"ListLootTable",
			"GET",
			"/loottable",
			a.listLootTable,
		},

		{
			"GetLootDropEntry",
			"GET",
			"/lootdrop/{lootDropID}/{itemID}",
			a.getLootDropEntry,
		},

		{
			"ListLootDropEntry",
			"GET",
			"/lootdrop/{lootDropID}",
			a.listLootDropEntry,
		},

		{
			"ListPost",
			"GET",
			"/topic/{topicID}",
			a.listPost,
		},

		{
			"GetPost",
			"GET",
			"/post/{postID}",
			a.getPost,
		},
		{
			"ListSpawn",
			"GET",
			"/spawn",
			a.listSpawn,
		},
		{
			"GetSpawn",
			"GET",
			"/spawn/{spawnID}",
			a.getSpawn,
		},
		{
			"ListForum",
			"GET",
			"/forum",
			a.listForum,
		},

		{
			"GetForum",
			"GET",
			"/forum/{forumID}/details",
			a.getForum,
		},

		{
			"CreateForum",
			"GET",
			"/forum/create",
			a.createForum,
		},

		{
			"GetTask",
			"GET",
			"/task/{taskID}/details",
			a.getTask,
		},

		{
			"ListTask",
			"GET",
			"/task",
			a.listTask,
		},

		{
			"GetZone",
			"GET",
			"/zone/{zoneID}",
			a.getZone,
		},

		{
			"ListZone",
			"GET",
			"/zone",
			a.listZone,
		},

		{
			"ListZoneByHotzone",
			"GET",
			"/zone/byhotzone",
			a.listZoneByHotzone,
		},

		{
			"GetItem",
			"GET",
			"/item/{itemID}",
			a.getItem,
		},

		{
			"SearchItem",
			"GET",
			"/item/search",
			a.searchItem,
		},

		{
			"SearchItemByAccount",
			"GET",
			"/item/search/byaccount",
			a.searchItemByAccount,
		},

		{
			"ListItemBySlot",
			"GET",
			"/item/byslot",
			a.listItemBySlot,
		},

		{
			"ListItemByZone",
			"GET",
			"/item/byzone",
			a.listItemByZone,
		},

		{
			"GetItemByZone",
			"GET",
			"/item/byzone/{zoneID}",
			a.getItemByZone,
		},

		{
			"GetItemBySlot",
			"GET",
			"/item/byslot/{slotID}",
			a.getItemBySlot,
		},

		{
			"ListItem",
			"GET",
			"/item",
			a.listItem,
		},
	}

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(rootPath + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	router.NotFoundHandler = http.HandlerFunc(a.notFound)
	return
}
