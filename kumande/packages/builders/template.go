package builders

import (
	"kumande/packages/helpers/typography"
	"strings"
)

func GetTemplateSelect(name string, firstTable, secondTable *string) string {
	if name == "tag_info" {
		return "tags.slug_name, tag_name, dictionaries.dct_name as tag_category"
	} else if name == "dct_info" {
		return "slug_name, dct_name, dct_type, dct_desc, type_name"
	} else if name == "properties" {
		return "created_at, created_by, updated_at, updated_by"
	}
	return ""
}

func GetTemplateConcat(name, col string) string {
	if name == "value_group" {
		return "GROUP_CONCAT(" + col + " SEPARATOR ', ') as " + col
	}
	return ""
}

func GetTemplateGeneralSelect(name string, firstTable *string) string {
	ftable := typography.RemoveLastChar(*firstTable, "s")
	if name == "info" {
		return *firstTable + ".slug_name, " + ftable + "_name, " + ftable + "_desc"
	}
	return ""
}

func GetTemplateOrder(name, tableName, ext string) string {
	if name == "permanent_data" {
		return tableName + ".created_at DESC, " + tableName + "." + ext + " DESC "
	} else if name == "dynamic_data" {
		return tableName + ".updated_at DESC, " + tableName + ".created_at DESC, " + tableName + "." + ext + " DESC "
	} else if name == "most_used_normal" {
		return " COUNT(1) DESC"
	}
	return ""
}

func GetTemplateJoin(typeJoin, firstTable, firstCol, secondTable, secondCol string, nullable bool) string {
	var join = ""
	if nullable {
		join = "LEFT JOIN "
	} else {
		join = "JOIN "
	}
	if typeJoin == "same_col" {
		return join + secondTable + " USING(" + firstCol + ") "
	} else if typeJoin == "total" {
		return join + secondTable + " ON " + secondTable + "." + secondCol + " = " + firstTable + "." + firstCol + " "
	}
	return ""
}

func GetTemplateGroup(is_multi_where bool, col string) string {
	var ext = " WHERE "
	if is_multi_where {
		ext = " AND "
	}

	return ext + col + " IS NOT NULL AND trim(" + col + ") != '' GROUP BY " + col + " "
}

func GetTemplateLogic(name string) string {
	if name == "active" {
		return ".deleted_at IS NULL "
	} else if name == "trash" {
		return ".deleted_at IS NOT NULL "
	}
	return ""
}

func GetWhereMine(token string) string {
	return "users_tokens.token ='" + token + "'"
}

// Stats
func GetTemplateStats(ctx, firstTable, name string, ord string, joinArgs *string) string {
	// Nullable args
	var args string
	if joinArgs == nil {
		args = ""
	} else {
		args = *joinArgs
	}
	// Notes :
	// Full query
	if name == "most_appear" {
		return "SELECT " + ctx + " as context, COUNT(1) AS total FROM " + firstTable + " " + args + " GROUP BY " + ctx + " ORDER BY total " + ord
	}

	return ""
}

func GetTemplateJobs(ctx, mainCol, key, res string) string {
	if ctx == "json_search" {
		return "REPLACE(JSON_EXTRACT(" + mainCol + ", '$[0]." + key + "'), '\"', '') as " + res
	}

	return ""
}

func GetTemplateCommand(name, tableName, colName string) string {
	if name == "soft_delete" {
		return "UPDATE " + tableName + " SET deleted_at = ?, deleted_by = ? WHERE " + tableName + "." + colName + " = ?"
	} else if name == "hard_delete" {
		return "DELETE FROM " + tableName + " WHERE " + tableName + "." + colName + " = ?"
	} else if name == "recover" {
		return "UPDATE " + tableName + " SET deleted_at = null, deleted_by = null WHERE " + tableName + "." + colName + " = ?"
	} else if name == "filter_tag" {
		tags := strings.Split(colName, ",")
		sytx := tableName + `_tag like `

		if len(tags) == 1 {
			sytx += `'%"slug_name":"` + tags[0] + `"%'`
		} else {
			for i := 0; i < len(tags); i++ {
				sytx += `'%"slug_name":"` + tags[i] + `"%'`

				if i != len(tags)-1 {
					sytx += ` or ` + tableName + `_tag like `
				}
			}
		}
		return sytx
	}
	return ""
}
