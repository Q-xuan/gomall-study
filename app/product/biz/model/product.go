package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	Price       float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (p ProductQuery) GetById(productId int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, productId).Error
	return
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products, "name like ? or description like ?",
		"%"+q+"%", "%"+q+"%",
	).Error
	return
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

func NewCacheProductQuery(ctx context.Context, db *gorm.DB, cacheClient *redis.Client) *CacheProductQuery {
	return &CacheProductQuery{
		productQuery: *NewProductQuery(ctx, db),
		cacheClient:  cacheClient,
		prefix:       "shop",
	}
}

type CacheProductQuery struct {
	productQuery ProductQuery
	cacheClient  *redis.Client
	prefix       string
}

func (c CacheProductQuery) GetById(productId int) (product Product, err error) {
	cachedKey := fmt.Sprintf("%s_%s_%d", c.prefix, "product_by_id", productId)
	cachedResult := c.cacheClient.Get(c.productQuery.ctx, cachedKey)

	err = func() error {
		if err := cachedResult.Err(); err != nil {
			return nil
		}
		crb, err := cachedResult.Bytes()
		if err != nil {
			return err
		}

		err = json.Unmarshal(crb, &product)
		if err != nil {
			return err
		}
		return nil
	}()

	if err != nil {
		product, err = c.productQuery.GetById(productId)
		if err != nil {
			return Product{}, err
		}
		encoded, err := json.Marshal(product)
		if err != nil {
			return product, err
		}

		_ = c.cacheClient.Set(c.productQuery.ctx, cachedKey, encoded, time.Hour)
	}
	return
}

func (c CacheProductQuery) SearchProducts(q string) (products []*Product, err error) {
	return c.productQuery.SearchProducts(q)
}

//读写分离 这用于写
// type ProductMutation struct{
// 	ctx context.Context
// 	db *gorm.DB
// }
