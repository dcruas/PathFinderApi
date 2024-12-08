package skillcheckhdl

import (
	"net/http"
	"strconv"

	"pathfinderapi/models/requests/skillcheckreq"

	"pathfinderapi/services/skillchecksvc"

	"github.com/gin-gonic/gin"
)

func GetCheckStatics(c *gin.Context) {

	modifierText := c.Query("modifier")
	dcText := c.Query("dc")

	modifier, err := strconv.Atoi(modifierText)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid modifier value"})
		return
	}

	dc, err := strconv.Atoi(dcText)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DC value"})
		return
	}

	if dc < -100 || dc > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DC value, please enter a valid value between -100 and 100"})
		return
	}

	if modifier < -100 || modifier > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Modifier value, please enter a value between -100 and 100"})
		return
	}

	request := skillcheckreq.SkillCheckRequest{
		Modifier: modifier,
		Dc:       dc,
	}

	result := skillchecksvc.CalculatePossibleOutcomes(request)

	c.JSON(http.StatusOK, result)
}
