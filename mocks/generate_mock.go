package mocks

//go:generate mockgen -source=../internal/repository.go -destination=mock_repository.go -package=mocks
//go:generate mockgen -source=../internal/usecase.go -destination=mock_usecase.go -package=mocks

//go:generate mockgen -source=../pkg/authentication/authentication.go -destination=mock_pkg_authentication.go -package=mocks
