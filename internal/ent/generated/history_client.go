// Code generated by entx.history, DO NOT EDIT.

// Code generated by ent, DO NOT EDIT.

package generated

import "github.com/theopenlane/entx/history"

// withHistory adds the history hooks to the appropriate schemas - generated by entx.history
func (c *Client) WithHistory() {
	for _, hook := range history.HistoryHooks[*ContactMutation]() {
		c.Contact.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*DocumentDataMutation]() {
		c.DocumentData.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*EntitlementMutation]() {
		c.Entitlement.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*EntitlementPlanMutation]() {
		c.EntitlementPlan.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*EntitlementPlanFeatureMutation]() {
		c.EntitlementPlanFeature.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*EntityMutation]() {
		c.Entity.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*EntityTypeMutation]() {
		c.EntityType.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*EventMutation]() {
		c.Event.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*FeatureMutation]() {
		c.Feature.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*FileMutation]() {
		c.File.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*GroupMutation]() {
		c.Group.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*GroupMembershipMutation]() {
		c.GroupMembership.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*GroupSettingMutation]() {
		c.GroupSetting.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*HushMutation]() {
		c.Hush.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*IntegrationMutation]() {
		c.Integration.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*NoteMutation]() {
		c.Note.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*OauthProviderMutation]() {
		c.OauthProvider.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*OrgMembershipMutation]() {
		c.OrgMembership.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*OrganizationMutation]() {
		c.Organization.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*OrganizationSettingMutation]() {
		c.OrganizationSetting.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*TemplateMutation]() {
		c.Template.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*UserMutation]() {
		c.User.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*UserSettingMutation]() {
		c.UserSetting.Use(hook)
	}
	for _, hook := range history.HistoryHooks[*WebhookMutation]() {
		c.Webhook.Use(hook)
	}
}