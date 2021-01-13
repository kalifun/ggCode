package font

import "fmt"

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func NewColorFont(str string) *FontColor {
	return &FontColor{Str: str}
}

type FontColor struct {
	Str string
}

func (fc *FontColor) Green() string {
	return SetColor(fc.Str, 0, 0, TextGreen)
}

func (fc *FontColor) Red() string {
	return SetColor(fc.Str, 0, 0, TextRed)
}

func (fc *FontColor) Yellow() string {
	return SetColor(fc.Str, 0, 0, TextYellow)
}

func (fc *FontColor) Blue() string {
	return SetColor(fc.Str, 0, 0, TextBlue)
}

func (fc *FontColor) Magenta() string {
	return SetColor(fc.Str, 0, 0, TextMagenta)
}

func (fc *FontColor) Cyan() string {
	return SetColor(fc.Str, 0, 0, TextCyan)
}

func (fc *FontColor) White() string {
	return SetColor(fc.Str, 0, 0, TextWhite)
}

func (fc *FontColor) Black() string {
	return SetColor(fc.Str, 0, 0, TextBlack)
}

func SetColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
