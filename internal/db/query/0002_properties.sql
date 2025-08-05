-- name: CreateProperty :one
insert into properties (
    agent_id, title, description, price, address, city, state,
    bedrooms, bathrooms, square_feet, property_type, 
    images, property_type, sale_type
) VALUES (
   @agent_id, @title, @description, @price, @address, @city, @state,
   @bedrooms, @bathrooms, @square_feet, @property_type,
   @images, @property_type, @sale_type
) RETURNING *;

-- name: GetProperties :many
select 
    id, property_type, images, title, description, price, 
    address, city, state, bedrooms, bathrooms, square_feet, created_at
from properties
where 
    (@name = 'none' or title ilike '%' || @name || '%')
    and (@city = 'none' or city = @city)
    and (@state = 'none' or state = @state)
    and (@min_price = 0 or price >= @min_price)
    and (@max_price = 0 or price <= @max_price)
    and (@property_type::text = 'none' or property_type::text = @property_type::text)
    and (@sale_type::text = 'none' or sale_type::text = @sale_type::text)
order by created_at desc
offset @skip limit 15;

-- name: GetPropertyInfo :one
select 
    p.id, p.property_type, p.images, p.title, p.description, p.price, 
    p.address, p.city, p.state, p.bedrooms, p.bathrooms, p.square_feet, p.created_at,
    json_build_object(
        'id', a.id,
        'picture', a.picture,
        'name', a.name,
        'description', a.description,
        'phone_number', a.phone_number,
        'email', a.email
    ) as agent
from properties p
inner join user_agents a on p.agent_id = a.id
where p.id = @id;
