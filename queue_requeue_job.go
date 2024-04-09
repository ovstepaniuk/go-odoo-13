package odoo

// QueueRequeueJob represents queue.requeue.job model.
type QueueRequeueJob struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	JobIds      *Relation `xmlrpc:"job_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// QueueRequeueJobs represents array of queue.requeue.job model.
type QueueRequeueJobs []QueueRequeueJob

// QueueRequeueJobModel is the odoo model name.
const QueueRequeueJobModel = "queue.requeue.job"

// Many2One convert QueueRequeueJob to *Many2One.
func (qrj *QueueRequeueJob) Many2One() *Many2One {
	return NewMany2One(qrj.Id.Get(), "")
}

// CreateQueueRequeueJob creates a new queue.requeue.job model and returns its id.
func (c *Client) CreateQueueRequeueJob(qrj *QueueRequeueJob) (int64, error) {
	ids, err := c.CreateQueueRequeueJobs([]*QueueRequeueJob{qrj})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQueueRequeueJob creates a new queue.requeue.job model and returns its id.
func (c *Client) CreateQueueRequeueJobs(qrjs []*QueueRequeueJob) ([]int64, error) {
	var vv []interface{}
	for _, v := range qrjs {
		vv = append(vv, v)
	}
	return c.Create(QueueRequeueJobModel, vv, nil)
}

// UpdateQueueRequeueJob updates an existing queue.requeue.job record.
func (c *Client) UpdateQueueRequeueJob(qrj *QueueRequeueJob) error {
	return c.UpdateQueueRequeueJobs([]int64{qrj.Id.Get()}, qrj)
}

// UpdateQueueRequeueJobs updates existing queue.requeue.job records.
// All records (represented by ids) will be updated by qrj values.
func (c *Client) UpdateQueueRequeueJobs(ids []int64, qrj *QueueRequeueJob) error {
	return c.Update(QueueRequeueJobModel, ids, qrj, nil)
}

// DeleteQueueRequeueJob deletes an existing queue.requeue.job record.
func (c *Client) DeleteQueueRequeueJob(id int64) error {
	return c.DeleteQueueRequeueJobs([]int64{id})
}

// DeleteQueueRequeueJobs deletes existing queue.requeue.job records.
func (c *Client) DeleteQueueRequeueJobs(ids []int64) error {
	return c.Delete(QueueRequeueJobModel, ids)
}

// GetQueueRequeueJob gets queue.requeue.job existing record.
func (c *Client) GetQueueRequeueJob(id int64) (*QueueRequeueJob, error) {
	qrjs, err := c.GetQueueRequeueJobs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*qrjs)[0]), nil
}

// GetQueueRequeueJobs gets queue.requeue.job existing records.
func (c *Client) GetQueueRequeueJobs(ids []int64) (*QueueRequeueJobs, error) {
	qrjs := &QueueRequeueJobs{}
	if err := c.Read(QueueRequeueJobModel, ids, nil, qrjs); err != nil {
		return nil, err
	}
	return qrjs, nil
}

// FindQueueRequeueJob finds queue.requeue.job record by querying it with criteria.
func (c *Client) FindQueueRequeueJob(criteria *Criteria) (*QueueRequeueJob, error) {
	qrjs := &QueueRequeueJobs{}
	if err := c.SearchRead(QueueRequeueJobModel, criteria, NewOptions().Limit(1), qrjs); err != nil {
		return nil, err
	}
	return &((*qrjs)[0]), nil
}

// FindQueueRequeueJobs finds queue.requeue.job records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueRequeueJobs(criteria *Criteria, options *Options) (*QueueRequeueJobs, error) {
	qrjs := &QueueRequeueJobs{}
	if err := c.SearchRead(QueueRequeueJobModel, criteria, options, qrjs); err != nil {
		return nil, err
	}
	return qrjs, nil
}

// FindQueueRequeueJobIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueRequeueJobIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(QueueRequeueJobModel, criteria, options)
}

// FindQueueRequeueJobId finds record id by querying it with criteria.
func (c *Client) FindQueueRequeueJobId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueRequeueJobModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
