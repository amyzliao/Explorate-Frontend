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

type New_opp struct {
	ID                       int
	Org_name                 string
	Org_location             string
	Mission_statement        string
	Impact_space             string
	Un_goals                 string
	Projects                 string
	Values                   string
	Benefits                 string
	Ethical_criteria         string
	Ethical_criteria_details string
	Volunteer_location       string
	Volunteer_skills         string
	Skill_projects           string
	Language_requirements    string
	Volunteer_qualities      string
	Work_env                 string
	Accomodations            string
	Cost                     string
	Website                  string
	Contact                  string
	Photos                   string
	Refer_others             string
	Additional_info          string
	Finding_volunteers       string
	Org_logo                 string
	Program_duration         string
	Project_timeframe        string
}

func NewInput(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// GetAll
func GetAll(conn *pgx.Conn, table string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rows, err := conn.Query(context.Background(),
			"SELECT * FROM "+table)
		if err != nil {
			fmt.Println("error:", err)
		}

		var opps []models.Ngo_opp
		for rows.Next() {

			var opp models.Ngo_opp
			if err := rows.Scan(
				&opp.ID,
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
		var opp models.Ngo_opp

		err := conn.QueryRow(context.Background(),
			"SELECT * FROM "+table+" WHERE ID = $1 LIMIT 1", search_id).Scan(
			&opp.ID,
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
		)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(opp)
		c.IndentedJSON(http.StatusOK, opp)
	}

	return gin.HandlerFunc(fn)
}

// Insert
func Insert(conn *pgx.Conn, table string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input New_opp
		// fmt.Println("works")
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		// fmt.Println(input)
		// opp := models.Ngo_opp{
		// 	// Org_name:                 NewInput(input.Org_name),
		// 	// Org_location:             NewInput(input.Org_location),
		// 	// Mission_statement:        NewInput(input.Mission_statement),
		// 	// Impact_space:             NewInput(input.Impact_space),
		// 	// Un_goals:                 NewInput(input.Un_goals),
		// 	// Projects:                 NewInput(input.Projects),
		// 	// Values:                   NewInput(input.Values),
		// 	// Benefits:                 NewInput(input.Benefits),
		// 	// Ethical_criteria:         NewInput(input.Ethical_criteria),
		// 	// Ethical_criteria_details: NewInput(input.Ethical_criteria_details),
		// 	// Volunteer_location:       NewInput(input.Volunteer_location),
		// 	// Volunteer_skills:         NewInput(input.Volunteer_skills),
		// 	// Skill_projects:           NewInput(input.Skill_projects),
		// 	// Language_requirements:    NewInput(input.Language_requirements),
		// 	// Volunteer_qualities:      NewInput(input.Volunteer_qualities),
		// 	// Work_env:                 NewInput(input.Work_env),
		// 	// Accomodations:            NewInput(input.Accomodations),
		// 	// Cost:                     NewInput(input.Cost),
		// 	// Website:                  NewInput(input.Website),
		// 	// Contact:                  NewInput(input.Contact),
		// 	// Photos:                   NewInput(input.Photos),
		// 	// Refer_others:             NewInput(input.Refer_others),
		// 	// Additional_info:          NewInput(input.Additional_info),
		// 	// Finding_volunteers:       NewInput(input.Finding_volunteers),
		// 	// Org_logo:                 NewInput(input.Org_logo),
		// 	// Program_duration:         NewInput(input.Program_duration),
		// 	// Project_timeframe:        NewInput(input.Project_timeframe),
		// 	ID:                       input.ID,
		// 	Org_name:                 input.Org_name,
		// 	Org_location:             input.Org_location,
		// 	Mission_statement:        input.Mission_statement,
		// 	Impact_space:             input.Impact_space,
		// 	Un_goals:                 input.Un_goals,
		// 	Projects:                 input.Projects,
		// 	Values:                   input.Values,
		// 	Benefits:                 input.Benefits,
		// 	Ethical_criteria:         input.Ethical_criteria,
		// 	Ethical_criteria_details: input.Ethical_criteria_details,
		// 	Volunteer_location:       input.Volunteer_location,
		// 	Volunteer_skills:         input.Volunteer_skills,
		// 	Skill_projects:           input.Skill_projects,
		// 	Language_requirements:    input.Language_requirements,
		// 	Volunteer_qualities:      input.Volunteer_qualities,
		// 	Work_env:                 input.Work_env,
		// 	Accomodations:            input.Accomodations,
		// 	Cost:                     input.Cost,
		// 	Website:                  input.Website,
		// 	Contact:                  input.Contact,
		// 	Photos:                   input.Photos,
		// 	Refer_others:             input.Refer_others,
		// 	Additional_info:          input.Additional_info,
		// 	Finding_volunteers:       input.Finding_volunteers,
		// 	Org_logo:                 input.Org_logo,
		// 	Program_duration:         input.Program_duration,
		// 	Project_timeframe:        input.Project_timeframe,
		// }
		// fmt.Println(opp.ID)
		conn.Exec(context.Background(),
			"INSERT INTO "+table+" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)",
			// opp.ID,
			// opp.Org_name,
			// opp.Org_location,
			// opp.Mission_statement,
			// opp.Impact_space,
			// opp.Un_goals,
			// opp.Projects,
			// opp.Values,
			// opp.Benefits,
			// opp.Ethical_criteria,
			// opp.Ethical_criteria_details,
			// opp.Volunteer_location,
			// opp.Volunteer_skills,
			// opp.Skill_projects,
			// opp.Language_requirements,
			// opp.Volunteer_qualities,
			// opp.Work_env,
			// opp.Accomodations,
			// opp.Cost,
			// opp.Website,
			// opp.Contact,
			// opp.Photos,
			// opp.Refer_others,
			// opp.Additional_info,
			// opp.Finding_volunteers,
			// opp.Org_logo,
			// opp.Program_duration,
			// opp.Project_timeframe,
			input.ID,
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
			"DELETE FROM "+table+" WHERE ID = $1", id)
	}
	return gin.HandlerFunc(fn)
}

// Update
func Update(conn *pgx.Conn, table string) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input New_opp
		id := c.Param("id")

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		fmt.Println(input)
		conn.Exec(context.Background(),
			"UPDATE "+table+" SET org_name = $1, org_location = $2, mission_statement = $3, impact_space = $4, un_goals = $5, projects = $6, values = $7, benefits = $8, ethical_criteria = $9, ethical_criteria_details = $10, volunteer_location = $11, volunteer_skills = $12, skill_projects = $13, language_requirements = $14, volunteer_qualities = $15, work_env = $16, accomodations = $17, cost = $18, website = $19, contact = $20, photos = $21, refer_others = $22, additional_info = $23, finding_volunteers = $24, org_logo = $25, program_duration = $26, project_timeframe = $27 WHERE ID = $28",
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
			id)
	}
	return gin.HandlerFunc(fn)
}
