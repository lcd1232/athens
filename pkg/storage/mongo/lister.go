package mongo

import (
	"strings"

	"github.com/globalsign/mgo/bson"
	"github.com/gomods/athens/pkg/storage"
)

func (s *MongoModuleStore) List(baseURL, module string) ([]string, error) {
	c := s.s.DB(s.d).C(s.c)
	result := make([]storage.Module, 0)
	err := c.Find(bson.M{"base_url": baseURL, "module": module}).All(&result)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = storage.ErrNotFound{Module: module, BasePath: baseURL}
		}
		return nil, err
	}

	versions := make([]string, len(result))
	for i, r := range result {
		versions[i] = r.Version
	}

	return versions, nil
}
