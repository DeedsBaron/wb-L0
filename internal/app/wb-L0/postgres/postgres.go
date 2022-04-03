package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"time"
	"wb-L0/internal/app/wb-L0/config"
	"wb-L0/internal/app/wb-L0/logger"
	"wb-L0/internal/app/wb-L0/storage"
	"wb-L0/pkg/utils"
)

var (
	Pool *pgxpool.Pool
)

func init() {
	Pool = NewPostgres()
}

func NewPostgres() *pgxpool.Pool {
	Pool, err := NewClient(context.Background(), &config.Config)
	if err != nil {
		logrus.Fatal(err)
	}
	return Pool
}

func NewClient(ctx context.Context, config *config.Cfg) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.Storage.Username,
		config.Storage.Password,
		config.Storage.Host,
		config.Storage.Port,
		config.Storage.Database)
	var pool *pgxpool.Pool

	err := utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		var err error
		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, config.Storage.Attempts, 5*time.Second)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	logrus.Info("Successfully connected to database")
	return pool, nil
}

func AddPayment(model *storage.ModelJSON) (int, error) {
	var fkPay int

	q := `INSERT INTO payment (transaction, request_id, currency, provider, 
			amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id;`
	err := Pool.QueryRow(context.Background(), q,
		model.Payment.Transaction,
		model.Payment.RequestID,
		model.Payment.Currency,
		model.Payment.Provider,
		model.Payment.Amount,
		model.Payment.PaymentDt,
		model.Payment.Bank,
		model.Payment.DeliveryCost,
		model.Payment.GoodsTotal,
		model.Payment.CustomFee,
	).Scan(&fkPay)
	if err != nil {
		return 0, err
	}
	logger.Log.Debug("Successfully added info to Payment table")
	return fkPay, nil
}

func AddDelivery(model *storage.ModelJSON) (int, error) {
	var fkDel int

	q := `INSERT INTO delivery (name, phone, zip, city, 
			address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;`
	err := Pool.QueryRow(context.Background(), q,
		model.Delivery.Name,
		model.Delivery.Phone,
		model.Delivery.Zip,
		model.Delivery.City,
		model.Delivery.Address,
		model.Delivery.Region,
		model.Delivery.Email,
	).Scan(&fkDel)
	if err != nil {
		return 0, err
	}
	logger.Log.Debug("Successfully added info to Delivery table")
	return fkDel, nil
}

func AddItems(model *storage.ModelJSON) error {
	for _, val := range model.Items {
		q := `INSERT INTO items (chrt_id, track_number, price, rid, 
			name, sale, size, total_price, nm_id, brand, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id;`
		_, err := Pool.Exec(context.Background(), q,
			val.ChrtID,
			val.TrackNumber,
			val.Price,
			val.Rid,
			val.Name,
			val.Sale,
			val.Size,
			val.TotalPrice,
			val.NmID,
			val.Brand,
			val.Status,
		)
		if err != nil {
			return err
		}
	}
	logger.Log.Debug("Successfully added info to Items table")
	return nil
}

func AddWb(fkPay, fkDel int, model *storage.ModelJSON) (int, error) {
	var fkWb int

	q := `INSERT INTO wb (order_uid, track_number ,entry, delivery_id,
			payment_id , locale, internal_signature,
			customer_id, delivery_service, shardkey,
			sm_id, date_created ,oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id;`
	err := Pool.QueryRow(context.Background(), q,
		model.OrderUID,
		model.TrackNumber,
		model.Entry,
		fkDel,
		fkPay,
		model.Locale,
		model.InternalSignature,
		model.CustomerID,
		model.DeliveryService,
		model.Shardkey,
		model.SmID,
		model.DateCreated,
		model.OofShard,
	).Scan(&fkWb)
	if err != nil {
		return 0, err
	}
	logger.Log.Debug("Successfully added info to Wb table")
	return fkWb, nil
}

func AddWbItems(model *storage.ModelJSON) error {
	for _, val := range model.Items {
		q := `INSERT INTO wb_items (order_uid, chrt_id)
		VALUES ($1, $2)`
		_, err := Pool.Exec(context.Background(), q, model.OrderUID, val.ChrtID)
		if err != nil {
			return err
		}
	}
	logger.Log.Debug("Successfully added info to Wb_Items table")
	return nil
}

func AddToDb(model *storage.ModelJSON) {
	fkPay, err := AddPayment(model)
	fmt.Println("fkPay = ", fkPay)
	if err != nil {
		logger.Log.Error(err)
		return
	}

	fkDel, err := AddDelivery(model)
	fmt.Println("fkPay = ", fkPay)
	if err != nil {
		logger.Log.Error(err)
		return
	}

	err = AddItems(model)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	fkWb, err := AddWb(fkPay, fkDel, model)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	err = AddWbItems(model)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	fmt.Println("fkPay = ", fkPay, "fkDel = ", fkDel, "fkWb = ", fkWb)
}

func GetWbTable() error {
	q := `SELECT wb.order_uid, wb.track_number, wb.entry,
		   d.name, d.phone, d.zip, d.city, d.address, d.region, d.email,
		   p.transaction, p.request_id, p.currency, p.provider, p.amount,
		   p.payment_dt, p.bank, p.delivery_cost, p.goods_total, p.custom_fee,
		   wb.locale, wb.internal_signature, wb.customer_id, wb.customer_id,
		   wb.delivery_service, wb.shardkey, wb.sm_id, wb.date_created, wb.oof_shard
		FROM wb INNER JOIN delivery d on d.id = wb.delivery_id
		INNER JOIN payment p on p.id = wb.payment_id`
	rows, err := Pool.Query(context.Background(), q)
	if err != nil {
		return err
	}
}
