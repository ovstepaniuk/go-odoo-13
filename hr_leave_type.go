package odoo

// HrLeaveType represents hr.leave.type model.
type HrLeaveType struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omitempty"`
	Active                   *Bool      `xmlrpc:"active,omitempty"`
	AllocationNotifSubtypeId *Many2One  `xmlrpc:"allocation_notif_subtype_id,omitempty"`
	AllocationType           *Selection `xmlrpc:"allocation_type,omitempty"`
	Code                     *String    `xmlrpc:"code,omitempty"`
	ColorName                *Selection `xmlrpc:"color_name,omitempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateCalendarMeeting    *Bool      `xmlrpc:"create_calendar_meeting,omitempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName              *String    `xmlrpc:"display_name,omitempty"`
	GroupDaysAllocation      *Float     `xmlrpc:"group_days_allocation,omitempty"`
	GroupDaysLeave           *Float     `xmlrpc:"group_days_leave,omitempty"`
	Id                       *Int       `xmlrpc:"id,omitempty"`
	LeaveNotifSubtypeId      *Many2One  `xmlrpc:"leave_notif_subtype_id,omitempty"`
	LeavesTaken              *Float     `xmlrpc:"leaves_taken,omitempty"`
	MaxLeaves                *Float     `xmlrpc:"max_leaves,omitempty"`
	Name                     *String    `xmlrpc:"name,omitempty"`
	RemainingLeaves          *Float     `xmlrpc:"remaining_leaves,omitempty"`
	RequestUnit              *Selection `xmlrpc:"request_unit,omitempty"`
	ResponsibleId            *Many2One  `xmlrpc:"responsible_id,omitempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omitempty"`
	TimeType                 *Selection `xmlrpc:"time_type,omitempty"`
	TimesheetGenerate        *Bool      `xmlrpc:"timesheet_generate,omitempty"`
	TimesheetProjectId       *Many2One  `xmlrpc:"timesheet_project_id,omitempty"`
	TimesheetTaskId          *Many2One  `xmlrpc:"timesheet_task_id,omitempty"`
	Unpaid                   *Bool      `xmlrpc:"unpaid,omitempty"`
	Valid                    *Bool      `xmlrpc:"valid,omitempty"`
	ValidationType           *Selection `xmlrpc:"validation_type,omitempty"`
	ValidityStart            *Time      `xmlrpc:"validity_start,omitempty"`
	ValidityStop             *Time      `xmlrpc:"validity_stop,omitempty"`
	VirtualRemainingLeaves   *Float     `xmlrpc:"virtual_remaining_leaves,omitempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrLeaveTypes represents array of hr.leave.type model.
type HrLeaveTypes []HrLeaveType

// HrLeaveTypeModel is the odoo model name.
const HrLeaveTypeModel = "hr.leave.type"

// Many2One convert HrLeaveType to *Many2One.
func (hlt *HrLeaveType) Many2One() *Many2One {
	return NewMany2One(hlt.Id.Get(), "")
}

// CreateHrLeaveType creates a new hr.leave.type model and returns its id.
func (c *Client) CreateHrLeaveType(hlt *HrLeaveType) (int64, error) {
	ids, err := c.CreateHrLeaveTypes([]*HrLeaveType{hlt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveType creates a new hr.leave.type model and returns its id.
func (c *Client) CreateHrLeaveTypes(hlts []*HrLeaveType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlts {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveTypeModel, vv, nil)
}

// UpdateHrLeaveType updates an existing hr.leave.type record.
func (c *Client) UpdateHrLeaveType(hlt *HrLeaveType) error {
	return c.UpdateHrLeaveTypes([]int64{hlt.Id.Get()}, hlt)
}

// UpdateHrLeaveTypes updates existing hr.leave.type records.
// All records (represented by ids) will be updated by hlt values.
func (c *Client) UpdateHrLeaveTypes(ids []int64, hlt *HrLeaveType) error {
	return c.Update(HrLeaveTypeModel, ids, hlt, nil)
}

// DeleteHrLeaveType deletes an existing hr.leave.type record.
func (c *Client) DeleteHrLeaveType(id int64) error {
	return c.DeleteHrLeaveTypes([]int64{id})
}

// DeleteHrLeaveTypes deletes existing hr.leave.type records.
func (c *Client) DeleteHrLeaveTypes(ids []int64) error {
	return c.Delete(HrLeaveTypeModel, ids)
}

// GetHrLeaveType gets hr.leave.type existing record.
func (c *Client) GetHrLeaveType(id int64) (*HrLeaveType, error) {
	hlts, err := c.GetHrLeaveTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hlts)[0]), nil
}

// GetHrLeaveTypes gets hr.leave.type existing records.
func (c *Client) GetHrLeaveTypes(ids []int64) (*HrLeaveTypes, error) {
	hlts := &HrLeaveTypes{}
	if err := c.Read(HrLeaveTypeModel, ids, nil, hlts); err != nil {
		return nil, err
	}
	return hlts, nil
}

// FindHrLeaveType finds hr.leave.type record by querying it with criteria.
func (c *Client) FindHrLeaveType(criteria *Criteria) (*HrLeaveType, error) {
	hlts := &HrLeaveTypes{}
	if err := c.SearchRead(HrLeaveTypeModel, criteria, NewOptions().Limit(1), hlts); err != nil {
		return nil, err
	}
	return &((*hlts)[0]), nil
}

// FindHrLeaveTypes finds hr.leave.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveTypes(criteria *Criteria, options *Options) (*HrLeaveTypes, error) {
	hlts := &HrLeaveTypes{}
	if err := c.SearchRead(HrLeaveTypeModel, criteria, options, hlts); err != nil {
		return nil, err
	}
	return hlts, nil
}

// FindHrLeaveTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrLeaveTypeModel, criteria, options)
}

// FindHrLeaveTypeId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
