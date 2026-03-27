package scopes

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// GroupCatalog defines the comprehensive tenant/user scopes for a group.
// Used for generating Lark app batch-import JSON — not for OAuth or runtime validation.
type GroupCatalog struct {
	TenantScopes []string
	UserScopes   []string
}

// Catalog contains comprehensive scope data for each group, derived from Lark API docs.
// These are all the scopes a Lark app needs configured to use the CLI's features.
var Catalog = map[string]GroupCatalog{
	"calendar": {
		TenantScopes: []string{
			"calendar:calendar",
			"calendar:calendar.acl:create",
			"calendar:calendar.acl:delete",
			"calendar:calendar.acl:read",
			"calendar:calendar.event:create",
			"calendar:calendar.event:delete",
			"calendar:calendar.event:read",
			"calendar:calendar.event:update",
			"calendar:calendar.free_busy:read",
			"calendar:calendar:create",
			"calendar:calendar:delete",
			"calendar:calendar:read",
			"calendar:calendar:readonly",
			"calendar:calendar:subscribe",
			"calendar:calendar:update",
			"calendar:time_off:create",
			"calendar:time_off:delete",
			"calendar:timeoff",
		},
		UserScopes: []string{
			"calendar:calendar",
			"calendar:calendar.acl:create",
			"calendar:calendar.acl:delete",
			"calendar:calendar.acl:read",
			"calendar:calendar.event:create",
			"calendar:calendar.event:delete",
			"calendar:calendar.event:read",
			"calendar:calendar.event:update",
			"calendar:calendar.free_busy:read",
			"calendar:calendar:create",
			"calendar:calendar:delete",
			"calendar:calendar:read",
			"calendar:calendar:readonly",
			"calendar:calendar:subscribe",
			"calendar:calendar:update",
		},
	},
	"contacts": {
		TenantScopes: []string{
			"contact:contact.base:readonly",
			"contact:contact:readonly_as_app",
			"contact:department.base:readonly",
			"contact:department.organize:readonly",
			"contact:user.email:readonly",
			"contact:user.employee:readonly",
			"contact:user.employee_id:readonly",
			"contact:user.id:readonly",
		},
		UserScopes: []string{
			"contact:contact.base:readonly",
			"contact:contact:access_as_app",
			"contact:contact:readonly_as_app",
			"contact:department.base:readonly",
			"contact:department.organize:readonly",
			"contact:user.base:readonly",
			"contact:user.department:readonly",
			"contact:user.email:readonly",
			"contact:user.employee:readonly",
			"contact:user.employee_id:readonly",
			"contact:user:search",
		},
	},
	"documents": {
		TenantScopes: []string{
			"docs:doc:readonly",
			"docs:document.comment:read",
			"docs:document.content:read",
			"docs:document.media:download",
			"docs:document:export",
			"docs:permission.member:auth",
			"docs:permission.setting:readonly",
			"docx:document",
			"docx:document.block:convert",
			"docx:document:create",
			"docx:document:readonly",
			"drive:drive",
			"drive:drive.metadata:readonly",
			"drive:drive.search:readonly",
			"drive:drive:readonly",
			"drive:drive:version:readonly",
			"drive:export:readonly",
			"drive:file:readonly",
			"space:document:retrieve",
			"wiki:node:read",
			"wiki:node:retrieve",
			"wiki:setting:read",
			"wiki:space:read",
			"wiki:space:retrieve",
			"wiki:wiki",
			"wiki:wiki:readonly",
		},
		UserScopes: []string{
			"docs:doc",
			"docs:doc:readonly",
			"docs:document.comment:create",
			"docs:document.comment:read",
			"docs:document.comment:update",
			"docs:document.content:read",
			"docs:document:export",
			"docx:document",
			"docx:document.block:convert",
			"docx:document:create",
			"docx:document:readonly",
			"docx:document:write_only",
			"drive:drive",
			"drive:drive:readonly",
			"drive:drive:version",
			"drive:export:readonly",
			"search:docs:read",
			"space:document:retrieve",
			"wiki:wiki:readonly",
		},
	},
	"bitable": {
		TenantScopes: []string{
			"base:app:copy",
			"base:app:create",
			"base:app:read",
			"base:field:create",
			"base:field:delete",
			"base:field:read",
			"base:field:update",
			"base:field_group:create",
			"base:form:update",
			"base:record:create",
			"base:record:delete",
			"base:record:retrieve",
			"base:record:update",
			"base:table:create",
			"base:table:delete",
			"base:table:read",
			"base:table:update",
			"base:view:read",
			"base:view:write_only",
			"bitable:app",
			"bitable:app:readonly",
		},
		UserScopes: []string{
			"base:app:copy",
			"base:app:create",
			"base:app:read",
			"base:field:create",
			"base:field:delete",
			"base:field:read",
			"base:field:update",
			"base:field_group:create",
			"base:form:update",
			"base:record:create",
			"base:record:delete",
			"base:record:retrieve",
			"base:record:update",
			"base:table:create",
			"base:table:delete",
			"base:table:read",
			"base:table:update",
			"base:view:read",
			"base:view:write_only",
			"bitable:app",
			"bitable:app:readonly",
		},
	},
	"messages": {
		TenantScopes: []string{
			"im:chat",
			"im:chat:create",
			"im:message",
			"im:message.reactions:read",
			"im:message.reactions:write_only",
			"im:message:readonly",
			"im:message:send_as_bot",
			"im:resource",
		},
		UserScopes: []string{
			"im:chat",
			"im:message",
			"im:message.reactions:read",
			"im:message.reactions:write_only",
			"im:message:readonly",
			"im:message:send_as_bot",
			"search:message",
		},
	},
	"mail": {
		TenantScopes: []string{
			"mail:user_mailbox.message.address:read",
			"mail:user_mailbox.message.body:read",
			"mail:user_mailbox.message.subject:read",
			"mail:user_mailbox.message:readonly",
		},
		UserScopes: []string{
			"mail:user_mailbox.message.address:read",
			"mail:user_mailbox.message.body:read",
			"mail:user_mailbox.message.subject:read",
			"mail:user_mailbox.message:readonly",
		},
	},
	"minutes": {
		TenantScopes: []string{
			"minutes:minutes:readonly",
		},
		UserScopes: []string{
			"minutes:minutes",
			"minutes:minutes.basic:read",
			"minutes:minutes.media:export",
			"minutes:minutes.statistics:read",
			"minutes:minutes.transcript:export",
			"minutes:minutes:readonly",
		},
	},
}

