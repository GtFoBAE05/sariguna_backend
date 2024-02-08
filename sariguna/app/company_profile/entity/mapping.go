package entity

import "sariguna_backend/sariguna/app/company_profile/model"

func CompanyProfileCoreToCompanyProfileModel(data CompanyProfileCore) model.CompanyProfile {
	return model.CompanyProfile{
		Sejarah: data.Sejarah,
		Visi:    data.Visi,
		Misi:    data.Misi,
	}
}

func CompanyProfileModelToCompanyProfileCore(data model.CompanyProfile) CompanyProfileCore {
	return CompanyProfileCore{
		Sejarah: data.Sejarah,
		Visi:    data.Visi,
		Misi:    data.Misi,
	}
}
