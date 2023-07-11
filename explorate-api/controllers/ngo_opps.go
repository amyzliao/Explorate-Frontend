package controllers

import (
	"context"
	"database/sql"
	"explorate/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type getAll struct {
	ID          int
	Name        sql.NullString
	Title       sql.NullString
	Description sql.NullString
	Location    sql.NullString
	Contact     sql.NullString
	Website     sql.NullString
}

// GetAll
func GetAll(conn *pgx.Conn, table string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		// rows, err := conn.Query(context.Background(),
		// 	"SELECT org_id, org_name, org_location, impact_space, projects, contact, website FROM "+table)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// var opps []getAll
		// for rows.Next() {
		// 	var opp getAll
		// 	if err := rows.Scan(
		// 		&opp.ID,
		// 		&opp.Name,
		// 		&opp.Title,
		// 		&opp.Description,
		// 		&opp.Location,
		// 		&opp.Contact,
		// 		&opp.Website,
		// 	); err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	opps = append(opps, opp)
		// }

		rows, err := conn.Query(context.Background(),
			"SELECT * FROM "+table)
		if err != nil {
			fmt.Println("error:", err)
		}

		var opps []models.Ngo_opp
		for rows.Next() {

			var opp models.Ngo_opp
			if err := rows.Scan(
				&opp.Org_ID,
				&opp.Org_name,
				&opp.Org_location,
				&opp.Mission_statement,
				&opp.Impact_space,
				&opp.Un_goals,
				&opp.Projects,
				&opp.Values,
				&opp.Benefits,
				&opp.Ethical_criteria,
				&opp.Ethical_criteria_details,
				&opp.Volunteer_location,
				&opp.Volunteer_skills,
				&opp.Skill_projects,
				&opp.Language_requirements,
				&opp.Volunteer_qualities,
				&opp.Work_env,
				&opp.Accomodations,
				&opp.Cost,
				&opp.Website,
				&opp.Contact,
				&opp.Photos,
				&opp.Refer_others,
				&opp.Additional_info,
				&opp.Finding_volunteers,
				&opp.Org_logo,
				&opp.Program_duration,
				&opp.Project_timeframe,
			); err != nil {
				fmt.Println(err)
			}

			opps = append(opps, opp)
		}
		c.IndentedJSON(http.StatusOK, opps)
	}

	return gin.HandlerFunc(fn)
}

// Get
func Find(conn *pgx.Conn, table string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		search_id := c.Param("id")
		var opp getAll

		err := conn.QueryRow(context.Background(),
			"SELECT org_id, org_name, org_location, impact_space, projects, contact, website FROM "+table+" WHERE org_id = $1 LIMIT 1", search_id).Scan(
			&opp.ID,
			&opp.Name,
			&opp.Title,
			&opp.Description,
			&opp.Location,
			&opp.Contact,
			&opp.Website,
		)

		if err != nil {
			fmt.Println(err)
		}

		c.IndentedJSON(http.StatusOK, opp)
	}

	return gin.HandlerFunc(fn)
}

// Insert
func Insert(conn *pgx.Conn, table string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input models.Ngo_opp

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		conn.Exec(context.Background(),
			"INSERT INTO "+table+" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)",
			input.Org_ID,
			input.Org_name,
			input.Org_location,
			input.Mission_statement,
			input.Impact_space,
			input.Un_goals,
			input.Projects,
			input.Values,
			input.Benefits,
			input.Ethical_criteria,
			input.Ethical_criteria_details,
			input.Volunteer_location,
			input.Volunteer_skills,
			input.Skill_projects,
			input.Language_requirements,
			input.Volunteer_qualities,
			input.Work_env,
			input.Accomodations,
			input.Cost,
			input.Website,
			input.Contact,
			input.Photos,
			input.Refer_others,
			input.Additional_info,
			input.Finding_volunteers,
			input.Org_logo,
			input.Program_duration,
			input.Project_timeframe,
		)
	}
	return gin.HandlerFunc(fn)
}

// Delete
func Delete(conn *pgx.Conn, table string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		id := c.Param("id")
		conn.Exec(context.Background(),
			"DELETE FROM "+table+" WHERE org_id = $1", id)
	}
	return gin.HandlerFunc(fn)
}

// // Update
