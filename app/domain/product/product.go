package product

import (
	"unicode/utf8"

	errDomain "github.com/t-shimpo/go-mysql-docker/app/domain/error"
	"github.com/t-shimpo/go-mysql-docker/pkg/ulid"
)

type Product struct {
	id          string
	ownerID     string
	name        string
	description string
	price       int64
	stock       int
}

// 商品の出品に関するビジネスルール。
// このビジネスルールは商品生成に関する制約であり、これを表現するためにファクトリー関数を利用。
func newProduct (
	id          string,
	ownerID     string,
	name        string,
	description string,
	price       int64,
	stock       int,
) (*Product, error) {
	if !ulid.IsValid(ownerID) {
		return nil, errDomain.NewError("オーナーIDの値が不正です。")
	}
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("商品の値が不正です。")
	}
	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.NewError("商品説明の値が不正です。")
	}
	if price < 1 {
		return nil, errDomain.NewError("価格の値が不正です。")
	}
	if stock < 0 {
		return nil, errDomain.NewError("在庫数の値が不正です 。")
	}

	return &Product{
		id: id,
		ownerID: ownerID,
		name: name,
		description: description,
		price: price,
		stock: stock,
	}, nil
}

// newProductで定義したものを外部から利用可能な形にする関数を定義する
// 商品オブジェクトの生成には、NewProductとReconstruct関数という2つのexportされたファクトリー関数を定義
func Reconstruct(
	id          string,
	ownerID     string,
	name        string,
	description string,
	price       int64,
	stock       int,
) (*Product, error) {
	return newProduct(
		id,
		ownerID,
		name,
		description,
		price,
		stock,
	)
}

func NewProduct(
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
		return newProduct(
		ulid.NewULID(),
		ownerID,
		name,
		description,
		price,
		stock,
	)
}

const (
	nameLengthMin = 1
	nameLengthMax = 100
	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)
