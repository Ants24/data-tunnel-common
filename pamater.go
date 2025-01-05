package datatunnelcommon

// TableObjectConfig 表对象配置
type TableObjectConfig struct {
	Index             bool `json:"index"`
	PrimaryKey        bool `json:"primaryKey"`
	ForeignKey        bool `json:"foreignKey"`
	DropIfExists      bool `json:"dropIfExists"`
	Comment           bool `json:"comment"`
	CreateIfNotExists bool `json:"createIfNotExists"`
}

type GlobalObjectConfig struct {
	Parallel    int  `json:"parallel"`
	IsErrorNext bool `json:"isErrorNext"` // 错误是否继续
	TableObjectConfig
}

type TableFullConfig struct {
	Parallel        int  `json:"parallel"`
	Split           bool `json:"split"`
	SplitBatchSize  int  `json:"splitBatchSize"`
	CommitBatchTime int  `json:"commitBatchTime"`
	CommitBatchSize int  `json:"commitBatchSize"`
	Partition       bool `json:"partition"`
	Truncate        bool `json:"truncate"`
	ChannelSize     int  `json:"channelSize"`
	IsUseDestColumn bool `json:"isUseDestColumn"` // 是否以目标数据源字段为准,来同步数据
}

type GlobalFullConfig struct {
	TableFullConfig
	IsErrorNext bool `json:"isErrorNext"` // 错误是否继续
}

type TableVerifyConfig struct {
	Parallel        int  `json:"parallel"`
	Split           bool `json:"split"`
	SplitBatchSize  int  `json:"splitBatchSize"`
	Partition       bool `json:"partition"`
	IsUseDestColumn bool `json:"isUseDestColumn"` // 是否以目标数据源字段为准,来校验数据,当IsData为true时,该值有效
	IsCount         bool `json:"isCount"`         // 是否总数校验
	IsData          bool `json:"isData"`          // 是否数据校验
}

type GlobalVerifyConfig struct {
	TableVerifyConfig
	IsErrorNext bool `json:"isErrorNext"` // 错误是否继续
}
