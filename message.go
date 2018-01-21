package gabel

import (
	"fmt"
	"unsafe"
)

/*
baseMessageTmpl is the colored base message to be displayed for labeling
ex)
 「text(colored)」
Please enter label.
*/
var baseMessageTmpl = "\x1b[33m 「%s」\x1b[0m\nPlease enter label. "

/*
messageTmpl returns the colored message to be displayed for labeling
ex)
1
 「text(colored)」
Please enter label {exist == [1]} or {not exist == [0]} or {Modify == mod}:
*/
func messageTmpl(id int, text string, labels []Label) string {
	return fmt.Sprintf("\n%d\n"+baseMessageTmpl+generateAvailableChoicesMessage(labels)+"or \x1b[4m{Modify == mod}\x1b[0m:", id, text)
}

/*
modfityLabelingMessageTmp returns the colored modify message to be displayed for labeling
ex)
1(colored)
 「csv text(colored)」
Please enter Label {exist == [1]} or {not exist == [0]}:
*/
func modifyMessageTmpl(id int, text string, labels []Label) string {
	return fmt.Sprintf("\nx1b[41m%d\x1b[0m\n"+baseMessageTmpl+generateAvailableChoicesMessage(labels)+":", id, text)
}

/*
generateAvailableChoicesMessage create available choices
ex)
{message == [values]} or {message == [values]} or ....
*/
func generateAvailableChoicesMessage(labels []Label) string {
	tmpl := "\x1b[4m{%s\x1b[0m == %d} "

	var msg []byte
	for i, label := range labels {
		msg = append(msg, fmt.Sprintf(tmpl, label.Name, label.Value)...)
		if i < len(labels)-1 {
			msg = append(msg, "or "...)
		}
	}

	return *(*string)(unsafe.Pointer(&msg))
}
