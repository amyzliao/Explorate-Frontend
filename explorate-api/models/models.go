package models

import "database/sql"

type Ngo_opp struct {
	ID                       int
	Org_name                 sql.NullString
	Org_location             sql.NullString
	Mission_statement        sql.NullString
	Impact_space             sql.NullString
	Un_goals                 sql.NullString
	Projects                 sql.NullString
	Values                   sql.NullString
	Benefits                 sql.NullString
	Ethical_criteria         sql.NullString
	Ethical_criteria_details sql.NullString
	Volunteer_location       sql.NullString
	Volunteer_skills         sql.NullString
	Skill_projects           sql.NullString
	Language_requirements    sql.NullString
	Volunteer_qualities      sql.NullString
	Work_env                 sql.NullString
	Accomodations            sql.NullString
	Cost                     sql.NullString
	Website                  sql.NullString
	Contact                  sql.NullString
	Photos                   sql.NullString
	Refer_others             sql.NullString
	Additional_info          sql.NullString
	Finding_volunteers       sql.NullString
	Org_logo                 sql.NullString
	Program_duration         sql.NullString
	Project_timeframe        sql.NullString
}
