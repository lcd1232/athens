package mongo

import (
	"io/ioutil"
)

func (m *MongoTests) TestGetSaveListRoundTrip() {
	r := m.Require()
	m.storage.Save(baseURL, module, version, mod, zip)
	listedVersions, err := m.storage.List(baseURL, module)
	r.NoError(err)
	r.Equal(1, len(listedVersions))
	retVersion := listedVersions[0]
	r.Equal(version, retVersion)
	gotten, err := m.storage.Get(baseURL, module, version)
	r.NoError(err)
	defer gotten.Zip.Close()
	r.Equal(version, gotten.RevInfo.Version)
	r.Equal(version, gotten.RevInfo.Name)
	r.Equal(version, gotten.RevInfo.Short)
	// TODO: test the time
	r.Equal(gotten.Mod, mod)
	zipContent, err := ioutil.ReadAll(gotten.Zip)
	r.NoError(err)
	r.Equal(zipContent, zip)
}

func (m *MongoTests) TestNewMongoStorage() {
	r := m.Require()
	url := "mongodb://127.0.0.1:27017"
	getterSaver := NewMongoStorage(url)
	getterSaver.Connect()

	r.NotNil(getterSaver.c)
	r.NotNil(getterSaver.d)
	r.NotNil(getterSaver.s)
	r.Equal(getterSaver.url, url)
}
