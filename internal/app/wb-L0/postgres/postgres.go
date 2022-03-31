package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"strconv"
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

func AddPayment() (int, error) {
	var fkPay int

	q := `INSERT INTO payment (transaction, request_id, currency, provider, 
			amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id;`
	err := Pool.QueryRow(context.Background(), q,
		storage.Model.Payment.Transaction,
		storage.Model.Payment.RequestID,
		storage.Model.Payment.Currency,
		storage.Model.Payment.Provider,
		storage.Model.Payment.Amount,
		storage.Model.Payment.PaymentDt,
		storage.Model.Payment.Bank,
		storage.Model.Payment.DeliveryCost,
		storage.Model.Payment.GoodsTotal,
		storage.Model.Payment.CustomFee,
	).Scan(&fkPay)
	if err != nil {
		return 0, err
	}
	logger.Log.Debug("Successfully added info to Payment table")
	return fkPay, nil
}

func AddDelivery() (int, error) {
	var fkDel int

	q := `INSERT INTO delivery (name, phone, zip, city, 
			address, region, email)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id;`
	err := Pool.QueryRow(context.Background(), q,
		storage.Model.Delivery.Name,
		storage.Model.Delivery.Phone,
		storage.Model.Delivery.Zip,
		storage.Model.Delivery.City,
		storage.Model.Delivery.Address,
		storage.Model.Delivery.Region,
		storage.Model.Delivery.Email,
	).Scan(&fkDel)
	if err != nil {
		return 0, err
	}
	logger.Log.Debug("Successfully added info to Delivery table")
	return fkDel, nil
}

func AddItems() ([]int, error) {
	var fkItems []int
	var fkItem int

	for i, val := range storage.Model.Items {
		q := `INSERT INTO items (chrt_id, track_number, price, rid, 
			name, sale, size, total_price, nm_id, brand, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id;`
		err := Pool.QueryRow(context.Background(), q,
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
		).Scan(&fkItem)
		if err != nil {
			logger.Log.Error("Wrong Item[", strconv.Itoa(i), "] :", err)
			return fkItems, err
		}
		fkItems = append(fkItems, fkItem)
	}
	logger.Log.Debug("Successfully added info to Items table")
	return fkItems, nil
}

func AddWb(fkPay, fkDel int) (int, error) {
	var fkWb int

	q := `INSERT INTO wb (order_uid, track_number ,entry, delivery_id,
			payment_id , locale, internal_signature,
			customer_id, delivery_service, shardkey,
			sm_id, date_created ,oof_shard)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id;`
	err := Pool.QueryRow(context.Background(), q,
		storage.Model.OrderUID,
		storage.Model.TrackNumber,
		storage.Model.Entry,
		fkDel,
		fkPay,
		storage.Model.Locale,
		storage.Model.InternalSignature,
		storage.Model.CustomerID,
		storage.Model.DeliveryService,
		storage.Model.Shardkey,
		storage.Model.SmID,
		storage.Model.DateCreated,
		storage.Model.OofShard,
	).Scan(&fkWb)
	if err != nil {
		return 0, err
	}
	logger.Log.Debug("Successfully added info to Wb table")
	return fkWb, nil
}

func AddToDb() {
	fkPay, err := AddPayment()
	fmt.Println("fkPay = ", fkPay)
	if err != nil {
		logger.Log.Error(err)
		return
	}

	fkDel, err := AddDelivery()
	fmt.Println("fkPay = ", fkPay)
	if err != nil {
		logger.Log.Error(err)
		return
	}

	fkItems, err := AddItems()
	if err != nil {
		logger.Log.Error("Wrong Item[", strconv.Itoa(len(fkItems)-1), "] :", err)
		return
	}
	fkWb, err := AddWb(fkPay, fkDel)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	fmt.Println("fkPay = ", fkPay, "fkDel = ", fkDel, "fkItems = ", fkItems, "fkWb = ", fkWb)
}
