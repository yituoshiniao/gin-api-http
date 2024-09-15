// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/yituoshiniao/gin-api-http/gen/dao/model"
)

func newAppleConfig(db *gorm.DB, opts ...gen.DOOption) appleConfig {
	_appleConfig := appleConfig{}

	_appleConfig.appleConfigDo.UseDB(db, opts...)
	_appleConfig.appleConfigDo.UseModel(&model.AppleConfig{})

	tableName := _appleConfig.appleConfigDo.TableName()
	_appleConfig.ALL = field.NewAsterisk(tableName)
	_appleConfig.ID = field.NewUint32(tableName, "id")
	_appleConfig.Channel = field.NewString(tableName, "channel")
	_appleConfig.KeyFile = field.NewString(tableName, "key_file")
	_appleConfig.PartnerKey = field.NewString(tableName, "partner_key")
	_appleConfig.AppID = field.NewInt64(tableName, "app_id")
	_appleConfig.Bid = field.NewString(tableName, "bid")
	_appleConfig.SubscriptionGroupsID = field.NewInt32(tableName, "subscription_groups_id")
	_appleConfig.IssuerID = field.NewString(tableName, "issuer_id")
	_appleConfig.Audience = field.NewString(tableName, "audience")
	_appleConfig.Alg = field.NewString(tableName, "alg")
	_appleConfig.Kid = field.NewString(tableName, "kid")
	_appleConfig.IsDel = field.NewInt32(tableName, "is_del")
	_appleConfig.Typ = field.NewString(tableName, "typ")
	_appleConfig.CreateTime = field.NewTime(tableName, "create_time")
	_appleConfig.UpdateTime = field.NewTime(tableName, "update_time")

	_appleConfig.fillFieldMap()

	return _appleConfig
}

type appleConfig struct {
	appleConfigDo appleConfigDo

	ALL                  field.Asterisk
	ID                   field.Uint32 // 主键
	Channel              field.String // 渠道来源: cs (全能扫描王) | camexam (蜜蜂) 等
	KeyFile              field.String // pem 秘钥路径
	PartnerKey           field.String // 计算签名key
	AppID                field.Int64  // appid
	Bid                  field.String // 应用bid
	SubscriptionGroupsID field.Int32  // 订阅组id
	IssuerID             field.String // 苹果后台 IssuerID
	Audience             field.String // appstoreconnect-v1
	Alg                  field.String // JWTHeader  Alg
	Kid                  field.String // JWTHeader 苹果后台kid
	IsDel                field.Int32  // 是否删除: 0 否; 1 是;
	Typ                  field.String // JWTHeader Typ
	CreateTime           field.Time   // 创建时间
	UpdateTime           field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (a appleConfig) Table(newTableName string) *appleConfig {
	a.appleConfigDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a appleConfig) As(alias string) *appleConfig {
	a.appleConfigDo.DO = *(a.appleConfigDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *appleConfig) updateTableName(table string) *appleConfig {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewUint32(table, "id")
	a.Channel = field.NewString(table, "channel")
	a.KeyFile = field.NewString(table, "key_file")
	a.PartnerKey = field.NewString(table, "partner_key")
	a.AppID = field.NewInt64(table, "app_id")
	a.Bid = field.NewString(table, "bid")
	a.SubscriptionGroupsID = field.NewInt32(table, "subscription_groups_id")
	a.IssuerID = field.NewString(table, "issuer_id")
	a.Audience = field.NewString(table, "audience")
	a.Alg = field.NewString(table, "alg")
	a.Kid = field.NewString(table, "kid")
	a.IsDel = field.NewInt32(table, "is_del")
	a.Typ = field.NewString(table, "typ")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")

	a.fillFieldMap()

	return a
}

func (a *appleConfig) WithContext(ctx context.Context) *appleConfigDo {
	return a.appleConfigDo.WithContext(ctx)
}

func (a appleConfig) TableName() string { return a.appleConfigDo.TableName() }

func (a appleConfig) Alias() string { return a.appleConfigDo.Alias() }

func (a appleConfig) Columns(cols ...field.Expr) gen.Columns { return a.appleConfigDo.Columns(cols...) }

func (a *appleConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *appleConfig) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 15)
	a.fieldMap["id"] = a.ID
	a.fieldMap["channel"] = a.Channel
	a.fieldMap["key_file"] = a.KeyFile
	a.fieldMap["partner_key"] = a.PartnerKey
	a.fieldMap["app_id"] = a.AppID
	a.fieldMap["bid"] = a.Bid
	a.fieldMap["subscription_groups_id"] = a.SubscriptionGroupsID
	a.fieldMap["issuer_id"] = a.IssuerID
	a.fieldMap["audience"] = a.Audience
	a.fieldMap["alg"] = a.Alg
	a.fieldMap["kid"] = a.Kid
	a.fieldMap["is_del"] = a.IsDel
	a.fieldMap["typ"] = a.Typ
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
}

