package cases

import (
	"fmt"
	"sync"

	"github.com/xackery/xegony/storage"
)

var (
	readers      = make(map[string]storage.Reader)
	readLock     = sync.RWMutex{}
	writers      = make(map[string]storage.Writer)
	writeLock    = sync.RWMutex{}
	initializers = make(map[string]storage.Initializer)
	initLock     = sync.RWMutex{}
)

// InitializeAll readers, writers, and initializers
func InitializeAll(sr storage.Reader, sw storage.Writer, si storage.Initializer) (err error) {
	scopes := []string{
		"account",
		"character",
		/*"aa",
		"aaRank",*/
		/*"activity",
		"base",
		"bazaar",
		"character",
		"characterGraph",
		"class",
		"error",
		"faction",
		"fishing",
		"forage",
		"forum",
		"goal",
		"hacker",
		"item",
		"itemCategory",
		"itemSize",
		"itemSlot",
		"lootdrop",
		"lootdropEntry",
		"loottable",
		"loottableEntry",
		"mail",
		"merchant",
		"merchantEntry",
		"npc",
		"npcloot",
		"post",
		"race",
		"recipe",
		"recipeEntry",
		"rule",
		"sharedbank",
		"skill",
		"spawn",
		"spawnEntry",
		"spell",
		"task",
		"topic",
		"user",
		"variable",
		"zone",
		"zoneLevel",*/
	}

	for _, scope := range scopes {
		SetReader(scope, sr)
		SetWriter(scope, sw)
		SetInitializer(scope, si)
	}
	return
}

func getReader(scope string) (sr storage.Reader, err error) {
	readLock.RLock()
	sr, ok := readers[scope]
	if !ok {
		err = fmt.Errorf("Not initialized")
	}
	readLock.RUnlock()
	return
}

// SetReader sets a reader with scope
func SetReader(scope string, sr storage.Reader) {
	readLock.Lock()
	readers[scope] = sr
	readLock.Unlock()
	return
}

func getWriter(scope string) (sw storage.Writer, err error) {
	writeLock.RLock()
	sw, ok := writers[scope]
	if !ok {
		err = fmt.Errorf("Not initialized")
	}
	writeLock.RUnlock()
	return
}

// SetWriter sets a writer with scope
func SetWriter(scope string, sw storage.Writer) {
	writeLock.Lock()
	writers[scope] = sw
	writeLock.Unlock()
	return
}

func getInitializer(scope string) (si storage.Initializer, err error) {
	initLock.RLock()
	si, ok := initializers[scope]
	if !ok {
		err = fmt.Errorf("Not initialized")
	}
	initLock.RUnlock()
	return
}

// SetInitializer sets an initializer with scope
func SetInitializer(scope string, si storage.Initializer) {
	initLock.Lock()
	initializers[scope] = si
	initLock.Unlock()
	return
}
