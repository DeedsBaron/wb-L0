CREATE DATABASE test;
\c test;

CREATE TABLE delivery (
                          id      SERIAL PRIMARY KEY,

                          name    text NOT NULL,
                          phone   text NOT NULL,
                          zip     text NOT NULL,
                          city    text NOT NULL,
                          address text NOT NULL,
                          region  text NOT NULL,
                          email   text NOT NULL
);

CREATE TABLE payment (
                         id            SERIAL PRIMARY KEY,

                         transaction   text NOT NULL unique,
                         request_id    text NULL unique,
                         currency      text NOT NULL,
                         provider      text NOT NULL,
                         amount        integer NOT NULL,
                         payment_dt    integer NOT NULL,
                         bank          text NOT NULL,
                         delivery_cost integer NOT NULL,
                         goods_total   integer NOT NULL,
                         custom_fee    integer NOT NULL

);

CREATE TABLE items (
                       id                SERIAL PRIMARY KEY,
                       chrt_id           integer unique,

                       track_number text NOT NULL,
                       price        integer NOT NULL,
                       rid          text NOT NULL,
                       name         text NOT NULL,
                       sale         integer NOT NULL,
                       "size"       text NOT NULL,
                       total_price  integer NOT NULL,
                       nm_id        integer NOT NULL unique,
                       brand        text NOT NULL,
                       status       integer NOT NULL
);

CREATE TABLE wb (
                        id                SERIAL PRIMARY KEY,

                        order_uid          text unique,
                        track_number       text NOT NULL,
                        entry              text NOT NULL,

                        delivery_id        integer NOT NULL,
                        payment_id         integer NOT NULL,

                        locale             text NOT NULL,
                        internal_signature text NULL,
                        customer_id        text NOT NULL unique,
                        delivery_service   text NOT NULL,
                        shardkey           text NOT NULL,
                        sm_id              integer NOT NULL,
                        date_created       date NOT NULL,
                        oof_shard          text NOT NULL,

                        CONSTRAINT fk_delivery_id
                            FOREIGN KEY(delivery_id)
                                REFERENCES delivery(id)
                                ON DELETE CASCADE,
                        CONSTRAINT fk_payment_id
                            FOREIGN KEY(payment_id)
                                REFERENCES payment(id)
                                ON DELETE CASCADE
);


CREATE INDEX fk_delivery_id ON wb (
                                   delivery_id
);

CREATE INDEX fk_payment_id ON wb (
                                  payment_id
);

CREATE TABLE wb_items (
                          id        SERIAL PRIMARY KEY,

                          order_uid text NOT NULL,
                          chrt_id   integer NOT NULL,

                          CONSTRAINT fk_order_uid
                              FOREIGN KEY(order_uid)
                                  REFERENCES wb(order_uid)
                                  ON DELETE CASCADE,
                          CONSTRAINT fk_chrt_id
                              FOREIGN KEY(chrt_id)
                                  REFERENCES items(chrt_id)
                                  ON DELETE CASCADE
);

CREATE INDEX fk_chrt_id ON wb_items (
     chrt_id
);
CREATE INDEX fk_order_uid ON wb_items (
     order_uid
);