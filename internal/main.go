package internal

import wa "github.com/freddieline/warehouse-stocks-checker/internal/infrastructure/warehouse-api"

type Main struct{
	warehouseApi wa
}

func NewMain()