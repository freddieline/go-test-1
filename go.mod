module warehouse-stocks-checker

go 1.18

require internal/infrastructure/warehouse-api v0.0.0
replace internal/infrastructure/warehouse-api => ./internal/infrastructure/warehouse-api
