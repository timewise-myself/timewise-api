package workspace_service

import (
	"api/service/workspace"
	"api/unit_test/utils"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/timewise-team/timewise-models/models"
	"testing"
)

type MockDMSClient struct {
	mock.Mock
}

//func TestFunc1_UTCID01(t *testing.T) {
//	t.Log("Func1_UTCID01")
//	utils.InitConfig()
//	mockDMS := new(MockDMSClient)
//	service := workspace.NewCreateWorkspaceService()
//
//	mockSchedule := create_workspace_dtos.CreateWorkspaceResponse{
//		ID:          35,
//		Title:       "Business",
//		Description: "Task for business",
//		Type:        "workspace",
//		IsDeleted:   false,
//	}
//
//	request := models.TwWorkspace{
//		ID:          0,
//		Title:       "Business",
//		Description: "Task for business",
//		Type:        "workspace",
//		IsDeleted:   false,
//	}
//	a, err := service.CreateWorkspace(request)
//	if err != nil {
//		t.Logf("Received error: %v", err)
//		t.FailNow()
//	}
//
//	assert.NoError(t, err)
//	assert.Equal(t, mockSchedule.ID, a.ID)
//	assert.Equal(t, mockSchedule.Title, a.Title)
//	assert.Equal(t, mockSchedule.Description, a.Description)
//	assert.Equal(t, mockSchedule.Type, a.Type)
//	assert.Equal(t, mockSchedule.IsDeleted, a.IsDeleted)
//
//	mockDMS.AssertExpectations(t)
//}

func TestFunc2_UTCID02(t *testing.T) {
	t.Log("Func2_UTCID02")
	utils.InitConfig()
	mockDMS := new(MockDMSClient)
	service := workspace.NewCreateWorkspaceService()

	expectedErr := errors.New("Title is required")

	request := models.TwWorkspace{
		ID:          0,
		Title:       "",
		Description: "Task for business",
		Type:        "workspace",
		IsDeleted:   false,
	}
	a, err := service.CreateWorkspace(request)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, a)

	mockDMS.AssertExpectations(t)
}

func TestFunc2_UTCID03(t *testing.T) {
	t.Log("Func2_UTCID03")
	utils.InitConfig()
	mockDMS := new(MockDMSClient)
	service := workspace.NewCreateWorkspaceService()

	expectedErr := errors.New("Title must not exceed 50 characters")

	request := models.TwWorkspace{
		ID:          0,
		Title:       "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		Description: "Task for business",
		Type:        "workspace",
		IsDeleted:   false,
	}
	a, err := service.CreateWorkspace(request)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, a)

	mockDMS.AssertExpectations(t)
}
func TestFunc2_UTCID04(t *testing.T) {
	t.Log("Func2_UTCID04")
	utils.InitConfig()
	mockDMS := new(MockDMSClient)
	service := workspace.NewCreateWorkspaceService()

	expectedErr := errors.New("Description is required")

	request := models.TwWorkspace{
		ID:          0,
		Title:       "Business",
		Description: "",
		Type:        "workspace",
		IsDeleted:   false,
	}
	a, err := service.CreateWorkspace(request)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, a)

	mockDMS.AssertExpectations(t)
}
func TestFunc2_UTCID05(t *testing.T) {
	t.Log("Func1_UTCID05")
	utils.InitConfig()
	mockDMS := new(MockDMSClient)
	service := workspace.NewCreateWorkspaceService()

	expectedErr := errors.New("Description must not exceed 255 characters")

	request := models.TwWorkspace{
		ID:    0,
		Title: "Business",
		Description: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" +
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		Type:      "workspace",
		IsDeleted: false,
	}
	a, err := service.CreateWorkspace(request)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Nil(t, a)

	mockDMS.AssertExpectations(t)
}
