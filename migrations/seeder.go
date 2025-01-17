package migrations

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"math/rand"
	"sourceanddestination/models"
	"time"

	"gorm.io/gorm"
)

func GenerateSeedData(dbSource *gorm.DB, dbDestination *gorm.DB) {

	// jika reset data
	isTruncateData := os.Getenv("TRUNCATE_DATA")
	if isTruncateData == "true" {
		log.Println("Reset Dummy Data")
		truncateTable(dbSource, "source_products")
		truncateTable(dbDestination, "destination_products")
	}

	// Insert 500 data untuk source_product dan destination_product
	var sourceProducts []models.SourceProduct
	sourceData := dbSource.Find(&sourceProducts)

	if sourceData.Error != nil {
		fmt.Println("Error fetching data:", sourceData.Error)
		return
	}

	var destinationProducts []models.DestinationProduct
	destData := dbDestination.Find(&destinationProducts)

	if destData.Error != nil {
		fmt.Println("Error fetching data:", destData.Error)
		return
	}

	if len(sourceProducts) == 0 && len(destinationProducts) == 0 {
		for i := 1; i <= 500; i++ {
			productName := "Product_" + strconv.Itoa(i)

			currentTime := time.Now()

			// Membuat data untuk source_product
			sourceProduct := models.SourceProduct{
				ID:           int64(i),
				ProductName:  productName,
				Qty:          int64(rand.Intn(100) + 1),
				SellingPrice: rand.Float64() * 1000,
				PromoPrice:   rand.Float64() * 500,
				CreatedAt:    currentTime,
			}
			dbSource.Create(&sourceProduct)

			// Membuat data untuk destination_product
			destinationProduct := models.DestinationProduct{
				ID:           int64(i),
				ProductName:  productName,
				Qty:          0,
				SellingPrice: 0,
				PromoPrice:   0,
				CreatedAt:    currentTime,
			}
			dbDestination.Create(&destinationProduct)
		}

		log.Println("500 Dummy Data Generated Successfully!")
	} else {
		log.Println("500 Dummy Data Already Usage!")
	}
}

func truncateTable(db *gorm.DB, tableName string) {
	err := db.Exec("TRUNCATE TABLE " + tableName).Error
	if err != nil {
		log.Fatalf("Failed to truncate table %s: %v", tableName, err)
	}
	log.Printf("Table %s truncated successfully", tableName)
}
