document.body.addEventListener("htmx:afterRequest", function (evt) {
    const { response } = evt.detail.xhr;
    const cortitoUrl = JSON.parse(response);

    const protocol = window.location.protocol;
    const port = window.location.port ? `:${window.location.port}` : "";
    const url = `${protocol}//${window.location.hostname}${port}/r/${cortitoUrl.shortUrl}`;

    const cortitoUrlDisplay$ = document.querySelector("#cortito-url-display");
    cortitoUrlDisplay$.classList.add("show");
    cortitoUrlDisplay$.innerHTML = `<a href=${url} target="_blank" rel="noopener noreferrer">${url}</a>`;
});

document.body.addEventListener("htmx:beforeRequest", function (evt) {
    const cortitoUrlDisplay$ = document.querySelector("#cortito-url-display");
    cortitoUrlDisplay$.classList.remove("show");
});
