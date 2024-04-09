package odoo

// AccountTaxRepartitionLineTemplate represents account.tax.repartition.line.template model.
type AccountTaxRepartitionLineTemplate struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	AccountId          *Many2One  `xmlrpc:"account_id,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	FactorPercent      *Float     `xmlrpc:"factor_percent,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	InvoiceTaxId       *Many2One  `xmlrpc:"invoice_tax_id,omitempty"`
	MinusReportLineIds *Relation  `xmlrpc:"minus_report_line_ids,omitempty"`
	PlusReportLineIds  *Relation  `xmlrpc:"plus_report_line_ids,omitempty"`
	RefundTaxId        *Many2One  `xmlrpc:"refund_tax_id,omitempty"`
	RepartitionType    *Selection `xmlrpc:"repartition_type,omitempty"`
	TagIds             *Relation  `xmlrpc:"tag_ids,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountTaxRepartitionLineTemplates represents array of account.tax.repartition.line.template model.
type AccountTaxRepartitionLineTemplates []AccountTaxRepartitionLineTemplate

// AccountTaxRepartitionLineTemplateModel is the odoo model name.
const AccountTaxRepartitionLineTemplateModel = "account.tax.repartition.line.template"

// Many2One convert AccountTaxRepartitionLineTemplate to *Many2One.
func (atrlt *AccountTaxRepartitionLineTemplate) Many2One() *Many2One {
	return NewMany2One(atrlt.Id.Get(), "")
}

// CreateAccountTaxRepartitionLineTemplate creates a new account.tax.repartition.line.template model and returns its id.
func (c *Client) CreateAccountTaxRepartitionLineTemplate(atrlt *AccountTaxRepartitionLineTemplate) (int64, error) {
	ids, err := c.CreateAccountTaxRepartitionLineTemplates([]*AccountTaxRepartitionLineTemplate{atrlt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTaxRepartitionLineTemplate creates a new account.tax.repartition.line.template model and returns its id.
func (c *Client) CreateAccountTaxRepartitionLineTemplates(atrlts []*AccountTaxRepartitionLineTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range atrlts {
		vv = append(vv, v)
	}
	return c.Create(AccountTaxRepartitionLineTemplateModel, vv, nil)
}

// UpdateAccountTaxRepartitionLineTemplate updates an existing account.tax.repartition.line.template record.
func (c *Client) UpdateAccountTaxRepartitionLineTemplate(atrlt *AccountTaxRepartitionLineTemplate) error {
	return c.UpdateAccountTaxRepartitionLineTemplates([]int64{atrlt.Id.Get()}, atrlt)
}

// UpdateAccountTaxRepartitionLineTemplates updates existing account.tax.repartition.line.template records.
// All records (represented by ids) will be updated by atrlt values.
func (c *Client) UpdateAccountTaxRepartitionLineTemplates(ids []int64, atrlt *AccountTaxRepartitionLineTemplate) error {
	return c.Update(AccountTaxRepartitionLineTemplateModel, ids, atrlt, nil)
}

// DeleteAccountTaxRepartitionLineTemplate deletes an existing account.tax.repartition.line.template record.
func (c *Client) DeleteAccountTaxRepartitionLineTemplate(id int64) error {
	return c.DeleteAccountTaxRepartitionLineTemplates([]int64{id})
}

// DeleteAccountTaxRepartitionLineTemplates deletes existing account.tax.repartition.line.template records.
func (c *Client) DeleteAccountTaxRepartitionLineTemplates(ids []int64) error {
	return c.Delete(AccountTaxRepartitionLineTemplateModel, ids)
}

// GetAccountTaxRepartitionLineTemplate gets account.tax.repartition.line.template existing record.
func (c *Client) GetAccountTaxRepartitionLineTemplate(id int64) (*AccountTaxRepartitionLineTemplate, error) {
	atrlts, err := c.GetAccountTaxRepartitionLineTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*atrlts)[0]), nil
}

// GetAccountTaxRepartitionLineTemplates gets account.tax.repartition.line.template existing records.
func (c *Client) GetAccountTaxRepartitionLineTemplates(ids []int64) (*AccountTaxRepartitionLineTemplates, error) {
	atrlts := &AccountTaxRepartitionLineTemplates{}
	if err := c.Read(AccountTaxRepartitionLineTemplateModel, ids, nil, atrlts); err != nil {
		return nil, err
	}
	return atrlts, nil
}

// FindAccountTaxRepartitionLineTemplate finds account.tax.repartition.line.template record by querying it with criteria.
func (c *Client) FindAccountTaxRepartitionLineTemplate(criteria *Criteria) (*AccountTaxRepartitionLineTemplate, error) {
	atrlts := &AccountTaxRepartitionLineTemplates{}
	if err := c.SearchRead(AccountTaxRepartitionLineTemplateModel, criteria, NewOptions().Limit(1), atrlts); err != nil {
		return nil, err
	}
	return &((*atrlts)[0]), nil
}

// FindAccountTaxRepartitionLineTemplates finds account.tax.repartition.line.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxRepartitionLineTemplates(criteria *Criteria, options *Options) (*AccountTaxRepartitionLineTemplates, error) {
	atrlts := &AccountTaxRepartitionLineTemplates{}
	if err := c.SearchRead(AccountTaxRepartitionLineTemplateModel, criteria, options, atrlts); err != nil {
		return nil, err
	}
	return atrlts, nil
}

// FindAccountTaxRepartitionLineTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxRepartitionLineTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountTaxRepartitionLineTemplateModel, criteria, options)
}

// FindAccountTaxRepartitionLineTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxRepartitionLineTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxRepartitionLineTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
