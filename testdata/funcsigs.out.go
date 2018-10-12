package test

func someSignatureThatIs100Chars____________________________________(someArg, someOtherArg string) {
}

func someSignatureThatIs101Chars_____________________________________(
	someArg, someOtherArg string,
) {
}

func someSignatureWithResults(someArg, someOtherArg string) (string, string, string, string, bool) {
	return "", "", "", "", false
}

func someSignatureWithLongResults(
	someArg, someOtherArg string,
) (string, string, string, string, string) {
	return "", "", "", "", ""
}

func someSigWithLongArgs(
	someArg string,
	someOtherArg string,
	someLoooooooooooooooooooooooooooooooooooooooooooooooooooooooog int,
) {
}

func someSigWithLongArgsAndElidedTypeShorthand(
	someArg, someOtherArg string,
	someLoooooooooooooooooooooooooooooooooooooooooooooooooooooooog int,
) {
}

type fooObj struct{}

func (f *fooObj) finalTxnStatsLocked() (duration, restarts int64, status int64) {
	return
}
