package queries

const (
	// ==========================
	// IMAGE
	// ==========================
	QueryGetImageByID = `
		SELECT id, name, base64
		FROM images
		WHERE id = %1
	`

	QueryInsertImage = `
		INSERT INTO images (name, base64)
		VALUES (%1, %2)
	`

	QueryUpdateImage = `
		UPDATE images
		SET name = %1, base64 = %2
		WHERE id = %3
	`

	QueryDeleteImage = `
		DELETE FROM images
		WHERE id = %1
	`
)
