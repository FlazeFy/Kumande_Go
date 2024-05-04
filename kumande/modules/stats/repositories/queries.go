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

func GetBudgetSpendingYear(year string) (response.Response, error) {
	// Declaration
	var obj models.GetMostAppear
	var res response.Response
	var baseTable = "budget"
	var sqlStatement string

	// Converted column
	var totalStr string

	sqlStatement = `SELECT ` +
		`REPLACE(JSON_EXTRACT(budget_month_year, '$[0].month'), '\"', '') as context, budget_total as total ` +
		`FROM ` + baseTable + ` ` +
		`WHERE REPLACE(JSON_EXTRACT(budget_month_year, '$[0].year'), '\"', '') = ` + year

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
		res.Data = obj
	}

	return res, nil
}

func GetSpendingInfo() (response.Response, error) {
	// Declaration
	var obj models.GetSpendingInfo
	var res response.Response
	var baseTable = "payment"
	var sqlStatement string

	// Converted column
	var totalDays string
	var totalPayment string

	sqlStatement = `SELECT 
		COUNT(payment_date) as total_days, CAST(IFNULL(SUM(total_payment),0) as INT) as total_payment 
		FROM (
			SELECT DATE(created_at) as payment_date, SUM(payment_price) as total_payment
			FROM ` + baseTable + ` ` + `
			WHERE created_by = '".$user_id."'
			GROUP BY payment_date
		) q`

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
			&totalDays,
			&totalPayment)

		if err != nil {
			return res, err
		}

		// Converted
		totalDaysInt, err := strconv.Atoi(totalDays)
		totalPaymentInt, err := strconv.Atoi(totalPayment)
		if err != nil {
			return res, err
		}

		obj.TotalDays = totalDaysInt
		obj.TotalPayment = totalPaymentInt
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
		res.Data = obj
	}

	return res, nil
}

func GetBodyInfo() (response.Response, error) {
	// Declaration
	var obj models.GetBodyInfo
	var res response.Response
	var baseTable = "count_calorie"
	var sqlStatement string

	// Converted column
	var weight string
	var height string
	var result string

	sqlStatement = `SELECT weight, height, result, created_at
		FROM ` + baseTable + ` ORDER BY created_at DESC`

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
			&weight,
			&height,
			&result,
			&obj.CreatedAt)

		if err != nil {
			return res, err
		}

		// Converted
		weightInt, err := strconv.Atoi(weight)
		heightInt, err := strconv.Atoi(height)
		resultInt, err := strconv.Atoi(result)

		if err != nil {
			return res, err
		}

		obj.Weight = weightInt
		obj.Height = heightInt
		obj.Result = resultInt
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
		res.Data = obj
	}

	return res, nil
}
