package core

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

type Color struct {
	Foreground string `yaml:"foreground"`
	Background string `yaml:"background"`
}

type Border struct {
	Top    int `yaml:"top"`
	Bottom int `yaml:"bottom"`
	Left   int `yaml:"left"`
	Right  int `yaml:"right"`
}

type Separator struct {
	Enabled    bool `yaml:"enabled"`
	BlockWidth int  `yaml:"block_width"`
}

type Alternate struct {
	Appearance *Appearance `yaml:"appearance"`
	Regex      string      `yaml:"regex"`
}
