package interfaces

const (
	TagName = "database"
)

type DBConnection interface {
	// Open()
	// Close()
	SaveData(tableName string, data interface{}) error
	FindData(tableName string, data interface{}) (interface{}, error)
}
