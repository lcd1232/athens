package rdbms

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RDBMSTests struct {
	suite.Suite
	log *Log
}

func TestRDBMS(t *testing.T) {
	suite.Run(t, new(RDBMSTests))
}

func (rd *RDBMSTests) SetupTest() {
	log, err := NewLog("test_postgres")
	if err != nil {
		panic(err)
	}

	log.conn
	rd.log = log
}