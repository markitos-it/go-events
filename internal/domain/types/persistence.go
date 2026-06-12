package types

type GoldenRepository interface {
	Create(golden *Golden) error
	Delete(id *GoldenId) error
	One(id *GoldenId) (*Golden, error)
	Update(golden *Golden) error
	All() ([]*Golden, error)
	SearchAndPaginate(searchTerm string, pageNumber int, pageSize int) ([]*Golden, error)
}
