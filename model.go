package jmdict

// JMdict represents the top-level <JMdict> element in the JMdict file.
type JMdict struct {
	Entries []*Entry `xml:"entry"`
}

// Entry represents the <entry> XML element in the JMdict file.
type Entry struct {
	Number  int       `xml:"ent_seq"`
	Kanji   []Kanji   `xml:"k_ele"`
	Reading []Reading `xml:"r_ele"`
	Sense   []Sense   `xml:"sense"`
}

// Kanji represents the <k_ele> element in the JMdict file.
type Kanji struct {
	Kanji    string   `xml:"keb"`
	Info     []string `xml:"ke_inf"`
	Priority []string `xml:"ke_pri"`
}

// Reading represents the <r_ele> element in the JMdict file.
type Reading struct {
	Reading  string   `xml:"reb"`
	NoKanji  Presence `xml:"re_nokanji"`
	Restrict []string `xml:"re_restr"`
	Info     []string `xml:"re_inf"`
	Priority []string `xml:"re_pri"`
}

// Sense represents the <sense> element in the JMdict file.
type Sense struct {
	RestrictKanji   []string  `xml:"stagk"`
	RestrictReading []string  `xml:"stagr"`
	PartOfSpeech    []string  `xml:"pos"`
	Xref            []string  `xml:"xref"`
	Antonym         []string  `xml:"ant"`
	Field           []string  `xml:"field"`
	Misc            []string  `xml:"misc"`
	Info            []string  `xml:"s_inf"`
	LSource         []LSource `xml:"lsource"`
	Dialect         []string  `xml:"dial"`
	Gloss           []Gloss   `xml:"gloss"`
}

// LSource represents the <lsource> element in the JMdict file.
type LSource struct {
	Source    string `xml:",chardata"`
	Language  string `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Type      string `xml:"ls_type,attr"`
	WaseiEigo YesNo  `xml:"ls_wasei,attr"`
}

// Gloss represents the <gloss> element in the JMdict file.
type Gloss struct {
	Gloss    string   `xml:",chardata"`
	Gender   string   `xml:"g_gend,attr"`
	Language string   `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Priority []string `xml:"pri"`
}
