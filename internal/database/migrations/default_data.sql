INSERT INTO categories (name) VALUES
    ('produce'),
    ('dairy'),
    ('bread&bakery'),
    ('meat&poultry');

INSERT INTO grocery_items (name, unit_price, stock, category_id) VALUES
    ('Butter',2.5,60,2),
    ('Banana',0.3,150,1),
    ('Eggs',3,80,4),
    ('Flour',1.5,110,3),
    ('Cucumber',0.9,100,1),
    ('Chicken Breast',5,40,4),
    ('Rice',1.8,120,3),
    ('Broccoli',1.5,70,1),
    ('Milk',2,50,2),
    ('Oats',2,85,3),
    ('Potato',0.8,110,1),
    ('Bacon',4.5,60,4),
    ('Sour Cream',1.3,60,2),
    ('Beef Steak',6,50,4),
    ('Apple',0.5,100,1),
    ('Sausage',4.8,65,4),
    ('Pasta',1.2,90,3),
    ('Yogurt',1.8,70,2),
    ('Lettuce',1.2,80,1),
    ('Pork Chops',4,55,4),
    ('Cheese',3,40,2),
    ('Orange',0.7,120,1),
    ('Carrot',0.7,130,1),
    ('Salmon',7,30,4),
    ('Cereal',2.5,70,3),
    ('Ground Beef',5.5,45,4),
    ('Cream',1.5,55,2),
    ('Tomato',1,90,1),
    ('Onion',0.6,140,1)
;

INSERT INTO order_statuses (status) VALUES
    ('created'),
    ('completed'),
    ('cancelled')
    ;
