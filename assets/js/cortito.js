document.body.addEventListener("htmx:afterRequest", function (evt) {
    const { xhr } = evt.detail;
    const { response } = xhr;
    const cortitoUrlResponse = JSON.parse(response);
    const cortitoUrlDisplay$ = document.querySelector("#cortito-url-display");
    let elementToDisplay = "";
    if (xhr.status === 429) {
        cortitoUrlDisplay$.classList.add("error");
        elementToDisplay = `<p>The quota has been exceeded!<br/>Try it later</p>`;
        setTimeout(() => {
            cortitoUrlDisplay$.classList.remove("show");
            cortitoUrlDisplay$.classList.remove("error");
        }, 3000);
    } else {
        const protocol = window.location.protocol;
        const port = window.location.port ? `:${window.location.port}` : "";
        const url = `${protocol}//${window.location.hostname}${port}/r/${cortitoUrlResponse.data.shortUrl}`;
        elementToDisplay = `<a href=${url} target="_blank" rel="noopener noreferrer">${url}</a>`;
    }

    cortitoUrlDisplay$.classList.add("show");
    cortitoUrlDisplay$.innerHTML = elementToDisplay;
});

document.body.addEventListener("htmx:beforeRequest", function (evt) {
    const cortitoUrlDisplay$ = document.querySelector("#cortito-url-display");
    cortitoUrlDisplay$.classList.toggle("hide");
    cortitoUrlDisplay$.classList.toggle("show");
});

document
    .querySelector("#cortito-url-display")
    ?.addEventListener("click", function () {
        navigator.permissions
            .query({ name: "clipboard-write" })
            .then((result) => {
                if (result.state === "granted" || result.state === "prompt") {
                    const cortitoUrlDisplay$ = document.querySelector(
                        "#cortito-url-display"
                    );
                    const cortitoAnchor$ =
                        cortitoUrlDisplay$.querySelector("a");
                    navigator.clipboard
                        .writeText(cortitoAnchor$.href)
                        .then(() => {
                            // Update text to notify that the text was copied to the clipboard
                            cortitoAnchor$.text = "COPIED!";
                            cortitoUrlDisplay$.classList.add("success");
                            setTimeout(() => {
                                cortitoUrlDisplay$.classList.remove("show");
                                cortitoUrlDisplay$.classList.remove("success");
                            }, 1500);
                        });
                }
            });
    });
