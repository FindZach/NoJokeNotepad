<script>
    import { onMount, onDestroy } from 'svelte';
    import * as runtime from "../../wailsjs/runtime/runtime.js";

    export let wailsReady = false;
    let isMaximized = false;
    let isDragging = false;
    let dragStartX = 0;
    let dragStartY = 0;

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

    function handleTitleBarDoubleClick() {
        if (wailsReady && window.runtime) {
            toggleMaximizeWindow();
        }
    }

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
</script>

<div class="title-bar" on:dblclick={handleTitleBarDoubleClick} on:mousedown={handleMouseDown}>
    <div class="title">NoJoke Notepad</div>
    <div class="window-controls">
        <button class="window-control" on:click={minimizeWindow} disabled={!wailsReady}>−</button>
        <button class="window-control" on:click={toggleMaximizeWindow} disabled={!wailsReady}>{isMaximized ? '🗗' : '🗖'}</button>
        <button class="window-control close" on:click={closeWindow} disabled={!wailsReady}>×</button>
    </div>
</div>