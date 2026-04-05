-- 👤 USERS
INSERT INTO users (id, name, email, password_hash, role, created_at, updated_at, deleted_at) VALUES
    ('52087336-befc-4d15-8e55-7f01637389fe', 'Luthfi Indrawan', 'luthfiindrawan@gmail.com', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mqrq3QfW3Zd8z1B9ZV3KqL8mN5Y4X2a', 'admin', '2026-03-06T06:55:56.510776', NULL, NULL),
    ('fafea7be-3802-42f4-8166-d14eacd8fa22', 'Dewi Lestari', 'dewilestarigmail.com', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mqrq3QfW3Zd8z1B9ZV3KqL8mN5Y4X2b', 'staff', '2026-03-11T06:55:56.510776', '2026-03-31T06:55:56.510776', NULL),
    ('87049d27-53d8-4cc4-a35a-b54fe6fb7653', 'Ahmad Rizky', 'ahmadrizky@gmail.com', '$2a$10$N9qo8uLOickgx2ZMRZoMy.Mqrq3QfW3Zd8z1B9ZV3KqL8mN5Y4X2c', 'borrower', '2026-03-26T06:55:56.510776', NULL, NULL);

-- 🏷️ CATEGORIES
INSERT INTO categories (id, name, created_at, updated_at) VALUES
    ('ded63ddc-5e2f-40dc-b71d-3fa125e48509', 'Power Tools', '2026-02-04T06:55:56.510776', NULL),
    ('8d759af2-36b1-48fb-8d43-2842b6e29e05', 'Hand Tools', '2026-02-09T06:55:56.510776', NULL),
    ('451c160e-b5a4-4277-9ab0-67364e13f82f', 'Measuring Instruments', '2026-02-14T06:55:56.510776', '2026-03-26T06:55:56.510776'),
    ('f27ecaee-a294-4590-94fa-8a976b8c9d31', 'Safety Equipment', '2026-02-19T06:55:56.510776', NULL),
    ('c3b6bac9-cb57-42fe-b74e-ccfd0fdabb78', 'Cleaning Supplies', '2026-02-24T06:55:56.510776', NULL);

-- 🧰 TOOLS
INSERT INTO tools (id, name, category_id, stock, description, created_at, updated_at, deleted_at) VALUES
    ('24fc1295-bc4d-4f8d-bcbf-f5d268771a1a', 'Bosch GSB 550 Professional Impact Drill', 'ded63ddc-5e2f-40dc-b71d-3fa125e48509', 5, 'Mesin bor tembok 550W dengan fitur impact untuk pengeboran beton dan tembok. Dilengkapi dengan speed control variable.', '2026-03-06T06:55:56.510776', NULL, NULL),
    ('71986c01-04ea-4a75-99c4-3547db3ef877', 'Makita MT M0901B Angle Grinder 4"', 'ded63ddc-5e2f-40dc-b71d-3fa125e48509', 3, 'Gerinda tangan 540W dengan kecepatan 11,000 RPM. Cocok untuk memotong dan menggerinda besi, baja, dan material keras lainnya.', '2026-03-08T06:55:56.510776', '2026-04-03T06:55:56.510776', NULL),
    ('63b3dc81-46f2-42b6-90b4-4889ee2724ed', 'Tekiro Obeng Set 6 Pcs', '8d759af2-36b1-48fb-8d43-2842b6e29e05', 12, 'Set obeng lengkap dengan 3 obeng plus (+) dan 3 obeng minus (-). Handle ergonomis dengan magnet pada ujung.', '2026-03-11T06:55:56.510776', NULL, NULL),
    ('2677a665-810c-4b20-8e8f-be4eadf69593', 'Stanley Claw Hammer 16 oz', '8d759af2-36b1-48fb-8d43-2842b6e29e05', 8, 'Palu karpenter dengan kepala baja carbon steel dan handle fiberglass. Berat 16 oz, cocok untuk paku kayu umum.', '2026-03-12T06:55:56.510776', NULL, NULL),
    ('a6ea62ae-5753-4967-b8d8-3629e532393b', 'Lippro Waterpass 40 cm Aluminum', '451c160e-b5a4-4277-9ab0-67364e13f82f', 6, 'Waterpass aluminum 40 cm dengan 3 bubble level (horizontal, vertical, 45°). Presisi tinggi untuk pekerjaan konstruksi.', '2026-03-16T06:55:56.510776', NULL, NULL),
    ('a0df75b2-c9bd-4dc6-b8a4-60aceacbbe9b', 'Mitutoyo Digital Caliper 150mm', '451c160e-b5a4-4277-9ab0-67364e13f82f', 4, 'Jangka sorong digital presisi 0.01mm dengan range 0-150mm. Layar LCD jernih dengan fitur zero setting.', '2026-03-18T06:55:56.510776', '2026-03-31T06:55:56.510776', NULL),
    ('b0fad0ed-218b-41e2-a3d1-104165947b2e', '3M 8210 N95 Particulate Respirator Mask', 'f27ecaee-a294-4590-94fa-8a976b8c9d31', 50, 'Masker pelindung N95 untuk melindungi dari debu halus dan partikel. Cocok untuk pekerjaan grinding dan sanding.', '2026-03-21T06:55:56.510776', NULL, NULL),
    ('03d3c99d-a41e-433e-8ae7-af82dae91bf3', 'Kings Safety Glasses KW220', 'f27ecaee-a294-4590-94fa-8a976b8c9d31', 20, 'Kacamata safety dengan lensa polycarbonate anti-fog dan anti-scratch. Frame nyaman dengan side shield protection.', '2026-03-22T06:55:56.510776', NULL, NULL),
    ('c1736ebc-d1de-40b1-acb2-84a70fbe7a09', 'Scot-Brite Heavy Duty Scrub Sponge', 'c3b6bac9-cb57-42fe-b74e-ccfd0fdabb78', 30, 'Spons cuci heavy duty dengan sisi abrasive untuk membersihkan noda membandel pada peralatan kerja.', '2026-03-24T06:55:56.510776', NULL, NULL),
    ('df941365-d0e8-49d9-8ca6-4d510cf5a621', 'Tora Microfiber Cleaning Cloth 3 Pcs', 'c3b6bac9-cb57-42fe-b74e-ccfd0fdabb78', 25, 'Kain lap microfiber serat halus untuk membersihkan debu dan kotoran tanpa meninggalkan goresan. Isi 3 pcs warna-warni.', '2026-03-26T06:55:56.510776', NULL, NULL);