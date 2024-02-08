package route

import (
	"sariguna_backend/sariguna/app/company_profile/handler"
	"sariguna_backend/sariguna/app/company_profile/repository"
	"sariguna_backend/sariguna/app/company_profile/service"
	"sariguna_backend/sariguna/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RouteCompanyProfile(r *gin.RouterGroup, db *sqlx.DB) {
	companyProfileRepository := repository.NewCompanyProfileRepository(db)
	companyProfileService := service.NewCompanyProfileService(companyProfileRepository)
	companyProfileHandler := handler.NewCompanyProfileHandler(companyProfileService)

	cp := r.Group("/profile")

	cp.GET("", companyProfileHandler.GetCompanyProfile)

	cp.PUT("/update", middleware.Auth, companyProfileHandler.UpdateCompanyProfile)

}
