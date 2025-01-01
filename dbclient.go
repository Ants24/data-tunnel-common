package datatunnelcommon

import (
	"context"
	"database/sql"
)

type DBClient interface {
	GetTableNames(logger Logger, schema string) ([]string, error)
	GetSchemaNames(logger Logger) ([]string, error)
	ConventTableAndColumnName(name string, fromDBType DBType) string
	GetTableColumnNames(ctx context.Context, schema string, tableName string, columnNames []string, fromDBType DBType) ([]string, []string, error)
	ReadData(ctx context.Context, logger Logger, taskConfig TaskFullTable, columnNames []string, columnTypes []string, channel chan []sql.NullString) (uint64, error)
	WriteData(ctx context.Context, logger Logger, taskConfig TaskFullTable, columnNames []string, columnTypes []string, channel chan []sql.NullString) error
	GenerateSubTasks(ctx context.Context, logger Logger, taskConfig TaskFullTable) ([]TaskFullTable, error)
	GetTablePrimaryKeys(ctx context.Context, logger Logger, schemaName string, tableName string) ([]string, error)
	GetTablePartitionNames(ctx context.Context, logger Logger, schemaName string, tableName string) ([]string, error)
	SelectTableCount(ctx context.Context, logger Logger, tableBase TableBase) (uint64, error)
	TruncateTable(ctx context.Context, logger Logger, schemaName string, tableName string) error
	GetTableDetail(ctx context.Context, logger Logger, taskConfig TaskObjectTable) (TableDetail, error)
	GenerateTableDetail(ctx context.Context, logger Logger, tableDetail TableDetail) (TableDetail, error)
	GenerateTableSql(ctx context.Context, logger Logger, taskConfig TaskObjectTable, tableDetail TableDetail) ([]string, []string, error)
}
