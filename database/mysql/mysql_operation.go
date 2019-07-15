package mysql

type DBOperation interface {
	Insert(sql string, args ...string) error
	Delete(sql string, args ...string) error
	Update(sql string, args ...string) error
	Query(sql string, args ...string) ([]string, error)
}