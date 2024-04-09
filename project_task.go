package odoo

// ProjectTask represents project.task model.
type ProjectTask struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omitempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omitempty"`
	AccessUrl                   *String    `xmlrpc:"access_url,omitempty"`
	AccessWarning               *String    `xmlrpc:"access_warning,omitempty"`
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AllowTimesheets             *Bool      `xmlrpc:"allow_timesheets,omitempty"`
	AnalyticAccountActive       *Bool      `xmlrpc:"analytic_account_active,omitempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omitempty"`
	BillableType                *Selection `xmlrpc:"billable_type,omitempty"`
	ChildIds                    *Relation  `xmlrpc:"child_ids,omitempty"`
	Color                       *Int       `xmlrpc:"color,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateAssign                  *Time      `xmlrpc:"date_assign,omitempty"`
	DateDeadline                *Time      `xmlrpc:"date_deadline,omitempty"`
	DateDeadlineFormatted       *String    `xmlrpc:"date_deadline_formatted,omitempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omitempty"`
	DateLastStageUpdate         *Time      `xmlrpc:"date_last_stage_update,omitempty"`
	Description                 *String    `xmlrpc:"description,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	DisplayedImageId            *Many2One  `xmlrpc:"displayed_image_id,omitempty"`
	EffectiveHours              *Float     `xmlrpc:"effective_hours,omitempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omitempty"`
	EmailFrom                   *String    `xmlrpc:"email_from,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IsProjectMapEmpty           *Bool      `xmlrpc:"is_project_map_empty,omitempty"`
	KanbanState                 *Selection `xmlrpc:"kanban_state,omitempty"`
	KanbanStateLabel            *String    `xmlrpc:"kanban_state_label,omitempty"`
	LegendBlocked               *String    `xmlrpc:"legend_blocked,omitempty"`
	LegendDone                  *String    `xmlrpc:"legend_done,omitempty"`
	LegendNormal                *String    `xmlrpc:"legend_normal,omitempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds           *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread               *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter        *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	ParentId                    *Many2One  `xmlrpc:"parent_id,omitempty"`
	PartnerCity                 *String    `xmlrpc:"partner_city,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	PlannedHours                *Float     `xmlrpc:"planned_hours,omitempty"`
	Priority                    *Selection `xmlrpc:"priority,omitempty"`
	Progress                    *Float     `xmlrpc:"progress,omitempty"`
	ProjectId                   *Many2One  `xmlrpc:"project_id,omitempty"`
	RatingAvg                   *Float     `xmlrpc:"rating_avg,omitempty"`
	RatingCount                 *Int       `xmlrpc:"rating_count,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	RatingLastFeedback          *String    `xmlrpc:"rating_last_feedback,omitempty"`
	RatingLastImage             *String    `xmlrpc:"rating_last_image,omitempty"`
	RatingLastValue             *Float     `xmlrpc:"rating_last_value,omitempty"`
	RemainingHours              *Float     `xmlrpc:"remaining_hours,omitempty"`
	SaleLineId                  *Many2One  `xmlrpc:"sale_line_id,omitempty"`
	SaleOrderId                 *Many2One  `xmlrpc:"sale_order_id,omitempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omitempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omitempty"`
	SubtaskCount                *Int       `xmlrpc:"subtask_count,omitempty"`
	SubtaskEffectiveHours       *Float     `xmlrpc:"subtask_effective_hours,omitempty"`
	SubtaskPlannedHours         *Float     `xmlrpc:"subtask_planned_hours,omitempty"`
	SubtaskProjectId            *Many2One  `xmlrpc:"subtask_project_id,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	TimesheetIds                *Relation  `xmlrpc:"timesheet_ids,omitempty"`
	TotalHoursSpent             *Float     `xmlrpc:"total_hours_spent,omitempty"`
	UserEmail                   *String    `xmlrpc:"user_email,omitempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WorkingDaysClose            *Float     `xmlrpc:"working_days_close,omitempty"`
	WorkingDaysOpen             *Float     `xmlrpc:"working_days_open,omitempty"`
	WorkingHoursClose           *Float     `xmlrpc:"working_hours_close,omitempty"`
	WorkingHoursOpen            *Float     `xmlrpc:"working_hours_open,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProjectTasks represents array of project.task model.
type ProjectTasks []ProjectTask

// ProjectTaskModel is the odoo model name.
const ProjectTaskModel = "project.task"

// Many2One convert ProjectTask to *Many2One.
func (pt *ProjectTask) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProjectTask creates a new project.task model and returns its id.
func (c *Client) CreateProjectTask(pt *ProjectTask) (int64, error) {
	ids, err := c.CreateProjectTasks([]*ProjectTask{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTask creates a new project.task model and returns its id.
func (c *Client) CreateProjectTasks(pts []*ProjectTask) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskModel, vv, nil)
}

// UpdateProjectTask updates an existing project.task record.
func (c *Client) UpdateProjectTask(pt *ProjectTask) error {
	return c.UpdateProjectTasks([]int64{pt.Id.Get()}, pt)
}

// UpdateProjectTasks updates existing project.task records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProjectTasks(ids []int64, pt *ProjectTask) error {
	return c.Update(ProjectTaskModel, ids, pt, nil)
}

// DeleteProjectTask deletes an existing project.task record.
func (c *Client) DeleteProjectTask(id int64) error {
	return c.DeleteProjectTasks([]int64{id})
}

// DeleteProjectTasks deletes existing project.task records.
func (c *Client) DeleteProjectTasks(ids []int64) error {
	return c.Delete(ProjectTaskModel, ids)
}

// GetProjectTask gets project.task existing record.
func (c *Client) GetProjectTask(id int64) (*ProjectTask, error) {
	pts, err := c.GetProjectTasks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pts)[0]), nil
}

// GetProjectTasks gets project.task existing records.
func (c *Client) GetProjectTasks(ids []int64) (*ProjectTasks, error) {
	pts := &ProjectTasks{}
	if err := c.Read(ProjectTaskModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTask finds project.task record by querying it with criteria.
func (c *Client) FindProjectTask(criteria *Criteria) (*ProjectTask, error) {
	pts := &ProjectTasks{}
	if err := c.SearchRead(ProjectTaskModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	return &((*pts)[0]), nil
}

// FindProjectTasks finds project.task records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTasks(criteria *Criteria, options *Options) (*ProjectTasks, error) {
	pts := &ProjectTasks{}
	if err := c.SearchRead(ProjectTaskModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTaskIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectTaskModel, criteria, options)
}

// FindProjectTaskId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
