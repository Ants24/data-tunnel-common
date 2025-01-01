package datatunnelcommon

import "time"

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