func (a appleConfig) clone(db *gorm.DB) appleConfig {
	a.appleConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a appleConfig) replaceDB(db *gorm.DB) appleConfig {
	a.appleConfigDo.ReplaceDB(db)
	return a
}

type appleConfigDo struct{ gen.DO }

func (a appleConfigDo) Debug() *appleConfigDo {
	return a.withDO(a.DO.Debug())
}

func (a appleConfigDo) WithContext(ctx context.Context) *appleConfigDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a appleConfigDo) ReadDB() *appleConfigDo {
	return a.Clauses(dbresolver.Read)
}

func (a appleConfigDo) WriteDB() *appleConfigDo {
	return a.Clauses(dbresolver.Write)
}

func (a appleConfigDo) Session(config *gorm.Session) *appleConfigDo {
	return a.withDO(a.DO.Session(config))
}

func (a appleConfigDo) Clauses(conds ...clause.Expression) *appleConfigDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a appleConfigDo) Returning(value interface{}, columns ...string) *appleConfigDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a appleConfigDo) Not(conds ...gen.Condition) *appleConfigDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a appleConfigDo) Or(conds ...gen.Condition) *appleConfigDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a appleConfigDo) Select(conds ...field.Expr) *appleConfigDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a appleConfigDo) Where(conds ...gen.Condition) *appleConfigDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a appleConfigDo) Order(conds ...field.Expr) *appleConfigDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a appleConfigDo) Distinct(cols ...field.Expr) *appleConfigDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a appleConfigDo) Omit(cols ...field.Expr) *appleConfigDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a appleConfigDo) Join(table schema.Tabler, on ...field.Expr) *appleConfigDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a appleConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) *appleConfigDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a appleConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) *appleConfigDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a appleConfigDo) Group(cols ...field.Expr) *appleConfigDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a appleConfigDo) Having(conds ...gen.Condition) *appleConfigDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a appleConfigDo) Limit(limit int) *appleConfigDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a appleConfigDo) Offset(offset int) *appleConfigDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a appleConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *appleConfigDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a appleConfigDo) Unscoped() *appleConfigDo {
	return a.withDO(a.DO.Unscoped())
}

func (a appleConfigDo) Create(values ...*model.AppleConfig) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a appleConfigDo) CreateInBatches(values []*model.AppleConfig, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a appleConfigDo) Save(values ...*model.AppleConfig) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a appleConfigDo) First() (*model.AppleConfig, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleConfig), nil
	}
}

func (a appleConfigDo) Take() (*model.AppleConfig, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleConfig), nil
	}
}

func (a appleConfigDo) Last() (*model.AppleConfig, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleConfig), nil
	}
}

func (a appleConfigDo) Find() ([]*model.AppleConfig, error) {
	result, err := a.DO.Find()
	return result.([]*model.AppleConfig), err
}

func (a appleConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AppleConfig, err error) {
	buf := make([]*model.AppleConfig, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a appleConfigDo) FindInBatches(result *[]*model.AppleConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a appleConfigDo) Attrs(attrs ...field.AssignExpr) *appleConfigDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a appleConfigDo) Assign(attrs ...field.AssignExpr) *appleConfigDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a appleConfigDo) Joins(fields ...field.RelationField) *appleConfigDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a appleConfigDo) Preload(fields ...field.RelationField) *appleConfigDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a appleConfigDo) FirstOrInit() (*model.AppleConfig, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleConfig), nil
	}
}

func (a appleConfigDo) FirstOrCreate() (*model.AppleConfig, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AppleConfig), nil
	}
}

func (a appleConfigDo) FindByPage(offset int, limit int) (result []*model.AppleConfig, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a appleConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a appleConfigDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a appleConfigDo) Delete(models ...*model.AppleConfig) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *appleConfigDo) withDO(do gen.Dao) *appleConfigDo {
	a.DO = *do.(*gen.DO)
	return a
}
