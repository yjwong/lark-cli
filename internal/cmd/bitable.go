package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yjwong/lark-cli/internal/api"
	"github.com/yjwong/lark-cli/internal/output"
)

var bitableCmd = &cobra.Command{
	Use:   "bitable",
	Short: "Bitable (database) commands",
	Long:  "Access Lark Bitable databases - list tables, fields, and records",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		validateScopeGroup("bitable")
	},
}

// --- bitable tables ---

var bitableTablesCmd = &cobra.Command{
	Use:   "tables <app_token>",
	Short: "List tables in a Bitable",
	Long: `List all tables in a Lark Bitable (database).

The app_token is from the Bitable URL.
For example, if the URL is https://xxx.larksuite.com/base/ABC123xyz
then the app_token is ABC123xyz.

Examples:
  lark bitable tables ABC123xyz`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appToken := args[0]

		client := api.NewClient()

		tables, err := client.ListBitableTables(appToken)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		outputTables := make([]api.OutputBitableTable, len(tables))
		for i, t := range tables {
			outputTables[i] = api.OutputBitableTable{
				TableID: t.TableID,
				Name:    t.Name,
			}
		}

		result := api.OutputBitableTableList{
			AppToken: appToken,
			Tables:   outputTables,
			Count:    len(outputTables),
		}

		output.JSON(result)
	},
}

// --- bitable fields ---

var bitableFieldsCmd = &cobra.Command{
	Use:   "fields <app_token> <table_id>",
	Short: "List fields in a Bitable table",
	Long: `List all fields (columns) in a Bitable table.

Examples:
  lark bitable fields ABC123xyz tblXYZ789`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		appToken := args[0]
		tableID := args[1]

		client := api.NewClient()

		fields, err := client.ListBitableFields(appToken, tableID)
		if err != nil {
			output.Fatal("API_ERROR", err)
		}

		outputFields := make([]api.OutputBitableField, len(fields))
		for i, f := range fields {
			outputFields[i] = api.OutputBitableField{
				FieldID:   f.FieldID,
				FieldName: f.FieldName,
				Type:      bitableFieldTypeToString(f.Type),
				IsPrimary: f.IsPrimary,
			}
		}

		result := api.OutputBitableFieldList{
			AppToken: appToken,
			TableID:  tableID,
			Fields:   outputFields,
			Count:    len(outputFields),
		}

		output.JSON(result)
	},
}

// --- bitable records ---

var (
	bitableRecordsLimit  int
	bitableRecordsViewID string
	bitableRecordsFilter string
)

var bitableRecordsCmd = &cobra.Command{
	Use:   "records <app_token> <table_id>",
	Short: "List records in a Bitable table",
	Long: `List records (rows) in a Bitable table.

Examples:
  lark bitable records ABC123xyz tblXYZ789
  lark bitable records ABC123xyz tblXYZ789 --limit 50
  lark bitable records ABC123xyz tblXYZ789 --view vewABC123`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		appToken := args[0]
		tableID := args[1]

		client := api.NewClient()

		opts := &api.BitableRecordOptions{
			ViewID:   bitableRecordsViewID,
			Filter:   bitableRecordsFilter,
			PageSize: 100,
		}

		var allRecords []api.BitableRecord
		var pageToken string
		hasMore := true
		remaining := bitableRecordsLimit

		for hasMore {
			if remaining > 0 && remaining < opts.PageSize {
				opts.PageSize = remaining
			}
			opts.PageToken = pageToken

			records, more, nextToken, err := client.ListBitableRecords(appToken, tableID, opts)
			if err != nil {
				output.Fatal("API_ERROR", err)
			}

			allRecords = append(allRecords, records...)
			hasMore = more
			pageToken = nextToken

			if bitableRecordsLimit > 0 {
				remaining = bitableRecordsLimit - len(allRecords)
				if remaining <= 0 {
					break
				}
			}
		}

		// Trim to limit if needed
		if bitableRecordsLimit > 0 && len(allRecords) > bitableRecordsLimit {
			allRecords = allRecords[:bitableRecordsLimit]
		}

		outputRecords := make([]api.OutputBitableRecord, len(allRecords))
		for i, r := range allRecords {
			outputRecords[i] = api.OutputBitableRecord{
				RecordID: r.RecordID,
				Fields:   r.Fields,
			}
		}

		result := api.OutputBitableRecordList{
			AppToken: appToken,
			TableID:  tableID,
			Records:  outputRecords,
			Count:    len(outputRecords),
			HasMore:  hasMore,
		}

		output.JSON(result)
	},
}

// bitableFieldTypeToString converts field type int to human-readable string
func bitableFieldTypeToString(fieldType int) string {
	switch fieldType {
	case 1:
		return "text"
	case 2:
		return "number"
	case 3:
		return "select"
	case 4:
		return "multi_select"
	case 5:
		return "date"
	case 7:
		return "checkbox"
	case 11:
		return "person"
	case 13:
		return "phone"
	case 15:
		return "url"
	case 17:
		return "attachment"
	case 18:
		return "link"
	case 19:
		return "formula"
	case 20:
		return "duplex_link"
	case 21:
		return "location"
	case 22:
		return "group"
	case 23:
		return "created_time"
	case 1001:
		return "created_user"
	case 1002:
		return "modified_time"
	case 1003:
		return "modified_user"
	case 1004:
		return "auto_number"
	default:
		return "unknown"
	}
}

func init() {
	// bitable records flags
	bitableRecordsCmd.Flags().IntVar(&bitableRecordsLimit, "limit", 0,
		"Maximum number of records to retrieve (0 = no limit)")
	bitableRecordsCmd.Flags().StringVar(&bitableRecordsViewID, "view", "",
		"View ID to filter records")
	bitableRecordsCmd.Flags().StringVar(&bitableRecordsFilter, "filter", "",
		"Filter expression")

	// Register subcommands
	bitableCmd.AddCommand(bitableTablesCmd)
	bitableCmd.AddCommand(bitableFieldsCmd)
	bitableCmd.AddCommand(bitableRecordsCmd)
}
