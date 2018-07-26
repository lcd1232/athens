package rdbms

import "github.com/gobuffalo/pop"

type Registry struct {
	conn           *pop.Connection
	connectionName string // settings name from database.yml
}

func (r *Registry) LookupPointer(deploymentID string) (string, error) {
	panic("implement me")
}

func (r *Registry) SetPointer(deploymentID, pointer string) error {
	panic("implement me")
}

func (r *Registry) Connect() error {
	c, err := pop.Connect(r.connectionName)
	if err != nil {
		return err
	}
	r.conn = c
	return nil
}
