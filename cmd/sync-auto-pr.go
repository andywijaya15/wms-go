package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
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

func fetchData(url string, ch chan<- []byte) {
	response, err := http.Get(url)
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
	ch <- body
}

func unmarshalData(data []byte, ch chan<- []OrderDetail) {
	var outstandingAllocationPrs []OrderDetail
	err := json.Unmarshal(data, &outstandingAllocationPrs)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	ch <- outstandingAllocationPrs
}

func retryDBOperation(operation func() error, retries int, delay time.Duration) error {
	for i := 0; i < retries; i++ {
		err := operation()
		if err == nil {
			return nil
		}
		if i < retries-1 {
			time.Sleep(delay)
		}
	}
	return fmt.Errorf("operation failed after %d retries", retries)
}

func SyncAutoPr() {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal(err)
	}
	DB := models.DB
	start := time.Now().In(loc)
	fmt.Println("Start:", start.Format("15:04:05"))

	dataCh := make(chan []byte)
	go fetchData("http://localhost:8001/v1/get-auto-pr", dataCh)

	body := <-dataCh

	outstandingCh := make(chan []OrderDetail)
	go unmarshalData(body, outstandingCh)

	outstandingAllocationPrs := <-outstandingCh

	end := time.Now().In(loc)
	fmt.Println("End:", end.Format("15:04:05"))

	elapsed := end.Sub(start)
	fmt.Println("Elapsed Time of Fetch Data:", elapsed.Seconds(), "seconds")

	var allocations = []models.Allocation{}
	var factories []models.Factory
	var buyers []models.Buyer

	err = retryDBOperation(func() error {
		return DB.Where("is_active = ?", true).Find(&factories).Error
	}, 3, 2*time.Second)
	if err != nil {
		log.Fatalf("Error fetching factories: %v", err)
	}

	err = retryDBOperation(func() error {
		return DB.Where("erp_client_org_id = ?", "bima").Find(&buyers).Error
	}, 3, 2*time.Second)
	if err != nil {
		log.Fatalf("Error fetching buyers: %v", err)
	}

	start = time.Now().In(loc)
	fmt.Println("Start:", start.Format("15:04:05"))

	var wg sync.WaitGroup
	allocationCh := make(chan models.Allocation, len(outstandingAllocationPrs))

	for _, each := range outstandingAllocationPrs {
		wg.Add(1)
		go func(each OrderDetail) {
			defer wg.Done()
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

			var buyerName string
			for _, buyer := range buyers {
				if buyer.PurchaseDocumentTypeID == strconv.Itoa(each.PurchaseDocumentTypeID) {
					buyerName = buyer.BuyerName
					break
				}
			}

			var factoryName string
			for _, factory := range factories {
				if factory.ID == each.FactoryID {
					factoryName = factory.FactoryName
					break
				}
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
			allocationCh <- allocation
		}(each)
	}

	go func() {
		wg.Wait()
		close(allocationCh)
	}()

	for allocation := range allocationCh {
		allocations = append(allocations, allocation)
	}

	batchSize := 300
	var insertWg sync.WaitGroup
	for i := 0; i < len(allocations); i += batchSize {
		insertWg.Add(1)
		go func(start int) {
			defer insertWg.Done()
			end := start + batchSize
			if end > len(allocations) {
				end = len(allocations)
			}
			err := retryDBOperation(func() error {
				return DB.CreateInBatches(allocations[start:end], batchSize).Error
			}, 3, 2*time.Second)
			if err != nil {
				log.Fatalf("Error inserting allocations: %v", err)
			}
		}(i)
	}
	insertWg.Wait()

	end = time.Now().In(loc)
	elapsed = end.Sub(start)
	fmt.Println("End:", end.Format("15:04:05"))
	fmt.Println("Elapsed Time of Insert Data:", elapsed.Seconds(), "seconds")
}
