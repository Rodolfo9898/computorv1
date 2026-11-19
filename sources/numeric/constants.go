package numeric

// Unit is the unit roundoff of a double: a rounded result is within Unit times
// its size of the exact one.
const Unit = 0x1p-53

// MinNormal is the smallest normal double: below it, a result loses digits.
const MinNormal = 0x1p-1022
