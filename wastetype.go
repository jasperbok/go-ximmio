package ximmio

// configuration contains the settings for a configOption.
//
// The two companies this library has been tested with have a different
// subset of these fields. That means that it is likely that there are
// more fields available than what's coded here, but that they haven't
// been encountered yet.
type configuration struct {
	Desc          string `json:"desc"`
	Disposal      string `json:"disposal"`
	Evening       string `json:"evening"`
	Morning       string `json:"morning"`
	WasteName     string `json:"wasteName"`
	ContainerName string `json:"containerName"`
	MainColor     string `json:"mainColor"`
	EmailSubject  string `json:"emailSubject"`
	EmailEvening  string `json:"emailEvening"`
	EmailMorning  string `json:"EmailMorning"`
}

// configOption represents the configuration for a type of waste.
type configOption struct {
	ConfigName     string        `json:"ConfigName"`
	Configurations configuration `json:"Configurations"`
	Communities    []string      `json:"Communities"`
}

func (c configOption) toWasteType() WasteType {
	return WasteType{
		WasteName:     c.Configurations.WasteName,
		ContainerName: c.Configurations.ContainerName,
		ConfigName:    c.ConfigName,
		Description:   c.Configurations.Desc,
		Disposal:      c.Configurations.Disposal,
		Evening:       c.Configurations.Evening,
		Morning:       c.Configurations.Morning,
		MainColor:     c.Configurations.MainColor,
		EmailSubject:  c.Configurations.EmailSubject,
		EmailEvening:  c.Configurations.EmailEvening,
		EmailMorning:  c.Configurations.EmailMorning,
		Communities:   c.Communities,
	}
}

// WasteType represent a type of waste and its configuration.
type WasteType struct {
	WasteName     string // The 'friendly' name, e.g. 'Textiel'.
	ContainerName string // The 'friendly' name of the container, e.g. 'GFT-container'.
	ConfigName    string // The codename, e.g. 'TEXTIEL'.
	Description   string // HTML-formatted description.
	Disposal      string // Text describing how to dispose of the waste.
	Evening       string // Text message informing about collection tomorrow.
	Morning       string // Text message informing about collection today.
	MainColor     string // Hex color code.
	EmailSubject  string
	EmailEvening  string // HTML email body.
	EmailMorning  string // HTML email body.
	Communities   []string
}
