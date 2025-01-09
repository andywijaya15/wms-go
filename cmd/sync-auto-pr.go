package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	"wms-go/models"

	"github.com/google/uuid"
)

type OrderDetail struct {
	TableName              string
	ID                     int
	COrderlineID           string
	COrderID               string
	CBPartnerID            string
	ItemID                 string
	FactoryID              string
	SoID                   string
	CategoryID             string
	PurchaseNumber         string
	SupplierName           string
	Category               string
	ItemCode               string
	UOM                    string
	QtyAllocation          float64
	POBuyer                string
	LCDate                 time.Time
	IsRecycle              bool
	StatusLC               string
	StdPrecision           int
	SoOrderType            string
	Season                 string
	WarehousePlace         string
	SoDocTypeID            string
	PromiseDate            time.Time
	JobOrder               string
	PurchaseDocumentTypeID string
	IsFabric               bool
	Color                  string
	UOMID                  string
	ItemDesc               string
	LastUpdatePO           time.Time
	IsMRPExists            bool
}

var DB = models.DB

func SyncAutoPr() {
	response, err := http.Get("http://localhost:8001/v1/get-auto-pr")
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch data. HTTP Status: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var outstandingAllocationPrs []OrderDetail
	err = json.Unmarshal(body, &outstandingAllocationPrs)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	var allocations = []models.Allocation{}

	var factories []models.Factory

	err = DB.Where("is_active = ?", true).Find(&factories).Error

	if err != nil {
		log.Fatalf("Error fetching factories: %v", err)
	}

	for _, each := range outstandingAllocationPrs {
		IsAllocationPurchase := false
		if each.StatusLC == "PR" {
			IsAllocationPurchase = true
		}
		parsedDate, err := time.Parse("2006-01-02", each.LCDate.Format("2006-01-02"))
		if err != nil {
			log.Fatalf("Error parsing date: %v", err)
			return
		}

		dayLc := parsedDate.Format("02")
		monthLc := parsedDate.Format("01")

		warehouseSequence := "003"
		switch each.FactoryID {
		case "1000002":
			warehouseSequence = "001"
		case "1000013":
			warehouseSequence = "002"
		case "1000082":
			warehouseSequence = "004"
		}

		isAllocationNagai := false
		if each.PurchaseDocumentTypeID == "1000239" {
			isAllocationNagai = true
		}

		var batchNumberStock *string
		if each.IsFabric {
			batchNumberStockValue := "auto"
			batchNumberStock = &batchNumberStockValue
		} else {
			batchNumberStock = nil
		}

		var buyer models.Buyer

		err = DB.Where("purchase_document_type_id = ? AND erp_client_org_id = ?", each.PurchaseDocumentTypeID, "bima").First(&buyer).Error

		if err != nil {
			log.Fatalf("Error get buyer: %v", err)
		}

		buyerName := buyer.BuyerName

		var factoryName string

		parseIntFactory, err := strconv.Atoi(each.FactoryID)
		if err != nil {
			log.Fatalf("Error converting FactoryID to int: %v", err)
		}

		for _, factory := range factories {
			if factory.ID == parseIntFactory {
				factoryName = factory.FactoryName
				break
			}
		}

		if err != nil {
			log.Fatalf("Error get factory: %v", err)
		}

		allocation := models.Allocation{
			ID:                        uuid.New().String(),
			IsNewBima:                 true,
			ERPOrg:                    "bima",
			ERPCOrderLineID:           each.COrderlineID,
			SODocumentTypeID:          each.SoDocTypeID,
			SOID:                      each.SoID,
			ERPSoTypeStock:            each.SoOrderType,
			ERPSoTypeStockCode:        each.SoOrderType,
			ERPWarehousePlace:         each.WarehousePlace,
			ERPSeason:                 each.Season,
			State:                     "reguler",
			ERPWMSTMaterialStockID:    nil,
			AllocationNumber:          "ERP-" + each.StatusLC + "LC-" + dayLc + "-" + monthLc + "-" + warehouseSequence,
			CategoryIDSource:          each.CategoryID,
			CategoryIDBook:            each.CategoryID,
			COrderID:                  each.COrderID,
			CBPartnerID:               each.CBPartnerID,
			ItemIDSource:              each.ItemID,
			ItemIDBook:                each.ItemID,
			FactoryID:                 each.FactoryID,
			StdPrecision:              float64(each.StdPrecision),
			OrderDate:                 each.LCDate,
			PromiseDate:               each.PromiseDate,
			PurchaseNumber:            each.PurchaseNumber,
			SupplierName:              each.SupplierName,
			POBuyer:                   each.POBuyer,
			JobOrder:                  each.JobOrder,
			Lot:                       "-",
			ItemCodeBook:              each.ItemCode,
			ItemDescBook:              each.ItemDesc[:185],
			CategoryBook:              each.Category,
			ItemCodeSource:            each.ItemCode,
			ItemDescSource:            each.ItemDesc[:185],
			CategorySource:            each.Category,
			FactoryName:               factoryName,
			UOM:                       each.UOM,
			QtyAllocation:             each.QtyAllocation,
			QtyOutstanding:            each.QtyAllocation,
			QtyAllocated:              0,
			IsFabric:                  each.IsFabric,
			IsAdditional:              false,
			IsAllocationStockTransfer: false,
			NoteAdditional:            nil,
			IsAllocationPurchase:      IsAllocationPurchase,
			GenerateFormBooking:       nil,
			OldSOID:                   each.SoID,
			OldPOBuyer:                each.POBuyer,
			OldJobOrder:               each.JobOrder,
			OldLot:                    "-",
			StatusPOBuyer:             "active",
			UserID:                    0,
			UpdatedUserID:             0,
			CreatedAt:                 time.Now(),
			UpdatedAt:                 time.Now(),
			ConfirmDate:               time.Now(),
			BatchNumberStock:          batchNumberStock,
			IsAllocationNagai:         isAllocationNagai,
			PurchaseDocumentTypeID:    each.PurchaseDocumentTypeID,
			BuyerName:                 buyerName,
			ItemColorBook:             each.Color[:185],
			ItemColorSource:           each.Color[:185],
			ERPAllocationID:           each.TableName + "_" + strconv.Itoa(each.ID),
			UOMIDSource:               each.UOMID,
			UOMIDBook:                 each.UOMID,
			IsRecycleSource:           each.IsRecycle,
			POBuyerReference:          nil,
			ERPBrand:                  "-",
		}
		allocations = append(allocations, allocation)
	}
	if len(allocations) > 0 {
		err := DB.CreateInBatches(&allocations, 1000).Error
		if err != nil {
			log.Fatalf("Error inserting allocations: %v", err)
		}
	}
}
