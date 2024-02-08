package entity

type CompanyProfileRepositoryInterface interface {
	GetCompanyProfile() (CompanyProfileCore, error)
	UpdateCompanyProfile(data CompanyProfileCore) error
}

type CompanyProfileServiceInterface interface {
	GetCompanyProfile() (CompanyProfileCore, error)
	UpdateCompanyProfile(data CompanyProfileCore) error
}
