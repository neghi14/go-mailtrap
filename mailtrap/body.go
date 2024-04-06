package mailtrap

type (
	sendEmailBody struct {
		Headers           map[string]string `json:"headers,omitempty"`
		From              identifier        `json:"from"`
		To                []identifier      `json:"to"`
		CC                []identifier      `json:"cc,omitempty"`
		BCC               []identifier      `json:"bcc,omitempty"`
		Attachments       []attachment      `json:"attachments,omitempty"`
		Subject           string            `json:"subject,omitempty"`
		Text              string            `json:"text,omitempty"`
		HTML              string            `json:"html,omitempty"`
		Category          string            `json:"category,omitempty"`
		TemplateUUID      string            `json:"template_uuid,omitempty"`
		TemplateVariables map[string]string `json:"template_variables,omitempty"`
		CustomVariables   map[string]string `json:"custom_variables,omitempty"`
	}
	identifier struct {
		Email string `json:"email"`
		Name  string `json:"name,omitempty"`
	}
	attachment struct {
		Content     string `json:"content"`
		Type        string `json:"type"`
		Filename    string `json:"filename"`
		Disposition string `json:"disposition"`
		ContentID   string `json:"content_id"`
	}
)

func (b *sendEmailBody) AddHeader(name, value string) *sendEmailBody {
	b.Headers[name] = value
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) SetFrom(name, email string) *sendEmailBody {
	b.From = identifier{
		Name:  name,
		Email: email,
	}
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) AddTo(name, email string) *sendEmailBody {
	newVar := identifier{
		Name:  name,
		Email: email,
	}
	result := append(b.To, newVar)
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                result,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) AddCC(name, email string) *sendEmailBody {
	newVar := identifier{
		Name:  name,
		Email: email,
	}
	result := append(b.CC, newVar)
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                result,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) AddBCC(name, email string) *sendEmailBody {
	newVar := identifier{
		Name:  name,
		Email: email,
	}
	result := append(b.BCC, newVar)
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               result,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) AddAttachment(content, content_type, filename, disposition, content_id string) *sendEmailBody {
	newVar := attachment{
		Content:     content,
		Type:        content_type,
		Filename:    filename,
		Disposition: disposition,
		ContentID:   content_id,
	}
	result := append(b.Attachments, newVar)
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       result,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) SetSubject(subject string) *sendEmailBody {
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) SetText(text string) *sendEmailBody {
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) SetHTML(html string) *sendEmailBody {
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              html,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) SetCategory(category string) *sendEmailBody {
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) SetTemplateID(template_id string) *sendEmailBody {
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      template_id,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) AddTemplateVariable(name, value string) *sendEmailBody {
	b.TemplateVariables[name] = value
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}

func (b *sendEmailBody) AddCustomVariable(name, value string) *sendEmailBody {
	b.CustomVariables[name] = value
	return &sendEmailBody{
		Headers:           b.Headers,
		From:              b.From,
		To:                b.To,
		CC:                b.CC,
		BCC:               b.BCC,
		Attachments:       b.Attachments,
		Subject:           b.Subject,
		Text:              b.Text,
		HTML:              b.HTML,
		Category:          b.Category,
		TemplateUUID:      b.TemplateUUID,
		TemplateVariables: b.TemplateVariables,
		CustomVariables:   b.CustomVariables,
	}
}
