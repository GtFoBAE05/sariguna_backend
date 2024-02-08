package repository

import (
	"sariguna_backend/sariguna/app/company_profile/entity"
	"sariguna_backend/sariguna/app/company_profile/model"

	"github.com/jmoiron/sqlx"
)

type CompanyProfileRepository struct {
	db *sqlx.DB
}

func NewCompanyProfileRepository(db *sqlx.DB) entity.CompanyProfileRepositoryInterface {
	return &CompanyProfileRepository{
		db: db,
	}
}

// GetCompanyProfile implements entity.CompanyProfileRepositoryInterface.
func (cpr *CompanyProfileRepository) GetCompanyProfile() (entity.CompanyProfileCore, error) {
	companyProfile := model.CompanyProfile{}

	query := `SELECT sejarah, visi, misi from company_profile where id = 1`

	err := cpr.db.Get(&companyProfile, query)
	if err != nil {
		return entity.CompanyProfileCore{}, err
	}

	result := entity.CompanyProfileModelToCompanyProfileCore(companyProfile)

	return result, nil
}

// UpdateCompanyProfile implements entity.CompanyProfileRepositoryInterface.
func (cpr *CompanyProfileRepository) UpdateCompanyProfile(data entity.CompanyProfileCore) error {
	request := entity.CompanyProfileCoreToCompanyProfileModel(data)

	query := `UPDATE company_profile
	SET sejarah = $1,
	visi = $2,
	misi = $3
	WHERE id = 1`

	_, err := cpr.db.Exec(query, request.Sejarah, request.Visi, request.Misi)

	if err != nil {
		return err
	}

	return nil
}
