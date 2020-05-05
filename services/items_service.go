package services

var (
	// ItemsService ...
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct {}

type itemsServiceInterface interface {
	GetItems()
	SaveItems()
}

// GetItems ...
func (s *itemsService) GetItems() {

}
 
//SaveItems ...
func (s *itemsService) SaveItems() {

}