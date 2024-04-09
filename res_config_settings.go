package odoo

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	LastUpdate                            *Time      `xmlrpc:"__last_update,omitempty"`
	AccountBankReconciliationStart        *Time      `xmlrpc:"account_bank_reconciliation_start,omitempty"`
	ActiveUserCount                       *Int       `xmlrpc:"active_user_count,omitempty"`
	AliasDomain                           *String    `xmlrpc:"alias_domain,omitempty"`
	AuthSignupResetPassword               *Bool      `xmlrpc:"auth_signup_reset_password,omitempty"`
	AuthSignupTemplateUserId              *Many2One  `xmlrpc:"auth_signup_template_user_id,omitempty"`
	AuthSignupUninvited                   *Selection `xmlrpc:"auth_signup_uninvited,omitempty"`
	AutomaticInvoice                      *Bool      `xmlrpc:"automatic_invoice,omitempty"`
	ChartTemplateId                       *Many2One  `xmlrpc:"chart_template_id,omitempty"`
	CompanyCount                          *Int       `xmlrpc:"company_count,omitempty"`
	CompanyCurrencyId                     *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyInformations                   *String    `xmlrpc:"company_informations,omitempty"`
	CompanyName                           *String    `xmlrpc:"company_name,omitempty"`
	ConfirmationTemplateId                *Many2One  `xmlrpc:"confirmation_template_id,omitempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmAliasPrefix                        *String    `xmlrpc:"crm_alias_prefix,omitempty"`
	CurrencyExchangeJournalId             *Many2One  `xmlrpc:"currency_exchange_journal_id,omitempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omitempty"`
	DefaultInvoicePolicy                  *Selection `xmlrpc:"default_invoice_policy,omitempty"`
	DefaultPickingPolicy                  *Selection `xmlrpc:"default_picking_policy,omitempty"`
	DefaultPurchaseMethod                 *Selection `xmlrpc:"default_purchase_method,omitempty"`
	DefaultSaleOrderTemplateId            *Many2One  `xmlrpc:"default_sale_order_template_id,omitempty"`
	DepositDefaultProductId               *Many2One  `xmlrpc:"deposit_default_product_id,omitempty"`
	DigestEmails                          *Bool      `xmlrpc:"digest_emails,omitempty"`
	DigestId                              *Many2One  `xmlrpc:"digest_id,omitempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omitempty"`
	ExternalEmailServerDefault            *Bool      `xmlrpc:"external_email_server_default,omitempty"`
	ExternalReportLayoutId                *Many2One  `xmlrpc:"external_report_layout_id,omitempty"`
	FailCounter                           *Int       `xmlrpc:"fail_counter,omitempty"`
	GenerateLeadFromAlias                 *Bool      `xmlrpc:"generate_lead_from_alias,omitempty"`
	GroupAnalyticAccounting               *Bool      `xmlrpc:"group_analytic_accounting,omitempty"`
	GroupAnalyticTags                     *Bool      `xmlrpc:"group_analytic_tags,omitempty"`
	GroupAutoDoneSetting                  *Bool      `xmlrpc:"group_auto_done_setting,omitempty"`
	GroupCashRounding                     *Bool      `xmlrpc:"group_cash_rounding,omitempty"`
	GroupDiscountPerSoLine                *Bool      `xmlrpc:"group_discount_per_so_line,omitempty"`
	GroupDisplayIncoterm                  *Bool      `xmlrpc:"group_display_incoterm,omitempty"`
	GroupFiscalYear                       *Bool      `xmlrpc:"group_fiscal_year,omitempty"`
	GroupLotOnDeliverySlip                *Bool      `xmlrpc:"group_lot_on_delivery_slip,omitempty"`
	GroupLotOnInvoice                     *Bool      `xmlrpc:"group_lot_on_invoice,omitempty"`
	GroupMassMailingCampaign              *Bool      `xmlrpc:"group_mass_mailing_campaign,omitempty"`
	GroupMultiCurrency                    *Bool      `xmlrpc:"group_multi_currency,omitempty"`
	GroupProductPricelist                 *Bool      `xmlrpc:"group_product_pricelist,omitempty"`
	GroupProductVariant                   *Bool      `xmlrpc:"group_product_variant,omitempty"`
	GroupProformaSales                    *Bool      `xmlrpc:"group_proforma_sales,omitempty"`
	GroupProjectRating                    *Bool      `xmlrpc:"group_project_rating,omitempty"`
	GroupSaleDeliveryAddress              *Bool      `xmlrpc:"group_sale_delivery_address,omitempty"`
	GroupSaleOrderTemplate                *Bool      `xmlrpc:"group_sale_order_template,omitempty"`
	GroupSalePricelist                    *Bool      `xmlrpc:"group_sale_pricelist,omitempty"`
	GroupShowLineSubtotalsTaxExcluded     *Bool      `xmlrpc:"group_show_line_subtotals_tax_excluded,omitempty"`
	GroupShowLineSubtotalsTaxIncluded     *Bool      `xmlrpc:"group_show_line_subtotals_tax_included,omitempty"`
	GroupStockAdvLocation                 *Bool      `xmlrpc:"group_stock_adv_location,omitempty"`
	GroupStockMultiLocations              *Bool      `xmlrpc:"group_stock_multi_locations,omitempty"`
	GroupStockMultiWarehouses             *Bool      `xmlrpc:"group_stock_multi_warehouses,omitempty"`
	GroupStockPackaging                   *Bool      `xmlrpc:"group_stock_packaging,omitempty"`
	GroupStockProductionLot               *Bool      `xmlrpc:"group_stock_production_lot,omitempty"`
	GroupStockTrackingLot                 *Bool      `xmlrpc:"group_stock_tracking_lot,omitempty"`
	GroupStockTrackingOwner               *Bool      `xmlrpc:"group_stock_tracking_owner,omitempty"`
	GroupSubtaskProject                   *Bool      `xmlrpc:"group_subtask_project,omitempty"`
	GroupUom                              *Bool      `xmlrpc:"group_uom,omitempty"`
	GroupUseLead                          *Bool      `xmlrpc:"group_use_lead,omitempty"`
	GroupWarningAccount                   *Bool      `xmlrpc:"group_warning_account,omitempty"`
	GroupWarningPurchase                  *Bool      `xmlrpc:"group_warning_purchase,omitempty"`
	GroupWarningSale                      *Bool      `xmlrpc:"group_warning_sale,omitempty"`
	GroupWarningStock                     *Bool      `xmlrpc:"group_warning_stock,omitempty"`
	HasAccountingEntries                  *Bool      `xmlrpc:"has_accounting_entries,omitempty"`
	HasChartOfAccounts                    *Bool      `xmlrpc:"has_chart_of_accounts,omitempty"`
	HrEmployeeSelfEdit                    *Bool      `xmlrpc:"hr_employee_self_edit,omitempty"`
	HrPresenceControlEmail                *Bool      `xmlrpc:"hr_presence_control_email,omitempty"`
	HrPresenceControlEmailAmount          *Int       `xmlrpc:"hr_presence_control_email_amount,omitempty"`
	HrPresenceControlIp                   *Bool      `xmlrpc:"hr_presence_control_ip,omitempty"`
	HrPresenceControlIpList               *String    `xmlrpc:"hr_presence_control_ip_list,omitempty"`
	HrPresenceControlLogin                *Bool      `xmlrpc:"hr_presence_control_login,omitempty"`
	Id                                    *Int       `xmlrpc:"id,omitempty"`
	IncotermId                            *Many2One  `xmlrpc:"incoterm_id,omitempty"`
	InvoiceIsEmail                        *Bool      `xmlrpc:"invoice_is_email,omitempty"`
	InvoiceIsPrint                        *Bool      `xmlrpc:"invoice_is_print,omitempty"`
	InvoiceIsSnailmail                    *Bool      `xmlrpc:"invoice_is_snailmail,omitempty"`
	InvoiceTerms                          *String    `xmlrpc:"invoice_terms,omitempty"`
	IsInstalledSale                       *Bool      `xmlrpc:"is_installed_sale,omitempty"`
	LanguageCount                         *Int       `xmlrpc:"language_count,omitempty"`
	Ldaps                                 *Relation  `xmlrpc:"ldaps,omitempty"`
	LeadEnrichAuto                        *Selection `xmlrpc:"lead_enrich_auto,omitempty"`
	LeadMiningInPipeline                  *Bool      `xmlrpc:"lead_mining_in_pipeline,omitempty"`
	LeaveTimesheetProjectId               *Many2One  `xmlrpc:"leave_timesheet_project_id,omitempty"`
	LeaveTimesheetTaskId                  *Many2One  `xmlrpc:"leave_timesheet_task_id,omitempty"`
	LockConfirmedPo                       *Bool      `xmlrpc:"lock_confirmed_po,omitempty"`
	MassMailingMailServerId               *Many2One  `xmlrpc:"mass_mailing_mail_server_id,omitempty"`
	MassMailingOutgoingMailServer         *Bool      `xmlrpc:"mass_mailing_outgoing_mail_server,omitempty"`
	ModuleAccount3WayMatch                *Bool      `xmlrpc:"module_account_3way_match,omitempty"`
	ModuleAccountAccountant               *Bool      `xmlrpc:"module_account_accountant,omitempty"`
	ModuleAccountBankStatementImportCamt  *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omitempty"`
	ModuleAccountBankStatementImportCsv   *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omitempty"`
	ModuleAccountBankStatementImportOfx   *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omitempty"`
	ModuleAccountBankStatementImportQif   *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omitempty"`
	ModuleAccountBatchPayment             *Bool      `xmlrpc:"module_account_batch_payment,omitempty"`
	ModuleAccountBudget                   *Bool      `xmlrpc:"module_account_budget,omitempty"`
	ModuleAccountCheckPrinting            *Bool      `xmlrpc:"module_account_check_printing,omitempty"`
	ModuleAccountIntrastat                *Bool      `xmlrpc:"module_account_intrastat,omitempty"`
	ModuleAccountInvoiceExtract           *Bool      `xmlrpc:"module_account_invoice_extract,omitempty"`
	ModuleAccountPayment                  *Bool      `xmlrpc:"module_account_payment,omitempty"`
	ModuleAccountPlaid                    *Bool      `xmlrpc:"module_account_plaid,omitempty"`
	ModuleAccountReports                  *Bool      `xmlrpc:"module_account_reports,omitempty"`
	ModuleAccountSepa                     *Bool      `xmlrpc:"module_account_sepa,omitempty"`
	ModuleAccountSepaDirectDebit          *Bool      `xmlrpc:"module_account_sepa_direct_debit,omitempty"`
	ModuleAccountTaxcloud                 *Bool      `xmlrpc:"module_account_taxcloud,omitempty"`
	ModuleAccountYodlee                   *Bool      `xmlrpc:"module_account_yodlee,omitempty"`
	ModuleAuthLdap                        *Bool      `xmlrpc:"module_auth_ldap,omitempty"`
	ModuleAuthOauth                       *Bool      `xmlrpc:"module_auth_oauth,omitempty"`
	ModuleBaseGengo                       *Bool      `xmlrpc:"module_base_gengo,omitempty"`
	ModuleBaseGeolocalize                 *Bool      `xmlrpc:"module_base_geolocalize,omitempty"`
	ModuleBaseImport                      *Bool      `xmlrpc:"module_base_import,omitempty"`
	ModuleCrmIapLead                      *Bool      `xmlrpc:"module_crm_iap_lead,omitempty"`
	ModuleCrmIapLeadEnrich                *Bool      `xmlrpc:"module_crm_iap_lead_enrich,omitempty"`
	ModuleCrmIapLeadWebsite               *Bool      `xmlrpc:"module_crm_iap_lead_website,omitempty"`
	ModuleCurrencyRateLive                *Bool      `xmlrpc:"module_currency_rate_live,omitempty"`
	ModuleDelivery                        *Bool      `xmlrpc:"module_delivery,omitempty"`
	ModuleDeliveryBpost                   *Bool      `xmlrpc:"module_delivery_bpost,omitempty"`
	ModuleDeliveryDhl                     *Bool      `xmlrpc:"module_delivery_dhl,omitempty"`
	ModuleDeliveryEasypost                *Bool      `xmlrpc:"module_delivery_easypost,omitempty"`
	ModuleDeliveryFedex                   *Bool      `xmlrpc:"module_delivery_fedex,omitempty"`
	ModuleDeliveryUps                     *Bool      `xmlrpc:"module_delivery_ups,omitempty"`
	ModuleDeliveryUsps                    *Bool      `xmlrpc:"module_delivery_usps,omitempty"`
	ModuleGoogleCalendar                  *Bool      `xmlrpc:"module_google_calendar,omitempty"`
	ModuleGoogleDrive                     *Bool      `xmlrpc:"module_google_drive,omitempty"`
	ModuleGoogleSpreadsheet               *Bool      `xmlrpc:"module_google_spreadsheet,omitempty"`
	ModuleHrAttendance                    *Bool      `xmlrpc:"module_hr_attendance,omitempty"`
	ModuleHrOrgChart                      *Bool      `xmlrpc:"module_hr_org_chart,omitempty"`
	ModuleHrPresence                      *Bool      `xmlrpc:"module_hr_presence,omitempty"`
	ModuleHrSkills                        *Bool      `xmlrpc:"module_hr_skills,omitempty"`
	ModuleHrTimesheet                     *Bool      `xmlrpc:"module_hr_timesheet,omitempty"`
	ModuleInterCompanyRules               *Bool      `xmlrpc:"module_inter_company_rules,omitempty"`
	ModuleL10NEuService                   *Bool      `xmlrpc:"module_l10n_eu_service,omitempty"`
	ModulePad                             *Bool      `xmlrpc:"module_pad,omitempty"`
	ModulePartnerAutocomplete             *Bool      `xmlrpc:"module_partner_autocomplete,omitempty"`
	ModuleProcurementJit                  *Selection `xmlrpc:"module_procurement_jit,omitempty"`
	ModuleProductEmailTemplate            *Bool      `xmlrpc:"module_product_email_template,omitempty"`
	ModuleProductExpiry                   *Bool      `xmlrpc:"module_product_expiry,omitempty"`
	ModuleProductMargin                   *Bool      `xmlrpc:"module_product_margin,omitempty"`
	ModuleProjectForecast                 *Bool      `xmlrpc:"module_project_forecast,omitempty"`
	ModuleProjectTimesheetHolidays        *Bool      `xmlrpc:"module_project_timesheet_holidays,omitempty"`
	ModuleProjectTimesheetSynchro         *Bool      `xmlrpc:"module_project_timesheet_synchro,omitempty"`
	ModulePurchaseProductMatrix           *Bool      `xmlrpc:"module_purchase_product_matrix,omitempty"`
	ModulePurchaseRequisition             *Bool      `xmlrpc:"module_purchase_requisition,omitempty"`
	ModuleSaleAmazon                      *Bool      `xmlrpc:"module_sale_amazon,omitempty"`
	ModuleSaleCoupon                      *Bool      `xmlrpc:"module_sale_coupon,omitempty"`
	ModuleSaleMargin                      *Bool      `xmlrpc:"module_sale_margin,omitempty"`
	ModuleSaleProductConfigurator         *Bool      `xmlrpc:"module_sale_product_configurator,omitempty"`
	ModuleSaleProductMatrix               *Bool      `xmlrpc:"module_sale_product_matrix,omitempty"`
	ModuleSaleQuotationBuilder            *Bool      `xmlrpc:"module_sale_quotation_builder,omitempty"`
	ModuleSnailmailAccount                *Bool      `xmlrpc:"module_snailmail_account,omitempty"`
	ModuleStockBarcode                    *Bool      `xmlrpc:"module_stock_barcode,omitempty"`
	ModuleStockDropshipping               *Bool      `xmlrpc:"module_stock_dropshipping,omitempty"`
	ModuleStockLandedCosts                *Bool      `xmlrpc:"module_stock_landed_costs,omitempty"`
	ModuleStockPickingBatch               *Bool      `xmlrpc:"module_stock_picking_batch,omitempty"`
	ModuleStockSms                        *Bool      `xmlrpc:"module_stock_sms,omitempty"`
	ModuleVoip                            *Bool      `xmlrpc:"module_voip,omitempty"`
	ModuleWebUnsplash                     *Bool      `xmlrpc:"module_web_unsplash,omitempty"`
	ModuleWebsiteSaleDigital              *Bool      `xmlrpc:"module_website_sale_digital,omitempty"`
	PaperformatId                         *Many2One  `xmlrpc:"paperformat_id,omitempty"`
	PartnerAutocompleteInsufficientCredit *Bool      `xmlrpc:"partner_autocomplete_insufficient_credit,omitempty"`
	PoDoubleValidation                    *Selection `xmlrpc:"po_double_validation,omitempty"`
	PoDoubleValidationAmount              *Float     `xmlrpc:"po_double_validation_amount,omitempty"`
	PoLead                                *Float     `xmlrpc:"po_lead,omitempty"`
	PoLock                                *Selection `xmlrpc:"po_lock,omitempty"`
	PoOrderApproval                       *Bool      `xmlrpc:"po_order_approval,omitempty"`
	PortalConfirmationPay                 *Bool      `xmlrpc:"portal_confirmation_pay,omitempty"`
	PortalConfirmationSign                *Bool      `xmlrpc:"portal_confirmation_sign,omitempty"`
	PredictiveLeadScoringFields           *Relation  `xmlrpc:"predictive_lead_scoring_fields,omitempty"`
	PredictiveLeadScoringFieldsStr        *String    `xmlrpc:"predictive_lead_scoring_fields_str,omitempty"`
	PredictiveLeadScoringStartDate        *Time      `xmlrpc:"predictive_lead_scoring_start_date,omitempty"`
	PredictiveLeadScoringStartDateStr     *String    `xmlrpc:"predictive_lead_scoring_start_date_str,omitempty"`
	ProductPricelistSetting               *Selection `xmlrpc:"product_pricelist_setting,omitempty"`
	ProductVolumeVolumeInCubicFeet        *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omitempty"`
	ProductWeightInLbs                    *Selection `xmlrpc:"product_weight_in_lbs,omitempty"`
	ProjectTimeModeId                     *Many2One  `xmlrpc:"project_time_mode_id,omitempty"`
	PurchaseTaxId                         *Many2One  `xmlrpc:"purchase_tax_id,omitempty"`
	QrCode                                *Bool      `xmlrpc:"qr_code,omitempty"`
	QuotationValidityDays                 *Int       `xmlrpc:"quotation_validity_days,omitempty"`
	ReportFooter                          *String    `xmlrpc:"report_footer,omitempty"`
	ResourceCalendarId                    *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	SaleTaxId                             *Many2One  `xmlrpc:"sale_tax_id,omitempty"`
	SecurityLead                          *Float     `xmlrpc:"security_lead,omitempty"`
	ShowBlacklistButtons                  *Bool      `xmlrpc:"show_blacklist_buttons,omitempty"`
	ShowEffect                            *Bool      `xmlrpc:"show_effect,omitempty"`
	ShowLineSubtotalsTaxSelection         *Selection `xmlrpc:"show_line_subtotals_tax_selection,omitempty"`
	SnailmailColor                        *Bool      `xmlrpc:"snailmail_color,omitempty"`
	SnailmailCover                        *Bool      `xmlrpc:"snailmail_cover,omitempty"`
	SnailmailDuplex                       *Bool      `xmlrpc:"snailmail_duplex,omitempty"`
	StockMailConfirmationTemplateId       *Many2One  `xmlrpc:"stock_mail_confirmation_template_id,omitempty"`
	StockMoveEmailValidation              *Bool      `xmlrpc:"stock_move_email_validation,omitempty"`
	StockMoveSmsValidation                *Bool      `xmlrpc:"stock_move_sms_validation,omitempty"`
	StockSmsConfirmationTemplateId        *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omitempty"`
	TaxCalculationRoundingMethod          *Selection `xmlrpc:"tax_calculation_rounding_method,omitempty"`
	TaxCashBasisJournalId                 *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omitempty"`
	TaxExigibility                        *Bool      `xmlrpc:"tax_exigibility,omitempty"`
	TemplateId                            *Many2One  `xmlrpc:"template_id,omitempty"`
	TimesheetEncodeUomId                  *Many2One  `xmlrpc:"timesheet_encode_uom_id,omitempty"`
	UnsplashAccessKey                     *String    `xmlrpc:"unsplash_access_key,omitempty"`
	UseInvoiceTerms                       *Bool      `xmlrpc:"use_invoice_terms,omitempty"`
	UsePoLead                             *Bool      `xmlrpc:"use_po_lead,omitempty"`
	UseQuotationValidityDays              *Bool      `xmlrpc:"use_quotation_validity_days,omitempty"`
	UseSecurityLead                       *Bool      `xmlrpc:"use_security_lead,omitempty"`
	UserDefaultRights                     *Bool      `xmlrpc:"user_default_rights,omitempty"`
	VatCheckVies                          *Bool      `xmlrpc:"vat_check_vies,omitempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	ids, err := c.CreateResConfigSettingss([]*ResConfigSettings{rcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettingss(rcss []*ResConfigSettings) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcss {
		vv = append(vv, v)
	}
	return c.Create(ResConfigSettingsModel, vv, nil)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs, nil)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	return &((*rcss)[0]), nil
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResConfigSettingsModel, criteria, options)
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
