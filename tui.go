package rogue

import "fmt"

const (
	Escape           = "\033["
	Suffix           = "m"
	ResetCode        = "0"
	BoldCode         = "1"
	EndBoldCode      = "21"
	DimCode          = "2"
	EndDimCode       = "22"
	UnderlineCode    = "4"
	EndUnderlineCode = "24"
	ItalicCode       = "3"
	EndItalicCode    = "23"
	BlackCode        = "30"
	DarkBlueCode     = "90"
	WhiteCode        = "97"
	RedCode          = "31"
	OrangeCode       = "91"
	GreenYellowCode  = "32"
	GreenCode        = "33"
	DarkGrayCode     = "92"
	YellowCode       = "33"
	GrayCode         = "93"
	BlueCode         = "34"
	LightGrayCode    = "94"
	PurpleCode       = "35"
	LightPurpleCode  = "95"
	CyanCode         = "36"
	OffWhiteCode     = "96"
	ResetColorCode   = "39"
)

const (
	RESET       = Escape + ResetCode + Suffix
	BOLD        = Escape + BoldCode + Suffix
	DIM         = Escape + DimCode + Suffix
	UNDERLINE   = Escape + UnderlineCode + Suffix
	ITALIC      = Escape + ItalicCode + Suffix
	BLACK       = Escape + BlackCode + Suffix
	DARKBLUE    = Escape + DarkBlueCode + Suffix
	WHITE       = Escape + WhiteCode + Suffix
	RED         = Escape + RedCode + Suffix
	ORANGE      = Escape + OrangeCode + Suffix
	GREENYELLOW = Escape + GreenYellowCode + Suffix
	GREEN       = Escape + GreenCode + Suffix
	DARKGRAY    = Escape + DarkGrayCode + Suffix
	YELLOW      = Escape + YellowCode + Suffix
	LIGHTGRAY   = Escape + LightGrayCode + Suffix
	BLUE        = Escape + BlueCode + Suffix
	GRAY        = Escape + GrayCode + Suffix
	PURPLE      = Escape + PurpleCode + Suffix
	LIGHTPURPLE = Escape + LightPurpleCode + Suffix
	CYAN        = Escape + CyanCode + Suffix
	OFFWHITE    = Escape + OffWhiteCode + Suffix
)

func Bold(text string) string {
	return BOLD + text + RESET
}

func Dim(text string) string {
	return BOLD + text + RESET
}

func Underline(text string) string {
	return UNDERLINE + text + RESET
}

func Italic(text string) string {
	return ITALIC + text + RESET
}

func Black(text string) string {
	return BLACK + text + RESET
}

func DarkBlue(text string) string {
	return DARKBLUE + text + RESET
}

func White(text string) string {
	return WHITE + text + RESET
}

func Red(text string) string {
	return RED + text + RESET
}
func Orange(text string) string {
	return ORANGE + text + RESET
}

func GreenYellow(text string) string {
	return GREENYELLOW + text + RESET
}

func Green(text string) string {
	return GREEN + text + RESET
}

func DarkGray(text string) string {
	return DARKGRAY + text + RESET
}

func Yellow(text string) string {
	return YELLOW + text + RESET
}

func Gray(text string) string {
	return GRAY + text + RESET
}

func Blue(text string) string {
	return BLUE + text + RESET
}

func LightGray(text string) string {
	return LIGHTGRAY + text + RESET
}

func Purple(text string) string {
	return PURPLE + text + RESET
}

func Cyan(text string) string {
	return CYAN + text + RESET
}

func OffWhite(text string) string {
	return OFFWHITE + text + RESET
}

func PrintBanner() {
	fmt.Println(Blue("Rogue") + White(":") + Bold(OffWhite("GO")) + OffWhite(Bold(" v")) + White("0.1.0") + " -" + OffWhite(" an open-source") + White(" counter-strike:go") + OffWhite(" hack-engine"))
	fmt.Println(OffWhite(Bold("==============================================================")))
}
