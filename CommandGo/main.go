package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// Editor merupakan objek yang akan dimodifikasi oleh perintah
type Editor struct {
	text      *widget.Entry
	clipboard string
}

// Command merupakan antarmuka untuk semua perintah
type Command interface {
	Execute() bool
	Undo()
}

// CopyCommand merupakan implementasi perintah untuk menyalin teks
type CopyCommand struct {
	editor   *Editor
	backup   string
	executed bool
}

// Execute melakukan operasi penyalinan teks
func (c *CopyCommand) Execute() bool {
	if c.executed {
		return false
	}

	c.backup = c.editor.text.Text
	selected := c.editor.text.SelectedText()
	c.editor.clipboard = selected
	c.executed = true
	return true
}

// Undo membatalkan operasi penyalinan teks
func (c *CopyCommand) Undo() {
	c.editor.text.SetText(c.backup)
	c.executed = false
}

// PasteCommand merupakan implementasi perintah untuk menempelkan teks
type PasteCommand struct {
	editor   *Editor
	backup   string
	executed bool
}

// Execute melakukan operasi penempelan teks
func (c *PasteCommand) Execute() bool {
	if c.executed || c.editor.clipboard == "" {
		return false
	}

	c.backup = c.editor.text.Text
	c.editor.text.SetText(c.editor.text.Text + c.editor.clipboard)
	c.executed = true
	return true
}

// Undo membatalkan operasi penempelan teks
func (c *PasteCommand) Undo() {
	c.editor.text.SetText(c.backup)
	c.executed = false
}

// CommandHistory menyimpan riwayat perintah yang dijalankan
type CommandHistory struct {
	history []Command
}

// Push menambahkan perintah baru ke dalam riwayat
func (c *CommandHistory) Push(cmd Command) {
	c.history = append(c.history, cmd)
}

// Undo membatalkan perintah terakhir yang dijalankan
func (c *CommandHistory) Undo() {
	if len(c.history) > 0 {
		lastCmd := c.history[len(c.history)-1]
		lastCmd.Undo()
		c.history = c.history[:len(c.history)-1]
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Text Editor")

	textEntry := widget.NewMultiLineEntry()
	editor := &Editor{text: textEntry}
	history := &CommandHistory{}

	copyButton := widget.NewButton("Copy", func() {
		copyCmd := &CopyCommand{editor: editor}
		history.Push(copyCmd)
		copyCmd.Execute()
	})

	pasteButton := widget.NewButton("Paste", func() {
		pasteCmd := &PasteCommand{editor: editor}
		history.Push(pasteCmd)
		pasteCmd.Execute()
	})

	undoButton := widget.NewButton("Undo", func() {
		history.Undo()
	})

	content := widget.NewVBox(textEntry, copyButton, pasteButton, undoButton)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}