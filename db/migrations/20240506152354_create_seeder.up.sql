-- USER SEEDER
INSERT INTO "public"."users" ("id", "name", "phone_number", "password", "created_at", "updated_at", "deleted_at") 
VALUES 
('7a5da8e5-8f41-4e22-bdb9-d3c63f0b2f6e', 'John Doe', '1234567890', '$2a$08$byUr0FmFYtz8zVp7RzsU8.ASjdwAKGAwL6n.nPU6J4g6VNpDx/utu', NOW(), NULL, NULL),
('b19ab0d0-0c49-47ff-8575-4a34a72b0e17', 'Jane Smith', '9876543210', '$2a$08$byUr0FmFYtz8zVp7RzsU8.ASjdwAKGAwL6n.nPU6J4g6VNpDx/utu', NOW(), NULL, NULL),
('f2bb0d18-8ef3-4d7a-a2fc-0744f13e32b7', 'Alice Johnson', '5551234567', '$2a$08$byUr0FmFYtz8zVp7RzsU8.ASjdwAKGAwL6n.nPU6J4g6VNpDx/utu', NOW(), NULL, NULL);

-- PRODUCT SEEDER
INSERT INTO "public"."products" ("id", "name", "sku", "category", "image_url", "notes", "price", "stock", "location", "is_available", "created_at", "updated_at", "deleted_at") 
VALUES 
('0f8fad5b-d9cb-469f-a165-70867728950e', 'Blue T-Shirt', 'TS001', 'Clothing', 'https://example.com/image1.jpg', 'Comfortable cotton material', 15.99, 50, 'Warehouse A', true, '2024-05-06 12:00:00', NULL, NULL),
('7c9e6679-7425-40de-944b-e07fc1f90ae7', 'Black Leather Belt', 'BLB002', 'Accessories', 'https://example.com/image2.jpg', 'Genuine leather, adjustable buckle', 29.99, 30, 'Warehouse B', true, '2024-05-06 12:00:00', NULL, NULL),
('8f14e45fceea167a5a36dedd4bea2543', 'Running Shoes', 'SH003', 'Footwear', 'https://example.com/image3.jpg', 'Breathable mesh, cushioned sole', 49.99, 100, 'Warehouse C', true, '2024-05-06 12:00:00', NULL, NULL),
('c9f0f895fb98ab9159f51fd0297e236d', 'Green Tea', 'GT004', 'Beverages', 'https://example.com/image4.jpg', 'Organic green tea leaves', 7.99, 200, 'Warehouse D', true, '2024-05-06 12:00:00', NULL, NULL);
