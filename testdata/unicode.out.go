package test

// café café café café café café café café café café café café café café café
// café
func docstringWithMultibyteRunes() string

// Le café était délicieux et très chaud, mais le résumé était trop naïf pour
// moi ici.
func docstringWithAccents() string

// This signature's byte length exceeds the wrap column but its rune length does
// not.
func multibyteSignature(résumé, café, déjà, naïve, piñata, jalapeño, señor, müller int) error
