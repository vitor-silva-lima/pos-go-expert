create table products (
    product_id uuid primary key default gen_random_uuid(),
    name text,
    price float
);