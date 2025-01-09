package models

import "time"

type Buyer struct {
	ID                      string     `gorm:"type:char(36);primaryKey" json:"id"`
	PurchaseDocumentTypeID  *string    `gorm:"type:varchar(191)" json:"purchase_document_type_id"`
	PurchaseNumberCode      *string    `gorm:"type:varchar(191)" json:"purchase_number_code"`
	BuyerName               *string    `gorm:"type:varchar(191)" json:"buyer_name"`
	TotalMaximumPointDefect *int       `gorm:"type:int" json:"total_maximum_point_defect"`
	CreatedAt               *time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt               *time.Time `gorm:"type:timestamp" json:"updated_at"`
	FormulaFIR              *string    `gorm:"type:varchar(191);default:'general'" json:"formula_fir"`
	StockType               *string    `gorm:"type:varchar(191)" json:"stock_type"`
	ERPClientOrgID          *string    `gorm:"type:varchar(191)" json:"erp_client_org_id"`
	DeletedAt               *time.Time `gorm:"type:timestamp" json:"deleted_at"`
	IsShow                  bool       `gorm:"type:boolean;default:true" json:"is_show"`
}
