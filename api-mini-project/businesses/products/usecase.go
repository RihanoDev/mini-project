package products

type productUsecase struct {
	productRepository Repository
}

func NewProductUsecase(pr Repository) Usecase {
	return &productUsecase{
		productRepository: pr,
	}
}

func (pu *productUsecase) GetAll() []Domain {
	return pu.productRepository.GetAll()
}

func (pu *productUsecase) GetByID(id string) Domain {
	return pu.productRepository.GetByID(id)
}

func (pu *productUsecase) Create(productDomain *Domain) Domain {
	return pu.productRepository.Create(productDomain)
}

func (pu *productUsecase) Update(id string, productDomain *Domain) Domain {
	return pu.productRepository.Update(id, productDomain)
}

func (pu *productUsecase) Delete(id string) bool {
	return pu.productRepository.Delete(id)
}

func (pu *productUsecase) Restore(id string) Domain {
	return pu.productRepository.Restore(id)
}

func (pu *productUsecase) ForceDelete(id string) bool {
	return pu.productRepository.ForceDelete(id)
}
