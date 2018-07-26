package rdbms

import (
	"github.com/gobuffalo/pop"
	"github.com/gomods/athens/pkg/eventlog"
)

type Log struct {
	conn           *pop.Connection
	connectionName string // settings name from database.yml
}

// NewLog creates event log from backing rdbms database
func NewLog(connectionName string) (*Log, error) {
	m := &Log{
		connectionName: connectionName,
	}
	return m, m.Connect()
}

func (m *Log) Connect() error {
	c, err := pop.Connect(m.connectionName)
	if err != nil {
		return err
	}
	m.conn = c
	return nil
}

func (m *Log) Read() ([]eventlog.Event, error) {
	panic("implement me")
}

func (m *Log) ReadFrom(id string) ([]eventlog.Event, error) {
	panic("implement me")
}

func (m *Log) ReadSingle(module, version string) (eventlog.Event, error) {
	panic("implement me")
}

func (m *Log) Clear(id string) error {
	panic("implement me")
}

func (m *Log) Append(event eventlog.Event) (string, error) {
	panic("implement me")
}