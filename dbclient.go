package datatunnelcommon

import (
	"context"
	"database/sql"
)

type DBClient interface {
	GetTableNames(logger Logger, schema string) ([]string, error)
	GetSchemaNames(logger Logger, databaseName string) ([]string, error)
	ConventTableAndColumnName(name string, fromDBType DBType) string
	GetTableColumnNames(ctx context.Context, tableInfo TableBase, columnNames []string, fromDBType DBType) ([]string, []string, error)
	ReadData(ctx context.Context, logger Logger, tableInfo TableBase, taskConfig TaskFullTable, columnNames []string, columnTypes []string, channel chan []sql.NullString, fromDBType DBType) (uint64, error)
	MD5(ctx context.Context, logger Logger, tableInfo TableBase, taskConfig TaskFullTable, columnNames []string, columnTypes []string, fromDBType DBType) (string, error)
	WriteData(ctx context.Context, logger Logger, tableInfo TableBase, taskConfig TaskFullTable, columnNames []string, columnTypes []string, channel chan []sql.NullString) error
	GenerateSubTasks(ctx context.Context, logger Logger, tableInfo TableBase, taskConfig TaskFullTable) ([]TaskFullTable, error)
	GetTablePrimaryKeys(ctx context.Context, logger Logger, tableInfo TableBase) ([]string, error)
	GetTablePartitionNames(ctx context.Context, logger Logger, tableInfo TableBase) ([]string, error)
	SelectTableCount(ctx context.Context, logger Logger, tableBase TableBase) (uint64, error)
	TruncateTable(ctx context.Context, logger Logger, tableBase TableBase) error
	GetTableDetail(ctx context.Context, logger Logger, tableInfo TableBase, taskConfig TaskObjectTable) (TableDetail, error)
	GenerateTableDetail(ctx context.Context, logger Logger, tableDetail TableDetail) (TableDetail, error)
	GenerateTableSql(ctx context.Context, logger Logger, tableInfo TableBase, taskConfig TaskObjectTable, tableDetail TableDetail) ([]string, []string, error)
}
