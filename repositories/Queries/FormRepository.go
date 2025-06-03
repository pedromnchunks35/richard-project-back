package queries

const (
	// ====================
	// FORM QUERIES
	// ====================

	QueryInsertForm = `
		INSERT INTO Form (
			name, company, nif, phone_number, email,
			fiscal_address, address, with_delivery,
			observations, id_intervention, id_form_products
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8,
			$9, $10, $11
		) RETURNING id;
	`

	QueryUpdateForm = `
		UPDATE Form SET
			name = $1,
			company = $2,
			nif = $3,
			phone_number = $4,
			email = $5,
			fiscal_address = $6,
			address = $7,
			with_delivery = $8,
			observations = $9,
			id_intervention = $10,
			id_form_products = $11
		WHERE id = $12;
	`

	QueryDeleteForm = `
		DELETE FROM Form WHERE id = $1;
	`

	QueryGetFormById = `
		SELECT * FROM Form WHERE id = $1;
	`

	QueryGetAllForms = `
		SELECT * FROM Form;
	`

	// ==========================
	// FORM PRODUCTS QUERIES
	// ==========================

	QueryInsertFormProduct = `
		INSERT INTO FormProducts (
			id_form, id_product, id_image_attachment,
			quantity, x, y, z
		) VALUES (
			$1, $2, $3,
			$4, $5, $6, $7
		) RETURNING id;
	`

	QueryUpdateFormProduct = `
		UPDATE FormProducts SET
			id_form = $1,
			id_product = $2,
			id_image_attachment = $3,
			quantity = $4,
			x = $5,
			y = $6,
			z = $7
		WHERE id = $8;
	`

	QueryDeleteFormProduct = `
		DELETE FROM FormProducts WHERE id = $1;
	`

	QueryGetFormProductById = `
		SELECT * FROM FormProducts WHERE id = $1;
	`

	QueryGetAllFormProducts = `
		SELECT * FROM FormProducts;
	`

	// ==========================
	// JOINED
	// ==========================

	QueryGetFormWithProducts = `
		SELECT 
			f.*, 
			fp.id AS form_product_id,
			fp.id_product,
			fp.id_image_attachment,
			fp.quantity,
			fp.x, fp.y, fp.z
		FROM Form f
		LEFT JOIN FormProducts fp ON f.id = fp.id_form
		WHERE f.id = $1;
	`

	QueryGetAllFormsWithProductCount = `
		SELECT 
			f.id,
			f.name,
			f.company,
			COUNT(fp.id) AS total_products
		FROM Form f
		LEFT JOIN FormProducts fp ON f.id = fp.id_form
		GROUP BY f.id, f.name, f.company;
	`

	QueryGetAllFormProductsWithFormInfo = `
		SELECT 
			fp.*, 
			f.name AS form_name,
			f.company
		FROM FormProducts fp
		INNER JOIN Form f ON f.id = fp.id_form;
	`
)
