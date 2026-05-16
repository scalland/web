// Fetch wrapper
async function api(method, url, data) {
    const opts = {method: method, headers: {'Content-Type': 'application/json'}};
    if (data) opts.body = JSON.stringify(data);
    const res = await fetch(url, opts);
    return res.json();
}
