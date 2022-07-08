/*
  Copyright 2022 DSR Corporation, Denver, Colorado.
  https://www.dsr-corporation.com

  This file is part of ssi-medical-prescriptions-demo.

  ssi-medical-prescriptions-demo is free software: you can redistribute it
  and/or modify it under the terms of the GNU Affero General Public License
  as published by the Free Software Foundation, either version 3 of the License,
  or (at your option) any later version.

  ssi-medical-prescriptions-demo is distributed in the hope that it will be
  useful, but WITHOUT ANY WARRANTY; without even the implied warranty
  of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
  See the GNU Affero General Public License for more details.

  You should have received a copy of the GNU Affero General Public License along
  with ssi-medical-prescriptions-demo. If not, see <https://www.gnu.org/licenses/>.
*/

package leveldb

import (
	"encoding/json"
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	tmrand "github.com/tendermint/tendermint/libs/rand"
)

type LevelDB struct {
	path string
}

func NewLevelDB(path string) (*LevelDB, error) {
	return &LevelDB{path: path}, nil
}

func GenerateDBPath() string {
	return fmt.Sprintf("tmp/%s", tmrand.Str(5))
}

func (s *LevelDB) WriteAsJson(key string, value any) error {
	db, err := leveldb.OpenFile(s.path, nil)
	if err != nil {
		return fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error converting type Prescription to json bytes: %v", err)
	}

	if err = db.Put([]byte(key), data, nil); err != nil {
		return fmt.Errorf("error writing to database: %v", err)
	}

	return nil
}

func (s *LevelDB) ReadFromJson(key string, out any) error {
	db, err := leveldb.OpenFile(s.path, nil)
	if err != nil {
		return fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	data, err := db.Get([]byte(key), nil)
	if err != nil {
		return fmt.Errorf("error reading from database: %v", err)
	}

	if err = json.Unmarshal(data, &out); err != nil {
		return fmt.Errorf("error unmarshalling data: %v", err)
	}

	return nil
}

func (s *LevelDB) Has(key string) (bool, error) {
	var exist = false

	db, err := leveldb.OpenFile(s.path, nil)
	if err != nil {
		return exist, fmt.Errorf("error opening database file: %v", err)
	}
	defer db.Close()

	exist, err = db.Has([]byte(key), nil)
	if err != nil {
		return exist, fmt.Errorf("error reading from database: %v", err)
	}

	return exist, nil
}
