package internal

import wa "app/internal/infrastructure/warehouse-api"

type Main struct{
	warehouseApi wa
}

func NewMain()