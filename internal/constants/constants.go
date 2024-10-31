package constants

// Alignment represents the alignment of segments
type Alignment int

const (
	// AlignLeft aligns segments to the left
	AlignLeft Alignment = iota
	// AlignRight aligns segments to the right
	AlignRight
)

const (
	// MinUnsignedInteger minimum unsigned integer
	MinUnsignedInteger uint = 0
	// MaxUnsignedInteger maximum unsigned integer
	MaxUnsignedInteger = ^MinUnsignedInteger
	// MaxInteger maximum integer
	MaxInteger = int(MaxUnsignedInteger >> 1)
)
