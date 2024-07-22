package controllers

import (
	"context"
	"net/http"
	"usermanagementservice/models"

	"github.com/gin-gonic/gin"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
)

// List Members godoc
// @Summary List Members for an Org.
// @Description List Members for an org.
// @Tags UserManagement
// @Accept json
// @Produce json
// @Success 200
// @Router /listmembers [get]
// @Security Bearer
func ListMembership(c *gin.Context) {

	TenantId, _ := c.Get("tenant_id")
	OrgmembersList, err := usermanagement.ListUsers(
		context.Background(),
		usermanagement.ListUsersOpts{
			OrganizationID: TenantId.(string),
		},
	)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"members": OrgmembersList})

}

// Invite User godoc
// @Summary Invite an user for an Org.
// @Description Invite an user for an org.
// @Tags UserManagement
// @Accept json
// @Produce json
// @Success 200
// @Param {object} body models.UserInvite true "User Invite"
// @Router /inviteuser [post]
// @Security Bearer
func InviteUser(c *gin.Context) {

	var inviteuser models.UserInvite

	if err := c.ShouldBindJSON(&inviteuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	TenantId, _ := c.Get("tenant_id")
	OrgmembersList, err := usermanagement.ListUsers(
		c.Request.Context(),
		usermanagement.ListUsersOpts{
			OrganizationID: TenantId.(string),
			Email:          inviteuser.EmailID,
		},
	)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if len(OrgmembersList.Data) != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User already Present in the org"})
		return
	}

	InviteUser, err := usermanagement.SendInvitation(
		c.Request.Context(),
		usermanagement.SendInvitationOpts{
			OrganizationID: TenantId.(string),
			Email:          inviteuser.EmailID,
			RoleSlug:       inviteuser.Role,
		},
	)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"members": InviteUser})

}
