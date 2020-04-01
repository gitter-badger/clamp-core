package handlers

import (
	"clamp-core/models"
	"clamp-core/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func createWorkflowHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody, _ := c.GetRawData()
		// Create new Service Workflow
		request := models.Workflow{}
		json.Unmarshal([]byte(requestBody), &request)
		fmt.Printf("Workflow request : %v \n", request.ServiceFlow)

		serviceFlowRes := models.CreateWorkflow(request)
		services.SaveServiceFlow(serviceFlowRes)
		//TODO - handle error scenario. Currently it is always 200 ok
		c.JSON(http.StatusOK, serviceFlowRes)
	}
}

func fetchWorkflowBasedOnWorkflowName() gin.HandlerFunc {
	return func(c *gin.Context) {
		workflowName := c.Param("workflow")
		var result models.Workflow
		result, _ = services.FindWorkflowByName(uuid.MustParse(workflowName))
		//TODO - handle error scenario. Currently it is always 200 ok
		c.JSON(http.StatusOK, result)
	}
}
