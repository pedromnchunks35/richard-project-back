package queries

const (
	// ==========================
	// PRODUCT
	// ==========================
	QueryGetProductByID = `
		SELECT id, name, id_category, id_image
		FROM products
		WHERE id = %1
	`

	QueryInsertProduct = `
		INSERT INTO products (name, id_category, id_image)
		VALUES (%1, %2, %3)
	`

	QueryUpdateProduct = `
		UPDATE products
		SET name = %1, id_category = %2, id_image = %3
		WHERE id = %4
	`

	QueryDeleteProduct = `
		DELETE FROM products
		WHERE id = %1
 `
	// ==========================
	// JOINED
	// ==========================

	QueryGetProductWithDetailsAndImage = `
		SELECT 
			p.id as product_id, p.name as product_name, p.id_category, p.id_image,
			d.id as detail_id, d.name as detail_name, d.value as detail_value, d.description as detail_description, d.id_agent,
			i.id as image_id, i.name as image_name, i.base64 as image_base64
		FROM products p
		LEFT JOIN product_details pd ON p.id = pd.id_product
		LEFT JOIN details d ON pd.id_detail = d.id
		LEFT JOIN images i ON p.id_image = i.id
		WHERE p.id = %1
	`

	QueryGetAllProductsWithDetailsAndImage = `
		SELECT 
			p.id as product_id, p.name as product_name, p.id_category, p.id_image,
			d.id as detail_id, d.name as detail_name, d.value as detail_value, d.description as detail_description, d.id_agent,
			i.id as image_id, i.name as image_name, i.base64 as image_base64
		FROM products p
		LEFT JOIN product_details pd ON p.id = pd.id_product
		LEFT JOIN details d ON pd.id_detail = d.id
		LEFT JOIN images i ON p.id_image = i.id
	`

	QueryGetDetailsByProductID = `
		SELECT 
			d.id, d.name, d.value, d.description, d.id_agent
		FROM product_details pd
		INNER JOIN details d ON pd.id_detail = d.id
		WHERE pd.id_product = %1
	`
)
