<script>
    export let label = '';
    export let isOpen = false;
    export let onToggle = () => {};

    function handleKeydown(event) {
        if (event.key === 'Enter' || event.key === ' ') {
            event.preventDefault();
            onToggle();
        }
    }
</script>

<div class="menu-item" tabindex="0" on:click={onToggle} on:keydown={handleKeydown}>
    {label}
    <div class="dropdown" class:visible={isOpen}>
        <slot />
    </div>
</div>