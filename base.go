package datatunnelcommon

import (
	"context"
)

var (
	JobCancelList                = make(map[uint]context.CancelFunc, 0)
	TaskFullTableResultChannel   = make(chan TaskFullTableResult, 10000)
	TaskVerifyTableResultChannel = make(chan TaskVerifyTableResult, 10000)
	TaskObjectTableResultChannel = make(chan TaskObjectTableResult, 10000)
)

type MigrationTask[T interface{}] struct {
	JobCode      string
	JobId        uint
	Parallel     int
	Tables       []T
	SourceConfig DBConfig
	DestConfig   DBConfig
}

type TableBase struct {
	SchemeName string
	TableName  string
	Filter     string
}

type TableDetail struct {
	TableId     uint
	FromDBType  DBType
	Indexes     []TableIndex
	Columns     []TableColumn
	PrimaryKeys []string
	Comment     string
}
