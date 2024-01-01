package product

import (
	productDomain "github.com/t-shimpo/go-mysql-docker/app/domain/product"
)

func (uc *SaveProductUseCase) Run(/* 引数省略 */) (/* 返り値省略 */) {
	// 新しい商品オブジェクトを作成し、問題がなければリポジトリに保存する。
	p, err := productDomain.NewProduct(input.OwnerId, input.Name, input.Description, input.Price, input.Stock)
	if err != nil {
		return nil, err
	}
	// 商品の永続化
	err = uc.productRepo.Save(ctx, p)
	if err != nil {
		return nil, err
	}
}
