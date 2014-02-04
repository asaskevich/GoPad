package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "io/ioutil"
)

var edit *walk.TextEdit

func readFromFile() {
	contents,_ := ioutil.ReadFile("file.txt")
	edit.SetText(string(contents))
}

func saveToFile() {
	ioutil.WriteFile("file.txt", []byte(edit.Text()), 0x777)
}

func main() {
    MainWindow{
        Title:   "GoPad",
        MinSize: Size{600, 400},
		Layout:  VBox{},
        Children: []Widget{
            TextEdit{AssignTo: &edit},
            HSplitter{
        		MaxSize: Size{600, 30},
                Children: []Widget{
                	PushButton{
            		    Text: "Copy",
            		    OnClicked: func() {
                            walk.Clipboard().SetText(edit.Text());
                        },
            		},
            		PushButton{
                		Text: "Paste",
                		OnClicked: func() {
                            if text, err := walk.Clipboard().Text(); err == nil {
                                edit.SetText(text)
                            }
                        },
            		},
            		PushButton{
                		Text: "Load",
                		OnClicked: readFromFile,
            		},
            		PushButton{
                		Text: "Save",
                		OnClicked: saveToFile,
            		},
                },
            },
        },
    }.Run()
}