// batchImportOutput is the JSON structure for Lark's scope batch import
type batchImportOutput struct {
	Scopes struct {
		Tenant []string `json:"tenant"`
		User   []string `json:"user"`
	} `json:"scopes"`
}

// GenerateBatchImportJSON generates the JSON for Lark's scope batch import.
// It returns an error for unknown group names or empty selections.
func GenerateBatchImportJSON(groupNames []string) ([]byte, error) {
	if len(groupNames) == 0 {
		return nil, fmt.Errorf("no scope groups specified")
	}

	for _, name := range groupNames {
		if _, ok := Catalog[name]; !ok {
			return nil, fmt.Errorf("unknown scope group: %q (valid groups: %s)", name, strings.Join(AllGroupNames(), ", "))
		}
	}

	tenantSet := make(map[string]bool)
	userSet := make(map[string]bool)

	for _, name := range groupNames {
		cat := Catalog[name]
		for _, s := range cat.TenantScopes {
			tenantSet[s] = true
		}
		for _, s := range cat.UserScopes {
			userSet[s] = true
		}
	}

	// Always include offline_access in user scopes for token refresh
	userSet[BaseScope] = true

	var out batchImportOutput
	out.Scopes.Tenant = sortedKeys(tenantSet)
	out.Scopes.User = sortedKeys(userSet)

	return json.MarshalIndent(out, "", "  ")
}

func sortedKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
