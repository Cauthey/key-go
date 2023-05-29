package property

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"
	"key-go/internal/app/dao/util"
	"key-go/internal/app/schema"
	"key-go/pkg/errors"
	"key-go/pkg/util/snowflake"
)

var PropertySet = wire.NewSet(wire.Struct(new(PropertyRepo), "*"))

type PropertyRepo struct {
	DB *gorm.DB
}

func (a *PropertyRepo) getQueryOption(ctx context.Context, opts ...schema.PropertyQueryOptions) schema.PropertyQueryOptions {
	var opt schema.PropertyQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *PropertyRepo) Query(ctx context.Context, params schema.PropertyQueryParam, opts ...schema.PropertyQueryOptions) (*schema.PropertyQueryResult, error) {
	opt := a.getQueryOption(ctx, opts...)

	db := GetPropertyDB(ctx, a.DB)
	if v := params.Name; v != "" {
		db = db.Where("name=?", v)
	}
	if v := params.QueryValue; v != "" {
		v = "%" + v + "%"
		db = db.Where("name LIKE ? OR value LIKE ?", v, v)
	}

	if len(opt.SelectFields) > 0 {
		db = db.Select(opt.SelectFields)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list Properties
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &schema.PropertyQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaProperties(),
	}
	return qr, nil
}

func (a *PropertyRepo) FindPropertyMap(ctx context.Context, name string) map[string]string {
	properties := a.FindAu(ctx, name)
	propertyMap := make(map[string]string)
	for i := range properties {
		propertyMap[properties[i].Name] = properties[i].Value
	}
	return propertyMap
}

func (a *PropertyRepo) FindAu(ctx context.Context, name string) (o []Property) {
	if GetPropertyDB(ctx, a.DB).Model(Property{}).Where("name like ?", name+"%").Find(&o).Error != nil {
		return nil
	}
	return
}

func (a *PropertyRepo) Get(ctx context.Context, name string, opts ...schema.PropertyQueryOptions) (*schema.Property, error) {
	var item Property
	ok, err := util.FindOne(ctx, GetPropertyDB(ctx, a.DB).Where("name=?", name), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	return item.ToSchemaProperty(), nil
}

func (a *PropertyRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	err := GetPropertyDB(ctx, a.DB).Model(&Property{}).Count(&count).Error
	return count, err
}

func (a *PropertyRepo) Create(ctx context.Context, item schema.Property) error {
	eitem := SchemaProperty(item).ToProperty()
	result := GetPropertyDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *PropertyRepo) CreateByMap(ctx context.Context, m map[string]string) (err error) {
	var o Property
	for k, v := range m {
		o.Model.ID = snowflake.MustID()
		o.Name = k
		o.Value = v
		if err = GetPropertyDB(ctx, a.DB).Create(&o).Error; err != nil {
			return
		}
	}
	return
}

func (a *PropertyRepo) Update(ctx context.Context, name string, item schema.Property) error {
	eitem := SchemaProperty(item).ToProperty()
	result := GetPropertyDB(ctx, a.DB).Where("name=?", name).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *PropertyRepo) Delete(ctx context.Context, name string) error {
	result := GetPropertyDB(ctx, a.DB).Where("name=?", name).Delete(Property{})
	return errors.WithStack(result.Error)
}
