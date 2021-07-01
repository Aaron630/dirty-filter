package dirtyFilter

import "strings"

var InvalidWord = make(map[string]interface{})

func init() {
	const InvalidWords = " ,~,!,@,#,$,%,^,&,*,(,),_,-,+,=,?,<,>,.,—,，,。,/,\\,|,《,》,？,;,:,：,',‘,；,“,"
	var words []string = strings.Split(InvalidWords, ",")
	for _, v := range words {
		InvalidWord[v] = nil
	}
}

func SetInvalidWords(words []string) {
	if len(words) > 0 {
		InvalidWord = make(map[string]interface{})
	}
	for _, v := range words {
		InvalidWord[v] = nil
	}
}
