# 📝 NoJokeNotepad

**NoJokeNotepad** is a lightweight text editor built with [Wails](https://wails.io) (Go backend + Svelte frontend), inspired by the simplicity of Windows Notepad—but with a modern twist. It features a clean UI, basic file operations, and a modular design that makes it easy to maintain and extend.

---

## 🚀 Features

### 🎨 User Interface

#### 🔷 Custom Title Bar (`TitleBar.svelte`)
- **App Icon** in the top-left corner (`icon.ico`)
- **Dynamic Title**: Shows current file path (e.g., `C:\path\file.txt - NoJoke Notepad`) or just `NoJoke Notepad`
- **Window Controls**:
    - Minimize (−)
    - Maximize/Restore (🗖 / 🗗)
    - Close (×)
- **Double-click to Maximize/Restore**
- **Drag to Restore** from maximized (using Wails native drag)
- **Non-selectable Title** (prevents unwanted text highlighting)

#### 📋 Menu Bar (`MenuBar.svelte`, `MenuItem.svelte`, `Dropdown.svelte`)
- **File Menu**:
    - New
    - Open
    - Save
    - Save As
- **Edit Menu**:
    - Cut
    - Copy
    - Paste
- **Format Menu**:
    - Word Wrap toggle
- **Dropdown Behavior**:
    - Closes on item click, outside click, or new menu open
    - Accessible via keyboard (`Enter`, `Space`)

#### ✍️ Editor (`Editor.svelte`)
- `<textarea>` based editor
- **Word Wrap Toggle**: `white-space: pre-wrap` or `pre`
- **Cursor Behavior**: "I" beam over text, default elsewhere
- **No Resizing / Overflow**: Layout is fully responsive with no scrollbars

#### 📊 Status Bar (`StatusBar.svelte`)
- Always visible at bottom
- **Displays**:
    - Line Count
    - Character Length
    - File Type (`.txt`, etc.)
    - Encoding (currently `UTF-8` placeholder)

---

### 🗂 File Operations
- **New**: Clears editor and resets file state
- **Open**: Load `.txt` file using dialog
- **Save**: Save to current file or prompt for a new one
- **Save As**: Save as a new file and update path

---

### 📝 Text Editing
- Standard editing: typing, deleting
- **Cut**: Copy + Clear
- **Copy** / **Paste**
- **Word Wrap** toggle for long lines

---

### 🪟 Window Management
- **Frameless Window** (`frameless: true` in `main.go`)
- **Double-click** or **Button** to Maximize/Restore
- **Drag from Title Bar** to restore and reposition
- **Minimize / Close** supported

---

### 🖼 Icon Support
- App icon set in `wails.json` (`assets/appicon.ico`)
- Displayed in the custom title bar

---

### ♿ Accessibility
- **Keyboard Navigation** for menus (`Enter`, `Space`)
- **Focus Styles** for keyboard accessibility (`outline: 1px solid`)

---

### 🧱 Scalability & Structure

#### 🔧 Modular Components
- `TitleBar.svelte`
- `MenuBar.svelte`
- `MenuItem.svelte`
- `Dropdown.svelte`
- `Editor.svelte`
- `StatusBar.svelte`

#### 🗂 Organized File Structure
- Components: `frontend/src/components/`
- Styles: `frontend/src/styles/`
- Public Assets: `frontend/public/`

#### 🎨 Centralized Styling
- All global styles in `styles/app.css`

---

### ❗ Error Handling
- **Wails Runtime Check**:
    - Polls for `window.runtime`
    - Shows error message if runtime fails after timeout:
      > "Error: Wails runtime failed to load. Please restart the app."

---

### 🐞 Debugging Support
- Console logs for:
    - Runtime initialization
    - Window state changes
    - File operations
    - Drag events

---

## 📚 Summary

| Category          | Features |
|-------------------|----------|
| **UI**            | Custom title bar, icon, dynamic title, dropdown menus, editor, status bar |
| **File Ops**      | New, Open, Save, Save As |
| **Text Editing**  | Basic editing, Cut/Copy/Paste, Word Wrap |
| **Window Mgmt**   | Frameless, Maximize/Restore, Drag, Minimize, Close |
| **Icons**         | App icon in title bar & window |
| **Accessibility** | Keyboard navigation, focus styles |
| **Scalability**   | Modular components, organized structure, clean CSS |
| **Error Handling**| Runtime check + graceful fallback |
| **Debugging**     | Helpful logs in console |

---

## 🧠 Potential Future Features

- ✅ Keyboard Shortcuts (e.g., Ctrl+S, Ctrl+O, Ctrl+N)
- 📍 Status Bar Enhancements (cursor position, file size)
- 🔤 File Encoding Support (UTF-8, ANSI detection)
- 🧲 Snap-to-Edge (maximize on drag to screen top)
- ↩️ Undo/Redo support
- 🔎 Find/Replace feature
- 🔠 Font Customization (size, family)
- 🌙 Dark Mode toggle

---

## 📦 Built With
- **[Wails](https://wails.io)** – Go-powered desktop app framework
- **[Svelte](https://svelte.dev)** – Frontend UI framework

---

## 🖥 Screenshot (optional)
_Add a screenshot here if available to showcase the UI._

---

## 📄 License
MIT 

---

> _“NoJokeNotepad: Not a joke, just simple and powerful.”_
