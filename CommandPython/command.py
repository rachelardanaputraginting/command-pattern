import tkinter as tk
from tkinter import Text, Scrollbar, Button

class Command:
    def execute(self):
        pass

class CopyCommand(Command):
    def __init__(self, editor):
        self.editor = editor
    
    def execute(self):
        self.editor.copy()

class PasteCommand(Command):
    def __init__(self, editor):
        self.editor = editor
    
    def execute(self):
        self.editor.paste()

class EditorApp:
    def __init__(self, root):
        self.root = root
        self.root.title("Text Editor")

        self.text_area = Text(root, wrap="word")
        self.text_area.pack(expand=True, fill="both")

        scroll = Scrollbar(self.text_area, command=self.text_area.yview)
        scroll.pack(side="right", fill="y")
        self.text_area.config(yscrollcommand=scroll.set)

        buttons_frame = tk.Frame(root)
        buttons_frame.pack(pady=10)

        copy_command = CopyCommand(self)
        copy_button = Button(buttons_frame, text="Copy (Ctrl+C)", command=copy_command.execute)
        copy_button.pack(side="left", padx=5)

        paste_command = PasteCommand(self)
        paste_button = Button(buttons_frame, text="Paste (Ctrl+V)", command=paste_command.execute)
        paste_button.pack(side="left", padx=5)

        self.clipboard = ""

    def copy(self):
        self.clipboard = self.text_area.get("sel.first", "sel.last")

    def paste(self):
        self.text_area.insert("insert", self.clipboard)

    def run(self):
        self.root.mainloop()

root = tk.Tk()
app = EditorApp(root)
app.run()
