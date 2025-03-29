<script>
    import { onMount, onDestroy } from 'svelte';

    export let wailsReady = false;
    export let currentFile = '';

    let isMaximized = false;
    let isDragging = false;

    // Construct the dynamic title
    $: title = currentFile ? `${currentFile} - NoJoke Notepad` : 'NoJoke Notepad';

    onMount(() => {
        if (wailsReady && window.runtime) {
            window.runtime.WindowIsMaximised().then(maximized => {
                isMaximized = maximized;
                console.log('Initial window state:', isMaximized ? 'Maximized' : 'Restored');
            });
        }

        document.addEventListener('mousemove', handleMouseMove);
        document.addEventListener('mouseup', handleMouseUp);
    });

    onDestroy(() => {
        document.removeEventListener('mousemove', handleMouseMove);
        document.removeEventListener('mouseup', handleMouseUp);
    });

    function minimizeWindow() {
        if (wailsReady && window.runtime) {
            window.runtime.WindowMinimise();
            console.log('Minimize button clicked');
        } else {
            console.error('Wails runtime not ready for minimize');
            alert('Wails runtime not ready for minimize');
        }
    }

    let lastWindowSize = { width: 500, height: 400 };

    function toggleMaximizeWindow() {
        if (wailsReady && window.runtime) {
            if (isMaximized) {
                window.runtime.WindowUnmaximise();
                window.runtime.WindowSetSize(lastWindowSize.width, lastWindowSize.height);
                isMaximized = false;
                console.log('Restore button clicked');
            } else {
                // Store the current size before maximizing
                window.runtime.WindowGetSize().then(size => {
                    lastWindowSize = size;
                    window.runtime.WindowMaximise();
                    isMaximized = true;
                    console.log('Maximize button clicked');
                });
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

    function handleTitleBarDoubleClick() {
        if (wailsReady && window.runtime) {
            toggleMaximizeWindow();
        }
    }

    function handleMouseDown(event) {
        if (wailsReady && window.runtime) {
            // Prevent dragging if the click is on a non-draggable area (title or buttons)
            if (event.target.classList.contains('title') || event.target.classList.contains('window-control') || event.target.classList.contains('app-icon')) {
                return;
            }

            // If the window is maximized, restore it when dragging starts
            if (isMaximized) {
                window.runtime.WindowUnmaximise();
                isMaximized = false;
                console.log('Restoring window via drag');
            }

            // Let Wails handle the dragging natively via --wails-draggable: drag
            isDragging = true;
        }
    }

    function handleMouseMove(event) {
        // No need to manually set position; Wails handles dragging
        if (isDragging) {
            console.log('Dragging window');
        }
    }

    function handleMouseUp() {
        isDragging = false;
        console.log('Drag ended');
    }
</script>

<div class="title-bar" on:dblclick={handleTitleBarDoubleClick} on:mousedown={handleMouseDown}>
    <img src="/images/icon.ico" alt="App Icon" class="app-icon" />
    <div class="title">{title}</div>
    <div class="window-controls">
        <button class="window-control" on:click={minimizeWindow} disabled={!wailsReady}>−</button>
        <button class="window-control" on:click={toggleMaximizeWindow} disabled={!wailsReady}>{isMaximized ? '🗗' : '🗖'}</button>
        <button class="window-control close" on:click={closeWindow} disabled={!wailsReady}>×</button>
    </div>
</div>