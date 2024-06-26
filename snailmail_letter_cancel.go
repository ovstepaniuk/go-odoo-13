package odoo

// SnailmailLetterCancel represents snailmail.letter.cancel model.
type SnailmailLetterCancel struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	HelpMessage *String   `xmlrpc:"help_message,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Model       *String   `xmlrpc:"model,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SnailmailLetterCancels represents array of snailmail.letter.cancel model.
type SnailmailLetterCancels []SnailmailLetterCancel

// SnailmailLetterCancelModel is the odoo model name.
const SnailmailLetterCancelModel = "snailmail.letter.cancel"

// Many2One convert SnailmailLetterCancel to *Many2One.
func (slc *SnailmailLetterCancel) Many2One() *Many2One {
	return NewMany2One(slc.Id.Get(), "")
}

// CreateSnailmailLetterCancel creates a new snailmail.letter.cancel model and returns its id.
func (c *Client) CreateSnailmailLetterCancel(slc *SnailmailLetterCancel) (int64, error) {
	ids, err := c.CreateSnailmailLetterCancels([]*SnailmailLetterCancel{slc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSnailmailLetterCancel creates a new snailmail.letter.cancel model and returns its id.
func (c *Client) CreateSnailmailLetterCancels(slcs []*SnailmailLetterCancel) ([]int64, error) {
	var vv []interface{}
	for _, v := range slcs {
		vv = append(vv, v)
	}
	return c.Create(SnailmailLetterCancelModel, vv, nil)
}

// UpdateSnailmailLetterCancel updates an existing snailmail.letter.cancel record.
func (c *Client) UpdateSnailmailLetterCancel(slc *SnailmailLetterCancel) error {
	return c.UpdateSnailmailLetterCancels([]int64{slc.Id.Get()}, slc)
}

// UpdateSnailmailLetterCancels updates existing snailmail.letter.cancel records.
// All records (represented by ids) will be updated by slc values.
func (c *Client) UpdateSnailmailLetterCancels(ids []int64, slc *SnailmailLetterCancel) error {
	return c.Update(SnailmailLetterCancelModel, ids, slc, nil)
}

// DeleteSnailmailLetterCancel deletes an existing snailmail.letter.cancel record.
func (c *Client) DeleteSnailmailLetterCancel(id int64) error {
	return c.DeleteSnailmailLetterCancels([]int64{id})
}

// DeleteSnailmailLetterCancels deletes existing snailmail.letter.cancel records.
func (c *Client) DeleteSnailmailLetterCancels(ids []int64) error {
	return c.Delete(SnailmailLetterCancelModel, ids)
}

// GetSnailmailLetterCancel gets snailmail.letter.cancel existing record.
func (c *Client) GetSnailmailLetterCancel(id int64) (*SnailmailLetterCancel, error) {
	slcs, err := c.GetSnailmailLetterCancels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*slcs)[0]), nil
}

// GetSnailmailLetterCancels gets snailmail.letter.cancel existing records.
func (c *Client) GetSnailmailLetterCancels(ids []int64) (*SnailmailLetterCancels, error) {
	slcs := &SnailmailLetterCancels{}
	if err := c.Read(SnailmailLetterCancelModel, ids, nil, slcs); err != nil {
		return nil, err
	}
	return slcs, nil
}

// FindSnailmailLetterCancel finds snailmail.letter.cancel record by querying it with criteria.
func (c *Client) FindSnailmailLetterCancel(criteria *Criteria) (*SnailmailLetterCancel, error) {
	slcs := &SnailmailLetterCancels{}
	if err := c.SearchRead(SnailmailLetterCancelModel, criteria, NewOptions().Limit(1), slcs); err != nil {
		return nil, err
	}
	return &((*slcs)[0]), nil
}

// FindSnailmailLetterCancels finds snailmail.letter.cancel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailLetterCancels(criteria *Criteria, options *Options) (*SnailmailLetterCancels, error) {
	slcs := &SnailmailLetterCancels{}
	if err := c.SearchRead(SnailmailLetterCancelModel, criteria, options, slcs); err != nil {
		return nil, err
	}
	return slcs, nil
}

// FindSnailmailLetterCancelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailLetterCancelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SnailmailLetterCancelModel, criteria, options)
}

// FindSnailmailLetterCancelId finds record id by querying it with criteria.
func (c *Client) FindSnailmailLetterCancelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SnailmailLetterCancelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
