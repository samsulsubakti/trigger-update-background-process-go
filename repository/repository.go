package repository

import (
	"log"
	"net/http"
	"sourceanddestination/models"

	"gorm.io/gorm"
)

type ApiResponse struct {
	Data interface{} `json:"data"`
}

func GetList(db *gorm.DB, models interface{}, w http.ResponseWriter) ApiResponse {
	if db == nil {
		log.Println("Database connection not initialized")
		http.Error(w, "Database connection not initialized", http.StatusInternalServerError)
		return ApiResponse{}
	}
	data := db.Find(&models)

	if data.Error != nil {
		log.Println("Error fetching source products:", data.Error)
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return ApiResponse{}
	}

	response := ApiResponse{
		Data: models,
	}

	return response
}

func ProcessBackgroundUpdate(dbSource, dbDestination *gorm.DB) {

	var sourceProducts []models.SourceProduct
	err := dbSource.Find(&sourceProducts).Error
	if err != nil {
		log.Println("Error fetching data from source_product:", err)
		return
	}

	for _, sourceProduct := range sourceProducts {

		var destProduct models.DestinationProduct
		err := dbDestination.First(&destProduct, "id = ?", sourceProduct.ID).Error
		if err != nil {
			log.Println("Error finding destination_product with ID:", sourceProduct.ID)
			continue
		}

		destProduct.ProductName = sourceProduct.ProductName
		destProduct.Qty = sourceProduct.Qty
		destProduct.SellingPrice = sourceProduct.SellingPrice
		destProduct.PromoPrice = sourceProduct.PromoPrice

		// Save the updated destination product
		err = dbDestination.Save(&destProduct).Error
		if err != nil {
			log.Println("Error updating destination_product with ID:", sourceProduct.ID, err)
		}
	}

	log.Println("Background process completed successfully.")
}
