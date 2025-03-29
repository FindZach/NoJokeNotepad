<script>
  import { onMount } from 'svelte';
  import TitleBar from './components/TitleBar.svelte';
  import MenuBar from './components/MenuBar.svelte';
  import Editor from './components/Editor.svelte';
  import StatusBar from './components/StatusBar.svelte';

  let content = '';
  let wailsReady = false;
  let currentFile = '';
  let runtimeError = false;
  let wordWrapEnabled = true;

  onMount(() => {
    console.log('App.svelte mounted');
    checkWailsRuntime();
  });

  function checkWailsRuntime(attempts = 0) {
    const maxAttempts = 20; // 20 attempts * 500ms = 10 seconds
    if (window.runtime) {
      console.log('Wails runtime (window.runtime) detected');
      console.log('Wails runtime methods available:', Object.keys(window.runtime));
      wailsReady = true;
      initializeApp();
    } else if (attempts >= maxAttempts) {
      console.error('Wails runtime (window.runtime) not available after maximum attempts');
      runtimeError = true;
      content = 'Error: Wails runtime failed to load. Please restart the app.';
    } else {
      console.log('Wails runtime (window.runtime) not available, retrying in 500ms... (attempt', attempts + 1, 'of', maxAttempts, ')');
      console.log('window.runtime:', window.runtime);
      setTimeout(() => checkWailsRuntime(attempts + 1), 500);
    }
  }

  async function initializeApp() {
    try {
      content = await window.go.main.App.NewFile();
      currentFile = '';
      console.log('Initialized app with content:', content);
    } catch (err) {
      console.error('Error calling NewFile:', err);
      content = 'Error: Failed to initialize app';
    }
  }

  async function newFile() {
    if (!wailsReady) {
      console.error('Wails runtime not ready');
      alert('Wails runtime not ready');
      return;
    }
    try {
      content = await window.go.main.App.NewFile();
      currentFile = '';
      console.log('NewFile called, content reset');
    } catch (err) {
      console.error('Error calling NewFile:', err);
      alert('Failed to create new file: ' + err);
    }
  }

  async function openFile() {
    if (!wailsReady) {
      console.error('Wails runtime not ready');
      alert('Wails runtime not ready');
      return;
    }
    try {
      content = await window.go.main.App.OpenFile();
      currentFile = content ? 'Untitled' : '';
      console.log('OpenFile called, content set to:', content);
    } catch (err) {
      console.error('Error calling OpenFile:', err);
      alert('Failed to open file: ' + err);
    }
  }

  async function saveFile() {
    if (!wailsReady) {
      console.error('Wails runtime not ready');
      alert('Wails runtime not ready');
      return;
    }
    try {
      const filePath = await window.go.main.App.SaveFile(content);
      currentFile = filePath || 'Untitled';
      console.log('SaveFile called with content:', content);
      alert('File saved to: ' + currentFile);
    } catch (err) {
      console.error('Error calling SaveFile:', err);
      alert('Failed to save file: ' + err);
    }
  }

  async function saveFileAs() {
    if (!wailsReady) {
      console.error('Wails runtime not ready');
      alert('Wails runtime not ready');
      return;
    }
    try {
      const filePath = await window.go.main.App.SaveFileAs(content);
      currentFile = filePath || 'Untitled';
      console.log('SaveFileAs called with content:', content);
      alert('File saved to: ' + currentFile);
    } catch (err) {
      console.error('Error calling SaveFileAs:', err);
      alert('Failed to save file: ' + err);
    }
  }

  function cutText() {
    navigator.clipboard.writeText(content);
    content = '';
  }

  function copyText() {
    navigator.clipboard.writeText(content);
  }

  function pasteText() {
    navigator.clipboard.readText().then(text => {
      content += text;
    });
  }

  function toggleWordWrap() {
    wordWrapEnabled = !wordWrapEnabled;
  }
</script>

<div class="container">
  <TitleBar {wailsReady} />
  <MenuBar
          {wordWrapEnabled}
          onNewFile={newFile}
          onOpenFile={openFile}
          onSaveFile={saveFile}
          onSaveFileAs={saveFileAs}
          onCut={cutText}
          onCopy={copyText}
          onPaste={pasteText}
          onToggleWordWrap={toggleWordWrap}
  />
  <Editor {content} {wordWrapEnabled} on:change={(e) => (content = e.target.value)} />
  <StatusBar {content} {currentFile} />
</div>