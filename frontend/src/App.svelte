<script>
  import { onMount } from 'svelte';

  let content = '';
  let showFileDropdown = false;
  let showEditDropdown = false;
  let wailsReady = false;
  let currentFile = '';
  let isMaximized = false;

  // Wait for Wails runtime to be ready
  onMount(() => {
    console.log('App.svelte mounted');
    checkWailsRuntime();
  });

  function checkWailsRuntime() {
    if (window.wails && window.wails.runtime) {
      console.log('Wails runtime (window.wails.runtime) detected');
      wailsReady = true;
      initializeApp();
    } else {
      console.log('Wails runtime (window.wails.runtime) not available, retrying in 500ms...');
      setTimeout(checkWailsRuntime, 500);
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
      showFileDropdown = false; // Close the dropdown
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
      showFileDropdown = false; // Close the dropdown
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
      showFileDropdown = false; // Close the dropdown
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
      showFileDropdown = false; // Close the dropdown
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
    showEditDropdown = false; // Close the dropdown
  }

  function copyText() {
    navigator.clipboard.writeText(content);
    showEditDropdown = false; // Close the dropdown
  }

  function pasteText() {
    navigator.clipboard.readText().then(text => {
      content += text;
      showEditDropdown = false; // Close the dropdown
    });
  }

  // Window control functions
  function minimizeWindow() {
    if (wailsReady && window.wails && window.wails.runtime) {
      window.wails.runtime.WindowMinimise();
      console.log('Minimize button clicked');
    } else {
      console.error('Wails runtime not ready for minimize');
      alert('Wails runtime not ready for minimize');
    }
  }

  function toggleMaximizeWindow() {
    if (wailsReady && window.wails && window.wails.runtime) {
      if (isMaximized) {
        window.wails.runtime.WindowUnmaximise();
        isMaximized = false;
        console.log('Restore button clicked');
      } else {
        window.wails.runtime.WindowMaximise();
        isMaximized = true;
        console.log('Maximize button clicked');
      }
    } else {
      console.error('Wails runtime not ready for maximize/restore');
      alert('Wails runtime not ready for maximize/restore');
    }
  }

  function closeWindow() {
    if (wailsReady && window.wails && window.wails.runtime) {
      window.wails.runtime.Quit();
      console.log('Close button clicked');
    } else {
      console.error('Wails runtime not ready for close');
      alert('Wails runtime not ready for close');
    }
  }
</script>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  }

  .container {
    display: flex;
    flex-direction: column;
    height: 100vh;
    width: 100vw;
  }

  .title-bar {
    background-color: #f0f0f0;
    height: 30px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 10px;
    -webkit-app-region: drag; /* Make the title bar draggable */
  }

  .title {
    font-size: 14px;
    font-weight: bold;
    -webkit-app-region: no-drag; /* Prevent dragging on the title text */
  }

  .window-controls {
    display: flex;
    -webkit-app-region: no-drag; /* Prevent dragging on the controls */
  }

  .window-control {
    width: 30px;
    height: 30px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    font-size: 14px;
    border: none;
    background: none;
    color: #000;
  }

  .window-control:hover {
    background-color: #e0e0e0;
  }

  .window-control.close:hover {
    background-color: #ff4444;
    color: white;
  }

  .window-control:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .menu-bar {
    background-color: #f0f0f0;
    border-bottom: 1px solid #ccc;
    padding: 0;
    display: flex;
    align-items: center;
    height: 24px;
  }

  .menu-item {
    position: relative;
    padding: 2px 10px;
    cursor: pointer;
    font-size: 14px;
    user-select: none;
  }

  .menu-item:hover {
    background-color: #e0e0e0;
  }

  .dropdown {
    position: absolute;
    top: 24px;
    left: 0;
    background-color: #f0f0f0;
    border: 1px solid #ccc;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
    z-index: 1000;
    display: none;
    min-width: 150px;
  }

  .dropdown.visible {
    display: block;
  }

  .dropdown-item {
    display: block;
    padding: 5px 10px;
    font-size: 14px;
    cursor: pointer;
    border: none;
    background: none;
    text-align: left;
    width: 100%;
  }

  .dropdown-item:hover {
    background-color: #d0d0d0;
  }

  .editor {
    flex: 1;
    border: none;
    outline: none;
    resize: none;
    padding: 5px;
    font-size: 14px;
    font-family: 'Consolas', 'Courier New', monospace;
    background-color: #fff;
    width: 100%;
    box-sizing: border-box;
    line-height: 1.5;
  }
</style>

<div class="container">
  <div class="title-bar">
    <div class="title">NoJoke Notepad</div>
    <div class="window-controls">
      <button class="window-control" on:click={minimizeWindow} disabled={!wailsReady}>−</button>
      <button class="window-control" on:click={toggleMaximizeWindow} disabled={!wailsReady}>{isMaximized ? '🗗' : '🗖'}</button>
      <button class="window-control close" on:click={closeWindow} disabled={!wailsReady}>×</button>
    </div>
  </div>
  <div class="menu-bar">
    <div class="menu-item" on:click={() => { showFileDropdown = !showFileDropdown; showEditDropdown = false }}>
      File
    </div>
    <div class="dropdown" class:visible={showFileDropdown}>
      <button class="dropdown-item" on:click={newFile}>New</button>
      <button class="dropdown-item" on:click={openFile}>Open</button>
      <button class="dropdown-item" on:click={saveFile}>Save</button>
      <button class="dropdown-item" on:click={saveFileAs}>Save As</button>
    </div>
    <div class="menu-item" on:click={() => { showEditDropdown = !showEditDropdown; showFileDropdown = false }}>
      Edit
    </div>
    <div class="dropdown" class:visible={showEditDropdown}>
      <button class="dropdown-item" on:click={cutText}>Cut</button>
      <button class="dropdown-item" on:click={copyText}>Copy</button>
      <button class="dropdown-item" on:click={pasteText}>Paste</button>
    </div>
  </div>
  <textarea class="editor" bind:value={content}></textarea>
</div>