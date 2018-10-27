package test

func someSignatureThatIs100Chars____________________________________(someArg, someOtherArg string) {
}

func someSignatureThatIs101Chars_____________________________________(someArg, someOtherArg string) {
}

func someSignatureWithResults(someArg, someOtherArg string) (string, string, string, string, bool) {
	return "", "", "", "", false
}

func someSignatureWithLongResults(someArg, someOtherArg string) (string, string, string, string, string) {
	return "", "", "", "", ""
}

func someSignatureWithVeryLongResults(
	someArg, someOtherArg string,
) (iAmLong string, andIAmLongToo string, andIAmAlsoLong string, butEspeciallyThisParameterIsLong string) {
	return "", "", "", "", ""
}

func someSigWithLongArgs(someArg string, someOtherArg string, someLoooooooooooooooooooooooooooooooooooooooooooooooooooooooog int) {
}

func someSigWithLongArgsAndElidedTypeShorthand(
	someArg, someOtherArg string, someLoooooooooooooooooooooooooooooooooooooooooooooooooooooooog int,
) {
}

func someSigWithMisIndentedArgs________________________(
	  thereAreTwoSpacesBeforeThisArgWhenThereShouldBeJustATab string,
) (
	  thereAreTwoSpacesBeforeThisArgWhenThereShouldBeJustATab string,
	  weDefinitelyNeedAnotherLineForTheseReturnParams int,
)

func someSigWithMisIndentedArgsNoNames________________________(
	  func(iAmSuperLooooooooong_____________________________________________________________________),
) (
	  func(iAmSuperLooooooooong_____________________________________________________________________),
	  string,
)

type fooObj struct{}

func (f *fooObj) finalTxnStatsLocked() (duration, restarts int64, status int64) {
	return
}

func (f *fooObj) someSigWithMisIndentedArgs________________________(
	  thereAreTwoSpacesBeforeThisArgWhenThereShouldBeJustATab string,
) (
	  thereAreTwoSpacesBeforeThisArgWhenThereShouldBeJustATab string,
	  weDefinitelyNeedAnotherLineForTheseReturnParams int,
)

func (f *fooObj) someSigWithMisIndentedArgsNoNames________________________(
	  func(iAmSuperLooooooooong_____________________________________________________________________),
) (
	  func(iAmSuperLooooooooong_____________________________________________________________________),
	  string,
)

func updateStatsOnPut(
	key []byte,
	origMetaKeySize, origMetaValSize,
	metaKeySize, metaValSize int64,
	orig, meta *int,
) int

func mvccPutProto(
	ctx string,
	engine string,
	ms string, // can be nil as the key is unreplicated and doesn't affect stats
	key string,


	timestamp string,
	// the following blank line should not be preserved

	// long-form documentation about txn
	// that spans multiple lines
	txn string,
	// message has a comment both above
	message string, // and on the line
) error

func mvccPutProtoShort(
	// This comment should force a reflow so that it can more clearly refer to
	// just a and not the following parameter.
	a string, someLoooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong int,
)

func paramsComments1(/* ba */ a int /* aa */, b string /* ab */, c int /* ac */)

func paramsComments2(a int /* after a */, b string, /* after
	b
	with
	newlines
	*/

	c []byte,
)

func resultsComments1() (/* ba */ a int /* aa */, b string /* ab */, c int /* ac */)

func resultsComments2() (a int /* after a */, b string, /* after
	b
	with
	newlines
	*/

	c []byte,
)

// docstring should stay as a docstring
func docstringWithParam(a string)

// docstring should stay as a docstring
func docstringWithResult() string
