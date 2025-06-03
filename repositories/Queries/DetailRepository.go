package queries

const (

	// ==========================
	// DETAIL
	// ==========================

	QueryGetDetailByID = `
		SELECT id, name, value, description, id_agent
		FROM details
		WHERE id = %1
	`

	QueryInsertDetail = `
		INSERT INTO details (name, value, description, id_agent)
		VALUES (%1, %2, %3, %4)
	`

	QueryUpdateDetail = `
		UPDATE details
		SET name = %1, value = %2, description = %3, id_agent = %4
		WHERE id = %5
	`

	QueryDeleteDetail = `
		DELETE FROM details
		WHERE id = %1
	`
)
