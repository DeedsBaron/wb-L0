# :heavy_check_mark: wb-L0
Cервис получения модели, используя nats-streaming, добавления модели в бд, а также просмотра
информации модели по запросу
# Сервис получает от nats-streaming модель следующего вида:
 ```json
  {
  "order_uid": "b563feb7b2b84b6test",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
  ```

# Сервис принимает следующие запросы по http:
1. Метод `GET`, который возвращает информацию о моделе по id
----
* **URL**: /{id}
*  **URL Params**: id
* **Data Params**: None
* **Success Response:**
  * **Code:** 200 <br />
    **Content:**![image](https://user-images.githubusercontent.com/80648065/161759909-21daa44a-2f37-46c2-8acd-c83a07a464e9.png)

    
  OR
  
  * **Code:** 400 <br />
    **Content:** <br /> ![image](https://user-images.githubusercontent.com/80648065/161760165-7fea1b2b-5898-486f-9335-236be787e2f9.png)

# Конфиг
Конфиг находиться по умолчанию configs/

![image](https://user-images.githubusercontent.com/80648065/161534453-da3fb8d4-4172-48f6-abb8-3c3b23a9d241.png)
# Флаги
![image](https://user-images.githubusercontent.com/80648065/161534720-45f962d7-3f3b-4026-ae58-10af32f7868f.png)
# Схема базы данных
![image](https://user-images.githubusercontent.com/80648065/161535043-5321106b-9bf8-49ad-bf32-72c4b1bed38c.png)
# Usage
По умолчанию собирается и запускается весь кластер

    make

Для того чтобы опубликовать в канал сообщение, файл должен находиться nats-streaming-publish/

    make pub FILE="filename.json"

    
## Other
**Author:**  
:vampire:*[Deeds Baron](https://github.com/DeedsBaron)*  
