package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type OrderDetailsView struct {
	ID                     string            `json:"id,omitempty"`
	Hash                   string            `json:"hash"`
	ShippingCharge         int               `json:"shipping_charge"`
	PaymentProcessingFee   int               `json:"payment_processing_fee"`
	SubTotal               int               `json:"sub_total"`
	PaymentGateway         string            `json:"payment_gateway"`
	Nonce                  *string           `json:"nonce,omitempty"`          // Private
	TransactionID          *string           `json:"transaction_id,omitempty"` //Private
	GrandTotal             int               `json:"grand_total"`
	Status                 OrderStatus       `json:"status"`
	PaymentStatus          PaymentStatus     `json:"payment_status"`
	CreatedAt              *time.Time        `json:"created_at"`
	UpdatedAt              *time.Time        `json:"updated_at"`
	ShippingID             *string           `json:"shipping_id,omitempty"`
	ShippingName           *string           `json:"shipping_name,omitempty"`
	ShippingAddress        *string           `json:"shipping_address,omitempty"`
	ShippingCity           *string           `json:"shipping_city,omitempty"`
	ShippingCountry        *string           `json:"shipping_country,omitempty"`
	ShippingPostcode       *string           `json:"shipping_postcode,omitempty"`
	ShippingEmail          *string           `json:"shipping_email,omitempty"`
	ShippingPhone          *string           `json:"shipping_phone,omitempty"`
	BillingID              string            `json:"billing_id"`
	BillingName            string            `json:"billing_name"`
	BillingAddress         string            `json:"billing_address"`
	BillingCity            string            `json:"billing_city"`
	BillingCountry         string            `json:"billing_country"`
	BillingPostcode        string            `json:"billing_postcode"`
	BillingEmail           string            `json:"billing_email"`
	BillingPhone           string            `json:"billing_phone"`
	StoreID                string            `json:"store_id"`
	StoreName              string            `json:"store_name"`
	StoreAddress           string            `json:"store_address"`
	StoreCity              string            `json:"store_city"`
	StoreCountry           string            `json:"store_country"`
	StorePostcode          string            `json:"store_postcode"`
	StoreEmail             string            `json:"store_email"`
	StorePhone             string            `json:"store_phone"`
	StoreStatus            string            `json:"store_status"`
	PaymentMethodID        string            `json:"payment_method_id"`
	PaymentMethodName      string            `json:"payment_method_name"`
	PaymentMethodIsOffline bool              `json:"payment_method_is_offline"`
	Items                  []OrderedItemView `json:"items"`
	UserID                 string            `json:"user_id"`
	UserName               string            `json:"user_name"`
	UserEmail              string            `json:"user_email"`
	UserPhone              *string           `json:"user_phone,omitempty"`
	UserPicture            *string           `json:"user_picture,omitempty"`
}

func (odv *OrderDetailsView) TableName() string {
	return "order_details_views"
}

func (odv *OrderDetailsView) CreateView(tx *gorm.DB) error {
	sql := fmt.Sprintf("CREATE OR REPLACE VIEW %s AS SELECT o.id AS id, o.hash AS hash, o.user_id AS user_id,"+
		" u.name AS user_name, u.email AS user_email, u.phone AS user_phone, u.profile_picture AS user_picture,"+
		" o.shipping_charge AS shipping_charge, o.payment_processing_fee AS payment_processing_fee, o.sub_total AS sub_total,"+
		" o.payment_gateway AS payment_gateway, o.nonce AS nonce, o.transaction_id AS transaction_id, o.grand_total AS grand_total,"+
		" o.status AS status, o.payment_status AS payment_status, o.created_at AS created_at,"+
		" o.updated_at AS updated_at, sa.id AS shipping_id, sa.name AS shipping_name, sa.address AS shipping_address,"+
		" sa.city AS shipping_city, sa.country AS shipping_country, sa.postcode AS shipping_postcode,"+
		" sa.email AS shipping_email, sa.phone AS shipping_phone, ba.id AS billing_id, ba.name AS billing_name, ba.address AS billing_address,"+
		" ba.city AS billing_city, ba.country AS billing_country, ba.postcode AS billing_postcode, ba.email AS billing_email,"+
		" ba.phone AS billing_phone, s.id AS store_id, s.name AS store_name, s.address AS store_address, s.city AS store_city,"+
		" s.country AS store_country, s.postcode AS store_postcode, s.email AS store_email, s.phone AS store_phone, s.status AS store_status,"+
		" pm.id AS payment_method_id, pm.name AS payment_method_name, pm.is_offline_payment AS payment_method_is_offline"+
		" FROM orders AS o"+
		" LEFT JOIN addresses AS sa ON o.shipping_address_id = sa.id"+
		" LEFT JOIN addresses AS ba ON o.billing_address_id = ba.id"+
		" LEFT JOIN stores AS s ON o.store_id = s.id"+
		" LEFT JOIN users AS u ON o.user_id = u.id"+
		" LEFT JOIN payment_methods AS pm ON o.payment_method_id = pm.id;", odv.TableName())
	if err := tx.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func (odv *OrderDetailsView) DropView(tx *gorm.DB) error {
	sql := fmt.Sprintf("DROP VIEW %s", odv.TableName())

	if err := tx.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}