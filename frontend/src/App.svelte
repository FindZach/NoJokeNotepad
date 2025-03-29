<script>
  import { onMount, onDestroy } from 'svelte';

  let content = '';
  let showFileDropdown = false;
  let showEditDropdown = false;
  let showFormatDropdown = false;
  let wailsReady = false;
  let currentFile = '';
  let isMaximized = false;
  let runtimeError = false;
  let wordWrapEnabled = true;
  let isDragging = false;
  let dragStartX = 0;
  let dragStartY = 0;

  // Wait for Wails runtime to be ready
  onMount(() => {
    console.log('App.svelte mounted');
    checkWailsRuntime();

    // Add click event listener to detect outside clicks
    document.addEventListener('click', handleOutsideClick);
    // Add mousemove and mouseup listeners for drag-to-restore
    document.addEventListener('mousemove', handleMouseMove);
    document.addEventListener('mouseup', handleMouseUp);
  });

  onDestroy(() => {
    // Clean up event listeners when the component is destroyed
    document.removeEventListener('click', handleOutsideClick);
    document.removeEventListener('mousemove', handleMouseMove);
    document.removeEventListener('mouseup', handleMouseUp);
  });

  function checkWailsRuntime(attempts = 0) {
    const maxAttempts = 20; // 20 attempts * 500ms = 10 seconds
    if (window.runtime) {
      console.log('Wails runtime (window.runtime) detected');
      console.log('Wails runtime methods available:', Object.keys(window.runtime));
      wailsReady = true;
      // Check initial maximized state
      window.runtime.WindowIsMaximised().then(maximized => {
        isMaximized = maximized;
        console.log('Initial window state:', isMaximized ? 'Maximized' : 'Restored');
      });
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

  function toggleWordWrap() {
    wordWrapEnabled = !wordWrapEnabled;
    showFormatDropdown = false; // Close the dropdown
  }

  // Window control functions
  function minimizeWindow() {
    if (wailsReady && window.runtime) {
      window.runtime.WindowMinimise();
      console.log('Minimize button clicked');
    } else {
      console.error('Wails runtime not ready for minimize');
      alert('Wails runtime not ready for minimize');
    }
  }

  function toggleMaximizeWindow() {
    if (wailsReady && window.runtime) {
      if (isMaximized) {
        window.runtime.WindowUnmaximise();
        isMaximized = false;
        console.log('Restore button clicked');
      } else {
        window.runtime.WindowMaximise();
        isMaximized = true;
        console.log('Maximize button clicked');
      }
    } else {
      console.error('Wails runtime not ready for maximize/restore');
      alert('Wails runtime not ready for maximize/restore');
    }
  }

  function closeWindow() {
    if (wailsReady && window.runtime) {
      window.runtime.Quit();
      console.log('Close button clicked');
    } else {
      console.error('Wails runtime not ready for close');
      alert('Wails runtime not ready for close');
    }
  }

  // Double-click handler for the title bar
  function handleTitleBarDoubleClick() {
    if (wailsReady && window.runtime) {
      toggleMaximizeWindow();
    }
  }

  // Drag-to-restore handler
  function handleMouseDown(event) {
    if (wailsReady && window.runtime && isMaximized) {
      // Prevent dragging if the click is on a non-draggable area (title or buttons)
      if (event.target.classList.contains('title') || event.target.classList.contains('window-control')) {
        return;
      }

      // Restore the window
      window.runtime.WindowUnmaximise();
      isMaximized = false;
      console.log('Restoring window via drag');

      // Get the current mouse position
      dragStartX = event.screenX;
      dragStartY = event.screenY;
      isDragging = true;

      // Set the initial position of the window to follow the mouse
      const windowWidth = 500; // Default width from main.go
      const windowHeight = 400; // Default height from main.go
      const newX = event.screenX - windowWidth / 2;
      const newY = event.screenY - 30; // Account for the title bar height
      window.runtime.WindowSetPosition(newX, newY);
    }
  }

  function handleMouseMove(event) {
    if (isDragging && wailsReady && window.runtime) {
      const newX = event.screenX - dragStartX;
      const newY = event.screenY - dragStartY;
      window.runtime.WindowSetPosition(newX, newY);
    }
  }

  function handleMouseUp() {
    isDragging = false;
  }

  // Keyboard handler for menu items
  function handleKeydown(event, menuType) {
    if (event.key === 'Enter' || event.key === ' ') {
      event.preventDefault();
      if (menuType === 'file') {
        showFileDropdown = !showFileDropdown;
        showEditDropdown = false;
        showFormatDropdown = false;
      } else if (menuType === 'edit') {
        showEditDropdown = !showEditDropdown;
        showFileDropdown = false;
        showFormatDropdown = false;
      } else if (menuType === 'format') {
        showFormatDropdown = !showFormatDropdown;
        showFileDropdown = false;
        showEditDropdown = false;
      }
    }
  }

  // Handle clicks outside the menu to close dropdowns
  function handleOutsideClick(event) {
    const menuItems = document.querySelectorAll('.menu-item');
    const dropdowns = document.querySelectorAll('.dropdown');
    let isClickInsideMenu = false;

    // Check if the click is inside a menu item or dropdown
    menuItems.forEach(menuItem => {
      if (menuItem.contains(event.target)) {
        isClickInsideMenu = true;
      }
    });

    dropdowns.forEach(dropdown => {
      if (dropdown.contains(event.target)) {
        isClickInsideMenu = true;
      }
    });

    // If the click is outside the menu, close all dropdowns
    if (!isClickInsideMenu) {
      showFileDropdown = false;
      showEditDropdown = false;
      showFormatDropdown = false;
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
    --wails-draggable: drag; /* Make the title bar draggable */
  }

  .title {
    font-size: 14px;
    font-weight: bold;
    color: #000; /* Set text color to black for contrast */
    --wails-draggable: no-drag; /* Prevent dragging on the title text */
    cursor: default; /* Prevent cursor change on hover */
  }

  .window-controls {
    display: flex;
    --wails-draggable: no-drag; /* Prevent dragging on the controls */
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

  .menu-item:focus {
    outline: 1px solid #000;
  }

  .dropdown {
    position: absolute;
    top: 24px; /* Match the height of the menu bar to position the dropdown below the tab */
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
    white-space: pre-wrap; /* Default to wrapping */
  }

  .editor.no-wrap {
    white-space: pre; /* Disable wrapping */
    overflow-x: auto; /* Enable horizontal scrolling */
  }
</style>

<div class="container">
  <div class="title-bar" on:dblclick={handleTitleBarDoubleClick} on:mousedown={handleMouseDown}>
    <div class="title">NoJoke Notepad</div>
    <div class="window-controls">
      <button class="window-control" on:click={minimizeWindow} disabled={!wailsReady}>−</button>
      <button class="window-control" on:click={toggleMaximizeWindow} disabled={!wailsReady}>{isMaximized ? '🗗' : '🗖'}</button>
      <button class="window-control close" on:click={closeWindow} disabled={!wailsReady}>×</button>
    </div>
  </div>
  <div class="menu-bar">
    <div class="menu-item" tabindex="0" on:click={() => { showFileDropdown = !showFileDropdown; showEditDropdown = false; showFormatDropdown = false }} on:keydown={(e) => handleKeydown(e, 'file')}>
      File
    </div>
    <div class="dropdown" class:visible={showFileDropdown}>
      <button class="dropdown-item" on:click={newFile}>New</button>
      <button class="dropdown-item" on:click={openFile}>Open</button>
      <button class="dropdown-item" on:click={saveFile}>Save</button>
      <button class="dropdown-item" on:click={saveFileAs}>Save As</button>
    </div>
    <div class="menu-item" tabindex="0" on:click={() => { showEditDropdown = !showEditDropdown; showFileDropdown = false; showFormatDropdown = false }} on:keydown={(e) => handleKeydown(e, 'edit')}>
      Edit
    </div>
    <div class="dropdown" class:visible={showEditDropdown}>
      <button class="dropdown-item" on:click={cutText}>Cut</button>
      <button class="dropdown-item" on:click={copyText}>Copy</button>
      <button class="dropdown-item" on:click={pasteText}>Paste</button>
    </div>
    <div class="menu-item" tabindex="0" on:click={() => { showFormatDropdown = !showFormatDropdown; showFileDropdown = false; showEditDropdown = false }} on:keydown={(e) => handleKeydown(e, 'format')}>
      Format
    </div>
    <div class="dropdown" class:visible={showFormatDropdown}>
      <button class="dropdown-item" on:click={toggleWordWrap}>
        Word Wrap {wordWrapEnabled ? '✔' : ''}
      </button>
    </div>
  </div>
  <textarea class="editor" class:no-wrap={!wordWrapEnabled} bind:value={content}></textarea>
</div>