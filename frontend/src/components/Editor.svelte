<script>
    export let content = '';
    export let wordWrapEnabled = true;
</script>

<textarea
        class="editor"
        class:no-wrap={!wordWrapEnabled}
        bind:value={content}
        style="resize: none;"
></textarea>