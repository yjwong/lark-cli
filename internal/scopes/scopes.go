package scopes

import "strings"

// ScopeGroup defines a group of OAuth scopes required for a set of commands
type ScopeGroup struct {
	Name        string   // e.g., "calendar"
	Description string   // e.g., "Calendar events and scheduling"
	Scopes      []string // OAuth scope strings
	Commands    []string // CLI commands that require this group
}

// BaseScope is always required for token refresh
const BaseScope = "offline_access"

// Groups defines all available scope groups
var Groups = map[string]ScopeGroup{
	"calendar": {
		Name:        "calendar",
		Description: "Calendar events and scheduling",
		Scopes:      []string{"calendar:calendar", "calendar:calendar:readonly"},
		Commands:    []string{"cal"},
	},
	"contacts": {
		Name:        "contacts",
		Description: "Company directory lookup",
		Scopes:      []string{"contact:contact.base:readonly", "contact:department.base:readonly", "contact:department.organize:readonly", "contact:user:search"},
		Commands:    []string{"contact"},
	},
	"documents": {
		Name:        "documents",
		Description: "Lark Docs and Drive access",
		Scopes:      []string{"docx:document:readonly", "docs:doc:readonly", "docs:document.content:read", "docs:document.comment:read", "drive:drive:readonly", "wiki:wiki:readonly", "space:document:retrieve"},
		Commands:    []string{"doc"},
	},
	"bitable": {
		Name:        "bitable",
		Description: "Lark Bitable (database) access",
		Scopes:      []string{"bitable:app:readonly"},
		Commands:    []string{"bitable"},
	},
	"messages": {
		Name:        "messages",
		Description: "Chat and messaging",
		Scopes:      []string{"im:message:readonly", "im:message", "im:message:send_as_bot", "im:message.reactions:read", "im:message.reactions:write_only"},
		Commands:    []string{"msg", "chat"},
	},
	"mail": {
		Name:        "mail",
		Description: "Email via IMAP",
		Scopes:      []string{"mail:user_mailbox.message.address:read", "mail:user_mailbox.message.body:read", "mail:user_mailbox.message.subject:read", "mail:user_mailbox.message:readonly"},
		Commands:    []string{"mail"},
	},
	"minutes": {
		Name:        "minutes",
		Description: "Meeting recordings and transcripts",
		Scopes:      []string{"minutes:minutes:readonly", "minutes:minute:download"},
		Commands:    []string{"minutes"},
	},
}

// AllGroupNames returns all scope group names in a consistent order
func AllGroupNames() []string {
	return []string{"calendar", "contacts", "documents", "bitable", "messages", "mail", "minutes"}
}

// GetScopesForGroups returns the combined scopes for the given group names
func GetScopesForGroups(groupNames []string) []string {
	scopeSet := make(map[string]bool)
	scopeSet[BaseScope] = true // Always include base scope

	for _, name := range groupNames {
		if group, ok := Groups[name]; ok {
			for _, scope := range group.Scopes {
				scopeSet[scope] = true
			}
		}
	}

	scopes := make([]string, 0, len(scopeSet))
	for scope := range scopeSet {
		scopes = append(scopes, scope)
	}
	return scopes
}

// GetAllScopes returns all scopes for all groups (full permissions)
func GetAllScopes() []string {
	return GetScopesForGroups(AllGroupNames())
}

// GetScopeString returns scopes as a space-separated string for OAuth
func GetScopeString(groupNames []string) string {
	scopes := GetScopesForGroups(groupNames)
	return strings.Join(scopes, " ")
}

// GetAllScopeString returns all scopes as a space-separated string
func GetAllScopeString() string {
	return GetScopeString(AllGroupNames())
}

// GetGroupForCommand returns the scope group required by a command
func GetGroupForCommand(cmd string) (ScopeGroup, bool) {
	for _, group := range Groups {
		for _, c := range group.Commands {
			if c == cmd {
				return group, true
			}
		}
	}
	return ScopeGroup{}, false
}

// ParseGroups parses a comma-separated list of group names and validates them
func ParseGroups(input string) ([]string, []string) {
	if input == "" {
		return nil, nil
	}

	parts := strings.Split(input, ",")
	valid := make([]string, 0, len(parts))
	invalid := make([]string, 0)

	for _, part := range parts {
		name := strings.TrimSpace(part)
		if name == "" {
			continue
		}
		if _, ok := Groups[name]; ok {
			valid = append(valid, name)
		} else {
			invalid = append(invalid, name)
		}
	}

	return valid, invalid
}
