package muscle_test

import (
	"context"
	"errors"
	"testing"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockMuscleRepository is a mock implementation of the MuscleRepository interface
type MockMuscleRepository struct {
	mock.Mock
}

func (m *MockMuscleRepository) Add(ctx context.Context, userID uuid.UUID, muscle *muscle.Muscle) error {
	args := m.Called(ctx, userID, muscle)
	return args.Error(0)
}

func (m *MockMuscleRepository) GetByID(ctx context.Context, userID, muscleID uuid.UUID) (*muscle.Muscle, error) {
	args := m.Called(ctx, userID, muscleID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*muscle.Muscle), args.Error(1)
}

func (m *MockMuscleRepository) List(ctx context.Context, userID uuid.UUID) ([]*muscle.Muscle, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*muscle.Muscle), args.Error(1)
}

func (m *MockMuscleRepository) Delete(ctx context.Context, userID, muscleID uuid.UUID) error {
	args := m.Called(ctx, userID, muscleID)
	return args.Error(0)
}

// Test cases for Muscle domain model
func TestNewMuscle(t *testing.T) {
	tests := []struct {
		name          string
		params        muscle.MuscleParams
		expectedError error
	}{
		{
			name: "Valid muscle creation",
			params: muscle.MuscleParams{
				ID:   uuid.New(),
				Name: "Biceps",
			},
			expectedError: nil,
		},
		{
			name: "Empty name",
			params: muscle.MuscleParams{
				ID:   uuid.New(),
				Name: "",
			},
			expectedError: muscle.ErrInvalidMuscle,
		},
		{
			name: "Empty ID should generate new ID",
			params: muscle.MuscleParams{
				ID:   uuid.Nil,
				Name: "Triceps",
			},
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m, err := muscle.NewMuscle(tc.params)

			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, m)

				if tc.params.ID == uuid.Nil {
					assert.NotEqual(t, uuid.Nil, m.ID(), "Should generate new UUID when ID is nil")
				} else {
					assert.Equal(t, tc.params.ID, m.ID())
				}

				assert.Equal(t, tc.params.Name, m.Name())
			}
		})
	}
}

// Test cases for Muscle service
func TestMuscleService_AddMuscle(t *testing.T) {
	mockRepo := new(MockMuscleRepository)
	service := muscle.NewMuscleService(mockRepo)

	ctx := context.Background()
	userID := uuid.New()

	validMuscle, _ := muscle.NewMuscle(muscle.MuscleParams{
		ID:   uuid.Nil,
		Name: "Quadriceps",
	})

	t.Run("Successful add", func(t *testing.T) {
		mockRepo.On("Add", ctx, userID, validMuscle).Return(nil).Once()

		err := service.AddMuscle(ctx, userID, validMuscle)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository error", func(t *testing.T) {
		expectedErr := errors.New("database connection failed")
		mockRepo.On("Add", ctx, userID, validMuscle).Return(expectedErr).Once()

		err := service.AddMuscle(ctx, userID, validMuscle)

		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestMuscleService_GetMuscleByID(t *testing.T) {
	mockRepo := new(MockMuscleRepository)
	service := muscle.NewMuscleService(mockRepo)

	ctx := context.Background()
	userID := uuid.New()
	muscleID := uuid.New()

	validMuscle, _ := muscle.NewMuscle(muscle.MuscleParams{
		ID:   muscleID,
		Name: "Hamstrings",
	})

	t.Run("Successful get", func(t *testing.T) {
		mockRepo.On("GetByID", ctx, userID, muscleID).Return(validMuscle, nil).Once()

		result, err := service.GetMuscleByID(ctx, userID, muscleID)

		assert.NoError(t, err)
		assert.Equal(t, validMuscle, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not found", func(t *testing.T) {
		expectedErr := errors.New("muscle not found")
		mockRepo.On("GetByID", ctx, userID, muscleID).Return(nil, expectedErr).Once()

		result, err := service.GetMuscleByID(ctx, userID, muscleID)

		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestMuscleService_ListMuscles(t *testing.T) {
	mockRepo := new(MockMuscleRepository)
	service := muscle.NewMuscleService(mockRepo)

	ctx := context.Background()
	userID := uuid.New()

	muscle1, _ := muscle.NewMuscle(muscle.MuscleParams{Name: "Biceps"})
	muscle2, _ := muscle.NewMuscle(muscle.MuscleParams{Name: "Triceps"})
	muscleList := []*muscle.Muscle{muscle1, muscle2}

	t.Run("Successful list", func(t *testing.T) {
		mockRepo.On("List", ctx, userID).Return(muscleList, nil).Once()

		result, err := service.ListMuscles(ctx, userID)

		assert.NoError(t, err)
		assert.Equal(t, muscleList, result)
		assert.Len(t, result, 2)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty list", func(t *testing.T) {
		emptyList := []*muscle.Muscle{}
		mockRepo.On("List", ctx, userID).Return(emptyList, nil).Once()

		result, err := service.ListMuscles(ctx, userID)

		assert.NoError(t, err)
		assert.Empty(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Repository error", func(t *testing.T) {
		expectedErr := errors.New("database error")
		mockRepo.On("List", ctx, userID).Return(nil, expectedErr).Once()

		result, err := service.ListMuscles(ctx, userID)

		assert.Equal(t, expectedErr, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestMuscleService_RemoveMuscle(t *testing.T) {
	mockRepo := new(MockMuscleRepository)
	service := muscle.NewMuscleService(mockRepo)

	ctx := context.Background()
	userID := uuid.New()
	muscleID := uuid.New()

	t.Run("Successful delete", func(t *testing.T) {
		mockRepo.On("Delete", ctx, userID, muscleID).Return(nil).Once()

		err := service.RemoveMuscle(ctx, userID, muscleID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Not found", func(t *testing.T) {
		expectedErr := errors.New("muscle not found")
		mockRepo.On("Delete", ctx, userID, muscleID).Return(expectedErr).Once()

		err := service.RemoveMuscle(ctx, userID, muscleID)

		assert.Equal(t, expectedErr, err)
		mockRepo.AssertExpectations(t)
	})
}
