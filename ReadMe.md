## My template for starting a new golang app

```sql
INSERT INTO users (full_name, email, password)
VALUES (
    'Agent Charles',
    'charles@example.com',
    crypt('test12345', gen_salt('bf'))
);

```
