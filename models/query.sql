CREATE TABLE IF NOT EXISTS carts(
    ID BIGSERIAL,
	sku varchar(50) not null,
	quantity int not null,
	status int not null,
    PRIMARY KEY (ID)
)

CREATE TABLE IF NOT EXISTS products(
	ID BIGSERIAL,
	sku varchar(100) not null unique,
	"name" varchar(100) not null,
	price float not null,
	quantity int not null,
	PRIMARY KEY (ID)
)

CREATE TABLE IF NOT EXISTS promo(
	ID BIGSERIAL,
	sku varchar(100) not null unique,
	promo_type varchar(100) not null,
	bonus_sku varchar(100) null,
	min_qty int null,
	pay_only int null,
	discount int null,
	PRIMARY KEY (ID)
)