document.body.addEventListener("htmx:afterRequest", function (evt) {
    const { response } = evt.detail.xhr;
    const cortitoUrl = JSON.parse(response);

    const protocol = window.location.protocol;
    const port = window.location.port ? `:${window.location.port}` : "";
    const url = `${protocol}//${window.location.hostname}${port}/r/${cortitoUrl.shortUrl}`;

    const cortitoUrlDisplay$ = document.querySelector("#cortito-url-display");
    cortitoUrlDisplay$.innerHTML = `<a href=${url} target="_blank" rel="noopener noreferrer">${url}</a>`;
});
