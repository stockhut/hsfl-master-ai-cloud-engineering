package termcolor

// how terminal colors work: https://stackoverflow.com/questions/4842424/list-of-ansi-color-escape-sequences

// color codes: https://en.wikipedia.org/wiki/ANSI_escape_code#3-bit_and_4-bit

// TextColorizer is a function that turns a string into a colored string
type TextColorizer func(s string) string

type FgColor string

func (f FgColor) String() string {
	return string(f)
}

type BgColor string

func (f BgColor) String() string {
	return string(f)
}

const (
	FgUndefined     = FgColor("")
	FgBlack         = FgColor("30")
	FgRed           = FgColor("31")
	FgGreen         = FgColor("32")
	FgYellow        = FgColor("33")
	FgBlue          = FgColor("34")
	FgMagenta       = FgColor("35")
	FgCyan          = FgColor("36")
	FgWhite         = FgColor("37")
	FgBrightBlack   = FgColor("90")
	FgBrightRed     = FgColor("91")
	FgBrightGreen   = FgColor("92")
	FgBrightYellow  = FgColor("93")
	FgBrightBlue    = FgColor("94")
	FgBrightMagenta = FgColor("95")
	FgBrightCyan    = FgColor("96")
	FgBrightWhite   = FgColor("97")
)

const (
	BgUndefined     = BgColor("")
	BgBlack         = BgColor("40")
	BgRed           = BgColor("41")
	BgGreen         = BgColor("42")
	BgYellow        = BgColor("43")
	BgBlue          = BgColor("44")
	BgMagenta       = BgColor("45")
	BgCyan          = BgColor("46")
	BgWhite         = BgColor("47")
	BgBrightBlack   = BgColor("100")
	BgBrightRed     = BgColor("101")
	BgBrightGreen   = BgColor("102")
	BgBrightYellow  = BgColor("103")
	BgBrightBlue    = BgColor("104")
	BgBrightMagenta = BgColor("105")
	BgBrightCyan    = BgColor("106")
	BgBrightWhite   = BgColor("107")
)

const beginColorSequence = "\033["
const endColorSequence = "m"

const resetColor = beginColorSequence + "0" + endColorSequence

func setColor(c string) string {
	return beginColorSequence + c + endColorSequence
}

// Fg creates a TextColorizer that wraps a given string in the commands needed
// to set the text/foreground color to `color`
func Fg(color FgColor) TextColorizer {
	return func(s string) string {
		return setColor(color.String()) + s + resetColor
	}
}

// Bg creates a TextColorizer that wraps a given string in the commands needed
// to set the background color to `color`
func Bg(color BgColor) TextColorizer {
	return func(s string) string {
		return setColor(color.String()) + s + resetColor
	}
}

// ColorBuilder is used to create a
type ColorBuilder struct {
	bgColor BgColor
	fgColor FgColor
}

// NewBuilder creates a new ColorBuilder instance
func NewBuilder() *ColorBuilder {
	return &ColorBuilder{}
}

// Text sets the foreground color for the builder
func (cb *ColorBuilder) Text(color FgColor) *ColorBuilder {
	cb.fgColor = color

	return cb
}

// Background sets the background color for the builder
func (cb *ColorBuilder) Background(color BgColor) *ColorBuilder {
	cb.bgColor = color

	return cb
}

// Build creates the TextColorizer
func (cb *ColorBuilder) Build() TextColorizer {
	return func(s string) string {

		if cb.fgColor != FgUndefined {
			s = Fg(cb.fgColor)(s)
		}
		if cb.bgColor != BgUndefined {
			s = Bg(cb.bgColor)(s)
		}
		return s
	}
}
