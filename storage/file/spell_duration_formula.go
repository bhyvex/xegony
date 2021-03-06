package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

var (
	spellDurationFormulaLock = sync.RWMutex{}
)

//GetSpellDurationFormula will grab data from file. This is an expensive call, it is recommended to use List instead and minimize reads.
func (s *Storage) GetSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	spellDurationFormulasDatabase, err := s.readSpellDurationFormulaFile()
	if err != nil {
		return
	}
	for _, tmpSpellDurationFormula := range spellDurationFormulasDatabase {
		if tmpSpellDurationFormula.ID == spellDurationFormula.ID {
			*spellDurationFormula = *tmpSpellDurationFormula
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//CreateSpellDurationFormula will grab data from storage
func (s *Storage) CreateSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	spellDurationFormulasDatabase, err := s.readSpellDurationFormulaFile()
	if err != nil {
		return
	}
	for _, tmpSpellDurationFormula := range spellDurationFormulasDatabase {
		if tmpSpellDurationFormula.ID == spellDurationFormula.ID {
			err = fmt.Errorf("spellDurationFormula already exists")
			return
		}
	}
	spellDurationFormulasDatabase = append(spellDurationFormulasDatabase, spellDurationFormula)
	err = s.writeSpellDurationFormulaFile(spellDurationFormulasDatabase)
	if err != nil {
		return
	}
	return
}

//ListSpellDurationFormula will grab data from storage
func (s *Storage) ListSpellDurationFormula(page *model.Page) (spellDurationFormulas []*model.SpellDurationFormula, err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	spellDurationFormulasDatabase, err := s.readSpellDurationFormulaFile()
	if err != nil {
		return
	}

	spellDurationFormulas = make([]*model.SpellDurationFormula, len(spellDurationFormulasDatabase))

	spellDurationFormulas = spellDurationFormulasDatabase

	if page.OrderBy == "" {
		page.OrderBy = "name"
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellDurationFormulas, func(i, j int) bool {
			return spellDurationFormulas[i].Name < spellDurationFormulas[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	/*	if page.IsDescending > 0 {
			sort.Sort(sort.Reverse(spellDurationFormulas))
		}
	*/
	return
}

//ListSpellDurationFormulaTotalCount will grab data from storage
func (s *Storage) ListSpellDurationFormulaTotalCount() (count int64, err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	spellDurationFormulasDatabase, err := s.readSpellDurationFormulaFile()
	if err != nil {
		return
	}
	count = int64(len(spellDurationFormulasDatabase))
	return
}

//ListSpellDurationFormulaBySearch will grab data from storage
func (s *Storage) ListSpellDurationFormulaBySearch(page *model.Page, spellDurationFormula *model.SpellDurationFormula) (spellDurationFormulas []*model.SpellDurationFormula, err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	spellDurationFormulasDatabase, err := s.readSpellDurationFormulaFile()
	if err != nil {
		return
	}
	if len(spellDurationFormula.Name) > 0 {
		for i := range spellDurationFormulasDatabase {
			if strings.Contains(spellDurationFormulasDatabase[i].Name, spellDurationFormula.Name) {
				spellDurationFormulas = append(spellDurationFormulas, spellDurationFormulasDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}

	switch page.OrderBy {
	case "name":
		sort.Slice(spellDurationFormulas, func(i, j int) bool {
			return spellDurationFormulas[i].Name < spellDurationFormulas[j].Name
		})
	default:
		err = fmt.Errorf("Unsupported sort name")
		return
	}

	//if page.IsDescending > 0 {
	//	sort.Sort(sort.Reverse(spellDurationFormulas))
	//}
	return
}

//ListSpellDurationFormulaBySearchTotalCount will grab data from storage
func (s *Storage) ListSpellDurationFormulaBySearchTotalCount(spellDurationFormula *model.SpellDurationFormula) (count int64, err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	spellDurationFormulasDatabase, err := s.readSpellDurationFormulaFile()
	if err != nil {
		return
	}

	spellDurationFormulas := []*model.SpellDurationFormula{}
	if len(spellDurationFormula.Name) > 0 {
		for i := range spellDurationFormulasDatabase {
			if strings.Contains(spellDurationFormulasDatabase[i].Name, spellDurationFormula.Name) {
				spellDurationFormulas = append(spellDurationFormulas, spellDurationFormulasDatabase[i])
			}
		}
	} else {
		err = fmt.Errorf("Unsupported search (Need shortname)")
		return
	}
	count = int64(len(spellDurationFormulas))
	return
}

//EditSpellDurationFormula will grab data from storage
func (s *Storage) EditSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	spellDurationFormulasDatabase, err := s.readSpellDurationFormulaFile()
	if err != nil {
		return
	}

	for i := range spellDurationFormulasDatabase {
		if spellDurationFormulasDatabase[i].ID == spellDurationFormula.ID {
			*spellDurationFormulasDatabase[i] = *spellDurationFormula
			err = s.writeSpellDurationFormulaFile(spellDurationFormulasDatabase)
			if err != nil {
				return
			}
			return
		}
	}
	err = &model.ErrNoContent{}
	return
}

//DeleteSpellDurationFormula will grab data from storage
func (s *Storage) DeleteSpellDurationFormula(spellDurationFormula *model.SpellDurationFormula) (err error) {
	spellDurationFormulaLock.Lock()
	defer spellDurationFormulaLock.Unlock()
	spellDurationFormulasDatabase, err := s.readSpellDurationFormulaFile()
	if err != nil {
		return
	}
	indexToDelete := 0
	for i := range spellDurationFormulasDatabase {
		if spellDurationFormulasDatabase[i].ID == spellDurationFormula.ID {
			indexToDelete = i
			break
		}
	}
	if indexToDelete < 1 {
		err = &model.ErrNoContent{}
		return
	}

	spellDurationFormulasDatabase[len(spellDurationFormulasDatabase)-1], spellDurationFormulasDatabase[indexToDelete] = spellDurationFormulasDatabase[indexToDelete], spellDurationFormulasDatabase[len(spellDurationFormulasDatabase)-1]
	spellDurationFormulasDatabase = spellDurationFormulasDatabase[:len(spellDurationFormulasDatabase)-1]
	err = s.writeSpellDurationFormulaFile(spellDurationFormulasDatabase)
	if err != nil {
		return
	}
	return
}

func (s *Storage) readSpellDurationFormulaFile() (spellDurationFormulas []*model.SpellDurationFormula, err error) {
	yf, err := ioutil.ReadFile(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			spellDurationFormulas = loadSpellDurationFormulaDefault()
			err = os.MkdirAll(s.directory, 0744)
			if err != nil {
				err = errors.Wrapf(err, "failed to make directory %s", s.path)
				return
			}
			err = s.writeSpellDurationFormulaFile(spellDurationFormulas)
			if err != nil {
				err = errors.Wrapf(err, "failed to write default spellDurationFormula data to file %s", s.path)
				return
			}
			return
		}
		err = errors.Wrapf(err, "failed to read file %s", s.path)
		return
	}
	err = yaml.Unmarshal(yf, &spellDurationFormulas)
	if err != nil {
		err = errors.Wrapf(err, "failed to unmarshal file %s", s.path)
		return
	}

	return
}

func (s *Storage) writeSpellDurationFormulaFile(spellDurationFormulas []*model.SpellDurationFormula) (err error) {

	bData, err := yaml.Marshal(spellDurationFormulas)
	if err != nil {
		err = errors.Wrap(err, "failed to marshal spellDurationFormulas")
		return
	}
	err = ioutil.WriteFile(s.path, bData, 0744)
	if err != nil {
		err = errors.Wrapf(err, "failed to write file %s", s.path)
		return
	}
	return
}
