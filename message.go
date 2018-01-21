package gabel

import "fmt"

/*
baseLabelingMessageTmpl is the colored base message to be displayed for labeling
ex)
 「text(colored)」
Please Input Label.
*/
var baseLabelingMessageTmpl = "\x1b[33m 「%s」\x1b[0m\nPlease input label. "

/*
labelingMessageTmpl returns the colored message to be displayed for labeling
ex)
1
 「text(colored)」
Please Input Label {exist == [1]} or {not exist == [0]} or {Modify == mod}:
*/
func labelingMessageTmpl(labels []Label) string {
	return "\n%d\n" + baseLabelingMessageTmpl + availableChoicesMessage(labels) + "or \x1b[4m{Modify == mod}\x1b[0m:"
}

/*
modfityLabelingMessageTmp returns the colored modify message to be displayed for labeling
ex)
1(colored)
 「csv text(colored)」
Please Input Label {exist == [1]} or {not exist == [0]}:
*/

func modifyLabelingMessageTmpl(labels []Label) string {
	return "\nx1b[41m%d\x1b[0m\n" + baseLabelingMessageTmpl + availableChoicesMessage(labels) + ":"
}

/*
availableChoicesMessage create available choices
ex)
{message == [values]} or {message == [values]} or ....
*/
func availableChoicesMessage(labels []Label) string {
	tmpl := "\x1b[4m{%s\x1b[0m == %d} "

	var msg string
	for _, label := range labels {
		msg += fmt.Sprintf(tmpl, label.Name, label.Value) + "or "
	}

	return msg[:len(msg)-3]
}
