package services

import (
	"catalyst-players/internal/domain/entities"
	"testing"
)

// MockTagRepository is a mock implementation of TagRepository for testing
type MockTagRepository struct {
	tags map[uint]*entities.Tag
	nextID uint
}

// NewMockTagRepository creates a new mock tag repository
func NewMockTagRepository() *MockTagRepository {
	return &MockTagRepository{
		tags: make(map[uint]*entities.Tag),
		nextID: 1,
	}
}

func (m *MockTagRepository) Create(tag *entities.Tag) error {
	tag.ID = m.nextID
	m.tags[tag.ID] = tag
	m.nextID++
	return nil
}

func (m *MockTagRepository) GetByID(id uint) (*entities.Tag, error) {
	if tag, exists := m.tags[id]; exists {
		return tag, nil
	}
	return nil, nil
}

func (m *MockTagRepository) GetAll() ([]entities.Tag, error) {
	tags := make([]entities.Tag, 0, len(m.tags))
	for _, tag := range m.tags {
		tags = append(tags, *tag)
	}
	return tags, nil
}

func (m *MockTagRepository) Update(tag *entities.Tag) error {
	if _, exists := m.tags[tag.ID]; exists {
		m.tags[tag.ID] = tag
		return nil
	}
	return nil
}

func (m *MockTagRepository) Delete(id uint) error {
	delete(m.tags, id)
	return nil
}

func (m *MockTagRepository) GetByPlayerID(playerID uint) ([]entities.Tag, error) {
	// Mock implementation - return empty slice
	return []entities.Tag{}, nil
}

func (m *MockTagRepository) GetByTeamID(teamID uint) ([]entities.Tag, error) {
	// Mock implementation - return empty slice
	return []entities.Tag{}, nil
}

// TestTagService_CreateTag tests the CreateTag method
func TestTagService_CreateTag(t *testing.T) {
	mockRepo := NewMockTagRepository()
	service := NewTagService(mockRepo)

	tests := []struct {
		name    string
		tag     *entities.Tag
		wantErr bool
	}{
		{
			name: "Valid tag",
			tag: &entities.Tag{
				Name: "Forward",
			},
			wantErr: false,
		},
		{
			name: "Empty name",
			tag: &entities.Tag{
				Name: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.CreateTag(tt.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTag() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestTagService_GetTagByID tests the GetTagByID method
func TestTagService_GetTagByID(t *testing.T) {
	mockRepo := NewMockTagRepository()
	service := NewTagService(mockRepo)

	// Create a test tag
	testTag := &entities.Tag{Name: "Forward"}
	service.CreateTag(testTag)

	tests := []struct {
		name    string
		id      uint
		wantErr bool
	}{
		{
			name:    "Valid ID",
			id:      1,
			wantErr: false,
		},
		{
			name:    "Invalid ID",
			id:      0,
			wantErr: true,
		},
		{
			name:    "Non-existent ID",
			id:      999,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.GetTagByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTagByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestTagService_GetAllTags tests the GetAllTags method
func TestTagService_GetAllTags(t *testing.T) {
	mockRepo := NewMockTagRepository()
	service := NewTagService(mockRepo)

	// Create some test tags
	tags := []*entities.Tag{
		{Name: "Forward"},
		{Name: "Midfielder"},
		{Name: "Defender"},
	}

	for _, tag := range tags {
		service.CreateTag(tag)
	}

	// Test getting all tags
	result, err := service.GetAllTags()
	if err != nil {
		t.Errorf("GetAllTags() error = %v", err)
	}

	if len(result) != len(tags) {
		t.Errorf("GetAllTags() returned %d tags, want %d", len(result), len(tags))
	}
} 