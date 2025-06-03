package mapper

import (
	dto "richard-project-back/dtos"
	model "richard-project-back/repositories/Models"
)

// Auxiliares para procurar por ID
func findCategoryById(categories []model.Category, id int64) *model.Category {
	for _, c := range categories {
		if c.Id == id {
			return &c
		}
	}
	return nil
}

func findImageById(images []model.Image, id int64) *model.Image {
	for _, img := range images {
		if img.Id == id {
			return &img
		}
	}
	return nil
}

func findDetailById(details []model.Detail, id int64) *model.Detail {
	for _, d := range details {
		if d.Id == id {
			return &d
		}
	}
	return nil
}

func findAgentById(agents []model.Agent, id int64) *model.Agent {
	for _, a := range agents {
		if a.Id == id {
			return &a
		}
	}
	return nil
}

// Mapear Agent
func mapAgent(m *model.Agent) dto.Agent {
	if m == nil {
		return dto.Agent{}
	}
	return dto.Agent{
		Id:   uint(m.Id),
		Name: m.Name,
	}
}

// Mapear Detail
func mapDetail(m *model.Detail, agents []model.Agent) dto.Detail {
	if m == nil {
		return dto.Detail{}
	}

	agent := findAgentById(agents, m.IdAgent)

	return dto.Detail{
		Id:          uint(m.Id),
		Name:        m.Name,
		Value:       m.Value,
		Description: m.Description,
		Agent:       mapAgent(agent),
	}
}

// Mapear Category
func mapCategory(m *model.Category) dto.Category {
	if m == nil {
		return dto.Category{}
	}
	return dto.Category{
		Id:   uint(m.Id),
		Name: m.Name,
	}
}

// Mapear Image
func mapImage(m *model.Image) dto.Image {
	if m == nil {
		return dto.Image{}
	}
	return dto.Image{
		Id:     uint(m.Id),
		Name:   m.Name,
		Base64: m.Base64,
	}
}

// Mapear Product
func mapProduct(m *model.Product, categories []model.Category, images []model.Image, details []model.Detail, agents []model.Agent) dto.Product {
	if m == nil {
		return dto.Product{}
	}

	category := findCategoryById(categories, m.IdCategory)
	image := findImageById(images, m.IdImage)
	detail := findDetailById(details, m.IdDetail)

	return dto.Product{
		Id:       uint(m.Id),
		Name:     m.Name,
		Category: mapCategory(category),
		Image:    mapImage(image),
		Detatil:  mapDetail(detail, agents), // Corrigi typo para Detail
	}
}
