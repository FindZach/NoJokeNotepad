<script>
    import { onMount, onDestroy } from 'svelte';
    import MenuItem from './MenuItem.svelte';

    let showFileDropdown = false;
    let showEditDropdown = false;
    let showFormatDropdown = false;
    export let wordWrapEnabled = true;
    export let onNewFile = () => {};
    export let onOpenFile = () => {};
    export let onSaveFile = () => {};
    export let onSaveFileAs = () => {};
    export let onCut = () => {};
    export let onCopy = () => {};
    export let onPaste = () => {};
    export let onToggleWordWrap = () => {};

    onMount(() => {
        document.addEventListener('click', handleOutsideClick);
    });

    onDestroy(() => {
        document.removeEventListener('click', handleOutsideClick);
    });

    function handleOutsideClick(event) {
        const menuItems = document.querySelectorAll('.menu-item');
        const dropdowns = document.querySelectorAll('.dropdown');
        let isClickInsideMenu = false;

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

        if (!isClickInsideMenu) {
            showFileDropdown = false;
            showEditDropdown = false;
            showFormatDropdown = false;
        }
    }

    function toggleFileDropdown() {
        showFileDropdown = !showFileDropdown;
        showEditDropdown = false;
        showFormatDropdown = false;
    }

    function toggleEditDropdown() {
        showEditDropdown = !showEditDropdown;
        showFileDropdown = false;
        showFormatDropdown = false;
    }

    function toggleFormatDropdown() {
        showFormatDropdown = !showFormatDropdown;
        showFileDropdown = false;
        showEditDropdown = false;
    }

    function handleItemClick(callback) {
        return () => {
            callback();
            showFileDropdown = false;
            showEditDropdown = false;
            showFormatDropdown = false;
        };
    }
</script>

<div class="menu-bar">
    <MenuItem label="File" isOpen={showFileDropdown} onToggle={toggleFileDropdown}>
        <button class="dropdown-item" on:click={handleItemClick(onNewFile)}>New</button>
        <button class="dropdown-item" on:click={handleItemClick(onOpenFile)}>Open</button>
        <button class="dropdown-item" on:click={handleItemClick(onSaveFile)}>Save</button>
        <button class="dropdown-item" on:click={handleItemClick(onSaveFileAs)}>Save As</button>
    </MenuItem>

    <MenuItem label="Edit" isOpen={showEditDropdown} onToggle={toggleEditDropdown}>
        <button class="dropdown-item" on:click={handleItemClick(onCut)}>Cut</button>
        <button class="dropdown-item" on:click={handleItemClick(onCopy)}>Copy</button>
        <button class="dropdown-item" on:click={handleItemClick(onPaste)}>Paste</button>
    </MenuItem>

    <MenuItem label="Format" isOpen={showFormatDropdown} onToggle={toggleFormatDropdown}>
        <button class="dropdown-item" on:click={handleItemClick(onToggleWordWrap)}>
            Word Wrap {wordWrapEnabled ? '✔' : ''}
        </button>
    </MenuItem>
</div>