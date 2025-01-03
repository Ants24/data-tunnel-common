package datatunnelcommon

import (
	"context"
	"time"
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

type TaskFullTable struct {
	TableId      uint
	JobId        uint
	SourceSchema string
	SourceTable  string
	DestSchema   string
	DestTable    string
	Filter       string
	Config       TableFullConfig
	// 一下做切片的配置
	SplitId       int
	SplitColumn   string
	PartitionName string
	StartValue    interface{}
	EndValue      interface{}
}

type TaskFullTableResult struct {
	JobId       uint
	TableId     uint
	SplitId     int
	HandlerTime time.Time //处理时间
	Status      JobStatus
	TotalNum    uint64
	SuccessNum  uint64
	FailedNum   uint64
	TakeTime    float64 //耗时
	Err         error
}

type TaskObjectTable struct {
	TableId      uint
	JobId        uint
	SourceSchema string
	SourceTable  string
	DestSchema   string
	DestTable    string
	Config       TableObjectConfig
}

type TaskObjectTableResult struct {
	JobId            uint
	TableId          uint
	TableComment     string
	FromDBType       DBType
	TableIndexes     []TableIndex
	TableColumns     []TableColumn
	TablePrimaryKeys []string
	SuccessSqls      []string
	FailedSqls       []string
	Error            error
	Status           JobStatus
}

type TaskVerifyTable struct {
	TableId      uint
	JobId        uint
	SourceSchema string
	SourceTable  string
	DestSchema   string
	DestTable    string
	Filter       string
	Config       TableVerifyConfig
}

type TaskVerifyTableResult struct {
	JobId       uint
	TableId     uint
	Status      JobVerifyStatus
	SourceCount uint64
	DestCount   uint64
	Error       error
}
