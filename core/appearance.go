package core

// Appearance encapsulates the attributes that control how a module
// on is rendered on the Bar.
// Appearance is typically defined through the configuration file.
type Appearance struct {
	Color       *Color     `yaml:"color"`
	Border      *Border    `yaml:"border"`
	Separator   *Separator `yaml:"separator"`
	Format      string     `yaml:"format"`
	FormatShort string     `yaml:"format_short"`
	Align       string     `yaml:"align"`
	MinWidth    int        `yaml:"min_width"`
	Urgent      bool       `yaml:"urgent"`
}

// Color represents a color on the Bar, and must be
// specified using the #RRGGBBAA or #RRGGBB notation.
type Color struct {
	Foreground string `yaml:"foreground"`
	Background string `yaml:"background"`
}

// Border encapsulates the border settings for a module.
type Border struct {
	Top    int `yaml:"top"`
	Bottom int `yaml:"bottom"`
	Left   int `yaml:"left"`
	Right  int `yaml:"right"`
}

// Separator contains the swaybar seprarator related settings.
type Separator struct {
	Enabled    bool `yaml:"enabled"`
	BlockWidth int  `yaml:"block_width"`
}

// Variant contains an appearance variant for the module.
// The appearance Variant is only applied, when the Regex expression
// matches the actual status string of the module.
// This can be used to conditionally control the appearance of
// the modules, for example setting a different background color,
// or different format text based on the status string.
type Variant struct {
	Appearance *Appearance `yaml:"appearance"`
	Regex      string      `yaml:"regex"`
}
