-- type Product struct {
-- 	ProductId uuid.UUID
-- 	Name      string
-- 	Price     float64
-- }

create table products (
    product_id uuid primary key default gen_random_uuid(),
    name text,
    price float
);