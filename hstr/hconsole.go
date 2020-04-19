package hstr

import (
	"bytes"
	"strconv"
	"strings"
)

type FontStyle struct {
	effects []Effect
}

func NewFontStyle(effects ...Effect) *FontStyle {
	style := &FontStyle{
		effects: effects,
	}
	return style
}

func (s *FontStyle) Render(message string) string {
	buf := &bytes.Buffer{}
	buf.WriteString("\033")
	buf.WriteString("[")
	var effects []string
	for _, effect := range s.effects {
		effects = append(effects, strconv.Itoa(int(effect)))
	}
	buf.WriteString(strings.Join(effects, ";"))
	buf.WriteString("m")
	buf.WriteString(message)
	buf.WriteString("\033[0m")

	return buf.String()
}

type Effect int

const (
	FormatSetClose     Effect = 0
	FormatSetBold      Effect = 1
	FormatSetDim       Effect = 2
	FormatSetUnderline Effect = 4
	FormatSetBlink     Effect = 5
	FormatSetReverse   Effect = 7
	FormatSetHidden    Effect = 8
)

const (
	FormatResetAll       Effect = 0
	FormatResetBold      Effect = 21
	FormatResetDim       Effect = 22
	FormatResetUnderline Effect = 24
	FormatResetBlink     Effect = 25
	FormatResetReverse   Effect = 27
	FormatResetHidden    Effect = 28
)

const (
	ForegroundDefault      Effect = 39
	ForegroundBlack        Effect = 30
	ForegroundRed          Effect = 31
	ForegroundGreen        Effect = 32
	ForegroundYellow       Effect = 33
	ForegroundBlue         Effect = 34
	ForegroundMagenta      Effect = 35
	ForegroundCyan         Effect = 36
	ForegroundLightGray    Effect = 37
	ForegroundDarkGray     Effect = 90
	ForegroundLightRed     Effect = 91
	ForegroundLightGreen   Effect = 92
	ForegroundLightYellow  Effect = 93
	ForegroundLightBlue    Effect = 94
	ForegroundLightMagenta Effect = 95
	ForegroundLightCyan    Effect = 96
	ForegroundWhite        Effect = 97
)

const (
	BackgroundDefault      Effect = 49
	BackgroundBlack        Effect = 40
	BackgroundRed          Effect = 41
	BackgroundGreen        Effect = 42
	BackgroundYellow       Effect = 43
	BackgroundBlue         Effect = 44
	BackgroundMagenta      Effect = 45
	BackgroundCyan         Effect = 46
	BackgroundLightGray    Effect = 47
	BackgroundDarkGray     Effect = 100
	BackgroundLightRed     Effect = 101
	BackgroundLightGreen   Effect = 102
	BackgroundLightYellow  Effect = 103
	BackgroundLightBlue    Effect = 104
	BackgroundLightMagenta Effect = 105
	BackgroundLightCyan    Effect = 106
	BackgroundWhite        Effect = 107
)
