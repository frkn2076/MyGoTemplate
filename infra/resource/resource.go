package resource

import(
	"app/MyGoTemplate/infra/db"
	"app/MyGoTemplate/infra/db/localization"
	"app/MyGoTemplate/infra/cache"
	"app/MyGoTemplate/infra/logger"
)

//Cache feeder
func init(){
	localizationRepo := new(localization.Repository) 
	if localizations, err := localizationRepo.GetAll(db.GormDB); err != nil {
		logger.ErrorLog("An error occured while loading localizations to cache: ", err.Error())
	} else {
		for _, localization := range localizations {
			cache.Set(localization.Resource + localization.Language, localization.Message, 0)
		}
	}
}