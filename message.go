package gabel

/*
baseLabelingMessageTmpl is the colored base message to be displayed for labeling

ex)
 「text(colored)」
Please Input Label {exist == [1]} or {not exist == [0]}
*/
var baseLabelingMessageTmpl = "\x1b[33m 「%s」\x1b[0m\nPlease input label. \x1b[4m{%s == %v}\x1b[0m or \x1b[4m{%s == %v}\x1b[0m or \x1b[4m{Modify == mod}\x1b[0m "

/*
labelingMessageTmpl is the colored message to be displayed for labeling

ex)
1
 「text(colored)」
Please Input Label {exist == [1]} or {not exist == [0]} or {Modify == mod}:
*/
var labelingMessageTmpl = "\n%d\n" + baseLabelingMessageTmpl + "or \x1b[4m{Modify == mod}\x1b[0m: "

/*
modfityLabelingMessageTmp is the colored modify message to be displayed for labeling

ex)
1(colored)
 「csv text(colored)」
Please Input Label {exist == [1]} or {not exist == [0]}:
*/
var modfityLabelingMessageTmp = "\n\x1b[41m%d\x1b[0m\n" + baseLabelingMessageTmpl + ": "
