package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mohammad-ebrahimi-it/car-shoping/api/dto"
	"github.com/mohammad-ebrahimi-it/car-shoping/common"
	"github.com/mohammad-ebrahimi-it/car-shoping/config"
	"github.com/mohammad-ebrahimi-it/car-shoping/constans"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/db"
	"github.com/mohammad-ebrahimi-it/car-shoping/data/models"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/logging"
	"github.com/mohammad-ebrahimi-it/car-shoping/pkg/service_errors"
	"gorm.io/gorm"
	"math"
	"reflect"
	"strings"
	"time"
)

type preload struct {
	string
}

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger   logging.Logger
	Preloads []preload
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDb(),
		Logger:   logging.NewLogger(cfg),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, _ := common.TypeConverter[T](req)

	tx := s.Database.WithContext(ctx).Begin()

	err := tx.Create(&model).Error

	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}

	tx.Commit()

	bm, _ := common.TypeConverter[models.BaseModel](model)
	return s.GetById(ctx, bm.Id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {

	updateMap, _ := common.TypeConverter[map[string]interface{}](req)

	snakeMap := map[string]interface{}{}

	for k, v := range *updateMap {
		snakeMap[common.ToSnakeCase(k)] = v
	}

	snakeMap["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constans.UserIdKey).(float64)), Valid: true}
	snakeMap["modified_at"] = &sql.NullTime{Time: time.Now().UTC(), Valid: true}

	model := new(T)

	tx := s.Database.WithContext(ctx).Begin()

	err := tx.Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(snakeMap).
		Error

	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}

	tx.Commit()
	return s.GetById(ctx, id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.Database.WithContext(ctx).Begin()

	model := new(T)

	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constans.UserIdKey).(float64)), Valid: true},
		"deleted_at": &sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}

	if ctx.Value(constans.UserIdKey) == nil {
		return &service_errors.ServiceError{
			EndUserMessage: service_errors.PermissionDenied,
		}
	}

	cnt := tx.Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(deleteMap).
		RowsAffected

	if cnt == 0 {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Delete, service_errors.RecordNotFound, nil)

		return &service_errors.ServiceError{
			EndUserMessage: service_errors.RecordNotFound,
		}
	}

	tx.Commit()

	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	db := Preload(s.Database, s.Preloads)
	err := db.WithContext(ctx).
		Where("id = ? and deleted_by is null", id).
		First(model).
		Error

	if err != nil {
		return nil, err
	}

	return common.TypeConverter[Tr](model)
}

func GetQuery[T any](filter *dto.DynamicFilter) string {
	t := new(T)

	typeT := reflect.TypeOf(*t)

	query := make([]string, 0)

	query = append(query, "deleted_by is null")

	if filter.Filter != nil {
		for name, filter := range filter.Filter {
			fld, ok := typeT.FieldByName(name)

			if ok {
				fld.Name = common.ToSnakeCase(fld.Name)

				switch filter.Type {
				case "contains":
					query = append(query, fmt.Sprintf("%s ILike '%%%s%%'", fld.Name, filter.From))
				case "notContains":
					query = append(query, fmt.Sprintf("%s not ILike '%%%s%%'", fld.Name, filter.From))
				case "startWith":
					query = append(query, fmt.Sprintf("%s  ILike '%s%%'", fld.Name, filter.From))
				case "endWith":
					query = append(query, fmt.Sprintf("%s  ILike '%%%s'", fld.Name, filter.From))
				case "equal":
					query = append(query, fmt.Sprintf("%s = '%s'", fld.Name, filter.From))
				case "notEqual":
					query = append(query, fmt.Sprintf("%s != '%s'", fld.Name, filter.From))
				case "lassThan":
					query = append(query, fmt.Sprintf("%s < '%s'", fld.Name, filter.From))
				case "lassThanOrEqual":
					query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.From))
				case "greaterThan":
					query = append(query, fmt.Sprintf("%s > '%s'", fld.Name, filter.From))
				case "greaterThanOrEqual":
					query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
				case "inRange":
					if fld.Type.Kind() == reflect.String {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))
					} else {
						query = append(query, fmt.Sprintf("%s >= '%s'", fld.Name, filter.From))
						query = append(query, fmt.Sprintf("%s <= '%s'", fld.Name, filter.To))

					}
				}
			}

		}
	}

	return strings.Join(query, " AND ")

}

func GetSort[T any](filter *dto.DynamicFilter) string {
	t := new(T)

	typeT := reflect.TypeOf(*t)

	sort := make([]string, 0)

	if filter.Sort != nil {
		for _, tp := range *filter.Sort {

			fld, ok := typeT.FieldByName(tp.ColId)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}

		}
	}

	return strings.Join(sort, ", ")
}

func Preload(db *gorm.DB, preloads []preload) *gorm.DB {
	for _, preload := range preloads {
		db = db.Preload(preload.string)
	}

	return db
}

func NewPagedList[T any](items *[]T, count int64, pageNumber int, pageSize int64) *dto.PageList[T] {
	pl := &dto.PageList[T]{
		PageNumber: pageNumber,
		TotalRows:  count,
		Items:      items,
	}

	pl.TotalPages = int(math.Ceil(float64(pl.TotalRows) / float64(pageSize)))

	pl.HasNextPage = pl.PageNumber < pl.TotalPages
	pl.HasPreviousPage = pl.PageNumber > 1

	return pl
}

func Paginate[T, Tr any](pagination *dto.PaginationInputWithFilter, preloads []preload, db *gorm.DB) (*dto.PageList[Tr], error) {
	model := new(T)

	var items *[]T
	var rItems *[]Tr

	db = Preload(db, preloads)

	query := GetQuery[T](&pagination.DynamicFilter)
	sort := GetSort[T](&pagination.DynamicFilter)

	var totalRow int64 = 0

	db.
		Model(model).
		Where("deleted_by is null").
		Where(query).
		Count(&totalRow)

	err := db.
		Where(query).
		Offset(pagination.GetOffset()).
		Limit(pagination.GetPageSize()).
		Order(sort).
		Find(&items).
		Error

	if err != nil {
		return nil, err
	}

	rItems, err = common.TypeConverter[[]Tr](items)

	if err != nil {
		return nil, err
	}

	return NewPagedList(rItems, totalRow, pagination.PageNumber, int64(pagination.PageSize)), nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[Tr], error) {
	return Paginate[T, Tr](req, s.Preloads, s.Database)
}
