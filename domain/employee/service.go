package employee

import "context"

type Service interface {
	CreateEmployee(ctx context.Context, createParams CreateParamsWithPasswordStrings) (*Employee, error)
}
