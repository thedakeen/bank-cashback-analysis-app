package main

import (
	"bank-cashback-analysis/backend/pkg/models"
	"bank-cashback-analysis/backend/pkg/models/mongodb"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func (app *application) insertHalyk(promoModel *mongodb.PromoModel) {
	cityCodes, err := GetAllCityCodes()
	if err != nil {
		log.Fatal(err)
	}

	for _, cityCode := range cityCodes {
		categories, err := GetCategoriesByCity(cityCode)
		if err != nil {
			log.Printf("Error getting categories for city %s: %v\n", cityCode, err)
			continue
		}

		//fmt.Printf("Categories for city %s:\n", cityCode)
		for _, cat := range categories {
			//fmt.Println("Category:", cat.Code, "- Count:", cat.Count)
			if cat.Count == 0 {
				//fmt.Printf("No merchants found for category %s in city %s\n", cat.Code, cityCode)
				continue
			}

			shops, err := GetShopByCategory(cityCode, cat.Code)
			if err != nil {
				continue
			}

			//fmt.Printf("Merchants for category %s in city %s:\n", cat.Code, cityCode)
			for _, shop := range shops {
				bonus, err := extractBonusFromTags(shop.Tags)
				if err != nil {
					log.Printf("Error extracting bonus from tags: %v\n", err)
					continue
				}

				promotion := models.Promotion{
					BankName:     "Halyk Bank",
					CompanyName:  shop.CompanyName,
					CategoryName: shop.CategoryName,
					Type:         "Company",
					BonusRate:    bonus,
				}

				err = promoModel.SavePromotionToDB(promotion)
				if err != nil {
					log.Printf("Error saving promotion to database: %v\n", err)
				}
			}
		}
	}
}

func extractBonusFromTags(tags []models.Tag) (float64, error) {
	for _, tag := range tags {
		bonusStr := extractBonusValue(tag.Bonus)
		bonus, err := strconv.ParseFloat(bonusStr, 64)
		if err != nil {
			return 0, err
		}
		return bonus, nil
	}
	return 0, models.ErrNoRecord
}

func extractBonusValue(tagText string) string {

	re := regexp.MustCompile(`\d+\.*\d*`)
	match := re.FindString(tagText)
	return match
}

func GetShopByCategory(cityCode, categoryName string) ([]models.Shop, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://pelican-api.homebank.kz/halykclub-api/v1/terminal/merchants?category_code=%s&filter=", categoryName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("City_id", cityCode)
	req.Header.Set("Accept-Language", "ru")
	req.Header.Set("Authorization", "Bearer _XRM9CVQN0-O0P8JB6DBKA")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var shopResponse models.ShopResponse
	err = json.NewDecoder(res.Body).Decode(&shopResponse)
	if err != nil {
		return nil, err
	}

	var shops []models.Shop
	for _, shop := range shopResponse.Shops {
		newShop := models.Shop{
			CompanyName:  shop.CompanyName,
			CategoryName: categoryName,
			Tags:         shop.Tags,
		}
		shops = append(shops, newShop)
	}

	return shops, nil
}

func GetCategoriesByCity(cityCode string) ([]models.Category, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://pelican-api.homebank.kz/halykclub-api/v1/dictionary/categories", nil)

	req.Header.Set("City_id", cityCode)
	req.Header.Set("Accept-Language", "ru")
	req.Header.Set("Authorization", "Bearer _XRM9CVQN0-O0P8JB6DBKA")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(res.Body)

	var categories []models.Category

	if err := json.NewDecoder(res.Body).Decode(&categories); err != nil {
		return nil, err
	}

	return categories, nil

}

func GetAllCityCodes() ([]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pelican-api.homebank.kz/halykclub-api/v1/dictionary/cities", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Language", "ru")
	req.Header.Set("Authorization", "Bearer _XRM9CVQN0-O0P8JB6DBKA")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var cities []struct {
		CityCode string `json:"city_id"`
	}

	if err := json.NewDecoder(res.Body).Decode(&cities); err != nil {
		return nil, err
	}

	var cityCodes []string
	for _, cit := range cities {
		cityCodes = append(cityCodes, cit.CityCode)
	}

	return cityCodes, nil
}
