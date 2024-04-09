package odoo

// QueueJobChannel represents queue.job.channel model.
type QueueJobChannel struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	CompleteName   *String   `xmlrpc:"complete_name,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	JobFunctionIds *Relation `xmlrpc:"job_function_ids,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	ParentId       *Many2One `xmlrpc:"parent_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// QueueJobChannels represents array of queue.job.channel model.
type QueueJobChannels []QueueJobChannel

// QueueJobChannelModel is the odoo model name.
const QueueJobChannelModel = "queue.job.channel"

// Many2One convert QueueJobChannel to *Many2One.
func (qjc *QueueJobChannel) Many2One() *Many2One {
	return NewMany2One(qjc.Id.Get(), "")
}

// CreateQueueJobChannel creates a new queue.job.channel model and returns its id.
func (c *Client) CreateQueueJobChannel(qjc *QueueJobChannel) (int64, error) {
	ids, err := c.CreateQueueJobChannels([]*QueueJobChannel{qjc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQueueJobChannel creates a new queue.job.channel model and returns its id.
func (c *Client) CreateQueueJobChannels(qjcs []*QueueJobChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range qjcs {
		vv = append(vv, v)
	}
	return c.Create(QueueJobChannelModel, vv, nil)
}

// UpdateQueueJobChannel updates an existing queue.job.channel record.
func (c *Client) UpdateQueueJobChannel(qjc *QueueJobChannel) error {
	return c.UpdateQueueJobChannels([]int64{qjc.Id.Get()}, qjc)
}

// UpdateQueueJobChannels updates existing queue.job.channel records.
// All records (represented by ids) will be updated by qjc values.
func (c *Client) UpdateQueueJobChannels(ids []int64, qjc *QueueJobChannel) error {
	return c.Update(QueueJobChannelModel, ids, qjc, nil)
}

// DeleteQueueJobChannel deletes an existing queue.job.channel record.
func (c *Client) DeleteQueueJobChannel(id int64) error {
	return c.DeleteQueueJobChannels([]int64{id})
}

// DeleteQueueJobChannels deletes existing queue.job.channel records.
func (c *Client) DeleteQueueJobChannels(ids []int64) error {
	return c.Delete(QueueJobChannelModel, ids)
}

// GetQueueJobChannel gets queue.job.channel existing record.
func (c *Client) GetQueueJobChannel(id int64) (*QueueJobChannel, error) {
	qjcs, err := c.GetQueueJobChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qjcs)[0]), nil
}

// GetQueueJobChannels gets queue.job.channel existing records.
func (c *Client) GetQueueJobChannels(ids []int64) (*QueueJobChannels, error) {
	qjcs := &QueueJobChannels{}
	if err := c.Read(QueueJobChannelModel, ids, nil, qjcs); err != nil {
		return nil, err
	}
	return qjcs, nil
}

// FindQueueJobChannel finds queue.job.channel record by querying it with criteria.
func (c *Client) FindQueueJobChannel(criteria *Criteria) (*QueueJobChannel, error) {
	qjcs := &QueueJobChannels{}
	if err := c.SearchRead(QueueJobChannelModel, criteria, NewOptions().Limit(1), qjcs); err != nil {
		return nil, err
	}
	return &((*qjcs)[0]), nil
}

// FindQueueJobChannels finds queue.job.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobChannels(criteria *Criteria, options *Options) (*QueueJobChannels, error) {
	qjcs := &QueueJobChannels{}
	if err := c.SearchRead(QueueJobChannelModel, criteria, options, qjcs); err != nil {
		return nil, err
	}
	return qjcs, nil
}

// FindQueueJobChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueJobChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QueueJobChannelModel, criteria, options)
}

// FindQueueJobChannelId finds record id by querying it with criteria.
func (c *Client) FindQueueJobChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueJobChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
