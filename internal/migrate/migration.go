package migrate

type Migration struct {
	Version int64
	Name    string
	Path    string
}

func NewMigration(version int64, name, path string) Migration {
	return Migration{
		Version: version,
		Name:    name,
		Path:    path,
	}
}
