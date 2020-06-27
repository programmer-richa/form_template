package models

type Form struct {
	Name            string
	ID              string
	Action          string
	Method          string
	Fields          []*FormField
	PrefixFieldName string
	EncryptType     string
	ErrorMsg        string
}

// NewForm
func NewForm(Name string, ID string, Action string, Method string, Fields []*FormField, PrefixFieldName string, EncryptType string) *Form {
	return &Form{Name: Name,
		ID:              ID,
		Action:          Action,
		Method:          Method,
		Fields:          Fields,
		PrefixFieldName: PrefixFieldName,
		EncryptType:     EncryptType,
		ErrorMsg:        "",
	}
}

func (f *Form) SetErrorMsg(msg string) {
	f.ErrorMsg = msg
}

type FormField struct {
	Label       string
	ID          string
	Value       string
	Type        string
	Placeholder string
	IsRequired  bool
	MinLength   int
	MaxLength   int
}

// NewFormField accepts 2 parameters : event name and tasks to be performed
func NewFormField(Label string, ID string, Value string, Type string, Placeholder string, IsRequired bool, MinLength int, MaxLength int) *FormField {
	return &FormField{Label: Label,
		ID:          ID,
		Value:       Value,
		Type:        Type,
		Placeholder: Placeholder,
		IsRequired:  IsRequired,
		MinLength:   MinLength,
		MaxLength:   MaxLength,
	}
}

type Event struct {
	Name  string
	Value string
}

// NewEvent accepts 2 parameters : event name and tasks to be performed
func NewEvent(name string, value string) *Event {
	return &Event{Name: name, Value: value}
}
