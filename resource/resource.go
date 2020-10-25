package resource

import(
	"app/MyGoTemplate/db"
	"app/MyGoTemplate/db/repo"
	"app/MyGoTemplate/cache"
	"app/MyGoTemplate/logger"
)

//Cache feeder
func init(){
	if localizations, err := repo.Localization.GetAll(db.GormDB); err != nil {
		logger.ErrorLog("An error occured while loading localizations to cache: ", err.Error())
	} else {
		for _, localization := range localizations {
			cache.Set(localization.Resource + localization.Language, localization.Message, -1)
		}
	}
}