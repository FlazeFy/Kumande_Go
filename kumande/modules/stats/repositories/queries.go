package repositories

import (
	"fmt"
	"kumande/modules/stats/models"
	"kumande/packages/builders"
	"kumande/packages/database"
	"kumande/packages/helpers/generator"
	"kumande/packages/helpers/response"
	"net/http"
	"strconv"
)

func GetTotalConsumeMulti(path string, ord string, limit string, view string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = "consume"
	var mainCol = view
	var sqlStatement string

	// Converted column
	var totalStr string

	// Query builder
	selectTemplate := builders.GetTemplateStats(mainCol, baseTable, "most_appear", ord, nil)

	sqlStatement = selectTemplate + " LIMIT " + limit
	fmt.Printf(sqlStatement)

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Total
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}

func GetTotalConsumeByConsumeDetail(path string, ord string, limit string, view string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var arrobj []models.GetMostAppear
	var res response.Response
	var baseTable = "consume"
	var mainCol = view
	var sqlStatement string

	// Converted column
	var totalStr string

	// Query builder
	selectTemplate := builders.GetTemplateJobs("json_search", "consume_detail", mainCol, "context")

	sqlStatement = "SELECT " + selectTemplate + ", count(1) as total " +
		"FROM " + baseTable + " " +
		"GROUP BY 1 " +
		"ORDER BY 2 " + ord + " " +
		"LIMIT 8"
	fmt.Printf(sqlStatement)

	// Exec
	con := database.CreateCon()
	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	// Map
	for rows.Next() {
		err = rows.Scan(
			&obj.Context,
			&totalStr)

		if err != nil {
			return res, err
		}

		// Converted
		totalInt, err := strconv.Atoi(totalStr)
		if err != nil {
			return res, err
		}

		obj.Total = totalInt
		arrobj = append(arrobj, obj)
	}

	// Total
	total, err := builders.GetTotalCount(con, baseTable, nil)
	if err != nil {
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg("Stats", total)
	if total == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
