document.getElementById('convert-form').addEventListener('submit', async function(e) {
    e.preventDefault();
    const url = document.getElementById('yt-url').value;
    const resp = await fetch('/api/getMediaDetails', {
        method: 'POST',
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
        body: 'url=' + encodeURIComponent(url)
    });
    const formats = await resp.json();
    const resultBlock = document.getElementById('result-block');
    resultBlock.innerHTML = '';
    if (formats.length === 0) {
        resultBlock.innerHTML = '<div>Нет доступных форматов mp3/mp4</div>';
        return;
    }
    resultBlock.innerHTML = `<form id="download-form">
        <div style="margin-bottom:1rem;">Выберите формат:</div>
        ${formats.map((f, i) => {
            let quality = f.ext === 'mp3' ? (f.tbr ? f.tbr + ' kbps' : '') : (f.height ? f.height + 'p' : '');
            let size = f.filesize ? (f.filesize / 1024 / 1024).toFixed(2) + ' MB' : '';
            let label = `${f.ext.toUpperCase()} ${quality} ${size}`;
            return `<div><input type="radio" name="format" id="format${i}" value="${f.format_id}" data-ext="${f.ext}" data-label="${label}" data-formatid="${f.format_id}" required> <label for="format${i}">${label}</label></div>`;
        }).join('')}
        <button type="submit" style="margin-top:1rem;">Download</button>
    </form>`;

    document.getElementById('download-form').addEventListener('submit', async function(ev) {
        ev.preventDefault();
        const checked = document.querySelector('input[name="format"]:checked');
        if (!checked) return;
        const formatID = checked.value;
        const ext = checked.getAttribute('data-ext');
        const label = checked.getAttribute('data-label');
        const filename = `download.${ext}`;
        const formData = new URLSearchParams();
        formData.append('url', url);
        formData.append('format_id', formatID);
        formData.append('filename', filename);
        // Скачивание через создание временной ссылки
        const downloadUrl = '/api/download';
        const response = await fetch(downloadUrl, {
            method: 'POST',
            body: formData,
        });
        if (!response.ok) {
            alert('Ошибка при скачивании');
            return;
        }
        const blob = await response.blob();
        const link = document.createElement('a');
        link.href = window.URL.createObjectURL(blob);
        link.download = filename;
        document.body.appendChild(link);
        link.click();
        link.remove();
    });
});
