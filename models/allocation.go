package models

import (
	"time"
)

type Allocation struct {
	ID                             string     `gorm:"column:id;primaryKey" json:"id"`
	AllocationNumber               string     `gorm:"column:allocation_number" json:"allocation_number"`
	COrderID                       string     `gorm:"column:c_order_id" json:"c_order_id"`
	CBPartnerID                    string     `gorm:"column:c_bpartner_id" json:"c_bpartner_id"`
	ItemIDSource                   string     `gorm:"column:item_id_source" json:"item_id_source"`
	ItemIDBook                     string     `gorm:"column:item_id_book" json:"item_id_book"`
	FactoryID                      string     `gorm:"column:factory_id" json:"factory_id"`
	OrderDate                      time.Time  `gorm:"column:order_date" json:"order_date"`
	PromiseDate                    time.Time  `gorm:"column:promise_date" json:"promise_date"`
	PurchaseNumber                 string     `gorm:"column:purchase_number" json:"purchase_number"`
	SupplierName                   string     `gorm:"column:supplier_name" json:"supplier_name"`
	StatusPOBuyer                  string     `gorm:"column:status_po_buyer" json:"status_po_buyer"`
	POBuyer                        string     `gorm:"column:po_buyer" json:"po_buyer"`
	OldPOBuyer                     string     `gorm:"column:old_po_buyer" json:"old_po_buyer"`
	Lot                            string     `gorm:"column:lot" json:"lot"`
	ItemCodeBook                   string     `gorm:"column:item_code_book" json:"item_code_book"`
	ItemDescBook                   string     `gorm:"column:item_desc_book" json:"item_desc_book"`
	CategoryBook                   string     `gorm:"column:category_book" json:"category_book"`
	ItemCodeSource                 string     `gorm:"column:item_code_source" json:"item_code_source"`
	ItemDescSource                 string     `gorm:"column:item_desc_source" json:"item_desc_source"`
	CategorySource                 string     `gorm:"column:category_source" json:"category_source"`
	FactoryName                    string     `gorm:"column:factory_name" json:"factory_name"`
	UOM                            string     `gorm:"column:uom" json:"uom"`
	QtyAllocation                  float64    `gorm:"column:qty_allocation" json:"qty_allocation"`
	QtyOutstanding                 float64    `gorm:"column:qty_outstanding" json:"qty_outstanding"`
	QtyAllocated                   float64    `gorm:"column:qty_allocated" json:"qty_allocated"`
	IsFabric                       bool       `gorm:"column:is_fabric;default:false" json:"is_fabric"`
	IsAllocationPurchase           bool       `gorm:"column:is_allocation_purchase;default:false" json:"is_allocation_purchase"`
	IsAdditional                   bool       `gorm:"column:is_additional;default:false" json:"is_additional"`
	NoteAdditional                 *string    `gorm:"column:note_additional" json:"note_additional"`
	CancelPOBuyerDate              time.Time  `gorm:"column:cancel_po_buyer_date" json:"cancel_po_buyer_date"`
	GenerateFormBooking            *time.Time `gorm:"column:generate_form_booking" json:"generate_form_booking"`
	UserID                         int        `gorm:"column:user_id" json:"user_id"`
	DeletedUserID                  int        `gorm:"column:deleted_user_id" json:"deleted_user_id"`
	UpdatedUserID                  int        `gorm:"column:updated_user_id" json:"updated_user_id"`
	CancelAllocationDate           time.Time  `gorm:"column:cancel_allocation_date" json:"cancel_allocation_date"`
	CancelAllocationNote           string     `gorm:"column:cancel_allocation_note" json:"cancel_allocation_note"`
	CreatedAt                      time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt                      time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt                      time.Time  `gorm:"column:deleted_at" json:"deleted_at"`
	LockedDate                     time.Time  `gorm:"column:locked_date" json:"locked_date"`
	SOID                           string     `gorm:"column:so_id" json:"so_id"`
	JobOrder                       string     `gorm:"column:job_order" json:"job_order"`
	CategoryIDSource               string     `gorm:"column:category_id_source" json:"category_id_source"`
	CategoryIDBook                 string     `gorm:"column:category_id_book" json:"category_id_book"`
	OldSOID                        string     `gorm:"column:old_so_id" json:"old_so_id"`
	OldJobOrder                    string     `gorm:"column:old_job_order" json:"old_job_order"`
	OldLot                         string     `gorm:"column:old_lot" json:"old_lot"`
	NoteDelete                     string     `gorm:"column:note_delete" json:"note_delete"`
	PlanningDate                   string     `gorm:"column:planning_date" json:"planning_date"`
	Style                          string     `gorm:"column:style" json:"style"`
	ArticleNo                      string     `gorm:"column:article_no" json:"article_no"`
	ConfirmDate                    time.Time  `gorm:"column:confirm_date" json:"confirm_date"`
	NewStyle                       string     `gorm:"column:new_style" json:"new_style"`
	NewArticle                     string     `gorm:"column:new_article" json:"new_article"`
	OldStyle                       string     `gorm:"column:old_style" json:"old_style"`
	OldArticle                     string     `gorm:"column:old_article" json:"old_article"`
	BatchNumberStock               *string    `gorm:"column:batch_number_stock" json:"batch_number_stock"`
	IsAllocationNagai              bool       `gorm:"column:is_allocation_nagai" json:"is_allocation_nagai"`
	QtyAllocatedAdjustment         float64    `gorm:"column:qty_allocated_adjustment" json:"qty_allocated_adjustment"`
	ItemColorSource                string     `gorm:"column:item_color_source" json:"item_color_source"`
	ItemColorBook                  string     `gorm:"column:item_color_book" json:"item_color_book"`
	BuyerName                      *string    `gorm:"column:buyer_name" json:"buyer_name"`
	PurchaseDocumentTypeID         string     `gorm:"column:purchase_document_type_id" json:"purchase_document_type_id"`
	UOMIDSource                    string     `gorm:"column:uom_id_source" json:"uom_id_source"`
	UOMIDBook                      string     `gorm:"column:uom_id_book" json:"uom_id_book"`
	IsAllocationStockTransfer      bool       `gorm:"column:is_allocation_stock_transfer" json:"is_allocation_stock_transfer"`
	ERPAllocationID                string     `gorm:"column:erp_allocation_id" json:"erp_allocation_id"`
	SplitTemporaryID               string     `gorm:"column:split_temporary_id" json:"split_temporary_id"`
	IsAllocationSubcont            bool       `gorm:"column:is_allocation_subcont" json:"is_allocation_subcont"`
	ERPWMSTMaterialStockID         *string    `gorm:"column:erp_wms_material_stock_id" json:"erp_wms_material_stock_id"`
	IsReturnToAzuma                bool       `gorm:"column:is_return_to_azuma" json:"is_return_to_azuma"`
	QtyPlanned                     float64    `gorm:"column:qty_planned" json:"qty_planned"`
	AllocatedPurchaseNumber        string     `gorm:"column:allocated_purchase_number" json:"allocated_purchase_number"`
	AllocatedSupplierName          string     `gorm:"column:allocated_supplier_name" json:"allocated_supplier_name"`
	AllocatedStyle                 string     `gorm:"column:allocated_style" json:"allocated_style"`
	AllocatedSOID                  string     `gorm:"column:allocated_so_id" json:"allocated_so_id"`
	AllocatedArticleNo             string     `gorm:"column:allocated_article_no" json:"allocated_article_no"`
	AllocatedLot                   string     `gorm:"column:allocated_lot" json:"allocated_lot"`
	IsAllocationNonSO              bool       `gorm:"column:is_allocation_non_so" json:"is_allocation_non_so"`
	Pattern                        string     `gorm:"column:pattern" json:"pattern"`
	QtyRequiredMRP                 float64    `gorm:"column:qty_required_mrp" json:"qty_required_mrp"`
	SODocumentTypeID               string     `gorm:"column:so_document_type_id" json:"so_document_type_id"`
	AllocationStyle                string     `gorm:"column:allocation_style" json:"allocation_style"`
	ERPAdditionalType              string     `gorm:"column:erp_additional_type" json:"erp_additional_type"`
	ERPOrg                         string     `gorm:"column:erp_org" json:"erp_org"`
	IsNewBima                      bool       `gorm:"column:is_new_bima" json:"is_new_bima"`
	State                          string     `gorm:"column:state" json:"state"`
	ERPWarehousePlace              string     `gorm:"column:erp_warehouse_place" json:"erp_warehouse_place"`
	ERPSoTypeStock                 string     `gorm:"column:erp_so_type_stock" json:"erp_so_type_stock"`
	ERPSoTypeStockCode             string     `gorm:"column:erp_so_type_stock_code" json:"erp_so_type_stock_code"`
	ERPSeason                      string     `gorm:"column:erp_season" json:"erp_season"`
	IsRecycleSource                bool       `gorm:"column:is_recycle_source" json:"is_recycle_source"`
	POBuyerReference               *string    `gorm:"column:po_buyer_referrence" json:"po_buyer_referrence"`
	StdPrecision                   float64    `gorm:"column:std_precision" json:"std_precision"`
	ERPCOrderLineID                string     `gorm:"column:erp_c_order_line_id" json:"erp_c_order_line_id"`
	StockPurchaseID                string     `gorm:"column:stock_purchase_id" json:"stock_purchase_id"`
	QtyAllocatedPurchase           float64    `gorm:"column:qty_allocated_purchase" json:"qty_allocated_purchase"`
	MigrationID                    string     `gorm:"column:migration_id" json:"migration_id"`
	MigrationMaterialPreparationID string     `gorm:"column:migration_material_preparation_id" json:"migration_material_preparation_id"`
	MMNote                         string     `gorm:"column:mm_note" json:"mm_note"`
	ERPCStatisticalDate            time.Time  `gorm:"column:erp_statistical_date" json:"erp_statistical_date"`
	ETADate                        time.Time  `gorm:"column:eta_date" json:"eta_date"`
	ETDDate                        time.Time  `gorm:"column:etd_date" json:"etd_date"`
	ETAActual                      time.Time  `gorm:"column:eta_actual" json:"eta_actual"`
	ETDActual                      time.Time  `gorm:"column:etd_actual" json:"etd_actual"`
	QtyInHouse                     float64    `gorm:"column:qty_in_house" json:"qty_in_house"`
	QtyOnShip                      float64    `gorm:"column:qty_on_ship" json:"qty_on_ship"`
	QtyMRD                         float64    `gorm:"column:qty_mrd" json:"qty_mrd"`
	NoInvoice                      string     `gorm:"column:no_invoice" json:"no_invoice"`
	ReceiveDate                    time.Time  `gorm:"column:receive_date" json:"receive_date"`
	DDPI                           time.Time  `gorm:"column:dd_pi" json:"dd_pi"`
	IsClosing                      bool       `gorm:"column:is_closing" json:"is_closing"`
	MRD                            time.Time  `gorm:"column:mrd" json:"mrd"`
	ETADelay                       time.Time  `gorm:"column:eta_delay" json:"eta_delay"`
	QtyAllocatedDashboard          float64    `gorm:"column:qty_allocated_dashboard" json:"qty_allocated_dashboard"`
	RemarkDelay                    string     `gorm:"column:remark_delay" json:"remark_delay"`
	RemarkUpdate                   string     `gorm:"column:remark_update" json:"remark_update"`
	DateClosing                    time.Time  `gorm:"column:date_closing" json:"date_closing"`
	SOMigrationToBima              string     `gorm:"column:so_migration_to_bima" json:"so_migration_to_bima"`
	ItemSourceMigrationToBima      string     `gorm:"column:item_source_migration_to_bima" json:"item_source_migration_to_bima"`
	ItemBookMigrationToBima        string     `gorm:"column:item_book_migration_to_bima" json:"item_book_migration_to_bima"`
	IsCalculate                    bool       `gorm:"column:is_calculate" json:"is_calculate"`
	ERPBrand                       string     `gorm:"column:erp_brand" json:"erp_brand"`
	ERPStyleSO                     string     `gorm:"column:erp_style_so" json:"erp_style_so"`
	ERPTotalOrder                  string     `gorm:"column:erp_total_order" json:"erp_total_order"`
	ERPSupplierNameMRP             string     `gorm:"column:erp_supplier_name_mrp" json:"erp_supplier_name_mrp"`
	ERPSupplierCodeMRP             string     `gorm:"column:erp_supplier_code_mrp" json:"erp_supplier_code_mrp"`
	ERPDestination                 string     `gorm:"column:erp_destination" json:"erp_destination"`
	ERPFactoryCode                 string     `gorm:"column:erp_factory_code" json:"erp_factory_code"`
	CalculateDate                  time.Time  `gorm:"column:calculate_date" json:"calculate_date"`
	UnCalculateDate                time.Time  `gorm:"column:un_calculate_date" json:"un_calculate_date"`
	UserCalculateName              string     `gorm:"column:user_calculate_name" json:"user_calculate_name"`
	UserUnCalculateName            string     `gorm:"column:user_uncalculate_name" json:"user_uncalculate_name"`
	QtyOverMOQ                     float64    `gorm:"column:qty_over_moq" json:"qty_over_moq"`
	RejectMaterialQualityControlID string     `gorm:"column:reject_material_quality_control_id" json:"reject_material_quality_control_id"`
	IsForPiping                    string     `gorm:"column:is_for_piping" json:"is_for_piping"`
	Barcode                        string     `gorm:"column:barcode" json:"barcode"`
	SwitchDate                     time.Time  `gorm:"column:switch_date" json:"switch_date"`
}

// TableName sets the default table name to 'allocations'
func (Allocation) TableName() string {
	return "allocations"
}