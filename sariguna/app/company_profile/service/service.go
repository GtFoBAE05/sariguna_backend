package service

import "sariguna_backend/sariguna/app/company_profile/entity"

type CompanyProfileService struct {
	CompanyProfileRepository entity.CompanyProfileRepositoryInterface
}

func NewCompanyProfileService(CompanyProfileRepository entity.CompanyProfileRepositoryInterface) entity.CompanyProfileServiceInterface {
	return &CompanyProfileService{
		CompanyProfileRepository: CompanyProfileRepository,
	}
}

// GetCompanyProfile implements entity.CompanyProfileServiceInterface.
func (cps *CompanyProfileService) GetCompanyProfile() (entity.CompanyProfileCore, error) {
	result, err := cps.CompanyProfileRepository.GetCompanyProfile()

	if err != nil {
		return entity.CompanyProfileCore{}, err
	}

	return result, nil
}

// UpdateCompanyProfile implements entity.CompanyProfileServiceInterface.
func (cps *CompanyProfileService) UpdateCompanyProfile(data entity.CompanyProfileCore) error {
	err := cps.CompanyProfileRepository.UpdateCompanyProfile(data)

	if err != nil {
		return err
	}

	return nil
}
