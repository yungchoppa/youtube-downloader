document.getElementById('convert-form').addEventListener('submit', async function(e) {
    e.preventDefault();
    const url = document.getElementById('yt-url').value;
    await fetch('/api/getMediaDetails', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: 'url=' + encodeURIComponent(url)
    });
});
