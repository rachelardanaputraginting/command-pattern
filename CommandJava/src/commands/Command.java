package commands;

public abstract class Command {
    public Editor editor;
    private String backup;

    public Command(Editor editor) {
        this.editor = editor;
    }

    void backup() {
        backup = editor.textField.getText();
    }

    public void undo() {
        if (backup != null) {
            editor.textField.setText(backup);
        }
    }

    public abstract boolean execute();
}
