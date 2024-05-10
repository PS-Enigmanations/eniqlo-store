-- USER SEEDER
INSERT INTO "public"."users" ("id", "name", "phone_number", "password", "created_at", "updated_at", "deleted_at")
VALUES
('7a5da8e5-8f41-4e22-bdb9-d3c63f0b2f6e', 'John Doe', '1234567890', '$2a$08$byUr0FmFYtz8zVp7RzsU8.ASjdwAKGAwL6n.nPU6J4g6VNpDx/utu', NOW(), NULL, NULL),
('b19ab0d0-0c49-47ff-8575-4a34a72b0e17', 'Jane Smith', '9876543210', '$2a$08$byUr0FmFYtz8zVp7RzsU8.ASjdwAKGAwL6n.nPU6J4g6VNpDx/utu', NOW(), NULL, NULL),
('f2bb0d18-8ef3-4d7a-a2fc-0744f13e32b7', 'Alice Johnson', '5551234567', '$2a$08$byUr0FmFYtz8zVp7RzsU8.ASjdwAKGAwL6n.nPU6J4g6VNpDx/utu', NOW(), NULL, NULL);

-- PRODUCT SEEDER
INSERT INTO "public"."products" ("id", "name", "sku", "category", "image_url", "notes", "price", "stock", "location", "is_available", "created_at", "updated_at", "deleted_at")
VALUES
('0f8fad5b-d9cb-469f-a165-70867728950e', 'Blue T-Shirt', 'TS001', 'Clothing', 'https://example.com/image1.jpg', 'Comfortable cotton material', 15.99, 50, 'Warehouse A', true, NOW(), NULL, NULL),
('7c9e6679-7425-40de-944b-e07fc1f90ae7', 'Black Leather Belt', 'BLB002', 'Accessories', 'https://example.com/image2.jpg', 'Genuine leather, adjustable buckle', 29.99, 30, 'Warehouse B', true, NOW(), NULL, NULL),
('8f14e45fceea167a5a36dedd4bea2543', 'Running Shoes', 'SH003', 'Footwear', 'https://example.com/image3.jpg', 'Breathable mesh, cushioned sole', 49.99, 100, 'Warehouse C', true, NOW(), NULL, NULL),
('c9f0f895fb98ab9159f51fd0297e236d', 'Green Tea', 'GT004', 'Beverages', 'https://example.com/image4.jpg', 'Organic green tea leaves', 7.99, 200, 'Warehouse D', true, NOW(), NULL, NULL),
('42343331-7dd2-4610-9893-491fa485745d', 'Red Hoodie', 'HD005', 'Clothing', 'https://example.com/image5.jpg', 'Warm and cozy, with a hood', 39.99, 20, 'Warehouse A', true, NOW(), NULL, NULL),
('18a6c3c8-38bd-4227-bf08-0852ffc76c1b', 'Silver Necklace', 'NK006', 'Accessories', 'https://example.com/image6.jpg', 'Elegant silver necklace', 19.99, 40, 'Warehouse B', true, NOW(), NULL, NULL),
('d37a7fb8-38e3-4427-b43f-1298daf56eff', 'Running Shorts', 'SH007', 'Clothing', 'https://example.com/image7.jpg', 'Lightweight and breathable', 24.99, 60, 'Warehouse C', true, NOW(), NULL, NULL),
('062b7f43-f93f-4055-88ce-d2aadfed2acf', 'Coffee Mug', 'MG008', 'Accessories', 'https://example.com/image8.jpg', 'Durable ceramic mug', 9.99, 80, 'Warehouse D', true, NOW(), NULL, NULL),
('ff6c0953-e4af-452c-95d9-0e5f019c284b', 'White Sneakers', 'SN009', 'Footwear', 'https://example.com/image9.jpg', 'Classic white sneakers', 59.99, 90, 'Warehouse A', true, NOW(), NULL, NULL),
('f72554ef-ae1c-4103-a976-f5cd6f295f59', 'Chocolate Bar', 'CB010', 'Beverages', 'https://example.com/image10.jpg', 'Rich and creamy chocolate', 4.99, 150, 'Warehouse B', true, NOW(), NULL, NULL);

-- CUSTOMER SEEDER
INSERT INTO "public"."customers" ("id", "name", "phone_number", "created_at", "updated_at", "deleted_at")
VALUES
('c5c0e5c7-62c5-4e6e-8f6d-09cc12345678', 'John Doe', '1234567890', NOW(), NULL, NULL),
('2c1b08f9-3c43-42d2-b0f2-b55e98765432', 'Jane Smith', '9876543210', NOW(), NULL, NULL),
('6f704b33-eb4b-49b1-af32-0ae655432167', 'Alice Johnson', '5551234567', NOW(), NULL, NULL);

-- TRANSACTION SEEDER
INSERT INTO "public"."transactions" ("id", "customer_id", "total", "paid", "change", "created_at", "updated_at")
VALUES
('33e19244-91a0-4e7b-90c8-41bf8e0514e1', 'c5c0e5c7-62c5-4e6e-8f6d-09cc12345678', 150.99, 160.00, 9.01, '2024-05-06 12:00:00', NULL),
('de9fb3b0-7b98-4c90-bd6b-35b8f75432fc', '2c1b08f9-3c43-42d2-b0f2-b55e98765432', 75.50, 80.00, 4.50, '2024-05-06 12:00:00', NULL),
('d4d7b6f2-ae92-4ee4-bfa7-f789f0d12345', '6f704b33-eb4b-49b1-af32-0ae655432167', 200.00, 200.00, 0, '2024-05-06 12:00:00', NULL);

-- TRANSACTION DETAIL SEEDER
INSERT INTO "public"."transaction_details" ("id", "transaction_id", "product_id", "quantity", "total", "created_at", "updated_at")
VALUES
('7b52009b-8b3c-4dc4-8f11-345bea012345', '33e19244-91a0-4e7b-90c8-41bf8e0514e1', '0f8fad5b-d9cb-469f-a165-70867728950e', 2, 31.98, '2024-05-06 12:00:00', NULL),
('60e2dece-087b-40ab-a4c9-7865dc987654', 'de9fb3b0-7b98-4c90-bd6b-35b8f75432fc', '7c9e6679-7425-40de-944b-e07fc1f90ae7', 1, 29.99, '2024-05-06 12:00:00', NULL),
('3f7b709b-41b2-4c5c-ae5f-1f669d012345', 'd4d7b6f2-ae92-4ee4-bfa7-f789f0d12345', '8f14e45fceea167a5a36dedd4bea2543', 3, 149.97, '2024-05-06 12:00:00', NULL);
