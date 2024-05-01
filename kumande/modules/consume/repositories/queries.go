package repositories

import (
	"database/sql"
	"kumande/modules/consume/models"
	"kumande/packages/builders"
	"kumande/packages/database"
	"kumande/packages/helpers/converter"
	"kumande/packages/helpers/generator"
	"kumande/packages/helpers/response"

	"net/http"
)

func GetConsumeBySlug(slug string) (response.Response, error) {
	// Declaration
	var obj models.GetConsumeSearch
	var arrobj []models.GetConsumeSearch
	var res response.Response
	var baseTable = "consume"
	var sqlStatement string
	filterTemplate := builders.GetTemplateCommand("filter_tag", baseTable, slug)

	// Nullable
	var ConsumeComment sql.NullString

	sqlStatement = "SELECT slug_name, consume_type, consume_name, consume_detail, consume_from, is_favorite, consume_tag, consume_comment, created_at " +
		"FROM " + baseTable + " " +
		"WHERE " + filterTemplate + " " +
		"ORDER BY consume_name ASC"

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
			&obj.Slug,
			&obj.ConsumeType,
			&obj.ConsumeName,
			&obj.ConsumeDetail,
			&obj.ConsumeFrom,
			&obj.IsFavorite,
			&obj.ConsumeTag,
			&ConsumeComment,
			&obj.CreatedAt,
		)

		if err != nil {
			return res, err
		}

		// Nullable
		obj.ConsumeComment = converter.CheckNullString(ConsumeComment)

		arrobj = append(arrobj, obj)
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateQueryMsg(baseTable, len(arrobj))
	if len(arrobj) == 0 {
		res.Data = nil
	} else {
		res.Data = arrobj
	}

	return res, nil
}
