<script>
    export let content = '';
    export let currentFile = '';

    // Calculate the current line number and character length
    $: lines = content.split('\n').length;
    $: charLength = content.length;
    $: fileType = currentFile ? currentFile.split('.').pop() || 'txt' : 'txt';
    $: encoding = 'UTF-8'; // Placeholder for encoding
</script>

<div class="status-bar">
    <span>Line: {lines}</span>
    <span>Length: {charLength}</span>
    <span>File Type: .{fileType}</span>
    <span>Encoding: {encoding}</span>
</div>