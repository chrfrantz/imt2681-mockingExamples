package mockExample

type DBService interface {

	Init() error

	Insert(string) error

	Get(string) error

	Count() error

	DeleteAll() error

}