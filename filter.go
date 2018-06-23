package unifilter

// FilterClass is a class of characters to filter, pass these to the FilterInput
// function to get your string filtered. Use the OR operator to commbine classes
type FilterClass int

const (
	filterControlCodes FilterClass = 1 << iota
	filterInvisiblePunctuation
	filterIllegalUnixPaths
	filterIllegalWindowsPaths
)

const (
	// FilterControlCodes filters all unicode and ASCII control codes of control
	// block C0 and C1, in addition to the DELETE character
	FilterControlCodes = filterControlCodes

	// FilterInvisiblePunctuation filters most invisible punctuation characters
	// like zero-width spaces, byte-order marks, text direction indicators, etc
	FilterInvisiblePunctuation = filterInvisiblePunctuation

	// FilterIllegalUnixPaths filters characters which are not allowed in Unix
	// directory and file names. This includes only the NUL terminator and
	// forward slash
	FilterIllegalUnixPaths = filterIllegalUnixPaths

	// FilterIllegalWindowsPaths filters characters which are not allowed in
	// Windows path names. This consisists of a lot of punctuation
	FilterIllegalWindowsPaths = filterControlCodes | filterIllegalWindowsPaths
)

// IllegalRune is a rune that got filtered
type IllegalRune struct {
	Rune     rune
	Position int
	Class    FilterClass
}

// IllegalRuneMap is a map of illegal runes returned by the FilterInput
// function. The index represents the position in the string where the rune was
// found
type IllegalRuneMap map[int]IllegalRune

// FilterInput checks the input string for runes which are part of the
// filterClasses you passed. You can pass multiple filter classes by separating
// them with a pipe, like this:
//  FilterInput("yoghurt", unifilter.FilterControlCodes|unifilter.FilterInvisiblePunctuation)
func FilterInput(input string, filterClass FilterClass) (ok bool, frunes IllegalRuneMap) {
	frunes = make(IllegalRuneMap)
	if filterClass&filterControlCodes != 0 {
		mapAppend(frunes, filterRuneSlice(input, controlRunes, filterControlCodes))
	}
	if filterClass&filterInvisiblePunctuation != 0 {
		mapAppend(frunes, filterRuneSlice(input, controlRunes, filterInvisiblePunctuation))
	}
	if filterClass&filterIllegalUnixPaths != 0 {
		mapAppend(frunes, filterRuneSlice(input, controlRunes, filterIllegalUnixPaths))
	}
	if filterClass&filterIllegalWindowsPaths != 0 {
		mapAppend(frunes, filterRuneSlice(input, controlRunes, filterIllegalWindowsPaths))
	}
	if len(frunes) == 0 {
		return true, nil
	}
	return false, frunes
}

func filterRuneSlice(input string, runes []rune, class FilterClass) (irunes IllegalRuneMap) {
	irunes = make(IllegalRuneMap)
	for i, runeVal := range input {
		for i2 := range runes {
			if runeVal == runes[i2] {
				irunes[i] = IllegalRune{
					Rune:     runeVal,
					Position: i,
					Class:    class,
				}
			}
		}
	}
	return irunes
}

func mapAppend(target, values IllegalRuneMap) {
	for k, v := range values {
		target[k] = v
	}
}
