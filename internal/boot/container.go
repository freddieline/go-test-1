package boot

import (
	warehouse_api "github.com/freddieline/go-test-1/internal/infrastructure/warehouse-api"
	"net/http"
)

func (c *Container) WarehouseAPIClient() warehouse_api.Client {
	wc := &APIClient{
		client: http.Client{},
	}
}
