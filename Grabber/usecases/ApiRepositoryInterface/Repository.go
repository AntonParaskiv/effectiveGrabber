package ApiRepositoryInterface

type Repository interface {
	GetOne() (err error)
	GetTwo() (err error)
}
