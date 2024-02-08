package request

import "sariguna_backend/sariguna/app/company_profile/entity"

func CompanyProfileCreateToCompanyProfileCore(data CompanyProfileUpdate) entity.CompanyProfileCore {
	return entity.CompanyProfileCore{
		Sejarah: data.Sejarah,
		Visi:    data.Visi,
		Misi:    data.Misi,
	}
}
