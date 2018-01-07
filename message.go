package gabel

/*
labelingMessageTmpl returns the colored message to be displayed

ex)
1
 「csv text」
Please Input Your Label {exist == [0]} or {not exist == [1]} %
*/
var labelingMessageTmpl = "\n%d\n\x1b[33m 「%s」\x1b[0m\nPlease input label. \x1b[4m{%s == %v}\x1b[0m or \x1b[4m{%s == %v}\x1b[0m "
