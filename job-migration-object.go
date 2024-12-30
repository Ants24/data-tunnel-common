package datatunnelcommon

type TableIndex struct {
	IndexName    string   `json:"indexName"`
	IndexColumns []string `json:"indexColumns"`
	IsUnique     bool     `json:"isUnique"`
}

type TableColumn struct {
	ColumnType         ColumnType `json:"-"`
	SourceColumnName   string     `json:"sourceColumnName"`
	SourceColumnType   string     `json:"sourceColumnType"`
	SourceColumnLength string     `json:"sourceColumnLength"`
	DestColumnName     string     `json:"destColumnName"`
	DestColumnType     string     `json:"destColumnType"`
	DestColumnLength   string     `json:"destColumnLength"`
	ColumnDefault      string     `json:"columnDefault"`
	ColumnComment      string     `json:"columnComment"`
	IsNullable         bool       `json:"isNullable"`
	IsAutoIncrement    bool       `json:"isAutoIncrement"`
}
