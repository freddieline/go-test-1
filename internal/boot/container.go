package boot

import (
	warehouse_api "github.com/freddieline/warehouse-stocks-checker/internal/infrastructure/warehouse-api"
	"net/http"
)

func (c *Container) WarehouseAPIClient() warehouse_api.Client {
	wc := &APIClient{
		client: http.Client{},
	}
}
