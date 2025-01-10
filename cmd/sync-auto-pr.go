package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
	"wms-go/models"
	"wms-go/utils"

	"github.com/google/uuid"
)

type OrderDetail struct {
	TableName              string
	ID                     int
	COrderlineID           int
	COrderID               int
	CBPartnerID            int
	ItemID                 int
	FactoryID              int
	SoID                   int
	CategoryID             int
	PurchaseNumber         string
	SupplierName           string
	Category               string
	ItemCode               string
	UOM                    string
	QtyAllocation          float64
	POBuyer                string
	LCDate                 time.Time
	IsRecycle              string
	StatusLC               string
	StdPrecision           int
	SoOrderType            string
	Season                 string
	WarehousePlace         string
	SoDocTypeID            int
	PromiseDate            time.Time
	JobOrder               string
	PurchaseDocumentTypeID int
	IsFabric               string
	Color                  string
	UOMID                  int
	ItemDesc               string
	LastUpdatePO           time.Time
	IsMRPExists            bool
}

func SyncAutoPr() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
	}
	DB := models.DB
	start := time.Now().In(loc)
	fmt.Println("Start:", start.Format("15:04:05"))
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

	end := time.Now().In(loc)
	fmt.Println("End:", end.Format("15:04:05"))

	elapsed := end.Sub(start)

	fmt.Println("Elapsed Time of Fetch Data:", elapsed.Seconds(), "seconds")

	var allocations = []models.Allocation{}

	var factories []models.Factory

	err = DB.Where("is_active = ?", true).Find(&factories).Error

	if err != nil {
		log.Fatalf("Error fetching factories: %v", err)
	}

	start = time.Now().In(loc)
	fmt.Println("Start:", start.Format("15:04:05"))
	for i, each := range outstandingAllocationPrs {
		fmt.Printf("%d of %d -> %d\n", i+1, len(outstandingAllocationPrs), each.ID)
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
		case 1000002:
			warehouseSequence = "001"
		case 1000013:
			warehouseSequence = "002"
		case 1000082:
			warehouseSequence = "004"
		}

		isAllocationNagai := false
		if each.PurchaseDocumentTypeID == 1000239 {
			isAllocationNagai = true
		}

		var batchNumberStock *string
		if each.IsFabric == "Y" {
			batchNumberStockValue := "auto"
			batchNumberStock = &batchNumberStockValue
		} else {
			batchNumberStock = nil
		}

		isFabricBool := true
		if each.IsFabric == "N" {
			isFabricBool = false
		}

		isRecycleBool := true
		if each.IsRecycle == "N" {
			isRecycleBool = false
		}

		var buyer models.Buyer

		err = DB.Where("purchase_document_type_id = ? AND erp_client_org_id = ?", strconv.Itoa(each.PurchaseDocumentTypeID), "bima").First(&buyer).Error

		if err != nil {
			log.Fatalf("Error get buyer: %v", err)
		}

		buyerName := buyer.BuyerName

		var factoryName string

		for _, factory := range factories {
			if factory.ID == each.FactoryID {
				factoryName = factory.FactoryName
				break
			}
		}

		if err != nil {
			log.Fatalf("Error get factory: %v", err)
		}

		stringFactoryID := strconv.Itoa(each.FactoryID)

		allocation := models.Allocation{
			ID:                        uuid.New().String(),
			IsNewBima:                 true,
			ERPOrg:                    "bima",
			ERPCOrderLineID:           strconv.Itoa(each.COrderlineID),
			SODocumentTypeID:          strconv.Itoa(each.SoDocTypeID),
			SOID:                      strconv.Itoa(each.SoID),
			ERPSoTypeStock:            each.SoOrderType,
			ERPSoTypeStockCode:        each.SoOrderType,
			ERPWarehousePlace:         each.WarehousePlace,
			ERPSeason:                 each.Season,
			State:                     "reguler",
			ERPWMSTMaterialStockID:    nil,
			AllocationNumber:          "ERP-" + each.StatusLC + "LC-" + dayLc + "-" + monthLc + "-" + warehouseSequence,
			CategoryIDSource:          strconv.Itoa(each.CategoryID),
			CategoryIDBook:            strconv.Itoa(each.CategoryID),
			COrderID:                  strconv.Itoa(each.COrderID),
			CBPartnerID:               strconv.Itoa(each.CBPartnerID),
			ItemIDSource:              strconv.Itoa(each.ItemID),
			ItemIDBook:                strconv.Itoa(each.ItemID),
			FactoryID:                 stringFactoryID,
			StdPrecision:              float64(each.StdPrecision),
			OrderDate:                 each.LCDate,
			PromiseDate:               each.PromiseDate,
			PurchaseNumber:            each.PurchaseNumber,
			SupplierName:              each.SupplierName,
			POBuyer:                   each.POBuyer,
			JobOrder:                  each.JobOrder,
			Lot:                       "-",
			ItemCodeBook:              each.ItemCode,
			ItemDescBook:              utils.SafeSlice(each.ItemDesc, 185),
			CategoryBook:              each.Category,
			ItemCodeSource:            each.ItemCode,
			ItemDescSource:            utils.SafeSlice(each.ItemDesc, 185),
			CategorySource:            each.Category,
			FactoryName:               factoryName,
			UOM:                       each.UOM,
			QtyAllocation:             each.QtyAllocation,
			QtyOutstanding:            each.QtyAllocation,
			QtyAllocated:              0,
			IsFabric:                  isFabricBool,
			IsAdditional:              false,
			IsAllocationStockTransfer: false,
			NoteAdditional:            nil,
			IsAllocationPurchase:      IsAllocationPurchase,
			GenerateFormBooking:       nil,
			OldSOID:                   strconv.Itoa(each.SoID),
			OldPOBuyer:                each.POBuyer,
			OldJobOrder:               each.JobOrder,
			OldLot:                    "-",
			StatusPOBuyer:             "active",
			UserID:                    315,
			UpdatedUserID:             0,
			CreatedAt:                 time.Now(),
			UpdatedAt:                 time.Now(),
			ConfirmDate:               time.Now(),
			BatchNumberStock:          batchNumberStock,
			IsAllocationNagai:         isAllocationNagai,
			PurchaseDocumentTypeID:    strconv.Itoa(each.PurchaseDocumentTypeID),
			BuyerName:                 buyerName,
			ItemColorBook:             utils.SafeSlice(each.Color, 185),
			ItemColorSource:           utils.SafeSlice(each.Color, 185),
			ERPAllocationID:           each.TableName + "_" + strconv.Itoa(each.ID),
			UOMIDSource:               strconv.Itoa(each.UOMID),
			UOMIDBook:                 strconv.Itoa(each.UOMID),
			IsRecycleSource:           isRecycleBool,
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
	end = time.Now().In(loc)
	elapsed = end.Sub(start)
	fmt.Println("End:", end.Format("15:04:05"))
	fmt.Println("Elapsed Time of Insert Data:", elapsed.Seconds(), "seconds")
}
