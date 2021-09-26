package reqstash

type Storage interface {
	Put(req Request) error
	ListAll() ([]Request, error)
}

type Purgable interface {
	Purge() (int, error)
}
