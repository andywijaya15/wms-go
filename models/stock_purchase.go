package models

import "time"

type StockPurchase struct {
	ID                     string     `gorm:"column:id;primaryKey"`
	COrderID               *string    `gorm:"column:c_order_id;index"`
	COrderLineID           *string    `gorm:"column:c_order_line_id"`
	CBpartnerID            *string    `gorm:"column:c_bpartner_id"`
	ItemID                 *string    `gorm:"column:item_id"`
	FactoryShipmentID      *string    `gorm:"column:factory_shipment_id"`
	PurchaseNumber         *string    `gorm:"column:purchase_number"`
	ItemCode               *string    `gorm:"column:item_code"`
	Category               *string    `gorm:"column:category"`
	UOMOrder               *string    `gorm:"column:uom_order"`
	QtyPurchase            float64    `gorm:"column:qty_purchase"`
	QtyOnShipment          float64    `gorm:"column:qty_on_shipment"`
	QtyInHouse             float64    `gorm:"column:qty_in_house"`
	UOMConversion          *string    `gorm:"column:uom_conversion"`
	QtyPurchaseConversion  float64    `gorm:"column:qty_purchase_conversion"`
	QtyAllocatedConversion float64    `gorm:"column:qty_allocated_conversion"`
	QtyAvailableConversion float64    `gorm:"column:qty_available_conversion"`
	StdPrecision           float64    `gorm:"column:std_precision"`
	Deviderate             float64    `gorm:"column:deviderate"`
	Multiplyrate           float64    `gorm:"column:multiplyrate"`
	PurchaseDocumentType   *string    `gorm:"column:purchase_document_type"`
	IsMiscellaneousItem    bool       `gorm:"column:is_miscellaneous_item"`
	CreatedAt              *time.Time `gorm:"column:created_at"`
	UpdatedAt              *time.Time `gorm:"column:updated_at"`
	DeletedAt              *time.Time `gorm:"column:deleted_at"`
	StockAoi1              float64    `gorm:"column:stock_aoi_1"`
	StockAoi2              float64    `gorm:"column:stock_aoi_2"`
	StockBbi               float64    `gorm:"column:stock_bbi"`
	DocumentNo             *string    `gorm:"column:document_no"`
	IsFabric               bool       `gorm:"column:is_fabric"`
	SupplierName           *string    `gorm:"column:supplier_name"`
	POBuyerSource          *string    `gorm:"column:po_buyer_source"`
	NewOrderType           *string    `gorm:"column:new_order_type"`
	Description            *string    `gorm:"column:description"`
	Brand                  *string    `gorm:"column:brand"`
	Season                 *string    `gorm:"column:season"`
	OrderDate              *string    `gorm:"column:order_date"`
	DescPrLine             *string    `gorm:"column:desc_pr_line"`
	OrderType              *string    `gorm:"column:order_type"`
	DescPr                 *string    `gorm:"column:desc_pr"`
	PriceActual            float64    `gorm:"column:priceactual"`
	Currency               *string    `gorm:"column:currency"`
	IsClosed               bool       `gorm:"column:is_closed"`
}
