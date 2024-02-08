package response

import "sariguna_backend/sariguna/app/company_profile/entity"

func CompanyProfileCoreToCompanyProfileResponse(data entity.CompanyProfileCore) CompanyProfileResponse {
	return CompanyProfileResponse{
		Sejarah: data.Sejarah,
		Visi:    data.Visi,
		Misi:    data.Misi,
	}
}
