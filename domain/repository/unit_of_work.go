package repository

// UnitOfWork defines the interface for a unit of work.
type UnitOfWork interface {
	// Execute runs a function within a transaction.
	// If the function returns an error, the transaction is rolled back.
	// Otherwise, the transaction is committed.
	Execute(func(store UnitOfWorkStore) error) error
}

// UnitOfWorkStore defines the interface for a store that can be used within a unit of work.
// It provides access to all repositories.
type UnitOfWorkStore interface {
	Users() UserRepository
	Products() ProductRepository
	Orders() OrderRepository
}
