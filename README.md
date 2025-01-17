# trigger-update-background-process-go
Trigger Process update 500 data at db_source - db_destination using golang

# Golang Product Synchronization

This project synchronizes data between two MySQL databases (`db_source` and `db_destination`). The process involves copying data from `source_product` in the source database to `destination_product` in the destination database. It runs a background job to update data in `destination_product` based on changes in `source_product`.

## Requirements

- Go 1.23.4 or higher
- MySQL
- GORM (Go ORM)
- Gorilla Mux (HTTP router)
- .env file configuration

## Environment Configuration

Make sure to create a `.env` file in the root of your project and add the following configurations:

```bash
PORT=4000
DB_SOURCE=user:password@tcp(127.0.0.1:3306)/db_source?charset=utf8mb4&parseTime=True&loc=Local
DB_DESTINATION=user:password@tcp(127.0.0.1:3306)/db_destination?charset=utf8mb4&parseTime=True&loc=Local
TRUNCATE_DATA=false
