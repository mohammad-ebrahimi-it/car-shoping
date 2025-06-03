package migrations

import (
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/constans"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/models"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func UP_1() {
	database := db.GetDb()

	createTables(database)

	createDefaultInformation(database)
}

func createTables(database *gorm.DB) {

	var tables []interface{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNameTable(database, country, tables)
	tables = addNameTable(database, role, tables)
	tables = addNameTable(database, userRole, tables)
	tables = addNameTable(database, user, tables)
	tables = addNameTable(database, city, tables)

	tables = addNameTable(database, models.File{}, tables)
	tables = addNameTable(database, models.PersianYear{}, tables)

	tables = addNameTable(database, models.PropertyCategory{}, tables)
	tables = addNameTable(database, models.Property{}, tables)

	tables = addNameTable(database, models.Company{}, tables)
	tables = addNameTable(database, models.Gearbox{}, tables)
	tables = addNameTable(database, models.Color{}, tables)
	tables = addNameTable(database, models.CarType{}, tables)

	tables = addNameTable(database, models.CarModel{}, tables)
	tables = addNameTable(database, models.CarModelColor{}, tables)
	tables = addNameTable(database, models.CarModelYear{}, tables)
	tables = addNameTable(database, models.CarModelImage{}, tables)
	tables = addNameTable(database, models.CarModelPriceHistory{}, tables)
	tables = addNameTable(database, models.CarModelProperty{}, tables)
	tables = addNameTable(database, models.CarModelComment{}, tables)

	err := database.Migrator().CreateTable(tables...)

	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}

	logger.Info(logging.Postgres, logging.Migration, "migration succeeded", nil)

}

func addNameTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}

	return tables
}

func createDefaultInformation(database *gorm.DB) {

	adminRole := models.Role{Name: constans.AdminRoleName}
	createRoleIfExists(database, &adminRole)

	defaultRole := models.Role{Name: constans.DefaultRoleName}
	createRoleIfExists(database, &defaultRole)

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user := models.User{
		Username:  constans.DefaultUserName,
		FirstName: "test",
		LastName:  "test",
		Mobile:    "09900723195",
		Email:     "m@gmail.com",
		Password:  string(password),
	}
	createAdminUserIfNotExists(database, &user, adminRole.Id)

}

func createRoleIfExists(database *gorm.DB, role *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", role.Name).
		First(&exists)

	if exists == 0 {
		database.Create(role)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, user *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", user.Username).
		First(&exists)

	if exists == 0 {
		database.Create(user)
		ur := models.UserRole{UserId: user.Id, RoleId: roleId}

		database.Create(&ur)
	}
}
