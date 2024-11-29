SELECT 
    c.cus_firstname AS ชื่อ,
    c.cus_lastname AS นามสกุล,
    o.order_id AS รหัสสั่งซื้อ,
    o.order_date AS วันที่สั่งซื้อ,
    o.order_total AS จำนวนสินค้าที่สั่ง
FROM 
    tbl_customer c
JOIN 
    tbl_orders o
ON 
    c.cus_id = o.cus_id
WHERE 
    o.order_total >= 500
ORDER BY 
    o.order_date;

